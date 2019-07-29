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

package openstack

import (
	"github.com/pkg/errors"

	"github.com/kubermatic/kubeone/pkg/credentials/helpers"
)

const (
	OpenStackAuthURL    = "OS_AUTH_URL"
	OpenStackDomainName = "OS_DOMAIN_NAME"
	OpenStackPassword   = "OS_PASSWORD"
	OpenStackRegionName = "OS_REGION_NAME"
	OpenStackTenantID   = "OS_TENANT_ID"
	OpenStackTenantName = "OS_TENANT_NAME"
	OpenStackUserName   = "OS_USER_NAME"
)

// OpenStackCredentials is used to parse OpenStack credentials
type OpenStackCredentials struct{}

func NewOpenStackCredentials() *OpenStackCredentials {
	return &OpenStackCredentials{}
}

func (o OpenStackCredentials) Fetch() (map[string]string, error) {
	return helpers.ParseCredentialVariables([]helpers.ProviderEnvironmentVariable{
		{Name: "OS_AUTH_URL"},
		{Name: "OS_USERNAME", MachineControllerName: "OS_USER_NAME"},
		{Name: "OS_PASSWORD"},
		{Name: "OS_DOMAIN_NAME"},
		{Name: "OS_REGION_NAME"},
		{Name: "OS_TENANT_ID"},
		{Name: "OS_TENANT_NAME"},
	}, openstackValidationFunc)
}

func openstackValidationFunc(creds map[string]string) error {
	for k, v := range creds {
		if k == OpenStackTenantID || k == OpenStackTenantName {
			continue
		}
		if len(v) == 0 {
			return errors.Errorf("key %v is required but isn't present", k)
		}
	}

	if v, ok := creds[OpenStackTenantID]; !ok || len(v) == 0 {
		if v, ok := creds[OpenStackTenantName]; !ok || len(v) == 0 {
			return errors.Errorf("key %v or %v is required but isn't present", OpenStackTenantID, OpenStackTenantName)
		}
	}

	return nil
}
