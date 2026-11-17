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

package templates

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func render(t *testing.T, src string) string {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, md.Convert([]byte(src), &buf))
	return buf.String()
}

func TestMarkdownLinks(t *testing.T) {
	t.Parallel()

	const ext = ` target="_blank" rel="noopener">`

	cases := []struct {
		name, src, want string
	}{
		{
			"bare URL", "See https://docs.aws.amazon.com/ec2/placement.html for details.",
			`<a href="https://docs.aws.amazon.com/ec2/placement.html"` + ext +
				`https://docs.aws.amazon.com/ec2/placement.html</a> for details.`,
		},
		{
			"trailing comma excluded", "Reconciling (https://google.aip.dev/128#reconciliation), if true",
			`(<a href="https://google.aip.dev/128#reconciliation"` + ext +
				`https://google.aip.dev/128#reconciliation</a>), if true`,
		},
		{
			"trailing semicolon excluded", "See https://www.tencentcloud.com/document/product/1145/62478; then",
			`<a href="https://www.tencentcloud.com/document/product/1145/62478"` + ext +
				`https://www.tencentcloud.com/document/product/1145/62478</a>; then`,
		},
		{
			"trailing quote excluded", `More info: https://git.k8s.io/community/api-conventions.md#types-kinds"`,
			`<a href="https://git.k8s.io/community/api-conventions.md#types-kinds"` + ext,
		},
		{"markdown link", "[docs](https://example.org/docs)", `<a href="https://example.org/docs"` + ext + `docs</a>`},
		{"relative link", "[aws](/registry/packages/aws/)", `<a href="/registry/packages/aws/">aws</a>`},
		{"email", "For example, alice@gmail.com or joe@example.com.", `For example, alice@gmail.com or joe@example.com.`},
		{"code span", "Use `https://example.org/a` here.", `<code>https://example.org/a</code>`},
		{
			"placeholder", "See https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone} now",
			`See https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone} now`,
		},
		{
			"path parameter", "See https://gitlab.com/api/v4/projects/:projectId for this",
			`See https://gitlab.com/api/v4/projects/:projectId for this`,
		},
		{"example host", "Point it at https://splunkserver.example.com", `Point it at https://splunkserver.example.com`},
		{
			"e.g. cue", "The base URL, e.g. https://api.eu.onelogin.com or https://api.us.onelogin.com.",
			`The base URL, e.g. https://api.eu.onelogin.com or https://api.us.onelogin.com.`,
		},
		{"for example cue", "The GitLab URL. For example, https://gitlab.com", `For example, https://gitlab.com`},
		{
			"cue on previous line", "Given an example URL:\nhttps://storage.googleapis.com/b/o#1",
			"Given an example URL:\nhttps://storage.googleapis.com/b/o#1",
		},
		{
			"cue in earlier sentence", "Use a value such as 30s. See https://pkg.go.dev/time#ParseDuration",
			`<a href="https://pkg.go.dev/time#ParseDuration"` + ext,
		},
		{
			"cue with reference", "For examples, see https://docs.datadoghq.com/tagging",
			`<a href="https://docs.datadoghq.com/tagging"` + ext,
		},
		{
			"parenthesized citation", "Supported formats are Glue and Apache Iceberg (https://iceberg.apache.org/), such as",
			`(<a href="https://iceberg.apache.org/"` + ext,
		},
		{
			"cue in another paragraph", "For example, foo.\n\nhttps://kubernetes.io/docs/concepts/",
			`<a href="https://kubernetes.io/docs/concepts/"` + ext,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert.Contains(t, render(t, c.src), c.want)
		})
	}
}
