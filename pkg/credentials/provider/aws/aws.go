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

package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"
)

// AWSCredentials is used to parse AWS credentials
type AWSCredentials struct {
	ProfilePath string
	ProfileName string
}

func NewAWSCredentials(profilePath, profileName string) *AWSCredentials {
	return &AWSCredentials{
		ProfilePath: profilePath,
		ProfileName: profileName,
	}
}

func (a AWSCredentials) Fetch() (map[string]string, error) {
	creds := make(map[string]string)
	envCredsProvider := credentials.NewEnvCredentials()
	envCreds, err := envCredsProvider.Get()
	if err != nil {
		return nil, err
	}
	if envCreds.AccessKeyID != "" && envCreds.SecretAccessKey != "" {
		creds["AWS_ACCESS_KEY_ID"] = envCreds.AccessKeyID
		creds["AWS_SECRET_ACCESS_KEY"] = envCreds.SecretAccessKey
		return creds, nil
	}

	// If env fails resort to config file
	configCredsProvider := credentials.NewSharedCredentials("", "")
	configCreds, err := configCredsProvider.Get()
	if err != nil {
		return nil, err
	}
	if configCreds.AccessKeyID != "" && configCreds.SecretAccessKey != "" {
		creds["AWS_ACCESS_KEY_ID"] = configCreds.AccessKeyID
		creds["AWS_SECRET_ACCESS_KEY"] = configCreds.SecretAccessKey
		return creds, nil
	}

	return nil, errors.New("error parsing aws credentials")
}
