// Copyright 2016-2024, Pulumi Corporation.
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

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// normalizeDocs applies the content repairs resourcedocsgen performs on every
// docs file it writes: end-of-line normalization and malformed Hugo shortcode
// delimiter repair. It is lenient and idempotent and never fails on content
// shape, so callers on `set -e` paths can apply it unconditionally.
func normalizeDocs(content []byte) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	return sanitizeShortcodeDelimiters(content)
}

// SanitizeDocsCmd exposes normalizeDocs to callers that obtain docs by means
// other than resourcedocsgen's own fetch. The versioned-docs build
// (scripts/generate-versioned-docs.sh) downloads a package's _index.md with a
// raw curl straight into the Hugo content tree, so it never passes through
// readDocsFile's repair. Piping that file through this command applies the
// identical repair, so a malformed delimiter from any docs source is fixed
// once, in one place, rather than re-implemented per consumer.
func SanitizeDocsCmd() *cobra.Command {
	var inPlace bool
	cmd := &cobra.Command{
		Use:   "sanitize-docs [file]",
		Short: "Repair malformed Hugo shortcode delimiters in a docs file",
		Long: "Read a docs file (or stdin) and repair malformed Hugo shortcode opening " +
			"delimiters (\"{{ <\" / \"{{ %\"), which otherwise abort the Hugo site build. " +
			"Writes the result to stdout, or back to the file with --in-place.",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if inPlace && len(args) != 1 {
				return fmt.Errorf("--in-place requires a file argument")
			}

			var (
				input []byte
				err   error
			)
			if len(args) == 1 {
				input, err = os.ReadFile(args[0])
			} else {
				input, err = io.ReadAll(cmd.InOrStdin())
			}
			if err != nil {
				return err
			}

			output := normalizeDocs(input)

			if inPlace {
				return os.WriteFile(args[0], output, 0o600)
			}
			_, err = cmd.OutOrStdout().Write(output)
			return err
		},
	}
	cmd.Flags().BoolVar(&inPlace, "in-place", false, "Rewrite the file in place instead of writing to stdout")
	return cmd
}
