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

package api

// Prefix is a common namespace prefix for versions.
const Prefix = "envexec"

const (
	// KindVaultSecrets is a VaultSecrets kind.
	KindVaultSecrets = "VaultSecrets"
)

const (
	// VersionV1alpha1 is a v1alpha1 api version.
	VersionV1alpha1 = "v1alpha1"
)

// Spec is a versioned kind of spec.
type Spec struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}
