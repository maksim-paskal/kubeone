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

package hetzner

import "github.com/kubermatic/kubeone/pkg/credentials/helpers"

// HetznerCredentials is used to parse Hetzner credentials
type HetznerCredentials struct{}

func NewHetznerCredentials() *HetznerCredentials {
	return &HetznerCredentials{}
}

func (h HetznerCredentials) Fetch() (map[string]string, error) {
	return helpers.ParseCredentialVariables([]helpers.ProviderEnvironmentVariable{
		{Name: "HCLOUD_TOKEN", MachineControllerName: "HZ_TOKEN"},
	}, helpers.DefaultValidationFunc)
}
