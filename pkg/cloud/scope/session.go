/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scope

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/throttle"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ServiceEndpoint defines a tuple containing AWS Service resolution information
type ServiceEndpoint struct {
	ServiceID     string
	URL           string
	SigningRegion string
}

var sessionCache sync.Map
var providerCache sync.Map

type sessionCacheEntry struct {
	session         *session.Session
	serviceLimiters throttle.ServiceLimiters
}

func sessionForRegion(region string, endpoint []ServiceEndpoint) (*session.Session, throttle.ServiceLimiters, error) {
	if s, ok := sessionCache.Load(region); ok {
		entry := s.(*sessionCacheEntry)
		return entry.session, entry.serviceLimiters, nil
	}

	resolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		for _, s := range endpoint {
			if service == s.ServiceID {
				return endpoints.ResolvedEndpoint{
					URL:           s.URL,
					SigningRegion: s.SigningRegion,
				}, nil
			}
		}
		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}
	ns, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		EndpointResolver: endpoints.ResolverFunc(resolver),
	})
	if err != nil {
		return nil, nil, err
	}

	sl := newServiceLimiters()
	sessionCache.Store(region, &sessionCacheEntry{
		session:         ns,
		serviceLimiters: sl,
	})
	return ns, sl, nil
}

func sessionForClusterWithRegion(k8sClient client.Client, awsCluster *infrav1.AWSCluster, region string, endpoint []ServiceEndpoint, logger logr.Logger) (*session.Session, throttle.ServiceLimiters, error) {
	log := logger.WithName("identity")
	log.V(4).Info("Creating an AWS Session")

	resolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		for _, s := range endpoint {
			if service == s.ServiceID {
				return endpoints.ResolvedEndpoint{
					URL:           s.URL,
					SigningRegion: s.SigningRegion,
				}, nil
			}
		}
		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	providers, err := getProvidersForCluster(context.Background(), k8sClient, awsCluster, log)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to get providers for cluster")

	}

	isChanged := false
	awsProviders := make([]credentials.Provider, len(providers))
	for i, provider := range providers {
		// load an existing matching providers from the cache if such a providers exists
		providerHash, err := provider.Hash()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Failed to calculate provider hash.")
		}
		cachedProvider, ok := providerCache.Load(providerHash)
		if ok {
			provider = cachedProvider.(AWSPrincipalTypeProvider)
		} else {
			isChanged = true
			// add this providers to the cache
			providerCache.Store(providerHash, provider)
			// TODO: Remove old provider from cache
		}
		awsProviders[i] = provider.(credentials.Provider)
	}

	if !isChanged {
		if s, ok := sessionCache.Load(getSessionName(region, awsCluster)); ok {
			entry := s.(*sessionCacheEntry)
			return entry.session, entry.serviceLimiters, nil
		}
	}
	awsConfig := &aws.Config{
		Region:           aws.String(region),
		EndpointResolver: endpoints.ResolverFunc(resolver),
	}

	if len(providers) > 0 {
		// Check if principal credentials can be retrieved. One reason this will fail is that source principal is not authorized for assume role.
		_, err := providers[0].Retrieve()
		if err != nil {
			conditions.MarkFalse(awsCluster, infrav1.PrincipalCredentialRetrievedCondition, infrav1.PrincipalCredentialRetrievalFailedReason, clusterv1.ConditionSeverityError, err.Error())
			return nil, nil, errors.Wrap(err, "Failed to retrieve principal credentials")
		}
		awsConfig = awsConfig.WithCredentials(credentials.NewChainCredentials(awsProviders))
	}

	conditions.MarkTrue(awsCluster, infrav1.PrincipalCredentialRetrievedCondition)

	ns, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to create a new AWS session")
	}
	sl := newServiceLimiters()
	sessionCache.Store(getSessionName(region, awsCluster), &sessionCacheEntry{
		session:         ns,
		serviceLimiters: sl,
	})

	return ns, sl, nil
}

func getSessionName(region string, cluster *infrav1.AWSCluster) string{
	return fmt.Sprintf("%s-%s-%s", region, cluster.Name, cluster.Namespace)
}

func newServiceLimiters() throttle.ServiceLimiters {
	return throttle.ServiceLimiters{
		ec2.ServiceID:                      newEC2ServiceLimiter(),
		elb.ServiceID:                      newGenericServiceLimiter(),
		resourcegroupstaggingapi.ServiceID: newGenericServiceLimiter(),
		secretsmanager.ServiceID:           newGenericServiceLimiter(),
	}
}

func newGenericServiceLimiter() *throttle.ServiceLimiter {
	return &throttle.ServiceLimiter{
		{
			Operation:  throttle.NewMultiOperationMatch("Describe", "Get", "List"),
			RefillRate: 20.0,
			Burst:      100,
		},
		{
			Operation:  ".*",
			RefillRate: 5.0,
			Burst:      200,
		},
	}
}

func newEC2ServiceLimiter() *throttle.ServiceLimiter {
	return &throttle.ServiceLimiter{
		{
			Operation:  throttle.NewMultiOperationMatch("Describe", "Get"),
			RefillRate: 20.0,
			Burst:      100,
		},
		{
			Operation: throttle.NewMultiOperationMatch(
				"AuthorizeSecurityGroupIngress",
				"CancelSpotInstanceRequests",
				"CreateKeyPair",
				"RequestSpotInstances",
			),
			RefillRate: 20.0,
			Burst:      100,
		},
		{
			Operation:  "RunInstances",
			RefillRate: 2.0,
			Burst:      5,
		},
		{
			Operation:  "StartInstances",
			RefillRate: 2.0,
			Burst:      5,
		},
		{
			Operation:  ".*",
			RefillRate: 5.0,
			Burst:      200,
		},
	}
}

func buildProvidersForRef(ctx context.Context, providers []AWSPrincipalTypeProvider, k8sClient client.Client, awsCluster *infrav1.AWSCluster, ref *corev1.ObjectReference, log logr.Logger) ([]AWSPrincipalTypeProvider, error) {
	if ref == nil {
		log.V(4).Info("AWSCluster does not have a PrincipalRef specified")
		return providers, nil
	}

	var provider AWSPrincipalTypeProvider
	principalObjectKey := client.ObjectKey{Name: ref.Name}
	log.V(4).Info("Get Principal", "Key", principalObjectKey)
	switch ref.Kind {
	case "AWSClusterControllerPrincipal":
		principal := &infrav1.AWSClusterControllerPrincipal{}
		principal.Kind = "AWSClusterControllerPrincipal"

		err := k8sClient.Get(ctx, client.ObjectKey{Name: ref.Name}, principal)
		if err != nil {
			return providers, err
		}

		if !clusterIsPermittedToUsePrincipal(principal.Spec.AllowedNamespaces, awsCluster.Namespace) {
			if awsCluster.Spec.PrincipalRef.Name == principal.Name {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.PrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			} else {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.SourcePrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			}
			return providers, errors.Errorf("AWSCluster %s/%s is not permitted to use principal %s", awsCluster.Namespace, awsCluster.Name, principal.Name)
		}

		// returning empty provider list to default to Controller Principal.
		return []AWSPrincipalTypeProvider{}, nil
	case "AWSClusterStaticPrincipal":
		principal := &infrav1.AWSClusterStaticPrincipal{}
		err := k8sClient.Get(ctx, principalObjectKey, principal)
		if err != nil {
			return providers, err
		}
		secret := &corev1.Secret{}
		err = k8sClient.Get(ctx, client.ObjectKey{Name: principal.Spec.SecretRef.Name, Namespace: principal.Spec.SecretRef.Namespace}, secret)
		if err != nil {
			return providers, err
		}
		log.V(4).Info("Found an AWSClusterStaticPrincipal", "principal", principal.GetName())

		if !clusterIsPermittedToUsePrincipal(principal.Spec.AllowedNamespaces, awsCluster.Namespace) {
			if awsCluster.Spec.PrincipalRef.Name == principal.Name {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.PrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			} else {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.SourcePrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			}
			return providers, errors.Errorf("AWSCluster %s/%s is not permitted to use principal %s", awsCluster.Namespace, awsCluster.Name, principal.Name)
		}

		provider = NewAWSStaticPrincipalTypeProvider(principal, secret)
		providers = append(providers, provider)
	case "AWSClusterRolePrincipal":
		principal := &infrav1.AWSClusterRolePrincipal{}
		err := k8sClient.Get(ctx, principalObjectKey, principal)
		if err != nil {
			return providers, err
		}

		if !clusterIsPermittedToUsePrincipal(principal.Spec.AllowedNamespaces, awsCluster.Namespace) {
			if awsCluster.Spec.PrincipalRef.Name == principal.Name {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.PrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			} else {
				conditions.MarkFalse(awsCluster, infrav1.PrincipalUsageAllowedCondition, infrav1.SourcePrincipalUsageUnauthorizedReason, clusterv1.ConditionSeverityError, err.Error())
			}
			return providers, errors.Errorf("AWSCluster %s/%s is not permitted to use principal %s", awsCluster.Namespace, awsCluster.Name, principal.Name)
		}

		// TODO: SourcePrincipalRef should always be non-nil, add webhook defaulting.
		if principal.Spec.SourcePrincipalRef != nil {
			providers, err = buildProvidersForRef(ctx, providers, k8sClient, awsCluster, principal.Spec.SourcePrincipalRef, log)
			if err != nil {
				return providers, err
			}
		}
		var sourceProvider AWSPrincipalTypeProvider
		if len(providers) > 0 {
			sourceProvider = providers[len(providers)-1]
			// Remove last provider
			if len(providers) > 0 {
				providers = providers[:len(providers)-1]
			}
		}
		log.V(4).Info("Found an AWSClusterRolePrincipal", "principal", principal.GetName())
		if sourceProvider != nil {
			provider = NewAWSRolePrincipalTypeProvider(principal, &sourceProvider, log)
		} else {
			provider = NewAWSRolePrincipalTypeProvider(principal, nil, log)
		}
		providers = append(providers, provider)
	default:
		return providers, errors.Errorf("No such provider known: '%s'", ref.Kind)
	}
	conditions.MarkTrue(awsCluster, infrav1.PrincipalUsageAllowedCondition)
	return providers, nil
}

func getProvidersForCluster(ctx context.Context, k8sClient client.Client, awsCluster *infrav1.AWSCluster, log logr.Logger) ([]AWSPrincipalTypeProvider, error) {
	providers := make([]AWSPrincipalTypeProvider, 0)
	providers, err := buildProvidersForRef(ctx, providers, k8sClient, awsCluster, awsCluster.Spec.PrincipalRef, log)
	if err != nil {
		return nil, err
	}

	return providers, nil
}

func clusterIsPermittedToUsePrincipal(allowedNs *infrav1.AllowedNamespacesList, ns string) bool {
	// nil value does not match with any namespaces
	if allowedNs == nil {
		return false
	}

	// empty value matches with all namespaces
	if len(allowedNs.NamespacesList) == 0 {
		return true
	}

	for _, v := range (*allowedNs).NamespacesList {
		if v == ns {
			return true
		}
	}

	return false
}
