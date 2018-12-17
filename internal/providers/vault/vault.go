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
	"os"

	"github.com/hashicorp/vault/api"
)

// NewClient is a constructor for Vault client.
func NewClient() (*api.Client, error) {
	vc, err := defaultVaultClient()
	if err != nil {
		return nil, err
	}

	m := os.Getenv("VAULT_AUTH_METHOD")
	var method AuthMethod
	if err = method.UnmarshalText([]byte(m)); err != nil {
		return nil, err
	}

	var auth authMethod
	switch method {
	case AuthMethodToken:
		auth = newTokenAuth()
	case AuthMethodKubernetes:
		auth = newKubernetesAuth(vc)
	default:
		return nil, fmt.Errorf("unknown auth method: %v", method)
	}

	token, err := auth.GetToken()
	if err != nil {
		return nil, fmt.Errorf("get token: %v", err)
	}

	vc.SetToken(token)

	return vc, nil
}

func defaultVaultClient() (*api.Client, error) {
	// Read config from env (VAULT_ADDR, VAULT_SKIP_VERIFY, etc)
	config := api.DefaultConfig()
	err := config.ReadEnvironment()
	if err != nil {
		return nil, fmt.Errorf("read environment for vault config: %v", err)
	}

	return api.NewClient(config)
}
