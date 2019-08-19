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
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	"github.com/kubermatic/kubeone/pkg/apis/kubeone"
)

// The environment variable names with credential in them
const (
	// Variables that KubeOne (and Terraform) expect to see
	AWSAccessKeyID          = "AWS_ACCESS_KEY_ID"
	AWSSecretAccessKey      = "AWS_SECRET_ACCESS_KEY"
	AzureClientID           = "ARM_CLIENT_ID"
	AzureClientSecret       = "ARM_CLIENT_SECRET"
	AzureTenantID           = "ARM_TENANT_ID"
	AzureSubscribtionID     = "ARM_SUBSCRIPTION_ID"
	DigitalOceanTokenKey    = "DIGITALOCEAN_TOKEN"
	GoogleServiceAccountKey = "GOOGLE_CREDENTIALS"
	HetznerTokenKey         = "HCLOUD_TOKEN"
	OpenStackAuthURL        = "OS_AUTH_URL"
	OpenStackDomainName     = "OS_DOMAIN_NAME"
	OpenStackPassword       = "OS_PASSWORD"
	OpenStackRegionName     = "OS_REGION_NAME"
	OpenStackTenantID       = "OS_TENANT_ID"
	OpenStackTenantName     = "OS_TENANT_NAME"
	OpenStackUserName       = "OS_USERNAME"
	PacketAPIKey            = "PACKET_AUTH_TOKEN"
	PacketProjectID         = "PACKET_PROJECT_ID"
	VSphereAddress          = "VSPHERE_SERVER"
	VSpherePassword         = "VSPHERE_PASSWORD"
	VSphereUsername         = "VSPHERE_USER"

	// Variables that machine-controller expects
	AzureClientIDMC           = "AZURE_CLIENT_ID"
	AzureClientSecretMC       = "AZURE_CLIENT_SECRET"
	AzureTenantIDMC           = "AZURE_TENANT_ID"
	AzureSubscribtionIDMC     = "AZURE_SUBSCRIPTION_ID"
	DigitalOceanTokenKeyMC    = "DO_TOKEN"
	GoogleServiceAccountKeyMC = "GOOGLE_SERVICE_ACCOUNT"
	HetznerTokenKeyMC         = "HZ_TOKEN"
	OpenStackUserNameMC       = "OS_USER_NAME"
	PacketAPIKeyMC            = "PACKET_API_KEY"
	VSphereAddressMC          = "VSPHERE_ADDRESS"
	VSphereUsernameMC         = "VSPHERE_USERNAME"
)

// ProviderEnvironmentVariable is used to match environment variable used by KubeOne to environment variable used by
// machine-controller.
type ProviderEnvironmentVariable struct {
	Name                  string
	MachineControllerName string
}

type variableSource map[string]string

type variableFetcher func(variableSource, string) string

func defaultVariableFetcher(_ variableSource, name string) string {
	return os.Getenv(name)
}

// ProviderCredentials match the cloudprovider and parses its credentials from environment
func ProviderCredentials(p kubeone.CloudProviderName, credentialsFilePath string) (map[string]string, error) {
	fetcher := defaultVariableFetcher
	var source variableSource
	if credentialsFilePath != "" {
		b, err := ioutil.ReadFile(credentialsFilePath)
		if err != nil {
			return nil, errors.Wrap(err, "unable to load credentials file")
		}
		err = yaml.Unmarshal(b, &source)
		if err != nil {
			return nil, errors.Wrap(err, "unable to unmarshal credentials file")
		}
	}
	return fetchCredentials(p, fetcher, source)
}

// fetchCredentials implements fetching credentials for each supported provider
func fetchCredentials(p kubeone.CloudProviderName, fetcher variableFetcher, source variableSource) (map[string]string, error) {
	switch p {
	case kubeone.CloudProviderNameAWS:
		return parseAWSCredentials()
	case kubeone.CloudProviderNameAzure:
		return parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: AzureClientID, MachineControllerName: AzureClientIDMC},
			{Name: AzureClientSecret, MachineControllerName: AzureClientSecretMC},
			{Name: AzureTenantID, MachineControllerName: AzureTenantIDMC},
			{Name: AzureSubscribtionID, MachineControllerName: AzureSubscribtionIDMC},
		}, defaultValidationFunc, fetcher, source)
	case kubeone.CloudProviderNameOpenStack:
		return parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: OpenStackAuthURL},
			{Name: OpenStackUserName, MachineControllerName: OpenStackUserNameMC},
			{Name: OpenStackPassword},
			{Name: OpenStackDomainName},
			{Name: OpenStackRegionName},
			{Name: OpenStackTenantID},
			{Name: OpenStackTenantName},
		}, openstackValidationFunc, fetcher, source)
	case kubeone.CloudProviderNameHetzner:
		return parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: HetznerTokenKey, MachineControllerName: HetznerTokenKeyMC},
		}, defaultValidationFunc, fetcher, source)
	case kubeone.CloudProviderNameDigitalOcean:
		return parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: DigitalOceanTokenKey, MachineControllerName: DigitalOceanTokenKeyMC},
		}, defaultValidationFunc, fetcher, source)
	case kubeone.CloudProviderNameGCE:
		gsa, err := parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: GoogleServiceAccountKey, MachineControllerName: GoogleServiceAccountKeyMC},
		}, defaultValidationFunc, fetcher, source)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// encode it before sending to secret to be consumed by
		// machine-controller, as machine-controller assumes it will be double encoded
		gsa[GoogleServiceAccountKeyMC] = base64.StdEncoding.EncodeToString([]byte(gsa[GoogleServiceAccountKeyMC]))
		return gsa, nil
	case kubeone.CloudProviderNamePacket:
		return parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: PacketAPIKey, MachineControllerName: PacketAPIKeyMC},
			{Name: PacketProjectID},
		}, defaultValidationFunc, fetcher, source)
	case kubeone.CloudProviderNameVSphere:
		vscreds, err := parseCredentialVariables([]ProviderEnvironmentVariable{
			{Name: VSphereAddress, MachineControllerName: VSphereAddressMC},
			{Name: VSphereUsername, MachineControllerName: VSphereUsernameMC},
			{Name: VSpherePassword},
		}, defaultValidationFunc, fetcher, source)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// force scheme, as machine-controller requires it while terraform does not
		vscreds[VSphereAddressMC] = "https://" + vscreds[VSphereAddressMC]
		return vscreds, nil
	}

	return nil, errors.New("no provider matched")
}

func parseAWSCredentials() (map[string]string, error) {
	creds := make(map[string]string)
	envCredsProvider := credentials.NewEnvCredentials()
	envCreds, err := envCredsProvider.Get()
	if err != nil {
		return nil, err
	}
	if envCreds.AccessKeyID != "" && envCreds.SecretAccessKey != "" {
		creds[AWSAccessKeyID] = envCreds.AccessKeyID
		creds[AWSSecretAccessKey] = envCreds.SecretAccessKey
		return creds, nil
	}

	// If env fails resort to config file
	configCredsProvider := credentials.NewSharedCredentials("", "")
	configCreds, err := configCredsProvider.Get()
	if err != nil {
		return nil, err
	}
	if configCreds.AccessKeyID != "" && configCreds.SecretAccessKey != "" {
		creds[AWSAccessKeyID] = configCreds.AccessKeyID
		creds[AWSSecretAccessKey] = configCreds.SecretAccessKey
		return creds, nil
	}

	return nil, errors.New("error parsing aws credentials")
}

func parseCredentialVariables(envVars []ProviderEnvironmentVariable, validationFunc func(map[string]string) error, fetcher variableFetcher, source variableSource) (map[string]string, error) {
	// Validate credentials using given validation function
	creds := make(map[string]string)
	for _, env := range envVars {
		creds[env.Name] = strings.TrimSpace(fetcher(source, env.Name))
	}
	if err := validationFunc(creds); err != nil {
		return nil, errors.Wrap(err, "unable to validate credentials")
	}

	// Prepare credentials to be used by machine-controller
	mcCreds := make(map[string]string)
	for _, env := range envVars {
		name := env.MachineControllerName
		if len(name) == 0 {
			name = env.Name
		}
		mcCreds[name] = creds[env.Name]
	}

	return mcCreds, nil
}

func defaultValidationFunc(creds map[string]string) error {
	for k, v := range creds {
		if len(v) == 0 {
			return errors.Errorf("key %v is required but isn't present", k)
		}
	}
	return nil
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
