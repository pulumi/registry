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
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/spf13/cobra"
)

// docNeedsNormalizing reports whether pkg.NormalizeDocs would change a doc's
// shortcode delimiters. It compares against the input with only line endings
// normalized, so a pre-existing CRLF is not by itself flagged: mktutorial-generated
// how-to-guides carry CRLF and do not pass through NormalizeDocs, and Hugo handles
// CRLF fine — the delimiters are what break the build. --check and --in-place both
// use this predicate so they agree on what is non-canonical.
func docNeedsNormalizing(content []byte) bool {
	lineEndingsOnly := bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	return !bytes.Equal(pkg.NormalizeDocs(content), lineEndingsOnly)
}

// SanitizeDocsCmd exposes pkg.NormalizeDocs — the single canonical delimiter
// normalization — to callers that obtain docs by other means than resourcedocsgen's
// own fetch (e.g. scripts/generate-versioned-docs.sh curls _index.md straight into
// the content tree). The argument may be a file or a directory of .md files. With
// no flag it writes the normalized file/stdin to stdout; --in-place rewrites files;
// --check reports any non-canonical file and exits non-zero without writing.
func SanitizeDocsCmd() *cobra.Command {
	var inPlace, check bool
	cmd := &cobra.Command{
		Use:   "sanitize-docs [path]",
		Short: "Normalize Hugo shortcode delimiters in docs files",
		Long: "Apply the canonical docs normalization (repair malformed \"{{ <\" / \"{{ %\" " +
			"delimiters outside code fences, neutralize live delimiters inside them), which " +
			"otherwise abort the Hugo site build. The path may be a file or a directory of " +
			".md files. With no flag, writes the normalized file (or stdin) to stdout; " +
			"--in-place rewrites in place; --check exits non-zero if anything is not already " +
			"canonical without writing.",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if inPlace && check {
				return fmt.Errorf("--in-place and --check are mutually exclusive")
			}

			var path string
			if len(args) == 1 {
				path = args[0]
			}

			if inPlace && path == "" {
				return fmt.Errorf("--in-place requires a file or directory argument")
			}

			if path != "" {
				info, err := os.Stat(path)
				if err != nil {
					return err
				}
				if info.IsDir() {
					if !inPlace && !check {
						return fmt.Errorf("a directory argument requires --in-place or --check")
					}
					return sanitizeTree(path, check)
				}
			}

			var (
				input []byte
				err   error
			)
			if path != "" {
				input, err = os.ReadFile(path)
			} else {
				input, err = io.ReadAll(cmd.InOrStdin())
			}
			if err != nil {
				return err
			}

			// Run NormalizeDocs once; --check and the write paths use the same result.
			output := pkg.NormalizeDocs(input)

			if check {
				if docNeedsNormalizing(input) {
					src := path
					if src == "" {
						src = "<stdin>"
					}
					return fmt.Errorf("%s: malformed Hugo shortcode delimiter(s); "+
						"run `resourcedocsgen sanitize-docs --in-place %s`", src, src)
				}
				return nil
			}

			if inPlace {
				return os.WriteFile(path, output, 0o600)
			}
			_, err = cmd.OutOrStdout().Write(output)
			return err
		},
	}
	cmd.Flags().BoolVar(&inPlace, "in-place", false, "Rewrite file(s) in place instead of writing to stdout")
	cmd.Flags().BoolVar(&check, "check", false,
		"Report files whose delimiters are not canonical and exit non-zero, without writing")
	return cmd
}

// sanitizeTree walks root for .md files. With check it collects every file that is
// not already canonical and returns an error listing them; otherwise it rewrites
// each non-canonical file in place. Both branches gate on docNeedsNormalizing, so a
// canonical-but-CRLF file is neither reported nor rewritten.
func sanitizeTree(root string, check bool) error {
	var problems []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".md") {
			return nil
		}
		content, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if !docNeedsNormalizing(content) {
			return nil
		}
		if check {
			problems = append(problems, p)
			return nil
		}
		return os.WriteFile(p, pkg.NormalizeDocs(content), 0o600)
	})
	if err != nil {
		return err
	}
	if check && len(problems) > 0 {
		return fmt.Errorf("malformed Hugo shortcode delimiter(s) in %d file(s) "+
			"(run `resourcedocsgen sanitize-docs --in-place %s` to fix):\n  %s",
			len(problems), root, strings.Join(problems, "\n  "))
	}
	return nil
}
