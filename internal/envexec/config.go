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
	"os"
	"strings"

	"github.com/integrii/flaggy"
)

// Config defines configuration for envexec.
type Config struct {
	bin  string
	args []string

	specs []string
}

// LoadConfig configures envexec using environment variables and CLI flags.
func LoadConfig() (*Config, error) {
	c := &Config{}

	if err := c.parseEnvironment(); err != nil {
		return nil, err
	}

	p := c.newCLI()
	if err := c.parseCLI(p); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) newCLI() *flaggy.Parser {
	p := flaggy.NewParser("envexec")
	p.Description = `a tool that executes any command or application provisioning its
environment with variables from providers (like secrets from Vault).`
	p.AdditionalHelpPrepend = `
https://github.com/hypnoglow/envexec`
	p.AdditionalHelpAppend = `
  Environment Variables:
    ENVEXEC_SPEC_FILES - A list of specification files in YAML format (can specify multiple, delimited by comma).

  Examples:
    envexec --spec-file vaultsecrets.yaml /usr/bin/env
    envexec /usr/bin/myapp --spec-file spec.yaml -- --arg-for myapp`

	return p
}

func (c *Config) parseEnvironment() error {
	c.specs = strings.Split(os.Getenv("ENVEXEC_SPEC_FILES"), ",")
	return nil
}

func (c *Config) parseCLI(p *flaggy.Parser) error {
	var specFiles []string
	var binary string

	p.StringSlice(&specFiles, "", "spec-file", "A list of specification files in YAML format (can specify multiple). Takes precedence over ENVEXEC_SPEC_FILES.")
	p.AddPositionalValue(&binary, "cmd", 1, false, "A command or an application to exec.")

	err := p.Parse()
	if err != nil {
		return fmt.Errorf("parse command: %v", err)
	}

	args := p.TrailingArguments
	if binary == "" &&
		len(args) > 0 &&
		!strings.HasPrefix(args[0], "-") {
		binary = args[0]
		args = args[1:]
	}

	if binary == "" {
		return fmt.Errorf("binary to run not specified")
	}

	c.bin = binary
	c.args = args

	if len(specFiles) > 0 {
		// NOTE: this overwrites ENVEXEC_SPEC_FILES env var.
		c.specs = specFiles
	}

	return nil
}
