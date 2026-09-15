// Copyright 2026, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	stderrors "errors"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/spf13/cobra"
)

// schemaSource is the shared input of the gen-* commands: a package schema,
// read either from disk (the author's own provider repo, which is the usual
// case) or from a URL (how we check a package already in the registry -- its
// schema_file_url is in themes/default/data/registry/packages/<name>.yaml).
type schemaSource struct {
	file string
	url  string
}

func (s *schemaSource) addFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&s.file, "schemaFile", "s", "",
		"The path to the package's schema.json (or schema.yaml)")
	cmd.Flags().StringVar(&s.url, "schemaFileURL", "",
		"The URL from which the package's schema can be retrieved, instead of --schemaFile")
}

func (s *schemaSource) validate() error {
	switch {
	case s.file == "" && s.url == "":
		return stderrors.New("one of --schemaFile or --schemaFileURL is required")
	case s.file != "" && s.url != "":
		return stderrors.New("--schemaFile and --schemaFileURL are mutually exclusive")
	}
	return nil
}

func (s *schemaSource) read(client HTTPDoer) (*schema.PackageSpec, error) {
	if s.url != "" {
		return readRemoteSchemaFile(client, s.url, "")
	}
	return readLocalSchemaFile(s.file)
}

func readLocalSchemaFile(path string) (*schema.PackageSpec, error) {
	schemaBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "reading schema file %q", path)
	}

	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		schemaBytes, err = yaml.YAMLToJSON(schemaBytes)
		if err != nil {
			return nil, errors.Wrap(err, "reading YAML schema")
		}
	}

	spec := &schema.PackageSpec{}
	if err := json.Unmarshal(schemaBytes, spec); err != nil {
		return nil, errors.Wrap(err, "unmarshalling schema into a PackageSpec")
	}
	return spec, nil
}
