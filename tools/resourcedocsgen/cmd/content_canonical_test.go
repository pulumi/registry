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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCommittedContentIsDelimiterCanonical is the repo-wide guard for broken Hugo
// shortcode delimiters. It runs in resourcedocsgen's Go test job — where the Go
// toolchain and the canonical pkg.NormalizeDocs already live — and asserts that
// every committed .md in the Hugo content tree is already in canonical delimiter
// form. Because it reuses docNeedsNormalizing (hence pkg.NormalizeDocs), this lint
// can never drift from the generator that produces the files, which is exactly the
// failure mode of the previous standalone JS regex.
//
// The content tree lives outside this Go module; the resourcedocsgen check-go job
// must include it in its sparse checkout (see extra-sparse-checkout in
// pull-request.yml). findRepoContentDir fails rather than skips under CI so a
// broken checkout surfaces loudly instead of silently disabling the guard.
func TestCommittedContentIsDelimiterCanonical(t *testing.T) {
	t.Parallel()

	contentDir := findRepoContentDir(t)

	var problems []string
	err := filepath.WalkDir(contentDir, func(p string, d fs.DirEntry, err error) error {
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
		if docNeedsNormalizing(content) {
			problems = append(problems, p)
		}
		return nil
	})
	require.NoError(t, err)

	if len(problems) > 0 {
		t.Fatalf("%d committed docs are not in canonical Hugo shortcode-delimiter form.\n"+
			"Fix with:\n"+
			"  (cd tools/resourcedocsgen && go run . sanitize-docs --in-place ../../themes/default/content)\n"+
			"Offending files:\n  %s",
			len(problems), strings.Join(problems, "\n  "))
	}
}

// findRepoContentDir locates the Hugo content tree by walking up from the test's
// working directory until it finds themes/default/content. Under CI (GITHUB_ACTIONS)
// its absence is a hard failure, so a mis-scoped checkout can never silently disable
// the guard; outside CI it skips, so running the module's tests in isolation from
// the repo does not spuriously fail.
func findRepoContentDir(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	require.NoError(t, err)

	for {
		candidate := filepath.Join(dir, "themes", "default", "content")
		if info, statErr := os.Stat(candidate); statErr == nil && info.IsDir() {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			if os.Getenv("GITHUB_ACTIONS") == "true" {
				t.Fatal("themes/default/content not found under CI; the resourcedocsgen " +
					"check-go job must sparse-checkout it (extra-sparse-checkout in pull-request.yml)")
			}
			t.Skip("repository content tree (themes/default/content) not found; skipping guard")
		}
		dir = parent
	}
}
