/*
Copyright 2019 The KubeOne Authors.

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

package credentials

import (
	"github.com/kubermatic/kubeone/pkg/apis/kubeone"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/aws"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/azure"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/digitalocean"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/gce"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/hetzner"
	openstack2 "github.com/kubermatic/kubeone/pkg/credentials/provider/openstack"
	openstack "github.com/kubermatic/kubeone/pkg/credentials/provider/packet"
	"github.com/kubermatic/kubeone/pkg/credentials/provider/vsphere"
)

// The environment variable names with credential in them that machine-controller expects to see
const (
	DigitalOceanTokenKey = "DO_TOKEN"
	HetznerTokenKey      = "HZ_TOKEN"
	PacketAPIKey    = "PACKET_API_KEY"
	PacketProjectID = "PACKET_PROJECT_ID"
)

// Credentials is used to fetch credentials for given provider
type Credentials interface {
	Fetch() (map[string]string, error)
}

type Options struct {
	AWSProfilePath string
	AWSProfileName string
}

func FetcherForProvider(p kubeone.CloudProviderName, fetchOptions Options) Credentials {
	switch p {
	case kubeone.CloudProviderNameAWS:
		return aws.NewAWSCredentials(fetchOptions.AWSProfilePath, fetchOptions.AWSProfileName)
	case kubeone.CloudProviderNameAzure:
		return azure.NewAzureCredentials()
	case kubeone.CloudProviderNameOpenStack:
		return openstack2.NewOpenStackCredentials()
	case kubeone.CloudProviderNameHetzner:
		return hetzner.NewHetznerCredentials()
	case kubeone.CloudProviderNameDigitalOcean:
		return digitalocean.NewDigitalOceanCredentials()
	case kubeone.CloudProviderNameGCE:
		return gce.NewGCECredentials()
	case kubeone.CloudProviderNamePacket:
		return openstack.NewPacketCredentials()
	case kubeone.CloudProviderNameVSphere:
		return vsphere.NewVSphereCredentials()
	}
	return nil
}
