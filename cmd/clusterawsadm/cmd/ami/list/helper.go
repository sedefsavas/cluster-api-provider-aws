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

package copy

import (
	"io/ioutil"
	"net/http"
	"strings"

	"sigs.k8s.io/cluster-api/test/framework/kubernetesversions"
)

const latestStableReleaseURL = "https://dl.k8s.io/release/stable.txt"

func getSupportedOsList() []string {
	return []string{"centos-7", "ubuntu-18.04", "ubuntu-20.04", "amazon-2"}
}

func getimageRegionList() []string {
	return []string{
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-south-1",
		"ap-southeast-1",
		"ap-northeast-2",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"sa-east-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	}
}

func getSupportedKubernetesVersions() ([]string, error) {
	supportedVersions := make([]string, 0)
	latestVersion, err := latestStableRelease()
	if err != nil {
		return nil, err
	}
	supportedVersions = append(supportedVersions, latestVersion)
	nMinusOne, err := kubernetesversions.PreviousMinorRelease(latestVersion)
	if err != nil {
		return nil, err
	}
	supportedVersions = append(supportedVersions, nMinusOne)

	nMinusTwo, err := kubernetesversions.PreviousMinorRelease(nMinusOne)
	if err != nil {
		return nil, err
	}
	supportedVersions = append(supportedVersions, nMinusTwo)

	return supportedVersions, nil
}

// latestStableRelease fetches the latest stable Kubernetes version
func latestStableRelease() (string, error) {
	resp, err := http.Get(latestStableReleaseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
