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

package overview

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// cosignParameter is the base64 parameterization parameter chainguard-dev's
// cosign package carries, which decodes to
// {"remote":{"url":"registry.opentofu.org/chainguard-dev/cosign","version":"0.4.19"}}.
const cosignParameter = "eyJyZW1vdGUiOnsidXJsIjoicmVnaXN0cnkub3BlbnRvZnUub3JnL2NoYWluZ3VhcmQtZGV2" +
	"L2Nvc2lnbiIsInZlcnNpb24iOiIwLjQuMTkifX0="

func spec(t *testing.T, jsonSpec string) *schema.PackageSpec {
	t.Helper()

	var out schema.PackageSpec
	require.NoError(t, json.Unmarshal([]byte(jsonSpec), &out))
	return &out
}

func TestPascalCase(t *testing.T) {
	t.Parallel()

	tests := []struct{ in, want string }{
		{"ise", "Ise"},
		{"vault", "Vault"},
		{"aws-native", "AwsNative"},
		{"azure_native", "AzureNative"},
		{"docker-build", "DockerBuild"},
		{"newrelic", "Newrelic"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, pascalCase(tt.in))
		})
	}
}

func TestGoPackageName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "ise", goPackageName("ise"))
	assert.Equal(t, "awsnative", goPackageName("aws-native"))
	assert.Equal(t, "dockerbuild", goPackageName("docker_build"))
}

func TestTerraformProviderSource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		jsonSpec string
		want     string
		wantOK   bool
	}{
		{
			name: "terraform provider",
			// parameter is the base64 of
			// {"remote":{"url":"registry.opentofu.org/chainguard-dev/cosign","version":"0.4.19"}}
			jsonSpec: `{"name":"cosign","parameterization":{
				"baseProvider":{"name":"terraform-provider","version":"1.0.1"},
				"parameter":"` + cosignParameter + `"}}`,
			want:   "chainguard-dev/cosign",
			wantOK: true,
		},
		{
			name:     "not parameterized",
			jsonSpec: `{"name":"vault"}`,
			wantOK:   false,
		},
		{
			name: "parameterized on some other base provider",
			jsonSpec: `{"name":"other","parameterization":{
				"baseProvider":{"name":"something-else","version":"1.0.0"},"parameter":"e30="}}`,
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, ok := terraformProviderSource(spec(t, tt.jsonSpec))
			assert.Equal(t, tt.wantOK, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDefaultLanguagesSkipsSDKsForParameterizedPackages(t *testing.T) {
	t.Parallel()

	// cosign declares all five language blocks and publishes none of them. A
	// parameterized provider is consumed with `pulumi package add`, which is
	// the one case the schema does let us call.
	parameterized := spec(t, `{"name":"cosign",
		"language":{"csharp":{},"go":{},"java":{},"nodejs":{},"python":{}},
		"parameterization":{"baseProvider":{"name":"terraform-provider","version":"1.0.1"},
		"parameter":"eyJyZW1vdGUiOnsidXJsIjoicmVnaXN0cnkub3BlbnRvZnUub3JnL2NoYWluZ3VhcmQtZGV2L2Nvc2lnbiJ9fQ=="}}`)
	assert.Empty(t, DefaultLanguages(parameterized))

	// vault declares four blocks; the missing java one is why --languages exists.
	vault := spec(t, `{"name":"vault","language":{"csharp":{},"go":{},"nodejs":{},"python":{}}}`)
	assert.Equal(t, []string{LangTypeScript, LangPython, LangGo, LangCSharp}, DefaultLanguages(vault))
}

func TestDeriveInstallsDefaults(t *testing.T) {
	t.Parallel()

	// No language overrides at all: every identifier comes from a default.
	plan := DeriveInstalls(
		spec(t, `{"name":"ise","repository":"https://github.com/pulumi/pulumi-ise",
			"language":{"csharp":{},"go":{},"java":{},"nodejs":{},"python":{}}}`),
		"v0.5.0", SDKLanguages)

	byLanguage := map[string]Install{}
	for _, install := range plan.Languages {
		byLanguage[install.Language] = install
	}

	assert.Equal(t, "npm install @pulumi/ise", byLanguage[LangTypeScript].Command)
	assert.Equal(t, "pip install pulumi-ise", byLanguage[LangPython].Command)
	assert.Equal(t, "dotnet add package Pulumi.Ise", byLanguage[LangCSharp].Command)
	assert.Equal(t, "com.pulumi:ise", byLanguage[LangJava].Package)
	assert.Contains(t, byLanguage[LangJava].Gradle, "com.pulumi:ise:0.5.0")

	// No importBasePath, so the Go path is guessed from the repository and the
	// caller is told so.
	assert.Equal(t, "go get github.com/pulumi/pulumi-ise/sdk/go/ise", byLanguage[LangGo].Command)
	assert.Contains(t, plan.Warnings[0], "no go importBasePath")
}

func TestDeriveInstallsPrefersSchemaNames(t *testing.T) {
	t.Parallel()

	plan := DeriveInstalls(
		spec(t, `{"name":"logfire","language":{
			"nodejs":{"packageName":"@pydantic/pulumi-logfire"},
			"python":{"packageName":"pulumi_logfire"},
			"go":{"importBasePath":"github.com/pydantic/pulumi-logfire/sdk/go/logfire"}}}`),
		"v0.1.19", []string{LangTypeScript, LangPython, LangGo})

	require.Len(t, plan.Languages, 3)
	assert.Equal(t, "npm install @pydantic/pulumi-logfire", plan.Languages[0].Command)
	// PyPI treats the two spellings as one project; published pages use hyphens.
	assert.Equal(t, "pip install pulumi-logfire", plan.Languages[1].Command)
	assert.Equal(t, "go get github.com/pydantic/pulumi-logfire/sdk/go/logfire", plan.Languages[2].Command)
	assert.Empty(t, plan.Warnings)
}

func TestDeriveInstallsKeepsExistingGoMajorVersion(t *testing.T) {
	t.Parallel()

	plan := DeriveInstalls(
		spec(t, `{"name":"vault","language":{
			"go":{"importBasePath":"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"}}}`),
		"v7.12.0", []string{LangGo})

	require.Len(t, plan.Languages, 1)
	assert.Equal(t, "go get github.com/pulumi/pulumi-vault/sdk/v7/go/vault", plan.Languages[0].Command)
	assert.Empty(t, plan.Warnings)
}

func TestDeriveInstallsOmitsPluginInstall(t *testing.T) {
	t.Parallel()

	// A pluginDownloadURL is baked into the generated SDKs, so the engine
	// downloads the plugin on first use. Emitting `pulumi plugin install`
	// because the field is set would tell readers to do by hand the very thing
	// that field automates.
	plan := DeriveInstalls(
		spec(t, `{"name":"example","pluginDownloadURL":"github://api.github.com/example/pulumi-example",
			"language":{"nodejs":{}}}`),
		"v1.2.3", []string{LangTypeScript})

	assert.NotContains(t, RenderInstallation(plan), "pulumi plugin install")
}

func TestRenderInstallationCoversEveryLanguage(t *testing.T) {
	t.Parallel()

	// logfire publishes three SDKs. C# and Java readers must still be given a
	// working path rather than an empty panel: `pulumi package add` generates
	// an SDK locally, taking its language from the project's runtime, which is
	// why the command reads identically for them and for YAML and HCL.
	plan := DeriveInstalls(
		spec(t, `{"name":"logfire","language":{
			"nodejs":{"packageName":"@pydantic/pulumi-logfire"},
			"python":{"packageName":"pulumi_logfire"},
			"go":{"importBasePath":"github.com/pydantic/pulumi-logfire/sdk/go/logfire"}}}`),
		"v0.1.19", []string{LangTypeScript, LangPython, LangGo})

	got := RenderInstallation(plan)

	assert.Contains(t, got, `{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}`)
	for _, lang := range AllLanguages {
		assert.Contains(t, got, "{{% choosable language "+lang+" %}}",
			"every language needs a tab, including those with no published SDK")
	}
	assert.Contains(t, got, "No SDK is published for this language.")
	assert.Equal(t, 3, strings.Count(got, "pulumi package add logfire"),
		"csharp, java and yaml fall back to `pulumi package add`; hcl does not")
}

func TestRenderInstallationHCLDeclaresTheProvider(t *testing.T) {
	t.Parallel()

	// HCL has no `pulumi package add` step. A program names the provider in its
	// own required_providers block and `pulumi install` fetches it -- which is
	// also why the chooser is never collapsed to a single command, even for a
	// package with no SDKs at all: HCL's instructions always differ.
	native := RenderInstallation(DeriveInstalls(spec(t, `{"name":"logfire"}`), "v0.1.19", nil))

	assert.Contains(t, native, "{{% choosable language hcl %}}")
	assert.Contains(t, native, `source  = "pulumi/logfire"`)
	// A native Pulumi provider takes an exact semver version, not a constraint.
	assert.Contains(t, native, `version = "0.1.19"`)
	assert.Contains(t, native, "pulumi install")

	// A parameterized package is bridged from its Terraform provider, so it
	// reuses the upstream source and version rather than the pulumi/ prefix.
	bridged := RenderInstallation(DeriveInstalls(
		spec(t, `{"name":"cosign","parameterization":{
			"baseProvider":{"name":"terraform-provider","version":"1.0.1"},
			"parameter":"`+cosignParameter+`"}}`), "", nil))

	assert.Contains(t, bridged, `source  = "chainguard-dev/cosign"`)
	assert.Contains(t, bridged, `version = "0.4.19"`)
	assert.NotContains(t, bridged, "pulumi/cosign")
}

func TestRenderInstallationWarnsWhenHCLHasNoVersion(t *testing.T) {
	t.Parallel()

	plan := DeriveInstalls(spec(t, `{"name":"logfire"}`), "", nil)

	assert.Contains(t, RenderInstallation(plan), `version = "VERSION"`)
	assert.Contains(t, plan.Warnings[0], "exact semver version")
}

func TestReadSchemaConfig(t *testing.T) {
	t.Parallel()

	cfg := ReadSchemaConfig(spec(t, `{"name":"example","config":{
		"defaults":["apiToken"],
		"variables":{
			"region":{"type":"string","description":"The region.",
				"defaultInfo":{"environment":["EXAMPLE_REGION","EXAMPLE_DEFAULT_REGION"]}},
			"apiToken":{"type":"string","description":"The token.","secret":true},
			"authLogin":{"$ref":"#/types/example:config/authLogin:authLogin","description":"Login."}}},
		"types":{"example:config/authLogin:authLogin":{"type":"object","properties":{}}}}`))

	// Sorted by name, so that repeated runs produce identical output.
	require.Len(t, cfg.Vars, 3)
	assert.Equal(t, []string{"apiToken", "authLogin", "region"},
		[]string{cfg.Vars[0].Name, cfg.Vars[1].Name, cfg.Vars[2].Name})

	// Requiredness comes from the config block's "defaults" list -- note the
	// JSON key does not match the Go field name.
	assert.True(t, cfg.Vars[0].Required)
	assert.True(t, cfg.Vars[0].Secret)
	assert.False(t, cfg.Vars[2].Required)

	// A $ref config variable resolves to the referenced type's own type.
	assert.Equal(t, "object", cfg.Vars[1].Type)

	assert.Equal(t, []string{"EXAMPLE_REGION", "EXAMPLE_DEFAULT_REGION"}, cfg.Vars[2].EnvVars)
}

func TestReadSchemaConfigSanitizesDescriptions(t *testing.T) {
	t.Parallel()

	// Bridged provider descriptions carry language-specific spans that the
	// registry's own renderer strips; a pasted snippet must not show them.
	description := `Infers the endpoint from the` +
		`<span pulumi-lang-nodejs=\" apiKey \" pulumi-lang-python=\" api_key \">` +
		` api_key </span>region.`

	cfg := ReadSchemaConfig(spec(t, `{"name":"example","config":{"variables":{
		"baseUrl":{"type":"string","description":"`+description+`"}}}}`))

	require.Len(t, cfg.Vars, 1)
	assert.NotContains(t, cfg.Vars[0].Description, "pulumi-lang-")
	assert.Contains(t, cfg.Vars[0].Description, "api_key")
}

func TestRenderConfigurationFoldsInEnvironmentVariables(t *testing.T) {
	t.Parallel()

	cfg := SchemaConfig{Package: "example", Vars: []ConfigVar{
		{Name: "region", Description: "The region.", EnvVars: []string{"EXAMPLE_REGION"}},
		{Name: "token", Description: "The token.", Required: true, Secret: true},
		// Already named in the description, so it is not appended twice.
		{Name: "url", Description: "Read from `EXAMPLE_URL` when unset.", EnvVars: []string{"EXAMPLE_URL"}},
	}}

	got := RenderConfiguration(cfg, ConfigStyleList)

	assert.Contains(t, got,
		"- `region` (Optional) — The region. May also be set with the `EXAMPLE_REGION` environment variable.")
	assert.Contains(t, got, "- `token` (Required, Secret) — The token.")
	assert.Contains(t, got, "- `url` (Optional) — Read from `EXAMPLE_URL` when unset.\n")
	assert.Contains(t, got, "TODO: Environment-variable fallbacks")
	assert.Contains(t, got, "TODO: Mutually exclusive options")
}

func TestRenderConfigurationEscapesTableCells(t *testing.T) {
	t.Parallel()

	cfg := SchemaConfig{Package: "example", Vars: []ConfigVar{
		{Name: "mode", Description: "One of a | b | c.\nSpanning two lines."},
	}}

	row := RenderConfiguration(cfg, ConfigStyleTable)

	// A raw pipe would end the cell, and a newline would end the row.
	assert.Contains(t, row, `| `+"`mode`"+` | No | No | One of a \| b \| c. Spanning two lines. |`)
}
