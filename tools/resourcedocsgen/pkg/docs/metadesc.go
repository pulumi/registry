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
	s := sanitizeDescription(comment)

	if loc := metaDescCutRegex.FindStringIndex(s); loc != nil {
		s = s[:loc[0]]
	}

	s = markdownLinkRegex.ReplaceAllString(s, "$1")
	s = htmlTagRegex.ReplaceAllString(s, "")
	s = strings.NewReplacer("`", "", "**", "", "*", "").Replace(s)
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
