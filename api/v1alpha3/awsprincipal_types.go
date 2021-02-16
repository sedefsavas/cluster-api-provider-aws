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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AWSClusterPrincipalSpec struct {
	// +optional
	AllowedNamespaces *AllowedNamespacesList `json:"allowedNamespaces"`
}

type AllowedNamespacesList struct {
	NamespacesList []string `json:"list"`
}

type AWSRoleSpec struct {
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleArn string `json:"roleARN"`
	// An identifier for the assumed role session
	SessionName string `json:"sessionName,omitempty"`
	// The duration, in seconds, of the role session before it is renewed.
	// +kubebuilder:validation:Minimum:=900
	// +kubebuilder:validation:Maximum:=43200
	DurationSeconds int32 `json:"durationSeconds,omitempty"`
	// An IAM policy in JSON format that you want to use as an inline session policy.
	InlinePolicy string `json:"inlinePolicy,omitempty"`

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want
	// to use as managed session policies.
	// The policies must exist in the same account as the role.
	PolicyARNs []string `json:"policyARNs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=awsclusterstaticprincipals,scope=Cluster,categories=cluster-api
// +kubebuilder:storageversion

// AWSClusterStaticPrincipal represents a reference to an AWS access key ID and
// secret access key, stored in a secret.
type AWSClusterStaticPrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec for this AWSClusterStaticPrincipal.
	Spec AWSClusterStaticPrincipalSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// AWSClusterStaticPrincipalList contains a list of AWSClusterStaticPrincipal
type AWSClusterStaticPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSClusterStaticPrincipal `json:"items"`
}

type AWSClusterStaticPrincipalSpec struct {
	AWSClusterPrincipalSpec `json:",inline"`
	// Reference to a secret containing the credentials. The secret should
	// contain the following data keys:
	//  AccessKeyID: AKIAIOSFODNN7EXAMPLE
	//  SecretAccessKey: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
	//  SessionToken: Optional
	SecretRef corev1.SecretReference `json:"secretRef"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=awsclusterroleprincipals,scope=Cluster,categories=cluster-api
type AWSClusterRolePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec for this AWSClusterRolePrincipal.
	Spec AWSClusterRolePrincipalSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// AWSClusterRolePrincipalList contains a list of AWSClusterRolePrincipal
type AWSClusterRolePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSClusterRolePrincipal `json:"items"`
}

type AWSClusterRolePrincipalSpec struct {
	AWSClusterPrincipalSpec `json:",inline"`
	AWSRoleSpec             `json:",inline"`
	// A unique identifier that might be required when you assume a role in another account.
	// If the administrator of the account to which the role belongs provided you with an
	// external ID, then provide that value in the ExternalId parameter. This value can be
	// any string, such as a passphrase or account number. A cross-account role is usually
	// set up to trust everyone in an account. Therefore, the administrator of the trusting
	// account might send an external ID to the administrator of the trusted account. That
	// way, only someone with the ID can assume the role, rather than everyone in the
	// account. For more information about the external ID, see How to Use an External ID
	// When Granting Access to Your AWS Resources to a Third Party in the IAM User Guide.
	// +optional
	ExternalID string `json:"externalID,omitempty"`

	// SourcePrincipalRef is a reference to another principal which will be chained to do
	// role assumption. All principal types are accepted.
	SourcePrincipalRef *corev1.ObjectReference `json:"sourcePrincipalRef,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=awsclustercontrollerprincipals,scope=Cluster,categories=cluster-api
type AWSClusterControllerPrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec for this AWSClusterControllerPrincipal.
	Spec AWSClusterControllerPrincipalSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// AWSClusterControllerPrincipalList contains a list of AWSClusterControllerPrincipal
type AWSClusterControllerPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSClusterControllerPrincipal `json:"items"`
}

type AWSClusterControllerPrincipalSpec struct {
	AWSClusterPrincipalSpec `json:",inline"`
}

func init() {
	SchemeBuilder.Register(
		&AWSClusterStaticPrincipal{},
		&AWSClusterStaticPrincipalList{},
		&AWSClusterRolePrincipal{},
		&AWSClusterRolePrincipalList{},
		&AWSClusterControllerPrincipal{},
		&AWSClusterControllerPrincipalList{},
	)
}
