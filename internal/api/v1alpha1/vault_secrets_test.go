/*
Copyright 2018 Igor Zibarev
Copyright 2018 The envexec Authors

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

package v1alpha1

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestVaultSecretsSpec_fromYaml(t *testing.T) {
	testCases := []struct {
		input       string
		spec        VaultSecretsSpec
		expectError bool
	}{
		{
			input: `
apiVersion: envexec/v1alpha1
kind: VaultSecrets
secrets:
  - path: secret/namespace/service/some
    key: api_key
    env: SOME_API_KEY
  - path: secret/namespace/service/db
    key: password
    env: DB_PASSWORD

`,
			spec: VaultSecretsSpec{
				Secrets: []VaultSecret{
					{
						Path: "secret/namespace/service/some",
						Key:  "api_key",
						Env:  "SOME_API_KEY",
					},
					{
						Path: "secret/namespace/service/db",
						Key:  "password",
						Env:  "DB_PASSWORD",
					},
				},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		var spec VaultSecretsSpec
		err := yaml.Unmarshal([]byte(tc.input), &spec)

		if tc.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, tc.spec, spec)
	}
}
