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
	"github.com/spf13/cobra"

	"github.com/golang/glog"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/registry/tools/resourcedocsgen/cmd/docs"
)

var (
	logToStderr bool
	verbose     int
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "resourcedocsgen",
		Short: "Generate Pulumi resource docs",
		Long: "A tool to generate resource docs and package metadata for Pulumi provider and component packages. " +
			"This tool relies on a Pulumi package's schema spec. " +
			"This tool will not generate the schema.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logging.InitLogging(logToStderr, verbose, false)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			glog.Flush()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&logToStderr, "logtostderr", false, "Log to stderr instead of to files")
	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 0, "Enable verbose logging (e.g., v=3); anything >3 is very verbose")

	rootCmd.AddCommand(docs.ResourceDocsCmd())
	rootCmd.AddCommand(PackageMetadataCmd())
	rootCmd.AddCommand(CheckVersion())

	return rootCmd
}
