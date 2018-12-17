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

package spec

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoader_vaultSecrets(t *testing.T) {
	r := bytes.NewBufferString(`
apiVersion: envexec/v1alpha1
kind: VaultSecrets
secrets:
  - path: secret/namespace/service/some
    key: api_key
    env: SOME_API_KEY
  - path: secret/namespace/service/db
    key: password
    env: DB_PASSWORD
`)
	_, err := Loader(r)
	assert.NoError(t, err)
}
