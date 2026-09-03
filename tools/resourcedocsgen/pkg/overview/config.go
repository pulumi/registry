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

// Package overview derives the mechanically-generatable sections of a package's
// Overview page (docs/_index.md) from its schema. See docs/overview-page.md for
// the standard these snippets are written against.
package overview

import (
	"sort"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/docs"
)

// ConfigVar is one provider configuration parameter, flattened out of the
// schema's config block into the four fields docs/overview-page.md asks an
// author to document, plus the environment-variable fallbacks the schema
// happens to know about.
type ConfigVar struct {
	// Name is the bare option name, e.g. "apiToken". The standard is explicit
	// that it carries no "<package>:" prefix.
	Name string
	// Type is the JSON-schema type, or "object" for a $ref'd config type.
	Type string
	// Description has been through docs.SanitizeDescription.
	Description string
	// Required reports membership of the schema's config "defaults" list. Note
	// that many bridged providers omit that list entirely, in which case every
	// parameter here reads as optional.
	Required bool
	// Secret reports the schema's secret flag, which is likewise unreliable:
	// pulumi-vault does not mark its `token` secret.
	Secret bool
	// EnvVars comes from defaultInfo.environment. It is sparse in practice --
	// most environment-variable fallbacks are read by the vendor SDK a layer
	// beneath the schema and appear nowhere in it.
	EnvVars []string
	// Deprecated is the schema's deprecationMessage, if any.
	Deprecated string
}

// SchemaConfig is a package's provider configuration, normalized for rendering.
type SchemaConfig struct {
	Package string
	Vars    []ConfigVar
}

// ReadSchemaConfig flattens spec's provider config block. Variables are sorted
// by name so that output is stable across runs; Go map iteration is not.
func ReadSchemaConfig(spec *schema.PackageSpec) SchemaConfig {
	required := make(map[string]bool, len(spec.Config.Required))
	for _, name := range spec.Config.Required {
		required[name] = true
	}

	vars := make([]ConfigVar, 0, len(spec.Config.Variables))
	for name, prop := range spec.Config.Variables {
		vars = append(vars, ConfigVar{
			Name:        name,
			Type:        configVarType(spec, prop),
			Description: docs.SanitizeDescription(prop.Description),
			Required:    required[name],
			Secret:      prop.Secret,
			EnvVars:     envVars(prop),
			Deprecated:  docs.SanitizeDescription(prop.DeprecationMessage),
		})
	}

	sort.Slice(vars, func(i, j int) bool { return vars[i].Name < vars[j].Name })

	return SchemaConfig{Package: spec.Name, Vars: vars}
}

// configVarType names the parameter's type for the reader. A config variable
// may be a $ref to a complex type declared elsewhere in the same schema --
// pulumi-vault's authLogin parameters are the common example -- in which case
// the referenced type's own type is the useful answer, and "object" is the
// fallback when the reference points outside this document.
func configVarType(spec *schema.PackageSpec, prop schema.PropertySpec) string {
	if prop.Ref == "" {
		if prop.Type == "" {
			return "object"
		}
		return prop.Type
	}

	token, ok := strings.CutPrefix(prop.Ref, "#/types/")
	if !ok {
		return "object"
	}
	if typ, ok := spec.Types[token]; ok && typ.Type != "" {
		return typ.Type
	}
	return "object"
}

func envVars(prop schema.PropertySpec) []string {
	if prop.DefaultInfo == nil {
		return nil
	}
	return prop.DefaultInfo.Environment
}
