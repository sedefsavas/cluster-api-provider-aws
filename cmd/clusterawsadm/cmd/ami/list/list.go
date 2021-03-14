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
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/flags"
	ec2service "sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services/ec2"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
)

var (
	kubernetesVersion string
	opSystem          string
)

func ListAMICmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "list",
		Short: "List AMIs from the default AWS account where AMIs are stored",
		Long: cmd.LongDesc(`
			List AMIs based on Kubernetes version, OS, region. If no arguments are provided,
			it will print all AMIs in all regions, OS types for the supported Kubernetes versions.
            Supported Kubernetes versions start from the latest stable version and goes 2 release back:
			if the latest stable release is v1.20.4- v1.19.x and v1.18.x are supported.
			To list AMIs of unsupported Kubernetes versions, --kubernetes-version flag needs to be provided.
		`),
		Example: cmd.Examples(`
		# List AMIs from the default AWS account where AMIs are stored.
		# Available os options: centos-7, ubuntu-18.04, ubuntu-20.04, amazon-2
		clusterawsadm ami list --kubernetes-version=v1.18.12 --os=ubuntu-20.04  --region=us-west-2
		# To list all supported AMIs in all supported Kubernetes versions, regions, and linux distributions:
		clusterawsadm ami list
		`),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			supportedOsList := []string{}
			if opSystem == "" {
				supportedOsList = getSupportedOsList()
			} else {
				supportedOsList = append(supportedOsList, opSystem)
			}
			imageRegionList := []string{}
			region := cmd.Flags().Lookup("region").Value.String()
			if region == "" {
				imageRegionList = getimageRegionList()
			} else {
				imageRegionList = append(imageRegionList, region)
			}

			supportedVersions := []string{}
			if kubernetesVersion == "" {
				var err error
				supportedVersions, err = getSupportedKubernetesVersions()
				if err != nil {
					fmt.Println("Failed to calculate supported Kubernetes versions")
					return err
				}
			} else {
				supportedVersions = append(supportedVersions, kubernetesVersion)
			}

			for _, version := range supportedVersions {
				fmt.Println("+------------------------------------------------------------+")
				fmt.Printf(fmt.Sprintf("|               %-15s%-10s               |\n", "KUBERNETES VERSION: ", version))
				fmt.Println("+------------------------------------------------------------+")
				fmt.Printf(fmt.Sprintf("%-15s %-20s %-15s\n", "OS", "REGION", "AMI-ID"))
				for _, os := range supportedOsList {
					fmt.Println("-------------------------------------------------------------")
					for _, region := range imageRegionList {
						sess, err := session.NewSessionWithOptions(session.Options{
							SharedConfigState: session.SharedConfigEnable,
							Config:            aws.Config{Region: aws.String(region)},
						})
						if err != nil {
							fmt.Printf("Error: %v\n", err)
							return err
						}
						ec2Client := ec2.New(sess)
						image, err := ec2service.DefaultAMILookup(ec2Client, ec2service.DefaultMachineAMIOwnerID, os, version, "")
						if err != nil {
							return err
						}
						fmt.Printf(fmt.Sprintf("%-15s %-20s %-15s\n", os, region, *image.ImageId))
					}
				}
				fmt.Println()
				fmt.Println()
			}

			return nil
		},
	}

	flags.AddRegionFlag(newCmd)
	addOsFlag(newCmd)
	addKubernetesVersionFlag(newCmd)
	return newCmd
}

func addOsFlag(c *cobra.Command) {
	c.Flags().StringVar(&opSystem, "os", "", "Operating system of the AMI to be copied")
}

func addKubernetesVersionFlag(c *cobra.Command) {
	c.Flags().StringVar(&kubernetesVersion, "kubernetes-version", "", "Kubernetes version of the AMI to be copied")
}
