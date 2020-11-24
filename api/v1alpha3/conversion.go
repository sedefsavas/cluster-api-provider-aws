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

package v1alpha3

import (
	"sigs.k8s.io/cluster-api-provider-aws/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *AWSCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSCluster)

	return  Convert_v1alpha3_AWSCluster_To_v1alpha4_AWSCluster(src, dst, nil)
}

func (dst *AWSCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSCluster)

	return Convert_v1alpha4_AWSCluster_To_v1alpha3_AWSCluster(src, dst, nil)
}

func (src *AWSClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSClusterList)

	return Convert_v1alpha3_AWSClusterList_To_v1alpha4_AWSClusterList(src, dst, nil)
}

func (dst *AWSClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSClusterList)

	return Convert_v1alpha4_AWSClusterList_To_v1alpha3_AWSClusterList(src, dst, nil)
}

func (src *AWSMachine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachine)

	return Convert_v1alpha3_AWSMachine_To_v1alpha4_AWSMachine(src, dst, nil)
}

func (dst *AWSMachine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachine)

	return Convert_v1alpha4_AWSMachine_To_v1alpha3_AWSMachine(src, dst, nil)
}

func (src *AWSMachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineList)

	return Convert_v1alpha3_AWSMachineList_To_v1alpha4_AWSMachineList(src, dst, nil)
}

func (dst *AWSMachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineList)

	return Convert_v1alpha4_AWSMachineList_To_v1alpha3_AWSMachineList(src, dst, nil)
}

func (src *AWSMachineTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineTemplate)

	return Convert_v1alpha3_AWSMachineTemplate_To_v1alpha4_AWSMachineTemplate(src, dst, nil)
}

func (dst *AWSMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineTemplate)

	return Convert_v1alpha4_AWSMachineTemplate_To_v1alpha3_AWSMachineTemplate(src, dst, nil)
}

func (src *AWSMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineTemplateList)

	return Convert_v1alpha3_AWSMachineTemplateList_To_v1alpha4_AWSMachineTemplateList(src, dst, nil)
}

func (dst *AWSMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineTemplateList)

	return Convert_v1alpha4_AWSMachineTemplateList_To_v1alpha3_AWSMachineTemplateList(src, dst, nil)
}
