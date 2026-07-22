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
// megaport@1.11.1 failure: the upstream (registry.opentofu.org-mirrored) docs
// contained a stray, unpaired code fence, which desynced the fence tracker so
// that real `{{< chooser >}}` / `{{% choosable %}}` wrapping shortcodes further
// down the page were wrongly treated as fence interior and neutralized to
// "{{ <" / "{{ %". That both aborts the Hugo build (the paired closing tag is
// still recognized) and trips the shortcode-delimiter lint.
//
// A registry fence-wrapping shortcode never legitimately appears inside a code
// example, so it must never be neutralized regardless of the (possibly
// corrupt) fence state, and encountering one re-syncs the tracker so genuine
// code-example delimiters after it are still disarmed.
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
