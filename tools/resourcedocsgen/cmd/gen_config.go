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

	"github.com/spf13/cobra"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/overview"
)

// GenConfigCmd generates the configuration parameters reference of a package's
// Overview page. Like gen-install it writes a snippet to stdout: the schema
// carries the name, requiredness, secrecy and description of every parameter,
// but not the environment-variable fallbacks or mutual-exclusivity rules the
// standard also asks for, so the output is a starting point the author edits.
func GenConfigCmd(client HTTPDoer) *cobra.Command {
	var (
		source schemaSource
		style  string
	)

	cmd := &cobra.Command{
		Use:   "gen-config",
		Short: "[EXPERIMENTAL] Generate the ## Configuration section of a package's Overview page",
		Long: "EXPERIMENTAL: output shape and flags may change.\n\n" +
			"Generate the configuration parameters reference of a package's Overview page " +
			"(docs/_index.md) from its schema's config block, as specified by docs/overview-page.md. " +
			"Writes markdown to stdout for you to paste into your provider repo's docs/_index.md.\n\n" +
			"Two things the standard requires cannot come from a schema -- environment-variable " +
			"fallbacks and mutually exclusive options -- so the output ends with a reminder to add " +
			"them by hand.",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := source.validate(); err != nil {
				return err
			}
			if style != string(overview.ConfigStyleTable) && style != string(overview.ConfigStyleList) {
				return fmt.Errorf("unknown --style %q: expected %q or %q",
					style, overview.ConfigStyleTable, overview.ConfigStyleList)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			spec, err := source.read(client)
			if err != nil {
				return err
			}

			cfg := overview.ReadSchemaConfig(spec)
			if len(cfg.Vars) == 0 {
				fmt.Fprintf(cmd.ErrOrStderr(),
					"%s declares no provider configuration; nothing to generate\n", spec.Name)
				return nil
			}

			_, err = fmt.Fprint(cmd.OutOrStdout(),
				overview.RenderConfiguration(cfg, overview.ConfigStyle(style)))
			return err
		},
	}

	source.addFlags(cmd)
	cmd.Flags().StringVar(&style, "style", string(overview.ConfigStyleList),
		fmt.Sprintf("The shape of the generated reference: %q or %q",
			overview.ConfigStyleList, overview.ConfigStyleTable))

	return cmd
}
