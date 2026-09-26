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
	"regexp"
	"strings"
)

// maxMetaDescLength is the target ceiling for a generated meta description,
// chosen to stay clear of typical search engine snippet truncation.
const maxMetaDescLength = 155

// minMetaDescSourceLength is the shortest a sanitized schema description can
// be before it's considered too thin to stand on its own as a meta
// description; below this we fall back to the generic template.
const minMetaDescSourceLength = 40

// metaDescCutRegex finds the start of structured content (Hugo/Pulumi
// shortcodes, markdown headings, fenced code blocks) that a schema
// description embeds after its opening summary. Everything from the first
// match onward is dropped before we build a meta description, since that
// content reads as a fragment out of context.
var metaDescCutRegex = regexp.MustCompile(`(?:\{\{[%<]|\n#|` + "```" + `)`)

// markdownLinkRegex matches `[text](url)` and keeps only the link text.
var markdownLinkRegex = regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)

// htmlTagRegex strips any remaining HTML tags.
var htmlTagRegex = regexp.MustCompile(`<[^>]+>`)

// blockquoteMarkerRegex matches a markdown blockquote marker at the start of
// a line (optionally indented, optionally followed by one space). Terraform
// -bridged provider descriptions frequently open with a "> Note ..." or
// "> If ..." callout; left in place, the literal ">" survives into the meta
// description and gets HTML-escaped when the value is later rendered, so it
// must come out as plain prose rather than markdown syntax.
var blockquoteMarkerRegex = regexp.MustCompile(`(?m)^[ \t]*>[ \t]?`)

// metaDescWhitespaceRegex collapses any run of whitespace (including
// newlines) down to a single space.
var metaDescWhitespaceRegex = regexp.MustCompile(`\s+`)

// summarizeForMetaDescription extracts a search-snippet-friendly summary
// from a resource's or function's raw schema description. Most descriptions
// on Terraform-bridged providers open with a plain-English summary sentence
// or two before diving into `{{% examples %}}` blocks, "## Argument
// Reference" sections, or fenced code samples; this pulls just that opening
// summary, strips remaining markdown/HTML, and trims it to fit comfortably
// within search engine snippet limits.
//
// It returns an empty string when the description doesn't yield a usable
// summary (empty, or too short after cleanup), signaling the caller should
// fall back to a generic templated description instead.
func summarizeForMetaDescription(comment string) string {
	s := SanitizeDescription(comment)

	if loc := metaDescCutRegex.FindStringIndex(s); loc != nil {
		s = s[:loc[0]]
	}

	s = blockquoteMarkerRegex.ReplaceAllString(s, "")
	s = markdownLinkRegex.ReplaceAllString(s, "$1")
	s = htmlTagRegex.ReplaceAllString(s, "")
	s = strings.NewReplacer("`", "", "**", "", "*", "").Replace(s)

	// Some upstream (typically Terraform-bridged) descriptions carry
	// backslash-escaped quotes as literal text (e.g. `\"true\"`). MetaDesc
	// is written into the generated front matter unescaped (see header.tmpl,
	// which renders it through the htmlSafe helper so html/template doesn't
	// double-encode it against Hugo's own escaping), so a literal double
	// quote or a stray backslash here would otherwise reach the YAML
	// document as-is and could break the double-quoted scalar. Drop stray
	// backslashes and normalize straight quotes to a form that's always
	// safe inside a double-quoted YAML scalar.
	s = strings.ReplaceAll(s, `\`, "")
	s = strings.ReplaceAll(s, `"`, "'")

	s = metaDescWhitespaceRegex.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)

	if len(s) < minMetaDescSourceLength {
		return ""
	}

	return truncateMetaDescription(s, maxMetaDescLength)
}

// truncateMetaDescription trims s to at most maxLen characters, breaking on
// a word boundary and ending on a clean sentence-like stop rather than
// mid-word.
func truncateMetaDescription(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	truncated := s[:maxLen]
	if idx := strings.LastIndex(truncated, " "); idx > 0 {
		truncated = truncated[:idx]
	}
	truncated = strings.TrimRight(truncated, ".,;:- ")
	return truncated + "."
}
