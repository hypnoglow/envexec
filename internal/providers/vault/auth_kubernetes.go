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
	"io/ioutil"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

const defaultTokenFile = "/var/run/secrets/kubernetes.io/serviceaccount/token"

// newKubernetesAuth returns new Kubernetes auth method.
func newKubernetesAuth(vc *api.Client) *kubernetesAuth {
	return &kubernetesAuth{
		vault: vc,
	}
}

// kubernetesAuth implements authMethod using Kubernetes.
// See: https://www.vaultproject.io/docs/auth/kubernetes.html
type kubernetesAuth struct {
	vault *api.Client
}

// GetToken returns token using Kubernetes auth method.
func (k *kubernetesAuth) GetToken() (string, error) {
	role := os.Getenv("VAULT_AUTH_KUBERNETES_ROLE")
	if role == "" {
		return "", fmt.Errorf("vault Kubernetes auth role is empty, use VAULT_AUTH_KUBERNETES_ROLE to set")
	}

	fpath := os.Getenv("VAULT_AUTH_KUBERNETES_TOKEN_FILE")
	if fpath == "" {
		fpath = defaultTokenFile
	}

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", fmt.Errorf("read token file: %v", err)
	}
	jwt := strings.TrimSpace(string(b))

	token, err := k.login(jwt, role)
	if err != nil {
		return "", fmt.Errorf("login to vault: %v", err)
	}

	return token, nil
}

func (k *kubernetesAuth) login(jwt, role string) (token string, err error) {
	data := map[string]interface{}{
		"role": role,
		"jwt":  jwt,
	}

	response, err := k.vault.Logical().Write("auth/kubernetes/login", data)
	if err != nil {
		return "", fmt.Errorf("vault login: %v", err)
	}
	if response.Auth == nil {
		return "", fmt.Errorf("no authentication information attached to login response")
	}

	return response.Auth.ClientToken, nil
}
