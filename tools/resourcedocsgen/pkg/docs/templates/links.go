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
	"net/url"
	"regexp"
	"strings"

	"github.com/pgavlin/goldmark"
	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/extension"
	"github.com/pgavlin/goldmark/parser"
	"github.com/pgavlin/goldmark/text"
	"github.com/pgavlin/goldmark/util"
)

// linkifyURLRegexp is the fork's default URL pattern, except that a URL can't end in punctuation. The default only
// trims a trailing "." or unbalanced ")", so "see https://example.com/a, and" would link "https://example.com/a,".
// Placeholder characters ({}, [], <>) stay in the pattern so that a templated URL is matched whole and then rejected
// by isExampleURL, rather than truncated at the placeholder and linked.
var linkifyURLRegexp = regexp.MustCompile(
	`^(?:http|https|ftp)://(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{2,256}\.[a-z]+` +
		`(?:[/#?](?:[-a-zA-Z0-9@:%_+.~#$!?&/=();,'"{}\[\]<>^*]*[-a-zA-Z0-9/#=+~%@)}\]>])?)?`)

// neverMatch disables linkify's email detection. Addresses in provider schemas are almost always placeholders
// ("joe@example.com"), so none of them become mailto: links.
var neverMatch = regexp.MustCompile(`a^`)

// exampleLinkify wraps goldmark's linkify parser and leaves URLs that are examples as plain text. Schemas often show a
// sample value ("For example, https://myvault.vault.azure.net/keys/mykey"), and linking it sends readers to a page
// that doesn't exist.
type exampleLinkify struct {
	inner parser.InlineParser
}

func (p *exampleLinkify) Trigger() []byte { return p.inner.Trigger() }

func (p *exampleLinkify) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, segment := block.PeekLine()
	n := p.inner.Parse(parent, block, pc)
	link, ok := n.(*ast.AutoLink)
	if !ok {
		return n
	}
	source := block.Source()
	start := segment.Start
	// The inner parser consumes one leading trigger character and appends it to parent as text.
	if strings.IndexByte(" *_~(", line[0]) >= 0 {
		start++
	}
	label := link.Label(source)
	paraStart := bytes.LastIndex(source[:start], []byte("\n\n")) + 1
	if isExampleURL(string(link.URL(source)), string(source[paraStart:start])) {
		return ast.NewTextSegment(text.NewSegment(start, start+len(label)))
	}
	return link
}

func alternation(words ...string) string { return `(` + strings.Join(words, `|`) + `)` }

var (
	placeholderRegexp = regexp.MustCompile(`[{}\[\]<>*]|\$\{|/:[A-Za-z]|\.\.\.|(?i)xxx`)
	exampleHostRegexp = regexp.MustCompile(`(?i)` + strings.Join([]string{
		`(^|\.)(example|test)\.(com|org|net)$`,
		`\.(example|test|invalid|localhost)$`,
		`^localhost$`, `^127\.`, `^0\.0\.0\.0$`, `^10\.`, `^192\.168\.`,
		`contoso`, `fabrikam`, `(^|[.-])your`,
		`(^|\.)my` + alternation("account", "company", "org", "domain", "site", "app", "server", "host", "instance",
			"bucket", "vault", "tenant"),
		`^` + alternation("hostname", "host", "server", "domain", "bucketname", "accountname", "servername") + `\.`,
	}, `|`))
	exampleCueRegexp = regexp.MustCompile(`(?i)\b` + alternation(`e\.g\b`, `eg[.:,]`, `examples?\b`, "for instance",
		"such as", "of the form", "looks like", `like\b`))
	// referenceCueRegexp marks a URL as a page to read even when an example cue precedes it, as in
	// "For examples, see https://...".
	referenceCueRegexp = regexp.MustCompile(`(?i)\b` + alternation("see", "here", "refer", "described", "defined",
		"documentation", "docs", "more info", "information", "details", "learn more", "visit", "check", "available at",
		"listed at", "found at", "specified in", "according to", "follows?", "restrictions", "conventions",
		"specification") + `\b`)
	// sentenceEndRegexp finds sentence boundaries, so a cue in an earlier sentence doesn't apply to this URL.
	// A period ending "e.g." or "i.e." is not a boundary.
	sentenceEndRegexp = regexp.MustCompile(`[;!?]\s|[^.]\.\s`)
	abbrevRegexp      = regexp.MustCompile(`(?i)\b(e\.g|i\.e|eg|ie|etc|ex|vs)\.\s$`)
)

// isExampleURL reports whether a bare URL is a sample value rather than a page to visit. before is the text that
// precedes the URL in its paragraph.
func isExampleURL(rawURL, before string) bool {
	if placeholderRegexp.MatchString(rawURL) {
		return true
	}
	if u, err := url.Parse(rawURL); err == nil && exampleHostRegexp.MatchString(u.Hostname()) {
		return true
	}
	sentence := before
	for _, m := range sentenceEndRegexp.FindAllStringIndex(before, -1) {
		if !abbrevRegexp.MatchString(before[:m[1]]) {
			sentence = before[m[1]:]
		}
	}
	// A URL in parentheses right after prose, as in "key rotation (https://...)", cites a source.
	if strings.HasSuffix(strings.TrimSpace(sentence), "(") {
		return false
	}
	return exampleCueRegexp.MatchString(sentence) && !referenceCueRegexp.MatchString(sentence)
}

// externalLinkTargets opens links to absolute http(s) URLs in a new tab, so following a reference to a vendor's docs
// doesn't navigate away from the API docs.
type externalLinkTargets struct{}

func (externalLinkTargets) Transform(doc *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		var dest []byte
		switch n := n.(type) {
		case *ast.Link:
			dest = n.Destination
		case *ast.AutoLink:
			if n.AutoLinkType != ast.AutoLinkURL {
				return ast.WalkContinue, nil
			}
			dest = n.URL(source)
		default:
			return ast.WalkContinue, nil
		}
		lower := bytes.ToLower(dest)
		if bytes.HasPrefix(lower, []byte("http://")) || bytes.HasPrefix(lower, []byte("https://")) {
			n.SetAttributeString("target", []byte("_blank"))
			n.SetAttributeString("rel", []byte("noopener"))
		}
		return ast.WalkContinue, nil
	})
}

// newMarkdown returns the goldmark instance used for property, type, and enum descriptions.
// Tables are enabled because the default goldmark.Convert() renders upstream Markdown tables as plain text.
func newMarkdown() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(
			parser.WithInlineParsers(util.Prioritized(&exampleLinkify{inner: extension.NewLinkifyParser(
				extension.WithLinkifyURLRegexp(linkifyURLRegexp),
				extension.WithLinkifyEmailRegexp(neverMatch),
			)}, 999)),
			parser.WithASTTransformers(util.Prioritized(externalLinkTargets{}, 999)),
		),
	)
}
