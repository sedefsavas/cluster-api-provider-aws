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
	"sigs.k8s.io/cluster-api-provider-aws/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (r *AWSCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSCluster)

	return Convert_v1alpha3_AWSCluster_To_v1alpha4_AWSCluster(r, dst, nil)
}

func (r *AWSCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSCluster)

	return Convert_v1alpha4_AWSCluster_To_v1alpha3_AWSCluster(src, r, nil)
}

func (r *AWSClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSClusterList)

	return Convert_v1alpha3_AWSClusterList_To_v1alpha4_AWSClusterList(r, dst, nil)
}

func (r *AWSClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSClusterList)

	return Convert_v1alpha4_AWSClusterList_To_v1alpha3_AWSClusterList(src, r, nil)
}

func (r *AWSMachine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachine)

	return Convert_v1alpha3_AWSMachine_To_v1alpha4_AWSMachine(r, dst, nil)
}

func (r *AWSMachine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachine)

	return Convert_v1alpha4_AWSMachine_To_v1alpha3_AWSMachine(src, r, nil)
}

func (r *AWSMachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineList)

	return Convert_v1alpha3_AWSMachineList_To_v1alpha4_AWSMachineList(r, dst, nil)
}

func (r *AWSMachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineList)

	return Convert_v1alpha4_AWSMachineList_To_v1alpha3_AWSMachineList(src, r, nil)
}

func (r *AWSMachineTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineTemplate)

	return Convert_v1alpha3_AWSMachineTemplate_To_v1alpha4_AWSMachineTemplate(r, dst, nil)
}

func (r *AWSMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineTemplate)

	return Convert_v1alpha4_AWSMachineTemplate_To_v1alpha3_AWSMachineTemplate(src, r, nil)
}

func (r *AWSMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachineTemplateList)

	return Convert_v1alpha3_AWSMachineTemplateList_To_v1alpha4_AWSMachineTemplateList(r, dst, nil)
}

func (r *AWSMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachineTemplateList)

	return Convert_v1alpha4_AWSMachineTemplateList_To_v1alpha3_AWSMachineTemplateList(src, r, nil)
}
