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

// linkifyURLRegexp is the fork's default URL pattern, restricted to http(s) and changed so that a URL can't end in
// punctuation. The default only trims a trailing "." or unbalanced ")", so "see https://example.com/a, and" would
// link "https://example.com/a,". "{}" and "[]" stay in the pattern so that a templated URL is matched whole and then
// rejected by shouldLink, rather than truncated at the placeholder and linked. "<" and ">" are left out because
// descriptions converted from HTML often put a tag such as <br> directly after a URL.
var linkifyURLRegexp = regexp.MustCompile(
	`^https?://[-a-zA-Z0-9@:%._+~#=]{2,256}\.[a-z]+` +
		`(?:[/#?](?:[-a-zA-Z0-9@:%_+.~#$!?&/=();,'"{}\[\]^*]*[-a-zA-Z0-9/#=+~%@)}\]])?)?`)

// neverMatch disables linkify's www. and email detection. A bare "www." host is rarely a citation, and addresses in
// provider schemas are almost always placeholders ("joe@example.com").
var neverMatch = regexp.MustCompile(`a^`)

// citeLinkify wraps goldmark's linkify parser and links a bare URL only when the text around it cites it as a page to
// read. Schemas often show a sample value ("For example, https://myvault.vault.azure.net/keys/mykey"), and linking it
// sends readers to a page that doesn't exist, so a URL without that evidence stays plain text.
type citeLinkify struct {
	inner parser.InlineParser
}

func (p *citeLinkify) Trigger() []byte { return p.inner.Trigger() }

func (p *citeLinkify) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, segment := block.PeekLine()
	n := p.inner.Parse(parent, block, pc)
	link, ok := n.(*ast.AutoLink)
	if !ok {
		return n
	}
	source := block.Source()
	start := segment.Start
	// The inner parser consumes one leading trigger character and appends it to parent as text. This mirrors
	// linkifyParser.Parse in github.com/pgavlin/goldmark/extension/linkify.go.
	if strings.IndexByte(" *_~(", line[0]) >= 0 {
		start++
	}
	label := link.Label(source)
	end := start + len(label)
	paraStart := bytes.LastIndex(source[:start], []byte("\n\n")) + 1
	if !shouldLink(string(link.URL(source)), string(source[paraStart:start]), string(source[end:segment.Stop])) {
		return ast.NewTextSegment(text.NewSegment(start, end))
	}
	return link
}

func alternation(words ...string) string { return `(` + strings.Join(words, `|`) + `)` }

var (
	placeholderRegexp = regexp.MustCompile(`[{}\[\]*]|\$\{|/:[A-Za-z]|\.\.\.|(?i)xxx|(?i)/(my|your|sample|example)[-_]`)
	exampleHostRegexp = regexp.MustCompile(`(?i)` + strings.Join([]string{
		`(^|\.)(example|test)\.(com|org|net)$`,
		`\.(example|test|invalid|localhost)$`,
		`^localhost$`, `^127\.`, `^0\.0\.0\.0$`, `^10\.`, `^192\.168\.`,
		`contoso`, `fabrikam`, `(^|[.-])your`,
		`(^|\.)my` + alternation("account", "company", "org", "domain", "site", "app", "server", "host", "instance",
			"bucket", "vault", "tenant"),
		`^` + alternation("hostname", "host", "server", "domain", "bucketname", "accountname", "servername") + `\.`,
	}, `|`))
	// htmlTagRegexp matches the tags that schema descriptions converted from HTML put directly after a URL. Any other
	// "<" there starts a placeholder, as in "https://host/projects/<id>".
	htmlTagRegexp = regexp.MustCompile(`(?i)^</?` + alternation("a", "b", "br", "code", "div", "em", "i", "li", "ol",
		"p", "pre", "span", "strong", "sub", "sup", "table", "td", "th", "tr", "ul") + `\b`)
	// citeCueRegexp matches phrases that point the reader at a page: "See https://...", "More info: https://...",
	// "Documentation on CEL: https://...", "names that follow the conventions in https://...".
	citeCueRegexp = regexp.MustCompile(`(?i)\b` + alternation(`see\b`, `refer\b`, `ref\s*:`, `please\s+reference`,
		`check\b`, `more\s+info`, `information\s+(at|on|in|about)`, `for\s+(more\s+)?details`, `details\s+(at|in|on)`,
		`learn\s+more`, `visit\b`, `read\s+(more|about)`, `according\s+to`, `as\s+per`, `per\s*$`,
		`(available|found|listed|described|defined|documented|explained|specified)\s+(at|in|on|here)`,
		`(conventions|requirements|restrictions|rules|guidelines)\s+(at|in|of)\s*$`,
		`list\s+of\s+(available|supported|valid)\b[^.:]{0,40}:\s*$`,
		`(documentation|docs|reference|specification|spec|deprecation|guide)\b[^.:]{0,40}[-:]\s*$`))
	exampleCueRegexp = regexp.MustCompile(`(?i)\b` + alternation(`e\.g\b`, `eg[.:,]`, `examples?\b`, "for instance",
		"such as", "of the form", "looks like", `like\b`))
	// sentenceEndRegexp finds sentence boundaries, so a cue in an earlier sentence doesn't apply to this URL.
	// A period ending "e.g." or "i.e." is not a boundary.
	sentenceEndRegexp = regexp.MustCompile(`[;!?]\s|[^.]\.\s`)
	abbrevRegexp      = regexp.MustCompile(`(?i)\b(e\.g|i\.e|eg|ie|etc|ex|vs)\.\s$`)
)

// shouldLink reports whether a bare URL is cited as a page to visit. before is the text that precedes the URL in its
// paragraph, and after is the rest of its line.
func shouldLink(rawURL, before, after string) bool {
	if placeholderRegexp.MatchString(rawURL) || (strings.HasPrefix(after, "<") && !htmlTagRegexp.MatchString(after)) {
		return false
	}
	if u, err := url.Parse(rawURL); err != nil || exampleHostRegexp.MatchString(u.Hostname()) {
		return false
	}
	sentence := before
	for _, m := range sentenceEndRegexp.FindAllStringIndex(before, -1) {
		if !abbrevRegexp.MatchString(before[:m[1]]) {
			sentence = before[m[1]:]
		}
	}
	// The cue nearest the URL decides, so "For examples, see https://..." links and "See below, e.g. https://..." doesn't.
	citeAt := lastMatch(citeCueRegexp, sentence)
	// A URL in parentheses right after prose, as in "key rotation (https://...)", or one that starts a new sentence
	// after prose, as in "Defines the behavior of a node. https://...", cites a source.
	trimmed := strings.TrimSpace(sentence)
	if strings.HasSuffix(trimmed, "(") || (trimmed == "" && strings.TrimSpace(before) != "") {
		citeAt = len(sentence)
	}
	return citeAt >= 0 && citeAt > lastMatch(exampleCueRegexp, sentence)
}

// lastMatch returns the start of the last match of re in s, or -1.
func lastMatch(re *regexp.Regexp, s string) int {
	ms := re.FindAllStringIndex(s, -1)
	if len(ms) == 0 {
		return -1
	}
	return ms[len(ms)-1][0]
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
			parser.WithInlineParsers(util.Prioritized(&citeLinkify{inner: extension.NewLinkifyParser(
				extension.WithLinkifyAllowedProtocols([][]byte{[]byte("http:"), []byte("https:")}),
				extension.WithLinkifyURLRegexp(linkifyURLRegexp),
				extension.WithLinkifyWWWRegexp(neverMatch),
				extension.WithLinkifyEmailRegexp(neverMatch),
			)}, 999)),
			parser.WithASTTransformers(util.Prioritized(externalLinkTargets{}, 999)),
		),
	)
}
