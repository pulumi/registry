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

package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/pulumi/registry/tools/resourcedocsgen/cmd/docs"
)

var verbose int

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "resourcedocsgen",
		Short: "Generate Pulumi resource docs",
		Long: "A tool to generate resource docs and package metadata for Pulumi provider and component packages. " +
			"This tool relies on a Pulumi package's schema spec. " +
			"This tool will not generate the schema.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			level := slog.LevelInfo
			if verbose > 0 {
				level = slog.LevelDebug
			}
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})))
		},
	}

	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 0,
		"Enable verbose logging (any value > 0 enables debug-level logs)")

	rootCmd.AddCommand(docs.ResourceDocsCmd())
	rootCmd.AddCommand(PackageMetadataCmd())
	rootCmd.AddCommand(CheckVersion())

	return rootCmd
}
