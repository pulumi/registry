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

// docNeedsNormalizing reports whether a docs file's shortcode delimiters are not
// in canonical form (and would therefore be changed by pkg.NormalizeDocs).
//
// It is deliberately tolerant of line endings: a pre-existing "\r\n" in
// hand-authored content is not a broken-delimiter problem — Hugo handles CRLF
// fine — so it compares NormalizeDocs's output against the input with only line
// endings normalized. A file is flagged if and only if a shortcode delimiter is
// malformed (stray whitespace outside a fence) or un-neutralized (a live
// delimiter inside a code fence).
func docNeedsNormalizing(content []byte) bool {
	lineEndingsOnly := bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	return !bytes.Equal(pkg.NormalizeDocs(content), lineEndingsOnly)
}

// SanitizeDocsCmd exposes pkg.NormalizeDocs to callers that obtain docs by means
// other than resourcedocsgen's own fetch. The versioned-docs build
// (scripts/generate-versioned-docs.sh) downloads a package's _index.md with a
// raw curl straight into the Hugo content tree, so it never passes through
// readDocsFile's normalization. Piping that file through this command applies the
// identical canonical normalization, so broken delimiters from any docs source
// are handled once, in one place, rather than re-implemented per consumer.
//
// The argument may be a single file or a directory (which is walked for .md
// files). Without a flag it writes the normalized single file/stdin to stdout;
// --in-place rewrites files; --check reports any file that is not already
// canonical and exits non-zero (the delimiter linter, sharing the generator's
// exact logic so the two can never disagree).
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

			output := pkg.NormalizeDocs(input)
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

// sanitizeTree walks root for .md files. With check, it collects every file that
// is not already canonical and returns an error listing them; otherwise it
// rewrites each non-canonical file in place.
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
		if check {
			if docNeedsNormalizing(content) {
				problems = append(problems, p)
			}
			return nil
		}
		if output := pkg.NormalizeDocs(content); !bytes.Equal(output, content) {
			return os.WriteFile(p, output, 0o600)
		}
		return nil
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
