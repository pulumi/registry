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

func TestNeutralizeHugoShortcodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "percent opener inside a code fence is disarmed",
			input: "```java\nString.format(\"Bearer {{%s|token}}\", _id)\n```",
			want:  "```java\nString.format(\"Bearer {{ %s|token}}\", _id)\n```",
		},
		{
			name:  "angle opener inside a code fence is disarmed",
			input: "```go\ncfg {{<ref \"x\">}} end\n```",
			want:  "```go\ncfg {{ <ref \"x\">}} end\n```",
		},
		{
			name: "legitimate shortcodes outside fences are preserved",
			input: "{{< chooser language \"go\" >}}\n{{% choosable language go %}}\n" +
				"```go\nx := 1\n```\n{{% /choosable %}}\n{{< /chooser >}}",
			want: "{{< chooser language \"go\" >}}\n{{% choosable language go %}}\n" +
				"```go\nx := 1\n```\n{{% /choosable %}}\n{{< /chooser >}}",
		},
		{
			name:  "triple brace is not a Hugo opener and is left alone",
			input: "```python\nf\"Bearer {{{id}|token}}\"\n```",
			want:  "```python\nf\"Bearer {{{id}|token}}\"\n```",
		},
		{
			name:  "opener in prose outside a fence is left for Hugo to resolve",
			input: `See {{% notes %}} for details.`,
			want:  `See {{% notes %}} for details.`,
		},
		{
			name:  "idempotent: already-neutralized fence content is unchanged",
			input: "```java\nBearer {{ %s|token}}\n```",
			want:  "```java\nBearer {{ %s|token}}\n```",
		},
		{
			name: "shorter inner fence does not end the outer fence",
			input: "````markdown\n```\n{{%s|token}}\n```\nstill inside {{<ref>}}\n````\n" +
				"out {{% notes %}}",
			want: "````markdown\n```\n{{ %s|token}}\n```\nstill inside {{ <ref>}}\n````\n" +
				"out {{% notes %}}",
		},
		{
			name:  "tilde fence interior is disarmed",
			input: "~~~go\n{{<ref>}}\n~~~",
			want:  "~~~go\n{{ <ref>}}\n~~~",
		},
		{
			name:  "shorter and sibling fence lines inside a longer fence are disarmed",
			input: "````markdown\n``` {{%s|token}}\n~~~ {{<ref>}}\n````",
			want:  "````markdown\n``` {{ %s|token}}\n~~~ {{ <ref>}}\n````",
		},
		{
			name:  "braces ending an HTML code span do not form a closing shortcode with the tag",
			input: "<p>A placeholder is <code>{{</code>, optional whitespace, and <code>}}</code>.</p>",
			want:  "<p>A placeholder is <code>&#123;{</code>, optional whitespace, and <code>}}</code>.</p>",
		},
		{
			name:  "percent opener inside an HTML code span is disarmed",
			input: `<dd>Use <code>Bearer {{%s|token}}</code> or <code class="x">{{<ref>}}</code>.</dd>`,
			want:  `<dd>Use <code>Bearer &#123;{%s|token}}</code> or <code class="x">&#123;{<ref>}}</code>.</dd>`,
		},
		{
			name:  "template placeholders in an HTML code span are left alone",
			input: "<code>{{budget.name}}</code> and <code>{{ .Value }}</code>",
			want:  "<code>{{budget.name}}</code> and <code>{{ .Value }}</code>",
		},
		{
			name:  "shortcodes next to but outside an HTML code span are preserved",
			input: "{{% notes %}} <code>x</code> {{< /chooser >}}",
			want:  "{{% notes %}} <code>x</code> {{< /chooser >}}",
		},
		{
			name:  "idempotent: already-neutralized HTML code span is unchanged",
			input: "<code>&#123;{</code> <code>&#123;{%s}}</code>",
			want:  "<code>&#123;{</code> <code>&#123;{%s}}</code>",
		},
		{
			name:  "HTML code span inside a fence gets the fence treatment only",
			input: "```html\n<code>{{%s}}</code>\n```",
			want:  "```html\n<code>{{ %s}}</code>\n```",
		},
		{
			name:  "four-space indented delimiter inside a fence does not close it",
			input: "```markdown\n    ```\n{{%s|token}}\n```",
			want:  "```markdown\n    ```\n{{ %s|token}}\n```",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, string(neutralizeHugoShortcodes([]byte(tt.input))))
		})
	}
}

// TestNeutralizeHugoShortcodesResyncsOnMalformedUpstreamFence reproduces the
// megaport@1.11.1 failure: a stray, unpaired code fence in the upstream doc
// desynced the fence tracker, so real chooser/choosable wrapping shortcodes below
// it were wrongly neutralized to "{{ <" — aborting the Hugo build. They must be
// preserved (and re-sync the tracker) while genuine in-fence delimiters after them
// are still disarmed.
func TestNeutralizeHugoShortcodesResyncsOnMalformedUpstreamFence(t *testing.T) {
	t.Parallel()

	input := "```java\nx := 1\n```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}\n" +
		"```\n" + // stray, unpaired fence from the malformed upstream doc
		"\n" +
		"Examples:\n" +
		"{{< chooser language \"typescript,python\" >}}\n" + // real: must be preserved
		"{{% choosable language typescript %}}\n" + // real: must be preserved
		"```typescript\n" +
		"const t = \"Bearer {{%s|token}}\";\n" + // genuine garbage: must be neutralized
		"```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}"

	want := "```java\nx := 1\n```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}\n" +
		"```\n" +
		"\n" +
		"Examples:\n" +
		"{{< chooser language \"typescript,python\" >}}\n" +
		"{{% choosable language typescript %}}\n" +
		"```typescript\n" +
		"const t = \"Bearer {{ %s|token}}\";\n" +
		"```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}"

	assert.Equal(t, want, string(neutralizeHugoShortcodes([]byte(input))))
}
