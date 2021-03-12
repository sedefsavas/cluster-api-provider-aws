/*
Copyright 2021 The Kubernetes Authors.

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
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	"time"
)

type AWSPrincipalTypeProvider interface {
	credentials.Provider
	// Hash returns a unique hash of the data forming the credentials
	// for this Principal
	Hash() (string, error)
	Name() string
}

func NewAWSControllerPrincipalTypeProvider() *AWSControllerPrincipalTypeProvider {
	def := defaults.Get()
	creds, _ := def.Config.Credentials.Get()
	fmt.Println(creds)

	return &AWSControllerPrincipalTypeProvider{
		credentials:     credentials.NewStaticCredentials(creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken),
		accessKeyID:     creds.AccessKeyID,
		secretAccessKey: creds.SecretAccessKey,
		sessionToken:    creds.SessionToken,
	}
}

type AWSControllerPrincipalTypeProvider struct {
	credentials *credentials.Credentials
	// these are for tests :/
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
}

func (p *AWSControllerPrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	return p.credentials.Get()
}
func (p *AWSControllerPrincipalTypeProvider) IsExpired() bool {
	return p.credentials.IsExpired()
}
func (p *AWSControllerPrincipalTypeProvider) Name() string {
	return "controller-principal"
}
func (p *AWSControllerPrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}

func NewAWSStaticPrincipalTypeProvider(principal *infrav1.AWSClusterStaticPrincipal, secret *corev1.Secret) *AWSStaticPrincipalTypeProvider {
	accessKeyID := string(secret.Data["AccessKeyID"])
	secretAccessKey := string(secret.Data["SecretAccessKey"])
	sessionToken := string(secret.Data["SessionToken"])

	return &AWSStaticPrincipalTypeProvider{
		Principal:       principal,
		credentials:     credentials.NewStaticCredentials(accessKeyID, secretAccessKey, sessionToken),
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		sessionToken:    sessionToken,
	}
}

func GetAssumeRoleCredentials(principal *infrav1.AWSClusterRolePrincipal, awsConfig *aws.Config) *credentials.Credentials {
	sess := session.Must(session.NewSession(awsConfig))

	creds := stscreds.NewCredentials(sess, principal.Spec.RoleArn, func(p *stscreds.AssumeRoleProvider) {
		if principal.Spec.ExternalID != "" {
			p.ExternalID = aws.String(principal.Spec.ExternalID)
		}
		p.RoleSessionName = principal.Spec.SessionName
		if principal.Spec.InlinePolicy != "" {
			p.Policy = aws.String(principal.Spec.InlinePolicy)
		}
		p.Duration = time.Duration(principal.Spec.DurationSeconds) * time.Second
	})
	return creds
}

func NewAWSRolePrincipalTypeProvider(principal *infrav1.AWSClusterRolePrincipal, sourceProvider *AWSPrincipalTypeProvider, log logr.Logger) *AWSRolePrincipalTypeProvider {
	return &AWSRolePrincipalTypeProvider{
		credentials:    nil,
		Principal:      principal,
		sourceProvider: sourceProvider,
		log:            log.WithName("AWSRolePrincipalTypeProvider"),
	}
}

type AWSStaticPrincipalTypeProvider struct {
	Principal   *infrav1.AWSClusterStaticPrincipal
	credentials *credentials.Credentials
	// these are for tests :/
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
}

func (p *AWSStaticPrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}
func (p *AWSStaticPrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	return p.credentials.Get()
}
func (p *AWSStaticPrincipalTypeProvider) Name() string {
	return p.Principal.Name
}
func (p *AWSStaticPrincipalTypeProvider) IsExpired() bool {
	return p.credentials.IsExpired()
}

type AWSRolePrincipalTypeProvider struct {
	Principal      *infrav1.AWSClusterRolePrincipal
	credentials    *credentials.Credentials
	sourceProvider *AWSPrincipalTypeProvider
	log            logr.Logger
}

func (p *AWSRolePrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}

func (p *AWSRolePrincipalTypeProvider) Name() string {
	return p.Principal.Name
}
func (p *AWSRolePrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	if p.credentials == nil || p.IsExpired() {
		awsConfig := aws.NewConfig()
		if p.sourceProvider != nil {
			sourceCreds, err := (*p.sourceProvider).Retrieve()
			if err != nil {
				return credentials.Value{}, err
			}
			awsConfig = awsConfig.WithCredentials(credentials.NewStaticCredentialsFromCreds(sourceCreds))
		}

		creds := GetAssumeRoleCredentials(p.Principal, awsConfig)
		// Update credentials
		p.credentials = creds
	}
	return p.credentials.Get()
}

func (p *AWSRolePrincipalTypeProvider) IsExpired() bool {
	return p.credentials.IsExpired()
}
