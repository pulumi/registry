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
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/spf13/cobra"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/overview"
)

// noSDKLanguages is the --languages value for a package that publishes no
// per-language SDKs. About a hundred registry packages are in that category.
const noSDKLanguages = "none"

// GenInstallCmd generates the `## Installation` section of a package's Overview
// page. It writes a snippet to stdout for the author to paste into their own
// repo's docs/_index.md rather than generating the page: that page is authored
// in the provider repo and fetched from the release tag by `metadata
// from-github`, so this tool is advisory by design.
func GenInstallCmd(client HTTPDoer) *cobra.Command {
	var (
		source    schemaSource
		version   string
		languages string
	)

	cmd := &cobra.Command{
		Use:   "gen-install",
		Short: "[EXPERIMENTAL] Generate the ## Installation section of a package's Overview page",
		Long: "EXPERIMENTAL: output shape and flags may change.\n\n" +
			"Generate the ## Installation section of a package's Overview page (docs/_index.md) " +
			"from its schema, as specified by docs/overview-page.md. Writes markdown to stdout for " +
			"you to paste into your provider repo's docs/_index.md.\n\n" +
			"A schema cannot say which SDKs a package actually publishes -- the language blob is " +
			"written by the code generator, not by the release pipeline -- so use --languages when " +
			"the guess is wrong, and --languages none for a package that publishes no SDKs at all.",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return source.validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			spec, err := source.read(client)
			if err != nil {
				return err
			}

			langs, err := resolveLanguages(languages, spec)
			if err != nil {
				return err
			}

			if version == "" {
				version = spec.Version
			}

			plan := overview.DeriveInstalls(spec, version, langs)
			for _, warning := range plan.Warnings {
				fmt.Fprintln(cmd.ErrOrStderr(), "warning: "+warning)
			}

			_, err = fmt.Fprint(cmd.OutOrStdout(), overview.RenderInstallation(plan))
			return err
		},
	}

	source.addFlags(cmd)
	cmd.Flags().StringVar(&version, "version", "",
		"The package version, used for the Maven and Gradle coordinates when the schema omits one")
	cmd.Flags().StringVar(&languages, "languages", "",
		"Comma-separated languages this package publishes SDKs for ("+
			strings.Join(overview.SDKLanguages, ", ")+"), or \""+noSDKLanguages+
			"\" for a package that publishes none. Defaults to the languages the schema declares.")

	return cmd
}

// resolveLanguages turns the --languages flag into the set of SDK languages to
// render, defaulting to what the schema declares when the flag is absent.
func resolveLanguages(flag string, spec *schema.PackageSpec) ([]string, error) {
	if strings.TrimSpace(flag) == "" {
		return overview.DefaultLanguages(spec), nil
	}
	if strings.EqualFold(strings.TrimSpace(flag), noSDKLanguages) {
		return nil, nil
	}

	valid := map[string]bool{}
	for _, lang := range overview.SDKLanguages {
		valid[lang] = true
	}

	var langs []string
	for _, lang := range strings.Split(flag, ",") {
		lang = strings.ToLower(strings.TrimSpace(lang))
		if lang == "" {
			continue
		}
		// yaml and hcl always appear -- they install with `pulumi package add`
		// rather than an SDK -- so accepting them here would be misleading.
		if !valid[lang] {
			return nil, fmt.Errorf("unknown language %q: expected one of %s, or %q",
				lang, strings.Join(overview.SDKLanguages, ", "), noSDKLanguages)
		}
		langs = append(langs, lang)
	}
	return langs, nil
}
