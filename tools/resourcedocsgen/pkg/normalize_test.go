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

package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeDocs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normalizes CRLF line endings to LF",
			input: "line one\r\nline two\r\n",
			want:  "line one\nline two\n",
		},
		{
			name:  "repairs a malformed delimiter outside a code fence",
			input: `{{ < chooser language "go" >}}`,
			want:  `{{< chooser language "go" >}}`,
		},
		{
			name:  "neutralizes a well-formed delimiter inside a code fence",
			input: "```go\n{{<ref \"x\">}}\n```",
			want:  "```go\n{{ <ref \"x\">}}\n```",
		},
		{
			name: "repairs outside and neutralizes inside in a single pass",
			input: "{{ < chooser language \"go\" >}}\n" +
				"{{% choosable language go %}}\n" +
				"```go\ncfg {{<ref>}} end\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}",
			want: "{{< chooser language \"go\" >}}\n" +
				"{{% choosable language go %}}\n" +
				"```go\ncfg {{ <ref>}} end\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}",
		},
		{
			name: "a stray upstream fence does not corrupt real chooser shortcodes",
			input: "```\n" + // stray, unpaired fence
				"{{< chooser language \"go\" >}}\n" +
				"{{% choosable language go %}}\n" +
				"```go\nx := 1\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}",
			want: "```\n" +
				"{{< chooser language \"go\" >}}\n" +
				"{{% choosable language go %}}\n" +
				"```go\nx := 1\n```\n" +
				"{{% /choosable %}}\n" +
				"{{< /chooser >}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := string(NormalizeDocs([]byte(tt.input)))
			assert.Equal(t, tt.want, got)
			// NormalizeDocs must be idempotent: the canonical form is a fixpoint.
			assert.Equal(t, tt.want, string(NormalizeDocs([]byte(got))), "not idempotent")
		})
	}
}

// TestSanitizeShortcodeDelimiters was moved here from the cmd package together
// with sanitizeShortcodeDelimiters, which now lives in pkg as part of the single
// canonical NormalizeDocs pipeline.
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
			input:    "{{ < chooser language \"go\" >}}",
			expected: `{{< chooser language "go" >}}`,
		},
		{
			name:     "unicode en-quad (U+2000) between braces and marker",
			input:    "{{ % choosable language go %}}",
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
