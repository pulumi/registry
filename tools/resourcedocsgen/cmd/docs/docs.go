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
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/docs"
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v3/codegen/convert"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pkghost "github.com/pulumi/pulumi/pkg/v3/host"
	pkgWorkspace "github.com/pulumi/pulumi/pkg/v3/workspace"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
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

	loader, closeLoader, err := newSchemaLoader(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating schema loader")
	}
	defer func() { contract.IgnoreError(closeLoader()) }()

	pulPkg, err := bindSpec(mainSpec, loader)
	if err != nil {
		if pkg, ok := bindGoogleNativeToleratingCycles(mainSpec, loader, err); ok {
			pulPkg = pkg
		} else {
			if dErr, ok := err.(hcl.Diagnostics); ok {
				writer := hcl.NewDiagnosticTextWriter(os.Stderr, nil, 80, true)
				wErr := writer.WriteDiagnostics(dErr)
				if wErr != nil {
					err = wErr
				}
			}
			return nil, nil, fmt.Errorf("importing package spec: %w", err)
		}
	}

	return pulPkg, docs.NewContext(tool, pulPkg), nil
}

// newSchemaLoader builds the same plugin loader that pulumi's schema binder
// constructs when handed a nil loader. As of pulumi v3.247.0 that nil-loader
// fallback creates its plugin host with nil diagnostic sinks, which panics
// with a nil dereference as soon as binding downloads a referenced provider
// plugin and logs progress — so we construct the loader ourselves with
// (discarding, but non-nil) sinks. The returned close function releases the
// plugin context and host once binding is done.
func newSchemaLoader(ctx context.Context) (pschema.ReferenceLoader, func() error, error) {
	sink := diag.DefaultSink(io.Discard, io.Discard, diag.FormatOptions{Color: colors.Never})
	host, err := pkghost.New(ctx, sink, sink, nil, pkgWorkspace.EnsureLanguageInstalled)
	if err != nil {
		return nil, nil, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, nil, stderrors.Join(err, host.Close())
	}
	pctx, err := plugin.NewContext(ctx, sink, sink, host, nil, cwd, nil, false, nil,
		pschema.NewLoaderServerFromContext, convert.NewMapperServerFromContext)
	if err != nil {
		return nil, nil, stderrors.Join(err, host.Close())
	}
	return pschema.NewPluginLoader(pctx), func() error {
		return stderrors.Join(pctx.Close(), host.Close())
	}, nil
}

// bindSpec mirrors pschema.ImportSpec's contract (returns the binding
// diagnostics as the error when they contain errors) while binding with an
// explicit loader, which ImportSpec does not accept. Unlike ImportSpec,
// pschema.BindSpec also validates the spec against the package metaschema.
func bindSpec(spec pschema.PackageSpec, loader pschema.Loader) (*pschema.Package, error) {
	pkg, diags, err := pschema.BindSpec(spec, loader, pschema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	if err != nil {
		return nil, err
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return pkg, nil
}

// requiredCycleErrSummary is the fragment pulumi's schema binder puts in the
// diagnostic summary for an unsatisfiable required-property cycle.
const requiredCycleErrSummary = "unsatisfiable required-property cycle"

// bindGoogleNativeToleratingCycles is a targeted workaround for the google-native
// provider. As of pulumi/pulumi v3.237+ the schema binder rejects schemas with
// "unsatisfiable required-property cycles" — a required property whose type
// transitively requires itself, describing a value of infinite size (e.g.
// google-native's bigquery/v2:StandardSqlDataTypeResponse referencing itself).
// This is correct for SDK generation but fires even in ImportSpec, which docs
// generation uses to load already-published schemas. google-native is archived
// (last released v0.32.0 in 2023), so its schema will never be fixed upstream
// (see pulumi/pulumi#22890 and pulumi/pulumi-google-native#1246).
//
// Such cycles only prevent *constructing* a value; they are harmless for
// documenting types. So for google-native specifically — and only when every
// binding error is one of these cycles — we re-bind with BindSpec, which (unlike
// ImportSpec) returns a usable package alongside the non-fatal diagnostics, and
// continue generating docs. Any other error, or any other package, is left to
// fail as before so we never silently mask real schema problems.
func bindGoogleNativeToleratingCycles(
	spec pschema.PackageSpec, loader pschema.Loader, importErr error,
) (*pschema.Package, bool) {
	if spec.Name != "google-native" {
		return nil, false
	}
	if !onlyRequiredCycleErrors(importErr) {
		return nil, false
	}

	pkg, diags, err := pschema.BindSpec(spec, loader, pschema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	// BindSpec runs metaschema validation that ImportSpec skips; bail out if it
	// surfaces anything beyond the expected required-property cycles.
	if err != nil || pkg == nil || !onlyRequiredCycleErrors(diags) {
		return nil, false
	}

	slog.Warn("Tolerating unsatisfiable required-property cycles in schema; generating docs anyway",
		"package", spec.Name)
	return pkg, true
}

// onlyRequiredCycleErrors reports whether every error-severity diagnostic in the
// given value (an error returned by ImportSpec or an hcl.Diagnostics) is an
// unsatisfiable required-property cycle.
func onlyRequiredCycleErrors(v any) bool {
	diags, ok := v.(hcl.Diagnostics)
	if !ok {
		return false
	}
	sawError := false
	for _, d := range diags {
		if d.Severity != hcl.DiagError {
			continue
		}
		sawError = true
		if !strings.Contains(d.Summary, requiredCycleErrSummary) {
			return false
		}
	}
	return sawError
}

func ResourceDocsCmd() *cobra.Command {
	var schemaFile string
	var version string
	var docsOutDir string
	var packageTreeJSONOutDir string
	var cliDocsOutDir string

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

			if err := generateCLIDocs(cliDocsOutDir, genctx); err != nil {
				return errors.Wrap(err, "generating CLI docs")
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
	cmd.Flags().StringVar(&cliDocsOutDir, "cliDocsOutDir", "",
		"The directory path to write terminal-friendly CLI docs (optional)")

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

func generateCLIDocs(outDir string, genctx *docs.Context) error {
	if outDir == "" {
		return nil // CLI docs generation is optional
	}

	bundle, err := genctx.GenerateCLIPackage()
	if err != nil {
		return errors.Wrap(err, "generating CLI package docs")
	}

	b, err := json.Marshal(bundle)
	if err != nil {
		return errors.Wrap(err, "marshalling CLI docs bundle")
	}

	if err := pkg.EmitFile(outDir, "llm-docs.json", b); err != nil {
		return errors.Wrap(err, "emitting LLM docs bundle")
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
