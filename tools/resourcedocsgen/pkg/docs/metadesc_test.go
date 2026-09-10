// Copyright 2026, Pulumi Corporation.
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

package docs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSummarizeForMetaDescription(t *testing.T) {
	t.Parallel()

	t.Run("returns empty for too-short descriptions", func(t *testing.T) {
		t.Parallel()
		require.Empty(t, summarizeForMetaDescription("Too short."))
	})

	t.Run("cuts at shortcode, heading, and fenced code", func(t *testing.T) {
		t.Parallel()
		got := summarizeForMetaDescription(
			"This resource manages a widget with configurable behavior across regions." +
				"\n\n{{% examples %}}\n## Example Usage\n```typescript\nconst x = 1;\n```",
		)
		require.Equal(t, "This resource manages a widget with configurable behavior across regions.", got)
	})

	// Regression test: some upstream (typically Terraform-bridged) schema
	// descriptions carry backslash-escaped quotes as literal text, e.g.
	// `\"true\"`. The generated front matter is rendered through
	// html/template, which HTML-escapes the quote character to `&#34;`
	// but leaves a pre-existing backslash untouched, producing the
	// invalid YAML escape sequence `\&#34;` and breaking the Hugo build
	// (observed live on pulumi/registry#12192 CI for the std.alltrue
	// function: "... true or \&#34;true\&#34;."). The summary must never
	// contain a backslash or a raw double quote.
	t.Run("strips backslashes and normalizes quotes so the YAML front matter never breaks", func(t *testing.T) {
		t.Parallel()
		got := summarizeForMetaDescription(
			`Returns true if all elements in a given collection are true or \"true\". ` +
				`It also returns true if the collection is empty.`,
		)
		require.NotContains(t, got, `\`)
		require.NotContains(t, got, `"`)
		require.Contains(t, got, "'true'")
	})

	t.Run("truncates long descriptions on a word boundary", func(t *testing.T) {
		t.Parallel()
		long := strings.Repeat("word ", 60)
		got := summarizeForMetaDescription(long)
		require.LessOrEqual(t, len(got), maxMetaDescLength+1) // +1 for the trailing period
		require.True(t, strings.HasSuffix(got, "."))
	})
}

func TestTruncateMetaDescription(t *testing.T) {
	t.Parallel()

	t.Run("returns short strings unchanged", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, "short", truncateMetaDescription("short", 155))
	})

	t.Run("breaks on a word boundary and ends with a period", func(t *testing.T) {
		t.Parallel()
		got := truncateMetaDescription("one two three four five", 15)
		require.Equal(t, "one two three.", got)
	})
}
