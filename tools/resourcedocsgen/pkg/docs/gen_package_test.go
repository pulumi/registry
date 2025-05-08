// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"testing"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// This file tests GeneratePackage by generating package documentation for a number of schema. Test data is located in
// the testdata/TestGeneratePackage directory. See README.md for information on running, adding, and updating tests for
// this suite.

// testCase captures details about a test case for GeneratePackage
type testCase struct {
	// The subdirectory of testdata/TestGeneratePackage that contains the test case's inputs and expected outputs.
	directory string
	// A human-readable description for the test case.
	description string
}

// testCases represents a suite of test cases for GeneratePackage, housed under testdata/TestGeneratePackage.
var testCases = []testCase{
	{
		directory:   "assets-and-archives",
		description: "A schema with assets and archives",
	},
	{
		directory:   "config-variables",
		description: "Testing config variables.",
	},
	{
		directory:   "cyclic-types",
		description: "Cyclic object types",
	},
	{
		directory:   "dash-named-schema",
		description: "Simple schema with a two part name (foo-bar)",
	},
	{
		directory:   "dashed-import-schema",
		description: "Ensure that we handle all valid go import paths",
	},
	{
		directory:   "different-enum",
		description: "An enum in a different package namespace",
	},
	{
		directory:   "embedded-crd-types",
		description: "A schema with CRD types with package names different from the main package",
	},
	{
		directory:   "enum-reference",
		description: "Ensure referencing external types/resources with referenced enums import correctly",
	},
	{
		directory:   "external-enum",
		description: "Ensure we generate valid tokens for external enums",
	},
	{
		directory:   "external-resource-schema",
		description: "External resource schema",
	},
	{
		directory: "functions-secrets",
		// Secret properties for non-Output<T> returning functions cannot be secret because they are plain.
		description: "functions that have properties that are secrets in the schema",
	},
	{
		directory:   "hyphen-url",
		description: "A resource url with a hyphen in its path",
	},
	{
		directory:   "kubernetes20",
		description: "Testing the kubernetes20 compatibility mode.",
	},
	{
		directory:   "legacy-names",
		description: "Testing the use of snake_case names and tokens.",
	},
	{
		directory:   "methods-return-plain-resource",
		description: "Test returning plain Resource objects from methods",
	},
	{
		directory:   "naming-collisions",
		description: "Schema with types that could potentially produce collisions.",
	},
	{
		directory:   "nested-module-thirdparty",
		description: "Third-party nested module",
	},
	{
		directory:   "nested-module",
		description: "Nested module",
	},
	{
		directory:   "other-owned",
		description: "CSharp rootNamespaces",
	},
	{
		directory:   "output-funcs-edgeorder",
		description: "Regresses Node compilation issues on a subset of azure-native",
	},
	{
		directory:   "output-funcs-tfbridge20",
		description: "Similar to output-funcs, but with compatibility: tfbridge20, to simulate pulumi-aws use case",
	},
	{
		directory:   "output-funcs",
		description: "Tests targeting the $fn_output helper code generation feature",
	},
	{
		directory:   "overlay-supported-languages",
		description: "Testing restricting the languages an overlay supports.",
	},
	{
		directory:   "plain-and-default",
		description: "Ensure that a resource with a plain default property works correctly",
	},
	{
		directory:   "plain-object-defaults",
		description: "Ensure that object defaults are generated (repro #8132)",
	},
	{
		directory:   "plain-object-disable-defaults",
		description: "Ensure that we can still compile safely when defaults are disabled",
	},
	{
		directory:   "plain-schema-gh6957",
		description: "Repro for #6957",
	},
	{
		directory:   "provider-config-schema",
		description: "Simple provider config schema",
	},
	{
		directory:   "provider-type-schema",
		description: "A schema with a type called Provider schema",
	},
	{
		directory:   "regress-8403",
		description: "Regress pulumi/pulumi#8403",
	},
	{
		directory:   "replace-on-change",
		description: "Simple use of replaceOnChange in schema",
	},
	{
		directory:   "resource-args-python-case-insensitive",
		description: "Resource args with same named resource and type case insensitive",
	},
	{
		directory:   "resource-args-python",
		description: "Resource args with same named resource and type",
	},
	{
		directory:   "resource-property-overlap",
		description: "A resource with the same name as its property",
	},
	{
		directory:   "secrets",
		description: "Generate a resource with secret properties",
	},
	{
		directory:   "simple-enum-schema",
		description: "Simple schema with enum types",
	},
	{
		directory:   "simple-methods-schema",
		description: "Simple schema with methods",
	},
	{
		directory:   "simple-methods-schema-single-value-returns",
		description: "Simple schema with methods that return single values",
	},
	{
		directory:   "simple-plain-schema",
		description: "Simple schema with plain properties",
	},
	{
		directory:   "simple-plain-schema-with-root-package",
		description: "Simple schema with root package set",
	},
	{
		directory:   "simple-resource-schema",
		description: "Simple schema with local resource properties",
	},
	{
		directory:   "simple-resource-schema-custom-pypackage-name",
		description: "Simple schema with local resource properties and custom Python package name",
	},
	{
		directory:   "simple-resource-with-aliases",
		description: "Simple schema with a resource that has aliases",
	},
	{
		directory:   "simple-schema-pyproject",
		description: "A simple schema that generates a pyproject.toml file",
	},
	{
		directory:   "simple-yaml-schema",
		description: "Simple schema encoded using YAML",
	},
	{
		directory:   "simplified-invokes",
		description: "Simplified invokes",
	},
	{
		directory:   "unions-inline",
		description: "Testing the use of unions/oneOf in the schema inline with the property definition.",
	},
	{
		directory:   "unions-inside-arrays",
		description: "A schema with a union type inside an array",
	},
	{
		directory:   "urn-id-properties",
		description: "Testing urn and id properties in valid locations",
	},
	{
		directory:   "using-shared-types-in-config",
		description: "Tests that shared types can be used in config.",
	},
	{
		directory:   "example-op-code",
		description: "Check that {{% examples %}} works as expected; repros 6641",
	},
}

// testCaseProvider captures details about a provider to be used in a test case for GeneratePackage.
type testCaseProvider struct {
	// The name of the package whose schema will be offered by the provider.
	name string
	// The version of the package whose schema will be offered by the provider.
	version string
}

// testCaseProviders represents a suite of providers to be used in test cases for GeneratePackage.
var testCaseProviders = []testCaseProvider{
	{
		name:    "aws",
		version: "4.36.0",
	},
	{
		name:    "azure-native",
		version: "1.28.0",
	},
	{
		name:    "google-native",
		version: "0.18.2",
	},
	{
		name:    "kubernetes",
		version: "3.7.0",
	},
	{
		name:    "random",
		version: "4.3.1",
	},
	{
		name:    "tls",
		version: "4.10.0",
	},
}

func TestGeneratePackage(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip("TestGeneratePackage is skipped on Windows")
	}

	testDir := filepath.Join("testdata", "TestGeneratePackage")
	schemaDir := filepath.Join(testDir, "_schema")

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.directory, func(t *testing.T) {
			t.Parallel()

			t.Log(testCase.description)
			testCaseDir := filepath.Join(testDir, testCase.directory)

			// Pick up the test case's schema from its directory. If we cannot load a
			// schema.json, we'll attempt to load a schema.yaml instead.
			testCaseSchema := filepath.Join(testCaseDir, "schema.json")
			if _, err := os.Stat(testCaseSchema); err != nil && os.IsNotExist(err) {
				testCaseSchema = filepath.Join(testCaseDir, "schema.yaml")
			}

			// Read and decode the schema.
			schemaBytes, err := os.ReadFile(testCaseSchema)
			require.NoErrorf(t, err, "Failed to read %s", testCaseSchema)

			schemaExt := filepath.Ext(testCaseSchema)
			var pkgSpec schema.PackageSpec
			if schemaExt == ".yaml" || schemaExt == ".yml" {
				err = yaml.Unmarshal(schemaBytes, &pkgSpec)
			} else {
				err = json.Unmarshal(schemaBytes, &pkgSpec)
			}

			require.NoErrorf(t, err, "Failed to unmarshal schema %s", testCaseSchema)

			// Create a test host/loader that offers schema for the providers/packages defined in testCaseProviders.
			host := newHostWithTestProviders(schemaDir, testCaseProviders)
			loader := schema.NewPluginLoader(host)

			// Bind the package specification (schema) we decoded against the loader/set of test providers.
			pkg, diags, err := schema.BindSpec(pkgSpec, loader, schema.ValidationOptions{
				AllowDanglingReferences: true,
			})
			require.NoErrorf(t, err, "Failed to bind schema %s", testCaseSchema)
			require.Falsef(t, diags.HasErrors(), "Schema %s has errors", testCaseSchema)

			// Generate package documentation.
			dctx := NewContext("test", pkg)
			actual, err := dctx.GeneratePackage()
			require.NoErrorf(t, err, "Failed to generate package from schema %s", testCaseSchema)

			// If PULUMI_ACCEPT is set, overwrite the expected outputs with those we generated. If not, assert that the
			// outputs we generated match the golden expected outputs already present.
			outPath := filepath.Join(testCaseDir, "docs")
			manifest, err := newTestCaseManifest(outPath)
			require.NoErrorf(t, err, "Failed to load test case manifest for %s", outPath)

			if os.Getenv("PULUMI_ACCEPT") != "" {
				err = manifest.save(actual)
				require.NoErrorf(t, err, "Failed to save test case outputs under %s", outPath)

				return
			}

			expected, err := manifest.load()
			require.NoErrorf(t, err, "Failed to load expected test case outputs under %s", outPath)

			assertEqualFiles(t, expected, actual)
		})
	}
}

// assertEqualFiles asserts that two sets of files (represented as maps from file name to byte content) are equal.
func assertEqualFiles(t *testing.T, expected map[string][]byte, actual map[string][]byte) {
	for f, expectedBytes := range expected {
		actualBytes, ok := actual[f]
		assert.Truef(t, ok, "Expected file %s not found", f)
		assert.Equalf(t, string(expectedBytes), string(actualBytes), "File %s does not match", f)
	}

	for f := range actual {
		_, ok := expected[f]
		assert.Truef(t, ok, "Unexpected file %s", f)
	}
}

// testCaseManifest represents a set of files emitted by a test case.
type testCaseManifest struct {
	// The directory in which the test case's outputs reside. This is used to load existing outputs and to save new ones.
	dir string
	// A list of the files that exist in the test case's output directory. This field is public so that we can JSON
	// serialize it to a codegen-manifest.json file in the output directory.
	EmittedFiles []string `json:"emittedFiles"`
}

// newTestCaseManifest creates a new manifest for test case outputs in the given directory.
func newTestCaseManifest(dir string) (*testCaseManifest, error) {
	m := &testCaseManifest{dir: dir}

	bytes, err := os.ReadFile(filepath.Join(dir, "codegen-manifest.json"))
	if err != nil {
		// If we couldn't load a manifest, we'll assume that this is the first run and that a call to save() will be made
		// subsequently.
		return m, nil
	}

	err = json.Unmarshal(bytes, m)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s/codegen-manifest.json: %w", dir, err)
	}

	return m, nil
}

// load reads the last set of files emitted by the test case.
func (m *testCaseManifest) load() (map[string][]byte, error) {
	files := make(map[string][]byte)
	for _, f := range m.EmittedFiles {
		path := filepath.Join(m.dir, f)
		bytes, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", path, err)
		}

		files[f] = bytes
	}

	return files, nil
}

// save writes a new set of files for the test case.
func (m *testCaseManifest) save(files map[string][]byte) error {
	m.EmittedFiles = make([]string, 0, len(files))

	// Remove the baseline directory's current contents.
	_, err := os.ReadDir(m.dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to read %s: %w", m.dir, err)
		}

		// The directory doesn't exist, so we don't need to remove it.
	} else {
		// The directory exists, so we need to remove its contents.
		err = os.RemoveAll(m.dir)
		if err != nil {
			return fmt.Errorf("failed to remove %s: %w", m.dir, err)
		}
	}

	for f, bytes := range files {
		path := filepath.Join(m.dir, f)
		err := writeFileEnsuringDir(path, bytes)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", path, err)
		}

		m.EmittedFiles = append(m.EmittedFiles, f)
	}

	sort.Strings(m.EmittedFiles)
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err = enc.Encode(m)
	if err != nil {
		return fmt.Errorf("failed to marshal codegen-manifest.json: %w", err)
	}

	manifestBytes := buf.Bytes()
	return os.WriteFile(filepath.Join(m.dir, "codegen-manifest.json"), manifestBytes, 0o600)
}

// writeFileEnsuringDir writes a file to the given path, ensuring that the directory structure leading to the file
// exists.
func writeFileEnsuringDir(path string, bytes []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil && !os.IsExist(err) {
		return err
	}

	return os.WriteFile(path, bytes, 0o600)
}

// newHostWithTestProviders creates a new plugin host that offers schema for the providers/packages defined in the
// given list of testCaseProviders.
func newHostWithTestProviders(schemaDir string, providers []testCaseProvider) plugin.Host {
	providerFor := func(p testCaseProvider) *deploytest.PluginLoader {
		return deploytest.NewProviderLoader(
			tokens.Package(p.name),
			semver.MustParse(p.version),
			func() (plugin.Provider, error) {
				return &deploytest.Provider{
					GetSchemaF: func(context.Context, plugin.GetSchemaRequest) (plugin.GetSchemaResponse, error) {
						f := fmt.Sprintf("%s-%s.json", p.name, p.version)
						path := filepath.Join(schemaDir, f)

						data, err := os.ReadFile(path)
						if err != nil {
							return plugin.GetSchemaResponse{}, err
						}

						return plugin.GetSchemaResponse{Schema: data}, nil
					},
				}, nil
			},
		)
	}

	loaders := make([]*deploytest.PluginLoader, len(providers))
	for i, p := range providers {
		loaders[i] = providerFor(p)
	}

	return deploytest.NewPluginHost(
		nil, /*sink*/
		nil, /*statusSink*/
		nil, /*languageRuntime*/
		loaders...,
	)
}
