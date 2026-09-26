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

	// Regression test: Terraform-bridged provider descriptions frequently
	// open with a markdown blockquote callout ("> Note ..." or "> If
	// ..."). Left in place, the literal ">" gets HTML-escaped once when
	// MetaDesc is written into the generated front matter and escaped
	// again when Hugo renders it into the page's meta/og description
	// tags and JSON-LD, producing a literal "&amp;gt;" in front of
	// visitors and in what answer engines ingest (observed live on
	// registry pages for random.RandomPassword, snowflake.Database,
	// tls.PrivateKey, and others). The marker must come out as plain
	// prose.
	t.Run("strips a leading blockquote marker", func(t *testing.T) {
		t.Parallel()
		got := summarizeForMetaDescription(
			"> If the managed resource supports a write-only attribute for the password " +
				"(first introduced in Terraform 1.11), then the ephemeral variant of this resource " +
				"should be used instead.",
		)
		require.NotContains(t, got, ">")
		require.True(t, strings.HasPrefix(got, "If the managed resource"))
	})

	t.Run("strips a blockquote marker mid-description without disturbing the rest", func(t *testing.T) {
		t.Parallel()
		got := summarizeForMetaDescription(
			"Provides a S3 bucket resource.\n\n" +
				"> This resource provides functionality for managing S3 general purpose buckets " +
				"in an AWS Partition. To manage Amazon S3 directory buckets in the Availability " +
				"Zone or Local Zone, see the aws.s3.DirectoryBucket resource.",
		)
		require.NotContains(t, got, ">")
		require.Contains(t, got, "Provides a S3 bucket resource.")
		require.Contains(t, got, "This resource provides functionality for managing S3 general purpose buckets")
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
