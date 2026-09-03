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
	"fmt"
	"sort"
	"strings"

	dotnet "github.com/pulumi/pulumi-dotnet/pulumi-language-dotnet/v3/codegen"
	"github.com/pulumi/pulumi-java/pkg/codegen/java"
	go_gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// The chooser language keys, in the order docs/overview-page.md lists them.
const (
	LangTypeScript = "typescript"
	LangPython     = "python"
	LangGo         = "go"
	LangCSharp     = "csharp"
	LangJava       = "java"
	LangYAML       = "yaml"
	LangHCL        = "hcl"
)

// SDKLanguages are the languages that install a per-language SDK. YAML and HCL
// are excluded: they consume the package directly through `pulumi package add`.
var SDKLanguages = []string{LangTypeScript, LangPython, LangGo, LangCSharp, LangJava}

// AllLanguages is every key the chooser accepts, in rendering order.
var AllLanguages = append(append([]string{}, SDKLanguages...), LangYAML, LangHCL)

// schemaLanguageKeys maps a chooser key to the key the schema's language blob
// uses for it. They agree everywhere except TypeScript, which the schema calls
// nodejs.
var schemaLanguageKeys = map[string]string{
	LangTypeScript: "nodejs",
	LangPython:     "python",
	LangGo:         "go",
	LangCSharp:     "csharp",
	LangJava:       "java",
}

// Install is the installation instruction for a single language.
type Install struct {
	// Language is a chooser key.
	Language string
	// Package is the published artifact's identifier -- the npm name, the PyPI
	// project, the Go import path, the NuGet id, or the Maven coordinate.
	Package string
	// Command is the shell one-liner that installs it. Empty for Java, which
	// has no one-line install and uses Maven/Gradle instead.
	Command string
	// Maven and Gradle are set for Java only.
	Maven  string
	Gradle string
}

// InstallPlan is everything the ## Installation section needs.
type InstallPlan struct {
	Package string
	Version string
	// Languages is empty when the package publishes no SDKs, in which case
	// PackageAdd is the whole installation section.
	Languages []Install
	// PackageAdd is the `pulumi package add` command. It covers every language
	// except HCL, which declares providers in the program instead.
	PackageAdd string
	// HCLSource is the provider source an HCL program names in its
	// required_providers block, and HCLVersion the version it pins.
	HCLSource  string
	HCLVersion string
	// Warnings describe what could not be derived and needs the author's hand.
	Warnings []string
}

// DefaultLanguages guesses which languages a package publishes SDKs for.
//
// The guess is only that. A schema's language blob is populated by the code
// generator, not by whatever the release pipeline actually pushed, so it is
// wrong in both directions: pulumi-vault declares no java block yet publishes
// com.pulumi/vault, and every one of the ~100 bridged providers that ship no
// SDKs at all still declares four or five language blocks. A parameterized
// provider is the one case we can call confidently -- it is consumed through
// `pulumi package add` -- so it reports no SDK languages. Everywhere else the
// author overrides this with --languages.
func DefaultLanguages(spec *schema.PackageSpec) []string {
	if isTerraformParameterized(spec) {
		return nil
	}

	var langs []string
	for _, lang := range SDKLanguages {
		if _, ok := spec.Language[schemaLanguageKeys[lang]]; ok {
			langs = append(langs, lang)
		}
	}
	return langs
}

// DeriveInstalls builds the installation plan for the given SDK languages.
// Pass a nil or empty languages slice for a package that publishes no SDKs.
func DeriveInstalls(spec *schema.PackageSpec, version string, languages []string) InstallPlan {
	plan := InstallPlan{
		Package:    spec.Name,
		Version:    strings.TrimPrefix(version, "v"),
		PackageAdd: packageAddCommand(spec),
	}
	plan.HCLSource, plan.HCLVersion = hclProvider(spec, plan.Version)
	if plan.HCLVersion == "" {
		plan.HCLVersion = "VERSION"
		plan.Warnings = append(plan.Warnings,
			"HCL pins a native Pulumi provider to an exact semver version and the schema carries none; "+
				"pass --version, or fill in the VERSION placeholder by hand")
	}

	// Deliberately no `pulumi plugin install` line. A schema's
	// pluginDownloadURL is baked into every generated SDK, which is what lets
	// the engine fetch the plugin binary on first use -- so its presence is
	// evidence the manual step is *not* needed. The rare provider whose plugin
	// really must be installed by hand has to say so itself.
	for _, lang := range sortLanguages(languages) {
		install, warnings := deriveInstall(spec, plan.Version, lang)
		plan.Warnings = append(plan.Warnings, warnings...)
		if install != nil {
			plan.Languages = append(plan.Languages, *install)
		}
	}

	return plan
}

// sortLanguages puts the requested languages into the standard's order and
// drops duplicates and the non-SDK keys, which are rendered from PackageAdd.
func sortLanguages(languages []string) []string {
	order := make(map[string]int, len(SDKLanguages))
	for i, lang := range SDKLanguages {
		order[lang] = i
	}

	seen := map[string]bool{}
	var out []string
	for _, lang := range languages {
		lang = strings.ToLower(strings.TrimSpace(lang))
		if _, ok := order[lang]; !ok || seen[lang] {
			continue
		}
		seen[lang] = true
		out = append(out, lang)
	}
	sort.Slice(out, func(i, j int) bool { return order[out[i]] < order[out[j]] })
	return out
}

func deriveInstall(spec *schema.PackageSpec, version, lang string) (*Install, []string) {
	switch lang {
	case LangTypeScript:
		var info nodejs.NodePackageInfo
		decodeLanguageInfo(spec, "nodejs", &info)
		name := info.PackageName
		if name == "" {
			name = "@pulumi/" + spec.Name
		}
		return &Install{Language: lang, Package: name, Command: "npm install " + name}, nil

	case LangPython:
		var info python.PackageInfo
		decodeLanguageInfo(spec, "python", &info)
		name := info.PackageName
		if name == "" {
			name = "pulumi_" + spec.Name
		}
		// PyPI treats underscores and hyphens as the same project. Every
		// published Overview page spells the install with hyphens.
		name = strings.ReplaceAll(name, "_", "-")
		return &Install{Language: lang, Package: name, Command: "pip install " + name}, nil

	case LangGo:
		var info go_gen.GoPackageInfo
		decodeLanguageInfo(spec, "go", &info)
		path := info.ImportBasePath
		if path == "" {
			var warnings []string
			path, warnings = fallbackGoImportPath(spec)
			return &Install{Language: lang, Package: path, Command: "go get " + path}, warnings
		}
		// Never synthesize the /vN major-version element: a schema that needs
		// one already carries it (pulumi-vault's is .../sdk/v7/go/vault), and
		// guessing produces an import path that does not resolve.
		return &Install{Language: lang, Package: path, Command: "go get " + path}, nil

	case LangCSharp:
		var info dotnet.CSharpPackageInfo
		decodeLanguageInfo(spec, "csharp", &info)
		id := info.GetRootNamespace() + "." + pascalCase(spec.Name)
		return &Install{Language: lang, Package: id, Command: "dotnet add package " + id}, nil

	case LangJava:
		var info java.PackageInfo
		decodeLanguageInfo(spec, "java", &info)
		groupID := strings.TrimSuffix(info.BasePackageOrDefault(), ".")
		artifactID := spec.Name
		var warnings []string
		v := version
		if v == "" {
			v = "VERSION"
			warnings = append(warnings,
				"Maven and Gradle coordinates require a version and the schema carries none; "+
					"pass --version, or fill in the VERSION placeholder by hand")
		}
		return &Install{
			Language: lang,
			Package:  groupID + ":" + artifactID,
			Maven:    mavenBlock(groupID, artifactID, v),
			Gradle:   fmt.Sprintf("implementation '%s:%s:%s'", groupID, artifactID, v),
		}, warnings
	}

	return nil, nil
}

func mavenBlock(groupID, artifactID, version string) string {
	return fmt.Sprintf(`<dependency>
    <groupId>%s</groupId>
    <artifactId>%s</artifactId>
    <version>%s</version>
</dependency>`, groupID, artifactID, version)
}

// fallbackGoImportPath builds an import path from the schema's repository when
// the go language blob declares no importBasePath.
func fallbackGoImportPath(spec *schema.PackageSpec) (string, []string) {
	repo := spec.Repository
	repo = strings.TrimSuffix(strings.TrimSuffix(repo, "/"), ".git")
	repo = strings.TrimPrefix(strings.TrimPrefix(repo, "https://"), "http://")
	if repo == "" {
		return "<TODO: Go import path>", []string{
			"the schema declares neither a go importBasePath nor a repository, " +
				"so the Go import path could not be derived",
		}
	}
	return fmt.Sprintf("%s/sdk/go/%s", repo, goPackageName(spec.Name)), []string{
		"the schema declares no go importBasePath; the Go import path was guessed " +
			"from the repository and may need a /vN major-version element",
	}
}

// goPackageName reduces a package name to the identifier a Go SDK's leaf
// directory uses: aws-native becomes awsnative.
func goPackageName(name string) string {
	var b strings.Builder
	for _, r := range name {
		if r == '-' || r == '_' || r == '.' {
			continue
		}
		b.WriteRune(r)
	}
	return strings.ToLower(b.String())
}

// pascalCase renders a package name the way the .NET generator names its
// assembly: aws-native becomes AwsNative.
func pascalCase(name string) string {
	parts := strings.FieldsFunc(name, func(r rune) bool { return r == '-' || r == '_' || r == '.' })
	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		b.WriteString(strings.ToLower(part[1:]))
	}
	return b.String()
}

// packageAddCommand builds the `pulumi package add` line YAML and HCL use, and
// which is the whole installation section for a package with no SDKs.
func packageAddCommand(spec *schema.PackageSpec) string {
	if source, ok := terraformProviderSource(spec); ok {
		return "pulumi package add terraform-provider " + source
	}
	return "pulumi package add " + spec.Name
}

// hclProvider returns the source and version an HCL program uses to reference
// this package.
//
// HCL does not use `pulumi package add`. A program names the provider in its
// required_providers block and `pulumi install` fetches it. A source prefixed
// with "pulumi/" resolves to a native Pulumi provider and must be pinned to an
// exact semver version; any other source is bridged from its Terraform
// provider, so a parameterized package reuses the upstream address and version
// it already carries.
func hclProvider(spec *schema.PackageSpec, version string) (source, pinned string) {
	if remote, ok := terraformProviderRemote(spec); ok {
		if remote.Version != "" {
			return remote.source(), remote.Version
		}
		return remote.source(), version
	}
	return "pulumi/" + spec.Name, version
}

func isTerraformParameterized(spec *schema.PackageSpec) bool {
	_, ok := terraformProviderSource(spec)
	return ok
}

// terraformRemote is the upstream provider a parameterized bridged package
// wraps, as recorded in parameterization.parameter.
type terraformRemote struct {
	URL     string `json:"url"`
	Version string `json:"version"`
}

// source strips the registry host from the remote URL, leaving the
// "<namespace>/<name>" address both `pulumi package add terraform-provider`
// and an HCL required_providers block expect.
func (r terraformRemote) source() string {
	parts := strings.Split(strings.Trim(r.URL, "/"), "/")
	if len(parts) < 2 {
		return ""
	}
	return strings.Join(parts[len(parts)-2:], "/")
}

// terraformProviderRemote decodes the base64 parameter a parameterized bridged
// package carries -- {"remote":{"url":"<registry host>/<ns>/<name>","version":
// "..."}} -- which is what the ~100 registry packages whose Overview pages read
// `pulumi package add terraform-provider chainguard-dev/cosign` are built from.
func terraformProviderRemote(spec *schema.PackageSpec) (terraformRemote, bool) {
	if spec.Parameterization == nil || spec.Parameterization.BaseProvider.Name != "terraform-provider" {
		return terraformRemote{}, false
	}

	var parameter struct {
		Remote terraformRemote `json:"remote"`
	}
	if err := json.Unmarshal(spec.Parameterization.Parameter, &parameter); err != nil {
		return terraformRemote{}, false
	}
	if parameter.Remote.source() == "" {
		return terraformRemote{}, false
	}
	return parameter.Remote, true
}

// terraformProviderSource is the "<namespace>/<name>" address of the upstream
// provider, when this package is a parameterized bridge over one.
func terraformProviderSource(spec *schema.PackageSpec) (string, bool) {
	remote, ok := terraformProviderRemote(spec)
	if !ok {
		return "", false
	}
	return remote.source(), true
}

func decodeLanguageInfo(spec *schema.PackageSpec, key string, into any) {
	raw, ok := spec.Language[key]
	if !ok || len(raw) == 0 {
		return
	}
	// A malformed or unexpected language blob is not worth failing over: the
	// caller falls back to the generator's own defaults, which is what the
	// language's code generator would do anyway.
	_ = json.Unmarshal(raw, into)
}
