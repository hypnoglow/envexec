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

package envexec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/hypnoglow/envexec/internal/spec"
)

// New returns a new envexec application.
func New(conf *Config) *Envexec {
	return &Envexec{
		conf: conf,
	}
}

// Envexec defines the application.
type Envexec struct {
	conf *Config
}

// Run runs envexec.
func (e Envexec) Run() error {
	envs := make(map[string]string)

	for _, fpath := range e.conf.specs {
		if err := e.loadEnvsFromFile(fpath, envs); err != nil {
			return err
		}
	}

	if err := e.execBinary(e.conf.bin, e.conf.args, envs); err != nil {
		return err
	}

	return nil
}

func (e Envexec) loadEnvsFromFile(fpath string, envs map[string]string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return fmt.Errorf("open spec file: %v", err)
	}
	defer f.Close()

	l, err := spec.Loader(f)
	if err != nil {
		return fmt.Errorf("create loader for spec file %v: %s", fpath, err)
	}

	ee, err := l.Load()
	if err != nil {
		return fmt.Errorf("load envs by spec file %s: %v", fpath, err)
	}

	for k, v := range ee {
		if _, ok := envs[k]; ok {
			log.Printf("WARNING: duplicate environment variable name %s - will keep last value", k)
		}
		envs[k] = v
	}

	return nil
}

func (e Envexec) execBinary(bin string, args []string, envs map[string]string) error {
	binary, err := exec.LookPath(bin)
	if err != nil {
		return fmt.Errorf("look path: %v", err)
	}

	args = append([]string{binary}, args...)

	environ := os.Environ()
	for k, v := range envs {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}

	err = syscall.Exec(binary, args, environ)
	if err != nil {
		return fmt.Errorf("exec: %v", err)
	}

	return nil
}
