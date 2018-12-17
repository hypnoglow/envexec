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
	"os"
	"testing"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/vault"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hypnoglow/envexec/internal/api/v1alpha1"
)

func TestSecretsLoader_Load(t *testing.T) {

	// init vault & secrets

	vlt, teardown := setup(t)
	defer teardown()

	var err error
	_, err = vlt.Logical().Write("secret/foo", map[string]interface{}{
		"passwd": "k8jghl37s;fn2hs",
	})
	require.NoError(t, err)
	_, err = vlt.Logical().Write("secret/bar", map[string]interface{}{
		"a_token":       "super-secret-token",
		"secret_number": 123,
	})
	require.NoError(t, err)

	// prepare spec

	s := v1alpha1.VaultSecretsSpec{
		Secrets: []v1alpha1.VaultSecret{
			{
				Path: "secret/foo",
				Key:  "passwd",
				Env:  "PASSWD",
			},
			{
				Path: "secret/bar",
				Key:  "a_token",
				Env:  "TOKEN",
			},
			{
				Path: "secret/bar",
				Key:  "secret_number",
				Env:  "SECRET_NUMBER",
			},
		},
	}

	// get envs & compare

	l := NewSecretsLoader(s)
	envs, err := l.Load()
	assert.NoError(t, err)

	expected := map[string]string{
		"PASSWD":        "k8jghl37s;fn2hs",
		"TOKEN":         "super-secret-token",
		"SECRET_NUMBER": "123",
	}
	assert.Equal(t, expected, envs)
}

func setup(t *testing.T) (vlt *api.Client, teardown func()) {
	// Disable noisy logger of vault core.
	dn, err := os.Open(os.DevNull)
	assert.NoError(t, err)
	hclog.DefaultOutput = dn

	core, _, token := vault.TestCoreUnsealed(t)

	ln, addr := http.TestServer(t, core)
	http.TestServerAuth(t, addr, token)

	vlt, err = api.NewClient(&api.Config{
		Address: addr,
	})
	require.NoError(t, err)

	vlt.SetToken(token)

	os.Setenv("VAULT_ADDR", addr)
	os.Setenv("VAULT_AUTH_METHOD", "token")
	os.Setenv("VAULT_TOKEN", token)

	return vlt, func() {
		ln.Close()
	}
}
