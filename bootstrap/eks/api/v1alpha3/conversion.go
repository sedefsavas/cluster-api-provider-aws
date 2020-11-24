/*
Copyright 2020 The Kubernetes Authors.

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
	"sigs.k8s.io/cluster-api-provider-aws/bootstrap/eks/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *EKSConfig) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.EKSConfig)

	return Convert_v1alpha3_EKSConfig_To_v1alpha4_EKSConfig(src, dst, nil)
}

func (dst *EKSConfig) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.EKSConfig)

	return Convert_v1alpha4_EKSConfig_To_v1alpha3_EKSConfig(src, dst, nil)
}

func (src *EKSConfigList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.EKSConfigList)

	return Convert_v1alpha3_EKSConfigList_To_v1alpha4_EKSConfigList(src, dst, nil)
}

func (dst *EKSConfigList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.EKSConfigList)

	return Convert_v1alpha4_EKSConfigList_To_v1alpha3_EKSConfigList(src, dst, nil)
}

func (src *EKSConfigTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.EKSConfigTemplate)

	return Convert_v1alpha3_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate(src, dst, nil)
}

func (dst *EKSConfigTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.EKSConfigTemplate)

	return Convert_v1alpha4_EKSConfigTemplate_To_v1alpha3_EKSConfigTemplate(src, dst, nil)
}

func (src *EKSConfigTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.EKSConfigTemplateList)

	return Convert_v1alpha3_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList(src, dst, nil)
}

func (dst *EKSConfigTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.EKSConfigTemplateList)

	return Convert_v1alpha4_EKSConfigTemplateList_To_v1alpha3_EKSConfigTemplateList(src, dst, nil)
}
