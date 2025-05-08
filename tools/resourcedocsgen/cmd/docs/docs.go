// Copyright 2024, Pulumi Corporation.
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

package docs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/docs"
	"github.com/spf13/cobra"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const (
	tool                        = "Pulumi Docs Generator"
	registryRepo                = "https://github.com/pulumi/registry"
	defaultSchemaFilePathFormat = "/provider/cmd/pulumi-resource-%s/schema.json"
)

func getPulumiPackageFromSchema(
	docsOutDir string, mainSpec pschema.PackageSpec,
) (*pschema.Package, *docs.Context, error) {
	// Delete existing docs before generating new ones.
	if err := os.RemoveAll(docsOutDir); err != nil {
		return nil, nil, errors.Wrapf(err, "deleting provider directory %v", docsOutDir)
	}

	pulPkg, err := pschema.ImportSpec(mainSpec, nil, pschema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	if err != nil {
		if dErr, ok := err.(hcl.Diagnostics); ok {
			writer := hcl.NewDiagnosticTextWriter(os.Stderr, nil, 80, true)
			wErr := writer.WriteDiagnostics(dErr)
			if wErr != nil {
				err = wErr
			}
		}
		return nil, nil, fmt.Errorf("importing package spec: %w", err)
	}

	// THIS IS A TEMPORARY HACK!!
	// temporarily hacking this tool to accommodate multiple versions of the azure-native package,
	// since both v2 and v3 are techically the same package (just different versons) so the
	// schema.json file will have `azure-native` as the package name. Without this, the files genned
	// will be overwritten by the second version of the package, since they will end up being written
	// to the same azure-native directory.
	if strings.Contains(docsOutDir, "azure-native-v2") {
		pulPkg.Name = "azure-native-v2"
	}

	return pulPkg, docs.NewContext(tool, pulPkg), nil
}

func ResourceDocsCmd() *cobra.Command {
	var schemaFile string
	var version string
	var docsOutDir string
	var packageTreeJSONOutDir string

	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Generate resource docs from a Pulumi schema file",
		RunE: func(cmd *cobra.Command, args []string) error {
			schema, err := os.ReadFile(schemaFile)
			if err != nil {
				return errors.Wrap(err, "reading schema file from path")
			}

			// The source schema can be in YAML format. If that's the case
			// convert it to JSON first.
			if strings.HasSuffix(schemaFile, ".yaml") {
				schema, err = yaml.YAMLToJSON(schema)
				if err != nil {
					return errors.Wrap(err, "reading YAML schema")
				}
			}

			var mainSpec pschema.PackageSpec
			if err := json.Unmarshal(schema, &mainSpec); err != nil {
				return errors.Wrap(err, "unmarshalling schema into a PackageSpec")
			}
			mainSpec.Version = version

			pulPkg, genctx, err := getPulumiPackageFromSchema(docsOutDir, mainSpec)
			if err != nil {
				return errors.Wrap(err, "generating package from schema file")
			}

			if err := generateDocsFromSchema(docsOutDir, genctx); err != nil {
				return errors.Wrap(err, "generating docs from schema")
			}

			if err := generatePackageTree(packageTreeJSONOutDir, pulPkg.Name, genctx); err != nil {
				return errors.Wrap(err, "generating package tree")
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&schemaFile, "schemaFile", "s", "", "Path to the schema.json file")
	cmd.Flags().StringVar(&version, "version", "", "The version of the package")
	cmd.Flags().StringVar(&docsOutDir, "docsOutDir", "", "The directory path to where the docs will be written to")
	cmd.Flags().StringVar(&packageTreeJSONOutDir, "packageTreeJSONOutDir", "",
		"The directory path to write the package tree JSON file to")

	contract.AssertNoErrorf(cmd.MarkFlagRequired("docsOutDir"), "could not find docsOutDir")
	contract.AssertNoErrorf(cmd.MarkFlagRequired("packageTreeJSONOutDir"), "could not find packageTreeJSONOutDir")
	contract.AssertNoErrorf(cmd.MarkFlagRequired("schemaFile"), "could not find schemaFile")
	contract.AssertNoErrorf(cmd.MarkFlagRequired("version"), "could not find version")

	cmd.AddCommand(resourceDocsFromRegistryCmd())

	return cmd
}

func generateDocsFromSchema(outDir string, pulPkg *docs.Context) error {
	files, err := pulPkg.GeneratePackage()
	if err != nil {
		return errors.Wrap(err, "generating Pulumi package")
	}

	for f, contents := range files {
		if err := pkg.EmitFile(outDir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
		}
	}
	return nil
}

func generatePackageTree(outDir string, pkgName string, genctx *docs.Context) error {
	tree, err := genctx.GeneratePackageTree()
	if err != nil {
		return errors.Wrap(err, "generating the package tree")
	}

	b, err := json.Marshal(tree)
	if err != nil {
		return errors.Wrap(err, "marshalling the package tree")
	}

	filename := pkgName + ".json"
	if err := pkg.EmitFile(outDir, filename, b); err != nil {
		return errors.Wrap(err, "writing the package tree")
	}

	return nil
}
