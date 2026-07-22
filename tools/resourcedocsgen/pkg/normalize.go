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
	"bytes"
	"regexp"
)

// malformedShortcodeOpen matches a Hugo shortcode opening delimiter that has stray
// whitespace between the braces and the directive marker, e.g. "{{ <" or "{{  %".
//
// Hugo requires the marker ("<" for HTML/block shortcodes, "%" for markdown
// shortcodes) to immediately follow "{{" with no intervening whitespace. A space
// there makes Hugo stop recognizing the opening tag while still recognizing the
// paired closing tag, which aborts the entire site build with:
//
//	shortcode "X" does not evaluate .Inner or .InnerDeindent, yet a closing tag was provided
//
// We have observed upstream-generated docs occasionally emit this malformed form
// (e.g. registry.opentofu.org-mirrored provider docs). Because a single bad page
// fails the whole Hugo build, we defensively repair the delimiter.
//
// The pattern is intentionally narrow: it only collapses whitespace that sits
// between "{{" and a following "<" or "%". A space after "{{" followed by anything
// else (e.g. a Go template expression "{{ .Value }}") is left untouched, and the
// closing side (e.g. the legitimate space in `... yaml" >}}`) is never modified.
//
// The whitespace class is [\t\p{Zs}] rather than [ \t]: \p{Zs} matches every
// Unicode space separator, so a non-breaking space (U+00A0) or other exotic
// space that Hugo still rejects — and that renders indistinguishably from a
// plain space — is repaired too. \p{Zs} excludes newlines, so a match never
// spans lines.
var malformedShortcodeOpen = regexp.MustCompile(`\{\{[\t\p{Zs}]+([<%])`)

// sanitizeShortcodeDelimiters repairs malformed Hugo shortcode opening delimiters
// of the form "{{ <" / "{{ %" back to the well-formed "{{<" / "{{%". It is a
// no-op for content that is already well-formed.
func sanitizeShortcodeDelimiters(content []byte) []byte {
	return malformedShortcodeOpen.ReplaceAll(content, []byte("{{$1"))
}

// NormalizeDocs is the single canonical transform resourcedocsgen applies to
// every docs .md file it writes, regardless of how the file was obtained
// (fetched from a mirror, fetched from GitHub, generated from a schema, or piped
// through the sanitize-docs command). It is the one place broken shortcode
// delimiters are handled, so every consumer treats them identically.
//
// It performs, in order:
//
//  1. Line-ending normalization: "\r\n" -> "\n".
//  2. Delimiter repair: collapse stray whitespace in opening delimiters
//     ("{{ <" -> "{{<") everywhere, so real shortcodes Hugo must resolve are
//     well-formed.
//  3. Delimiter neutralization: disarm delimiters inside fenced code blocks
//     ("{{<" -> "{{ <"), so shortcode-shaped example text Hugo must NOT resolve
//     is left inert.
//
// The result is the canonical committed form — well-formed delimiters outside
// code fences, neutralized delimiters inside them. Steps 2 and 3 target disjoint
// regions (outside vs. inside fences), and registry wrapping shortcodes
// (chooser/choosable/example(s)) are never neutralized, so NormalizeDocs is
// idempotent: normalizing already-normalized content is a no-op. It is lenient
// and never fails on content shape, so callers on `set -e` paths can apply it
// unconditionally.
func NormalizeDocs(content []byte) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	content = sanitizeShortcodeDelimiters(content)
	return neutralizeHugoShortcodes(content)
}
