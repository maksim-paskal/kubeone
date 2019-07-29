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

package gce

import (
	"encoding/base64"

	"github.com/pkg/errors"

	"github.com/kubermatic/kubeone/pkg/credentials/helpers"
)

// GCECredentials is used to parse GCE credentials
type GCECredentials struct{}

func NewGCECredentials() *GCECredentials {
	return &GCECredentials{}
}

func (g GCECredentials) Fetch() (map[string]string, error) {
	gsa, err := helpers.ParseCredentialVariables([]helpers.ProviderEnvironmentVariable{
		{Name: "GOOGLE_CREDENTIALS", MachineControllerName: "GOOGLE_SERVICE_ACCOUNT"},
	}, helpers.DefaultValidationFunc)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// encode it before sending to secret to be consumed by
	// machine-controller, as machine-controller assumes it will be double encoded
	gsa["GOOGLE_SERVICE_ACCOUNT"] = base64.StdEncoding.EncodeToString([]byte(gsa["GOOGLE_SERVICE_ACCOUNT"]))
	return gsa, nil
}
