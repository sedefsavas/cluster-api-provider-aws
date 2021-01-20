// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	apiv1alpha4 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha4"
	v1alpha4 "sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1alpha4"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterapiapiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlane)(nil), (*v1alpha4.AWSManagedControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane(a.(*AWSManagedControlPlane), b.(*v1alpha4.AWSManagedControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.AWSManagedControlPlane)(nil), (*AWSManagedControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(a.(*v1alpha4.AWSManagedControlPlane), b.(*AWSManagedControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneList)(nil), (*v1alpha4.AWSManagedControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList(a.(*AWSManagedControlPlaneList), b.(*v1alpha4.AWSManagedControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.AWSManagedControlPlaneList)(nil), (*AWSManagedControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(a.(*v1alpha4.AWSManagedControlPlaneList), b.(*AWSManagedControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneSpec)(nil), (*v1alpha4.AWSManagedControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec(a.(*AWSManagedControlPlaneSpec), b.(*v1alpha4.AWSManagedControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.AWSManagedControlPlaneSpec)(nil), (*AWSManagedControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(a.(*v1alpha4.AWSManagedControlPlaneSpec), b.(*AWSManagedControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneStatus)(nil), (*v1alpha4.AWSManagedControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus(a.(*AWSManagedControlPlaneStatus), b.(*v1alpha4.AWSManagedControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.AWSManagedControlPlaneStatus)(nil), (*AWSManagedControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(a.(*v1alpha4.AWSManagedControlPlaneStatus), b.(*AWSManagedControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneLoggingSpec)(nil), (*v1alpha4.ControlPlaneLoggingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1alpha4_ControlPlaneLoggingSpec(a.(*ControlPlaneLoggingSpec), b.(*v1alpha4.ControlPlaneLoggingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.ControlPlaneLoggingSpec)(nil), (*ControlPlaneLoggingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(a.(*v1alpha4.ControlPlaneLoggingSpec), b.(*ControlPlaneLoggingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EncryptionConfig)(nil), (*v1alpha4.EncryptionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EncryptionConfig_To_v1alpha4_EncryptionConfig(a.(*EncryptionConfig), b.(*v1alpha4.EncryptionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.EncryptionConfig)(nil), (*EncryptionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EncryptionConfig_To_v1alpha3_EncryptionConfig(a.(*v1alpha4.EncryptionConfig), b.(*EncryptionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EndpointAccess)(nil), (*v1alpha4.EndpointAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess(a.(*EndpointAccess), b.(*v1alpha4.EndpointAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.EndpointAccess)(nil), (*EndpointAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess(a.(*v1alpha4.EndpointAccess), b.(*EndpointAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IAMAuthenticatorConfig)(nil), (*v1alpha4.IAMAuthenticatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_IAMAuthenticatorConfig_To_v1alpha4_IAMAuthenticatorConfig(a.(*IAMAuthenticatorConfig), b.(*v1alpha4.IAMAuthenticatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.IAMAuthenticatorConfig)(nil), (*IAMAuthenticatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(a.(*v1alpha4.IAMAuthenticatorConfig), b.(*IAMAuthenticatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubernetesMapping)(nil), (*v1alpha4.KubernetesMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(a.(*KubernetesMapping), b.(*v1alpha4.KubernetesMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.KubernetesMapping)(nil), (*KubernetesMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(a.(*v1alpha4.KubernetesMapping), b.(*KubernetesMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OIDCProviderStatus)(nil), (*v1alpha4.OIDCProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus(a.(*OIDCProviderStatus), b.(*v1alpha4.OIDCProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.OIDCProviderStatus)(nil), (*OIDCProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(a.(*v1alpha4.OIDCProviderStatus), b.(*OIDCProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleMapping)(nil), (*v1alpha4.RoleMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_RoleMapping_To_v1alpha4_RoleMapping(a.(*RoleMapping), b.(*v1alpha4.RoleMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.RoleMapping)(nil), (*RoleMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_RoleMapping_To_v1alpha3_RoleMapping(a.(*v1alpha4.RoleMapping), b.(*RoleMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserMapping)(nil), (*v1alpha4.UserMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_UserMapping_To_v1alpha4_UserMapping(a.(*UserMapping), b.(*v1alpha4.UserMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.UserMapping)(nil), (*UserMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_UserMapping_To_v1alpha3_UserMapping(a.(*v1alpha4.UserMapping), b.(*UserMapping), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane(in *AWSManagedControlPlane, out *v1alpha4.AWSManagedControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane(in *AWSManagedControlPlane, out *v1alpha4.AWSManagedControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane(in, out, s)
}

func autoConvert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in *v1alpha4.AWSManagedControlPlane, out *AWSManagedControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane is an autogenerated conversion function.
func Convert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in *v1alpha4.AWSManagedControlPlane, out *AWSManagedControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in, out, s)
}

func autoConvert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList(in *AWSManagedControlPlaneList, out *v1alpha4.AWSManagedControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha4.AWSManagedControlPlane)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList(in *AWSManagedControlPlaneList, out *v1alpha4.AWSManagedControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList(in, out, s)
}

func autoConvert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in *v1alpha4.AWSManagedControlPlaneList, out *AWSManagedControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]AWSManagedControlPlane)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in *v1alpha4.AWSManagedControlPlaneList, out *AWSManagedControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in, out, s)
}

func autoConvert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec(in *AWSManagedControlPlaneSpec, out *v1alpha4.AWSManagedControlPlaneSpec, s conversion.Scope) error {
	out.EKSClusterName = in.EKSClusterName
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.NetworkSpec, &out.NetworkSpec, 0); err != nil {
		return err
	}
	out.SecondaryCidrBlock = (*string)(unsafe.Pointer(in.SecondaryCidrBlock))
	out.Region = in.Region
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.RoleName = (*string)(unsafe.Pointer(in.RoleName))
	out.RoleAdditionalPolicies = (*[]string)(unsafe.Pointer(in.RoleAdditionalPolicies))
	out.Logging = (*v1alpha4.ControlPlaneLoggingSpec)(unsafe.Pointer(in.Logging))
	out.EncryptionConfig = (*v1alpha4.EncryptionConfig)(unsafe.Pointer(in.EncryptionConfig))
	out.AdditionalTags = *(*apiv1alpha4.Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMAuthenticatorConfig = (*v1alpha4.IAMAuthenticatorConfig)(unsafe.Pointer(in.IAMAuthenticatorConfig))
	if err := Convert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess(&in.EndpointAccess, &out.EndpointAccess, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, 0); err != nil {
		return err
	}
	out.ImageLookupFormat = in.ImageLookupFormat
	out.ImageLookupOrg = in.ImageLookupOrg
	out.ImageLookupBaseOS = in.ImageLookupBaseOS
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Bastion, &out.Bastion, 0); err != nil {
		return err
	}
	out.TokenMethod = (*v1alpha4.EKSTokenMethod)(unsafe.Pointer(in.TokenMethod))
	out.AssociateOIDCProvider = in.AssociateOIDCProvider
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec(in *AWSManagedControlPlaneSpec, out *v1alpha4.AWSManagedControlPlaneSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneSpec_To_v1alpha4_AWSManagedControlPlaneSpec(in, out, s)
}

func autoConvert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(in *v1alpha4.AWSManagedControlPlaneSpec, out *AWSManagedControlPlaneSpec, s conversion.Scope) error {
	out.EKSClusterName = in.EKSClusterName
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.NetworkSpec, &out.NetworkSpec, 0); err != nil {
		return err
	}
	out.SecondaryCidrBlock = (*string)(unsafe.Pointer(in.SecondaryCidrBlock))
	out.Region = in.Region
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.RoleName = (*string)(unsafe.Pointer(in.RoleName))
	out.RoleAdditionalPolicies = (*[]string)(unsafe.Pointer(in.RoleAdditionalPolicies))
	out.Logging = (*ControlPlaneLoggingSpec)(unsafe.Pointer(in.Logging))
	out.EncryptionConfig = (*EncryptionConfig)(unsafe.Pointer(in.EncryptionConfig))
	out.AdditionalTags = *(*apiv1alpha3.Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMAuthenticatorConfig = (*IAMAuthenticatorConfig)(unsafe.Pointer(in.IAMAuthenticatorConfig))
	if err := Convert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess(&in.EndpointAccess, &out.EndpointAccess, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, 0); err != nil {
		return err
	}
	out.ImageLookupFormat = in.ImageLookupFormat
	out.ImageLookupOrg = in.ImageLookupOrg
	out.ImageLookupBaseOS = in.ImageLookupBaseOS
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Bastion, &out.Bastion, 0); err != nil {
		return err
	}
	out.TokenMethod = (*EKSTokenMethod)(unsafe.Pointer(in.TokenMethod))
	out.AssociateOIDCProvider = in.AssociateOIDCProvider
	return nil
}

// Convert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec is an autogenerated conversion function.
func Convert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(in *v1alpha4.AWSManagedControlPlaneSpec, out *AWSManagedControlPlaneSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(in, out, s)
}

func autoConvert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus(in *AWSManagedControlPlaneStatus, out *v1alpha4.AWSManagedControlPlaneStatus, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Network, &out.Network, 0); err != nil {
		return err
	}
	out.FailureDomains = *(*clusterapiapiv1alpha4.FailureDomains)(unsafe.Pointer(&in.FailureDomains))
	out.Bastion = (*apiv1alpha4.Instance)(unsafe.Pointer(in.Bastion))
	if err := Convert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus(&in.OIDCProvider, &out.OIDCProvider, s); err != nil {
		return err
	}
	out.ExternalManagedControlPlane = (*bool)(unsafe.Pointer(in.ExternalManagedControlPlane))
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Conditions = *(*clusterapiapiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus(in *AWSManagedControlPlaneStatus, out *v1alpha4.AWSManagedControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneStatus_To_v1alpha4_AWSManagedControlPlaneStatus(in, out, s)
}

func autoConvert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(in *v1alpha4.AWSManagedControlPlaneStatus, out *AWSManagedControlPlaneStatus, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Network, &out.Network, 0); err != nil {
		return err
	}
	out.FailureDomains = *(*clusterapiapiv1alpha3.FailureDomains)(unsafe.Pointer(&in.FailureDomains))
	out.Bastion = (*apiv1alpha3.Instance)(unsafe.Pointer(in.Bastion))
	if err := Convert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(&in.OIDCProvider, &out.OIDCProvider, s); err != nil {
		return err
	}
	out.ExternalManagedControlPlane = (*bool)(unsafe.Pointer(in.ExternalManagedControlPlane))
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Conditions = *(*clusterapiapiv1alpha3.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(in *v1alpha4.AWSManagedControlPlaneStatus, out *AWSManagedControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(in, out, s)
}

func autoConvert_v1alpha3_ControlPlaneLoggingSpec_To_v1alpha4_ControlPlaneLoggingSpec(in *ControlPlaneLoggingSpec, out *v1alpha4.ControlPlaneLoggingSpec, s conversion.Scope) error {
	out.APIServer = in.APIServer
	out.Audit = in.Audit
	out.Authenticator = in.Authenticator
	out.ControllerManager = in.ControllerManager
	out.Scheduler = in.Scheduler
	return nil
}

// Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1alpha4_ControlPlaneLoggingSpec is an autogenerated conversion function.
func Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1alpha4_ControlPlaneLoggingSpec(in *ControlPlaneLoggingSpec, out *v1alpha4.ControlPlaneLoggingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_ControlPlaneLoggingSpec_To_v1alpha4_ControlPlaneLoggingSpec(in, out, s)
}

func autoConvert_v1alpha4_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in *v1alpha4.ControlPlaneLoggingSpec, out *ControlPlaneLoggingSpec, s conversion.Scope) error {
	out.APIServer = in.APIServer
	out.Audit = in.Audit
	out.Authenticator = in.Authenticator
	out.ControllerManager = in.ControllerManager
	out.Scheduler = in.Scheduler
	return nil
}

// Convert_v1alpha4_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec is an autogenerated conversion function.
func Convert_v1alpha4_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in *v1alpha4.ControlPlaneLoggingSpec, out *ControlPlaneLoggingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in, out, s)
}

func autoConvert_v1alpha3_EncryptionConfig_To_v1alpha4_EncryptionConfig(in *EncryptionConfig, out *v1alpha4.EncryptionConfig, s conversion.Scope) error {
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	out.Resources = *(*[]*string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha3_EncryptionConfig_To_v1alpha4_EncryptionConfig is an autogenerated conversion function.
func Convert_v1alpha3_EncryptionConfig_To_v1alpha4_EncryptionConfig(in *EncryptionConfig, out *v1alpha4.EncryptionConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_EncryptionConfig_To_v1alpha4_EncryptionConfig(in, out, s)
}

func autoConvert_v1alpha4_EncryptionConfig_To_v1alpha3_EncryptionConfig(in *v1alpha4.EncryptionConfig, out *EncryptionConfig, s conversion.Scope) error {
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	out.Resources = *(*[]*string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha4_EncryptionConfig_To_v1alpha3_EncryptionConfig is an autogenerated conversion function.
func Convert_v1alpha4_EncryptionConfig_To_v1alpha3_EncryptionConfig(in *v1alpha4.EncryptionConfig, out *EncryptionConfig, s conversion.Scope) error {
	return autoConvert_v1alpha4_EncryptionConfig_To_v1alpha3_EncryptionConfig(in, out, s)
}

func autoConvert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess(in *EndpointAccess, out *v1alpha4.EndpointAccess, s conversion.Scope) error {
	out.Public = (*bool)(unsafe.Pointer(in.Public))
	out.PublicCIDRs = *(*[]*string)(unsafe.Pointer(&in.PublicCIDRs))
	out.Private = (*bool)(unsafe.Pointer(in.Private))
	return nil
}

// Convert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess is an autogenerated conversion function.
func Convert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess(in *EndpointAccess, out *v1alpha4.EndpointAccess, s conversion.Scope) error {
	return autoConvert_v1alpha3_EndpointAccess_To_v1alpha4_EndpointAccess(in, out, s)
}

func autoConvert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess(in *v1alpha4.EndpointAccess, out *EndpointAccess, s conversion.Scope) error {
	out.Public = (*bool)(unsafe.Pointer(in.Public))
	out.PublicCIDRs = *(*[]*string)(unsafe.Pointer(&in.PublicCIDRs))
	out.Private = (*bool)(unsafe.Pointer(in.Private))
	return nil
}

// Convert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess is an autogenerated conversion function.
func Convert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess(in *v1alpha4.EndpointAccess, out *EndpointAccess, s conversion.Scope) error {
	return autoConvert_v1alpha4_EndpointAccess_To_v1alpha3_EndpointAccess(in, out, s)
}

func autoConvert_v1alpha3_IAMAuthenticatorConfig_To_v1alpha4_IAMAuthenticatorConfig(in *IAMAuthenticatorConfig, out *v1alpha4.IAMAuthenticatorConfig, s conversion.Scope) error {
	out.RoleMappings = *(*[]v1alpha4.RoleMapping)(unsafe.Pointer(&in.RoleMappings))
	out.UserMappings = *(*[]v1alpha4.UserMapping)(unsafe.Pointer(&in.UserMappings))
	return nil
}

// Convert_v1alpha3_IAMAuthenticatorConfig_To_v1alpha4_IAMAuthenticatorConfig is an autogenerated conversion function.
func Convert_v1alpha3_IAMAuthenticatorConfig_To_v1alpha4_IAMAuthenticatorConfig(in *IAMAuthenticatorConfig, out *v1alpha4.IAMAuthenticatorConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_IAMAuthenticatorConfig_To_v1alpha4_IAMAuthenticatorConfig(in, out, s)
}

func autoConvert_v1alpha4_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in *v1alpha4.IAMAuthenticatorConfig, out *IAMAuthenticatorConfig, s conversion.Scope) error {
	out.RoleMappings = *(*[]RoleMapping)(unsafe.Pointer(&in.RoleMappings))
	out.UserMappings = *(*[]UserMapping)(unsafe.Pointer(&in.UserMappings))
	return nil
}

// Convert_v1alpha4_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig is an autogenerated conversion function.
func Convert_v1alpha4_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in *v1alpha4.IAMAuthenticatorConfig, out *IAMAuthenticatorConfig, s conversion.Scope) error {
	return autoConvert_v1alpha4_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in, out, s)
}

func autoConvert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(in *KubernetesMapping, out *v1alpha4.KubernetesMapping, s conversion.Scope) error {
	out.UserName = in.UserName
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping is an autogenerated conversion function.
func Convert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(in *KubernetesMapping, out *v1alpha4.KubernetesMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(in, out, s)
}

func autoConvert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(in *v1alpha4.KubernetesMapping, out *KubernetesMapping, s conversion.Scope) error {
	out.UserName = in.UserName
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping is an autogenerated conversion function.
func Convert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(in *v1alpha4.KubernetesMapping, out *KubernetesMapping, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(in, out, s)
}

func autoConvert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus(in *OIDCProviderStatus, out *v1alpha4.OIDCProviderStatus, s conversion.Scope) error {
	out.ARN = in.ARN
	out.TrustPolicy = in.TrustPolicy
	return nil
}

// Convert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus is an autogenerated conversion function.
func Convert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus(in *OIDCProviderStatus, out *v1alpha4.OIDCProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_OIDCProviderStatus_To_v1alpha4_OIDCProviderStatus(in, out, s)
}

func autoConvert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in *v1alpha4.OIDCProviderStatus, out *OIDCProviderStatus, s conversion.Scope) error {
	out.ARN = in.ARN
	out.TrustPolicy = in.TrustPolicy
	return nil
}

// Convert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus is an autogenerated conversion function.
func Convert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in *v1alpha4.OIDCProviderStatus, out *OIDCProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in, out, s)
}

func autoConvert_v1alpha3_RoleMapping_To_v1alpha4_RoleMapping(in *RoleMapping, out *v1alpha4.RoleMapping, s conversion.Scope) error {
	out.RoleARN = in.RoleARN
	if err := Convert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_RoleMapping_To_v1alpha4_RoleMapping is an autogenerated conversion function.
func Convert_v1alpha3_RoleMapping_To_v1alpha4_RoleMapping(in *RoleMapping, out *v1alpha4.RoleMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_RoleMapping_To_v1alpha4_RoleMapping(in, out, s)
}

func autoConvert_v1alpha4_RoleMapping_To_v1alpha3_RoleMapping(in *v1alpha4.RoleMapping, out *RoleMapping, s conversion.Scope) error {
	out.RoleARN = in.RoleARN
	if err := Convert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_RoleMapping_To_v1alpha3_RoleMapping is an autogenerated conversion function.
func Convert_v1alpha4_RoleMapping_To_v1alpha3_RoleMapping(in *v1alpha4.RoleMapping, out *RoleMapping, s conversion.Scope) error {
	return autoConvert_v1alpha4_RoleMapping_To_v1alpha3_RoleMapping(in, out, s)
}

func autoConvert_v1alpha3_UserMapping_To_v1alpha4_UserMapping(in *UserMapping, out *v1alpha4.UserMapping, s conversion.Scope) error {
	out.UserARN = in.UserARN
	if err := Convert_v1alpha3_KubernetesMapping_To_v1alpha4_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_UserMapping_To_v1alpha4_UserMapping is an autogenerated conversion function.
func Convert_v1alpha3_UserMapping_To_v1alpha4_UserMapping(in *UserMapping, out *v1alpha4.UserMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_UserMapping_To_v1alpha4_UserMapping(in, out, s)
}

func autoConvert_v1alpha4_UserMapping_To_v1alpha3_UserMapping(in *v1alpha4.UserMapping, out *UserMapping, s conversion.Scope) error {
	out.UserARN = in.UserARN
	if err := Convert_v1alpha4_KubernetesMapping_To_v1alpha3_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_UserMapping_To_v1alpha3_UserMapping is an autogenerated conversion function.
func Convert_v1alpha4_UserMapping_To_v1alpha3_UserMapping(in *v1alpha4.UserMapping, out *UserMapping, s conversion.Scope) error {
	return autoConvert_v1alpha4_UserMapping_To_v1alpha3_UserMapping(in, out, s)
}