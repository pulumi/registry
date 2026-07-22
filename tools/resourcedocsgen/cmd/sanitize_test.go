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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDocsFileSanitizesShortcodeDelimiters(t *testing.T) {
	t.Parallel()

	// A docs file whose body contains a malformed chooser block, as has been
	// observed in upstream-generated (registry.opentofu.org-mirrored) provider docs.
	body := "---\ntitle: Example\n---\n\n" +
		"Examples:\n\n" +
		"{{ < chooser language \"typescript,python\" >}}\n" +
		"{{ % choosable language typescript %}}\n" +
		"```typescript\nconst x = 1;\n```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}\n"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(body))
		require.NoError(t, err)
	}))
	defer server.Close()

	got, err := readDocsFile(server.URL)
	require.NoError(t, err)

	assert.NotContains(t, string(got), "{{ <", "malformed opening delimiter must be repaired")
	assert.NotContains(t, string(got), "{{ %", "malformed opening delimiter must be repaired")
	assert.Contains(t, string(got), `{{< chooser language "typescript,python" >}}`)
	assert.Contains(t, string(got), "{{% choosable language typescript %}}")
	// Surrounding content must be preserved.
	assert.Contains(t, string(got), "```typescript\nconst x = 1;\n```")
}

func TestSanitizeShortcodeDelimiters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "opening chooser with stray space",
			input:    `{{ < chooser language "typescript,python" >}}`,
			expected: `{{< chooser language "typescript,python" >}}`,
		},
		{
			name:     "opening choosable with stray space",
			input:    `{{ % choosable language typescript %}}`,
			expected: `{{% choosable language typescript %}}`,
		},
		{
			name:     "closing chooser with stray space",
			input:    `{{ < /chooser >}}`,
			expected: `{{< /chooser >}}`,
		},
		{
			name:     "closing choosable with stray space",
			input:    `{{ % /choosable %}}`,
			expected: `{{% /choosable %}}`,
		},
		{
			name:     "multiple stray spaces collapse",
			input:    `{{   < chooser >}}`,
			expected: `{{< chooser >}}`,
		},
		{
			name:     "tab between braces and marker",
			input:    "{{\t% choosable language go %}}",
			expected: `{{% choosable language go %}}`,
		},
		{
			// A non-breaking space (U+00A0) renders identically to a plain space
			// but is not matched by [ \t]; Hugo still rejects it. This nbsp gap
			// motivated widening the whitespace class to \p{Zs}.
			name:     "non-breaking space (U+00A0) between braces and marker",
			input:    "{{\u00a0< chooser language \"go\" >}}",
			expected: `{{< chooser language "go" >}}`,
		},
		{
			name:     "unicode en-quad (U+2000) between braces and marker",
			input:    "{{\u2000% choosable language go %}}",
			expected: `{{% choosable language go %}}`,
		},
		{
			name:     "already well-formed is unchanged (idempotent)",
			input:    `{{< chooser language "typescript" >}}`,
			expected: `{{< chooser language "typescript" >}}`,
		},
		{
			name:     "legitimate trailing space before closing brace is preserved",
			input:    `{{< chooser language "typescript,python,go,csharp,java,yaml" >}}`,
			expected: `{{< chooser language "typescript,python,go,csharp,java,yaml" >}}`,
		},
		{
			name: "real megaport-style corrupted block is repaired, code preserved",
			input: "Examples:\n\n" +
				"{{ < chooser language \"typescript,python,go,csharp,java,yaml\" >}}\n" +
				"{{ % choosable language typescript %}}\n" +
				"```typescript\nconst x = 1; // {{ not a delimiter }}\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}\n",
			expected: "Examples:\n\n" +
				"{{< chooser language \"typescript,python,go,csharp,java,yaml\" >}}\n" +
				"{{% choosable language typescript %}}\n" +
				"```typescript\nconst x = 1; // {{ not a delimiter }}\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}\n",
		},
		{
			name:     "go template expression in prose is not a shortcode and untouched",
			input:    "Use the syntax {{ .Value }} in your template.",
			expected: "Use the syntax {{ .Value }} in your template.",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := string(sanitizeShortcodeDelimiters([]byte(tt.input)))
			assert.Equal(t, tt.expected, got)
		})
	}
}
