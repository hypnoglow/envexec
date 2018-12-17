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

package vault

import (
	"fmt"

	"github.com/hypnoglow/envexec/internal/api/v1alpha1"
	"github.com/hypnoglow/envexec/internal/providers/vault"
)

// NewSecretsLoader returns a new SecretsLoader.
func NewSecretsLoader(spec v1alpha1.VaultSecretsSpec) *SecretsLoader {
	return &SecretsLoader{spec: spec}
}

// SecretsLoader is a loader that loads secrets from Vault.
type SecretsLoader struct {
	spec v1alpha1.VaultSecretsSpec
}

// Load loads secrets from Vault and returns them as key-values.
func (l *SecretsLoader) Load() (map[string]string, error) {
	vlt, err := vault.NewClient()
	if err != nil {
		return nil, fmt.Errorf("create vault client: %v", err)
	}

	envs := make(map[string]string)

	for path, ss := range mapByPath(l.spec.Secrets) {
		secret, err := vlt.Logical().Read(path)
		if err != nil {
			return nil, fmt.Errorf("read secret from vault: %v", err)
		}
		if secret == nil {
			return nil, fmt.Errorf("no secret found by path %v", path)
		}

		for _, s := range ss {
			v, ok := secret.Data[s.Key]
			if !ok {
				return nil, fmt.Errorf("key %s not found in vault secret", s.Key)
			}

			envs[s.Env] = fmt.Sprintf("%v", v)
		}
	}

	return envs, nil
}

func mapByPath(secrets []v1alpha1.VaultSecret) map[string][]v1alpha1.VaultSecret {
	m := make(map[string][]v1alpha1.VaultSecret)
	for _, secret := range secrets {
		m[secret.Path] = append(m[secret.Path], secret)
	}
	return m
}
