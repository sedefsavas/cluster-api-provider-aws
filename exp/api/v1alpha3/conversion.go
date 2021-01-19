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
	"sigs.k8s.io/cluster-api-provider-aws/exp/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (r *AWSMachinePool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachinePool)

	return Convert_v1alpha3_AWSMachinePool_To_v1alpha4_AWSMachinePool(r, dst, nil)
}

func (r *AWSMachinePool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachinePool)

	return Convert_v1alpha4_AWSMachinePool_To_v1alpha3_AWSMachinePool(src, r, nil)
}

func (r *AWSMachinePoolList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSMachinePoolList)

	return Convert_v1alpha3_AWSMachinePoolList_To_v1alpha4_AWSMachinePoolList(r, dst, nil)
}

func (r *AWSMachinePoolList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSMachinePoolList)

	return Convert_v1alpha4_AWSMachinePoolList_To_v1alpha3_AWSMachinePoolList(src, r, nil)
}

func (r *AWSManagedCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedCluster)

	return Convert_v1alpha3_AWSManagedCluster_To_v1alpha4_AWSManagedCluster(r, dst, nil)
}

func (r *AWSManagedCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedCluster)

	return Convert_v1alpha4_AWSManagedCluster_To_v1alpha3_AWSManagedCluster(src, r, nil)
}

func (r *AWSManagedClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedClusterList)

	return Convert_v1alpha3_AWSManagedClusterList_To_v1alpha4_AWSManagedClusterList(r, dst, nil)
}

func (r *AWSManagedClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedClusterList)

	return Convert_v1alpha4_AWSManagedClusterList_To_v1alpha3_AWSManagedClusterList(src, r, nil)
}

func (r *AWSManagedMachinePool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedMachinePool)

	return Convert_v1alpha3_AWSManagedMachinePool_To_v1alpha4_AWSManagedMachinePool(r, dst, nil)
}

func (r *AWSManagedMachinePool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedMachinePool)

	return Convert_v1alpha4_AWSManagedMachinePool_To_v1alpha3_AWSManagedMachinePool(src, r, nil)
}

func (r *AWSManagedMachinePoolList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedMachinePoolList)

	return Convert_v1alpha3_AWSManagedMachinePoolList_To_v1alpha4_AWSManagedMachinePoolList(r, dst, nil)
}

func (r *AWSManagedMachinePoolList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedMachinePoolList)

	return Convert_v1alpha4_AWSManagedMachinePoolList_To_v1alpha3_AWSManagedMachinePoolList(src, r, nil)
}
