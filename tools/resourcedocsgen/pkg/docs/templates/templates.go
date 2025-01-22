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

// Package templates encapsulates **all** templates used by docs generation.
//
// All functions in this package are safe to call concurrently.
package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/golang/glog"
	"github.com/pgavlin/goldmark"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

//go:embed *.tmpl
var packagedTemplates embed.FS

type Template = func(io.Writer, any) error

func Index(wr io.Writer, a any) error { return templates.ExecuteTemplate(wr, "index.tmpl", a) }

func Resource(wr io.Writer, a any) error { return templates.ExecuteTemplate(wr, "resource.tmpl", a) }

func Function(wr io.Writer, a any) error { return templates.ExecuteTemplate(wr, "function.tmpl", a) }

func ParamSeparator(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "param_separator", a)
}

func PyParamSeparator(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "py_param_separator", a)
}

func TsFormalParam(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "ts_formal_param", a)
}

func GoFormalParam(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "go_formal_param", a)
}

func CSharpFormalParam(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "csharp_formal_param", a)
}

func JavaFormalParam(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "java_formal_param", a)
}

func PyFormalParam(wr io.Writer, a any) error {
	return templates.ExecuteTemplate(wr, "py_formal_param", a)
}

var templates *template.Template

func init() {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"lCSharp": func() language.Language { return language.CSharp },
		"lGo":     func() language.Language { return language.Go },
		"lJava":   func() language.Language { return language.Java },
		"lNodejs": func() language.Language { return language.Typescript },
		"lPython": func() language.Language { return language.Python },
		"lYAML":   func() language.Language { return language.YAML },

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
	templates = tmpl
}
