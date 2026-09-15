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
	"bytes"
	"path/filepath"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
)

func schemaFixture(name string) string {
	return filepath.Join("testdata", "gen", name+".json")
}

// runGen executes one of the gen-* commands against a fixture schema and
// returns its stdout and stderr. Every case reads a local file, so no test
// here touches the network.
func runGen(t *testing.T, cmd *cobra.Command, args ...string) (string, string) {
	t.Helper()

	var out, errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs(args)
	require.NoError(t, cmd.Execute())

	return out.String(), errOut.String()
}

func TestGenInstall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
	}{
		{
			// Five languages on pure defaults: @pulumi/ise, pulumi-ise,
			// Pulumi.Ise, com.pulumi:ise, and a Go path with no /vN element.
			name: "defaults",
			args: []string{"--schemaFile", schemaFixture("ise")},
		},
		{
			// ise's schema carries no version, so Maven and Gradle need one.
			name: "version-flag",
			args: []string{"--schemaFile", schemaFixture("ise"), "--version", "v0.5.0"},
		},
		{
			// The ~100 bridged providers that publish nothing: one command,
			// no chooser.
			name: "no-sdks",
			args: []string{"--schemaFile", schemaFixture("ise"), "--languages", "none"},
		},
		{
			// vault declares no java block but does publish com.pulumi/vault,
			// so the author names the languages. Its Go import path already
			// carries /sdk/v7 and must not be rewritten.
			name: "explicit-languages",
			args: []string{
				"--schemaFile", schemaFixture("vault"),
				"--languages", "typescript,python,go,csharp,java",
				"--version", "v7.12.0",
			},
		},
		{
			// A parameterized bridged provider: the `pulumi package add
			// terraform-provider <ns>/<name>` form, recovered from the
			// base64 parameterization parameter.
			name: "parameterized",
			args: []string{"--schemaFile", schemaFixture("cosign")},
		},
		{
			// A community package with a scoped npm name and only three SDKs.
			name: "community",
			args: []string{"--schemaFile", schemaFixture("logfire"), "--version", "v0.1.19"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out, _ := runGen(t, GenInstallCmd(pkg.NewHTTPClient()), tt.args...)
			autogold.ExpectFile(t, autogold.Raw(out))
		})
	}
}

func TestGenInstallWarnsWithoutVersion(t *testing.T) {
	t.Parallel()

	out, errOut := runGen(t, GenInstallCmd(pkg.NewHTTPClient()), "--schemaFile", schemaFixture("ise"))

	assert.Contains(t, out, "com.pulumi:ise:VERSION")
	assert.Contains(t, errOut, "Maven and Gradle coordinates require a version")
}

func TestGenInstallRejectsUnknownLanguage(t *testing.T) {
	t.Parallel()

	cmd := GenInstallCmd(pkg.NewHTTPClient())
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	// yaml is a chooser key but not an SDK language: it is always rendered
	// from `pulumi package add`, so naming it here is a mistake worth catching.
	cmd.SetArgs([]string{"--schemaFile", schemaFixture("ise"), "--languages", "yaml"})

	assert.ErrorContains(t, cmd.Execute(), `unknown language "yaml"`)
}

func TestGenConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "table",
			args: []string{"--schemaFile", schemaFixture("logfire"), "--style", "table"},
		},
		{
			// The default style. Most existing pages use a bullet list, and it
			// keeps long descriptions readable where a table column squeezes them.
			name: "list",
			args: []string{"--schemaFile", schemaFixture("logfire")},
		},
		{
			// vault is the honest worst case: no defaults list, so nothing
			// reads as required; `token` is not marked secret; authLogin is a
			// $ref to a complex type; two parameters carry environment-variable
			// fallbacks that get folded into their descriptions.
			name: "poor-fidelity",
			args: []string{"--schemaFile", schemaFixture("vault")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out, _ := runGen(t, GenConfigCmd(pkg.NewHTTPClient()), tt.args...)
			autogold.ExpectFile(t, autogold.Raw(out))
		})
	}
}

func TestGenConfigEmptySchemaConfig(t *testing.T) {
	t.Parallel()

	out, errOut := runGen(t, GenConfigCmd(pkg.NewHTTPClient()), "--schemaFile", schemaFixture("empty-config"))

	assert.Empty(t, out)
	assert.Contains(t, errOut, "declares no provider configuration")
}

func TestGenConfigRejectsUnknownStyle(t *testing.T) {
	t.Parallel()

	cmd := GenConfigCmd(pkg.NewHTTPClient())
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.SetArgs([]string{"--schemaFile", schemaFixture("logfire"), "--style", "prose"})

	assert.ErrorContains(t, cmd.Execute(), `unknown --style "prose"`)
}

func TestGenSchemaSourceValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
		err  string
	}{
		{
			name: "neither",
			// An explicitly empty slice, not nil: cobra falls back to
			// os.Args[1:] when SetArgs is given nil, which would let `go test
			// -update` leak into the command under test.
			args: []string{},
			err:  "one of --schemaFile or --schemaFileURL is required",
		},
		{
			name: "both",
			args: []string{"--schemaFile", schemaFixture("ise"), "--schemaFileURL", "https://example.com/schema.json"},
			err:  "mutually exclusive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cmd := GenInstallCmd(pkg.NewHTTPClient())
			cmd.SetOut(new(bytes.Buffer))
			cmd.SetErr(new(bytes.Buffer))
			cmd.SetArgs(tt.args)

			assert.ErrorContains(t, cmd.Execute(), tt.err)
		})
	}
}
