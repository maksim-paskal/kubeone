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

package helpers

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

// ProviderEnvironmentVariable is used to match environment variable used by KubeOne to environment variable used by
// machine-controller.
type ProviderEnvironmentVariable struct {
	Name                  string
	MachineControllerName string
}

func ParseCredentialVariables(envVars []ProviderEnvironmentVariable, validationFunc func(map[string]string) error) (map[string]string, error) {
	// Validate credentials using given validation function
	creds := make(map[string]string)
	for _, env := range envVars {
		creds[env.Name] = strings.TrimSpace(os.Getenv(env.Name))
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

func DefaultValidationFunc(creds map[string]string) error {
	for k, v := range creds {
		if len(v) == 0 {
			return errors.Errorf("key %v is required but isn't present", k)
		}
	}
	return nil
}
