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
	"fmt"
	"strings"
)

// The two constraints docs/overview-page.md calls out as impossible to derive
// from a schema. The generator emits them as reminders rather than pretending
// the generated reference is complete.
const (
	envVarPlaceholder = `<!-- TODO: Environment-variable fallbacks. A parameter's environment-variable ` +
		`fallback is often read by the vendor SDK a layer beneath the schema, so it cannot be generated. ` +
		`Where a parameter can be supplied by an environment variable, name the variable in its description. -->`

	exclusivityPlaceholder = `<!-- TODO: Mutually exclusive options. If setting one parameter forbids or ` +
		`overrides another, or if a group of parameters must be supplied together, say so in the description ` +
		`of every parameter involved. Nothing else in the docs will surface that constraint. -->`
)

// RenderInstallation renders the ## Installation section.
//
// Mind the delimiters: chooser takes angle brackets and choosable takes percent
// signs. Mixing them silently breaks the rendered page.
func RenderInstallation(plan InstallPlan) string {
	var b strings.Builder
	b.WriteString("## Installation\n\n")

	installs := make(map[string]Install, len(plan.Languages))
	for _, install := range plan.Languages {
		installs[install.Language] = install
	}

	// Every language gets a tab, not only the ones with a published SDK. A
	// reader who picks C# and finds an empty panel has been told nothing; a
	// reader told to run `pulumi package add` has a working path.
	fmt.Fprintf(&b, "{{< chooser language %q >}}\n", strings.Join(AllLanguages, ","))

	for _, lang := range AllLanguages {
		fmt.Fprintf(&b, "{{%% choosable language %s %%}}\n\n", lang)

		install, published := installs[lang]
		switch {
		case lang == LangHCL:
			// HCL has no `pulumi package add` step at all: a program names the
			// provider in its own required_providers block and `pulumi install`
			// fetches it.
			b.WriteString("Declare the provider in your program, then run `pulumi install`:\n\n")
			b.WriteString(fence("hcl", requiredProvidersBlock(plan)))
			b.WriteString("\n")
			b.WriteString(bashFence("pulumi install"))
		case published && lang == LangJava:
			b.WriteString("Maven:\n\n")
			b.WriteString(fence("xml", install.Maven))
			b.WriteString("\nGradle:\n\n")
			b.WriteString(fence("groovy", install.Gradle))
		case published:
			b.WriteString(bashFence(install.Command))
		case lang == LangYAML:
			// YAML consumes the package directly rather than through an SDK.
			b.WriteString(bashFence(plan.PackageAdd))
		default:
			// No SDK is published for this language, so generate one locally.
			// `pulumi package add` takes the SDK language from the project's
			// runtime, which is why the command reads the same here as it does
			// for YAML.
			b.WriteString(localSDKNote)
			b.WriteString(bashFence(plan.PackageAdd))
		}

		b.WriteString("\n{{% /choosable %}}\n")
	}

	b.WriteString("{{< /chooser >}}\n")

	return b.String()
}

// localSDKNote precedes the `pulumi package add` command shown for a language
// with no published SDK, so a reader understands they are generating one rather
// than installing a package from a feed.
const localSDKNote = "No SDK is published for this language. Run the following command from your " +
	"Pulumi project to generate one locally and record it in `Pulumi.yaml`:\n\n"

// requiredProvidersBlock renders the terraform block an HCL program uses to
// reference the package. A "pulumi/"-prefixed source resolves to the native
// Pulumi provider and must carry an exact semver version; any other source is
// bridged from its Terraform provider.
func requiredProvidersBlock(plan InstallPlan) string {
	return fmt.Sprintf(`terraform {
  required_providers {
    %s = {
      source  = %q
      version = %q
    }
  }
}`, plan.Package, plan.HCLSource, plan.HCLVersion)
}

// ConfigStyle selects the shape of the generated configuration reference.
type ConfigStyle string

const (
	// ConfigStyleTable renders a GFM table, one row per parameter.
	ConfigStyleTable ConfigStyle = "table"
	// ConfigStyleList renders a bullet list, which is what most existing
	// Overview pages use.
	ConfigStyleList ConfigStyle = "list"
)

// RenderConfiguration renders the ## Configuration section's parameter
// reference in the requested style.
func RenderConfiguration(cfg SchemaConfig, style ConfigStyle) string {
	var b strings.Builder
	b.WriteString("## Configuration\n\n")

	if style == ConfigStyleList {
		for _, v := range cfg.Vars {
			fmt.Fprintf(&b, "- `%s` (%s) — %s\n", v.Name, listMarkers(v), configDescription(v))
		}
	} else {
		b.WriteString("| Name | Required | Secret | Description |\n")
		b.WriteString("|---|---|---|---|\n")
		for _, v := range cfg.Vars {
			fmt.Fprintf(&b, "| `%s` | %s | %s | %s |\n",
				v.Name, yesNo(v.Required), yesNo(v.Secret), escapeCell(configDescription(v)))
		}
	}

	b.WriteString("\n")
	b.WriteString(envVarPlaceholder)
	b.WriteString("\n")
	b.WriteString(exclusivityPlaceholder)
	b.WriteString("\n")

	return b.String()
}

// listMarkers renders the required/secret pair as the parenthetical the bullet
// form uses, e.g. "Required, Secret".
func listMarkers(v ConfigVar) string {
	markers := []string{"Optional"}
	if v.Required {
		markers = []string{"Required"}
	}
	if v.Secret {
		markers = append(markers, "Secret")
	}
	return strings.Join(markers, ", ")
}

func yesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// configDescription is the parameter's description as a single line, with the
// environment-variable fallbacks and deprecation notice the schema does know
// about folded in.
func configDescription(v ConfigVar) string {
	description := flatten(v.Description)
	if description == "" {
		description = "<TODO: describe this parameter and what a valid value looks like.>"
	}

	if len(v.EnvVars) > 0 && !mentionsAny(description, v.EnvVars) {
		quoted := make([]string, len(v.EnvVars))
		for i, name := range v.EnvVars {
			quoted[i] = "`" + name + "`"
		}
		description = strings.TrimSuffix(description, ".") + ". May also be set with the " +
			joinWithOr(quoted) + " environment " + pluralize("variable", len(quoted)) + "."
	}

	if v.Deprecated != "" {
		description = strings.TrimSuffix(description, ".") + ". Deprecated: " + flatten(v.Deprecated)
	}

	return description
}

func mentionsAny(haystack string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(haystack, needle) {
			return true
		}
	}
	return false
}

func joinWithOr(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	case 2:
		return items[0] + " or " + items[1]
	}
	return strings.Join(items[:len(items)-1], ", ") + ", or " + items[len(items)-1]
}

func pluralize(word string, n int) string {
	if n == 1 {
		return word
	}
	return word + "s"
}

// flatten collapses a description onto one line. Every hand-written list item
// and table row in this repo is a single line, and a table cell cannot contain
// a newline at all.
func flatten(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func escapeCell(s string) string {
	return strings.ReplaceAll(s, "|", `\|`)
}

func fence(language, body string) string {
	return fmt.Sprintf("```%s\n%s\n```\n", language, body)
}

func bashFence(command string) string {
	return fence("bash", command)
}
