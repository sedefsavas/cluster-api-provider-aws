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

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var _ = logf.Log.WithName("awsclustercontrollerprincipal-resource")

func (r *AWSClusterControllerPrincipal) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha3-awsclustercontrollerprincipal,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=awsclustercontrollerprincipals,versions=v1alpha3,name=validation.awsclustercontrollerprincipal.infrastructure.cluster.x-k8s.io,sideEffects=None
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1alpha3-awsclustercontrollerprincipal,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=awsclustercontrollerprincipals,versions=v1alpha3,name=default.awsclustercontrollerprincipal.infrastructure.cluster.x-k8s.io,sideEffects=None

var (
	_ webhook.Validator = &AWSClusterControllerPrincipal{}
	_ webhook.Defaulter = &AWSClusterControllerPrincipal{}
)

func (r *AWSClusterControllerPrincipal) ValidateCreate() error {
	return nil
}

func (r *AWSClusterControllerPrincipal) ValidateDelete() error {
	return nil
}

func (r *AWSClusterControllerPrincipal) ValidateUpdate(old runtime.Object) error {
	return nil
	//var allErrs field.ErrorList
	//
	//oldP, ok := old.(*AWSClusterControllerPrincipal)
	//if !ok {
	//	return apierrors.NewBadRequest(fmt.Sprintf("expected an AWSClusterControllerPrincipal but got a %T", old))
	//}
	//
	//// Validate selector parses as Selector
	//selector, err := v1.LabelSelectorAsSelector(&r.Spec.AllowedNamespaces)
	//if err != nil {
	//	allErrs = append(
	//		allErrs,
	//		field.Invalid(field.NewPath("spec", "allowedNamespaces"), r.Spec.AllowedNamespaces, err.Error()),
	//	)
	//}
	//
	//// Validate selector parses as Selector
	//oldSelector, err := v1.LabelSelectorAsSelector(&oldP.Spec.AllowedNamespaces)
	//if err != nil {
	//	allErrs = append(
	//		allErrs,
	//		field.Invalid(field.NewPath("spec", "allowedNamespaces"), r.Spec.AllowedNamespaces, err.Error()),
	//	)
	//}
	//
	//// Validate that the selector isn't empty.
	//// Only allow updates if existing spec.AllowedNamespaces includes all namespaces to avoid increasing permissions accidentally during upgrade.
	//if selector != nil && !selector.Empty() && !reflect.DeepEqual(selector, oldSelector) {
	//	allErrs = append(
	//		allErrs,
	//		field.Invalid(field.NewPath("spec", "allowedNamespaces"), r.Spec.AllowedNamespaces, "allowedNamespaces must be empty to be updated"),
	//	)
	//}
	//
	//return aggregateObjErrors(r.GroupVersionKind().GroupKind(), r.Name, allErrs)
}

func (r *AWSClusterControllerPrincipal) Default() {
}
