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
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/hypnoglow/envexec/internal/api"
	"github.com/hypnoglow/envexec/internal/api/v1alpha1"
	"github.com/hypnoglow/envexec/internal/loaders"
	"github.com/hypnoglow/envexec/internal/loaders/vault"
)

// Loader returns a loader that can load values defined by spec in reader r.
func Loader(r io.Reader) (loaders.Loader, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var spec api.Spec
	if err = yaml.Unmarshal(b, &spec); err != nil {
		return nil, err
	}

	ver := strings.TrimPrefix(spec.ApiVersion, api.Prefix+"/")

	switch spec.Kind {
	case api.KindVaultSecrets:
		switch ver {
		case api.VersionV1alpha1:
			var s v1alpha1.VaultSecretsSpec
			if err = yaml.Unmarshal(b, &s); err != nil {
				return nil, err
			}
			return vault.NewSecretsLoader(s), nil
		}
	}

	return nil, fmt.Errorf("unknown kind and/or apiVersion: %s %s", spec.Kind, spec.ApiVersion)
}
