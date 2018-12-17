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

import "os"

// newEnvironmentAuth returns new token auth method.
func newTokenAuth() *tokenAuth {
	return &tokenAuth{}
}

// tokenAuth implements authMethod using static token.
// See: https://www.vaultproject.io/docs/auth/token.html
type tokenAuth struct{}

// GetToken returns token from the environment variable.
func (e *tokenAuth) GetToken() (string, error) {
	return os.Getenv("VAULT_TOKEN"), nil
}
