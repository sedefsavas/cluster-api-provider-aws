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

package v1alpha4

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"regexp"
)

var (
	sshKeyValidNameRegex = regexp.MustCompile(`^[[:graph:]]+([[:print:]]*[[:graph:]]+)*$`)
)

func aggregateObjErrors(gk schema.GroupKind, name string, allErrs field.ErrorList) error {
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		gk,
		name,
		allErrs,
	)
}

func isValidSSHKey(sshKey *string) field.ErrorList {
	var allErrs field.ErrorList
	if sshKey != nil {
		if sshKey != nil && !sshKeyValidNameRegex.Match([]byte(*sshKey)) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("sshKey"), sshKey, "Name is invalid. Must be specified in ASCII and must not start or end in whitespace"))
		}
	}

	return allErrs
}
