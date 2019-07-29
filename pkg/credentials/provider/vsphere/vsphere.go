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

package vsphere

import (
	"github.com/pkg/errors"

	"github.com/kubermatic/kubeone/pkg/credentials/helpers"
)

const (
	VSphereAddress  = "VSPHERE_ADDRESS"
	VSpherePassword = "VSPHERE_PASSWORD"
	VSphereUsername = "VSPHERE_USERNAME"
)

// VSphereCredentials is used to parse vSphere credentials
type VSphereCredentials struct{}

func NewVSphereCredentials() *VSphereCredentials {
	return &VSphereCredentials{}
}

func (v VSphereCredentials) Fetch() (map[string]string, error) {
	vscreds, err := helpers.ParseCredentialVariables([]helpers.ProviderEnvironmentVariable{
		{Name: "VSPHERE_SERVER", MachineControllerName: VSphereAddress},
		{Name: "VSPHERE_USER", MachineControllerName: VSphereUsername},
		{Name: VSpherePassword},
	}, helpers.DefaultValidationFunc)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// force scheme, as machine-controller requires it while terraform does not
	vscreds[VSphereAddress] = "https://" + vscreds[VSphereAddress]
	return vscreds, nil
}
