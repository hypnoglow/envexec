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

import "fmt"

// AuthMethod is vault authentication method.
// See: https://www.vaultproject.io/docs/auth/index.html
type AuthMethod string

const (
	// AuthMethodToken represents authentication method using static token.
	// See: https://www.vaultproject.io/docs/auth/token.html
	AuthMethodToken AuthMethod = "token"

	// AuthMethodKubernetes represents authentication method using kubernetes.
	// See: https://www.vaultproject.io/docs/auth/kubernetes.html
	AuthMethodKubernetes AuthMethod = "kubernetes"
)

// UnmarshalText implements encoding.TextUnmarshaler.
func (m *AuthMethod) UnmarshalText(text []byte) error {
	txt := string(text)
	switch txt {
	case "token":
		*m = AuthMethodToken
	case "kubernetes":
		*m = AuthMethodKubernetes
	case "": // by default use token method
		*m = AuthMethodToken
	default:
		return fmt.Errorf("unknown auth method: %q", txt)
	}

	return nil
}

// authMethod can get authentication token for Vault.
type authMethod interface {
	GetToken() (string, error)
}
