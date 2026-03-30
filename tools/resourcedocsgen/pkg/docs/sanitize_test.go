// Copyright 2024, Pulumi Corporation.
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripLangSpans(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "span with all language attributes",
			input: `<span pulumi-lang-nodejs="` + "`mongodbatlas.Cluster`" + `"` +
				` pulumi-lang-dotnet="` + "`mongodbatlas.Cluster`" + `"` +
				` pulumi-lang-go="` + "`Cluster`" + `"` +
				` pulumi-lang-python="` + "`Cluster`" + `"` +
				` pulumi-lang-yaml="` + "`mongodbatlas.Cluster`" + `"` +
				` pulumi-lang-java="` + "`mongodbatlas.Cluster`" + `">` +
				"`mongodbatlas.Cluster`</span> provides a resource.",
			expected: "`mongodbatlas.Cluster` provides a resource.",
		},
		{
			name: "span with spaces as content",
			input: `You may find<span pulumi-lang-nodejs=" groupId "` +
				` pulumi-lang-python=" group_id "> group_id </span>in the docs.`,
			expected: "You may find group_id in the docs.",
		},
		{
			name: "multiple spans in one string",
			input: "Use " +
				`<span pulumi-lang-nodejs="` + "`Foo`" + `" pulumi-lang-python="` + "`Foo`" + `">` +
				"`Foo`</span> and " +
				`<span pulumi-lang-nodejs="` + "`Bar`" + `" pulumi-lang-python="` + "`Bar`" + `">` +
				"`Bar`</span>.",
			expected: "Use `Foo` and `Bar`.",
		},
		{
			name: "span with nested double quotes in attributes",
			input: `values: <span pulumi-lang-nodejs=""customS2s""` +
				` pulumi-lang-python=""custom_s2s"">"custom_s2s"</span>.`,
			expected: `values: "custom_s2s".`,
		},
		{
			name:     "no spans leaves text unchanged",
			input:    "Just a normal description.",
			expected: "Just a normal description.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := stripLangSpans(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFixHTMLBullets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "mongodbatlas blockquote with html bullets",
			input: `> **IMPORTANT:**
<br> &#8226; First item about clusters.
<br> &#8226; Second item about billing.`,
			expected: `> **IMPORTANT:**

- First item about clusters.

- Second item about billing.`,
		},
		{
			name: "no space after entity",
			input: `> **NOTE:**
<br>&#8226;Item one.
<br>&#8226;Item two.`,
			expected: `> **NOTE:**

- Item one.

- Item two.`,
		},
		{
			name:     "no html bullets leaves text unchanged",
			input:    "> **IMPORTANT:** This is a normal blockquote.",
			expected: "> **IMPORTANT:** This is a normal blockquote.",
		},
		{
			name: "html bullets with code block between them",
			input: `> **IMPORTANT:**
<br> &#8226; Upgrade your config from:
` + "```" + `
old_config = "value"
` + "```" + `
To:
` + "```" + `
new_config = "value"
` + "```" + `
<br> &#8226; Changes can affect costs.`,
			expected: `> **IMPORTANT:**

- Upgrade your config from:
` + "```" + `
old_config = "value"
` + "```" + `
To:
` + "```" + `
new_config = "value"
` + "```" + `

- Changes can affect costs.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := fixHTMLBullets(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFixTableSeparators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "mongodbatlas malformed table",
			input: `| Host alerts         | Replica set alerts  |  Sharded cluster alerts |
|:----------           |:-------------       |:------                 |
| TYPE_NAME            | REPLICA_SET_NAME    | CLUSTER_NAME           |`,
			expected: `| Host alerts         | Replica set alerts  |  Sharded cluster alerts |
|:----------|:-------------|:------|
| TYPE_NAME            | REPLICA_SET_NAME    | CLUSTER_NAME           |`,
		},
		{
			name: "already valid table separator",
			input: `| A | B |
|:---|:---|
| 1 | 2 |`,
			expected: `| A | B |
|:---|:---|
| 1 | 2 |`,
		},
		{
			name: "separator with alignment markers",
			input: `| Left | Center | Right |
|:---    |:---:    |---:    |
| a | b | c |`,
			expected: `| Left | Center | Right |
|:---|:---:|---:|
| a | b | c |`,
		},
		{
			name:     "not a table separator row",
			input:    "| this is | regular | text |",
			expected: "| this is | regular | text |",
		},
		{
			name:     "no tables leaves text unchanged",
			input:    "Just a normal description.\nWith multiple lines.",
			expected: "Just a normal description.\nWith multiple lines.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := fixTableSeparators(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIsTableSeparatorRow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected bool
	}{
		{"|:---|:---|", true},
		{"|---|---|", true},
		{"|:---:|:---:|", true},
		{"|---:|:---|", true},
		{"|:------           |:------      |", true},
		{"| --- | --- |", true},
		{"|:----------           |:-------------       |:------                 |", true},
		{"|:---|", true},

		// Not separator rows:
		{"| text | text |", false},
		{"", false},
		{"||", false},
		{"not a table", false},
		{"|:--a--|:---|", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, isTableSeparatorRow(tt.input))
		})
	}
}

func TestStripTerraformArgSections(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "strips single argument reference section",
			input: `Some description.

## Argument Reference

The following arguments are supported:

* ` + "`name`" + ` - (Required) Name of the thing.

* ` + "`config`" + ` - (Optional) Config block.

## Import

You can import using the id.`,
			expected: `Some description.


## Import

You can import using the id.`,
		},
		{
			name: "strips multiple argument reference sections",
			input: `Description.

## Argument Reference - V1

* ` + "`name`" + ` - Name.

## Argument Reference - V2

* ` + "`name`" + ` - Name.

## Import

Import docs.`,
			expected: `Description.



## Import

Import docs.`,
		},
		{
			name: "strips attributes reference section",
			input: `Description.

## Attributes Reference

The following attributes are exported:

* ` + "`id`" + ` - The ID.

## Import

Import docs.`,
			expected: `Description.


## Import

Import docs.`,
		},
		{
			name: "no terraform sections leaves text unchanged",
			input: `Description.

## Import

Import docs.`,
			expected: `Description.

## Import

Import docs.`,
		},
		{
			name: "strips section at end of string",
			input: `Description.

## Argument Reference

* ` + "`name`" + ` - Name.`,
			expected: "Description.\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := stripTerraformArgSections(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSanitizeDescription(t *testing.T) {
	t.Parallel()

	// Integration test: all fixes applied together.
	input := `> **IMPORTANT:**
<br> &#8226; This table matters:

| Col A | Col B |
|:---    |:---    |
| 1 | 2 |

<br> &#8226; And that is all.`

	expected := `> **IMPORTANT:**

- This table matters:

| Col A | Col B |
|:---|:---|
| 1 | 2 |


- And that is all.`

	actual := sanitizeDescription(input)
	assert.Equal(t, expected, actual)
}
