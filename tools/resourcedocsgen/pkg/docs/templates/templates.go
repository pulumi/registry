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

package templates

import (
	"bytes"
	"embed"
	"html/template"
	"strings"

	"github.com/golang/glog"
	"github.com/pgavlin/goldmark"
)

//go:embed *.tmpl
var packagedTemplates embed.FS

var Templates *template.Template

func init() {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"htmlSafe": func(html string) template.HTML {
			// Markdown fragments in the templates need to be rendered as-is, so that html/template package doesn't try to
			// inject data into it, which will most certainly fail.
			//nolint:gosec
			return template.HTML(html)
		},
		"markdownify": func(html string) template.HTML {
			// Convert a string of Markdown into HTML.
			var buf bytes.Buffer
			if err := goldmark.Convert([]byte(html), &buf); err != nil {
				glog.Fatalf("rendering Markdown: %v", err)
			}
			rendered := buf.String()

			// Trim surrounding <p></p> tags.
			result := strings.TrimSpace(rendered)
			result = strings.TrimPrefix(result, "<p>")
			result = strings.TrimSuffix(result, "</p>")

			// If there are still <p> tags, there are multiple paragraphs, in which case use the original rendered string
			// (untrimmed).
			if strings.Contains(result, "<p>") {
				result = rendered
			}

			//nolint:gosec
			return template.HTML(result)
		},
	}).ParseFS(packagedTemplates, "*")
	if err != nil {
		glog.Fatal("initializing templates: %v", err)
	}
	Templates = tmpl
}
