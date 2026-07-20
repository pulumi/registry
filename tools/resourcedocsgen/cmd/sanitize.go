// Copyright 2016-2024, Pulumi Corporation.
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

package cmd

import "regexp"

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
// fails the whole Hugo build, we defensively repair the delimiter on fetch.
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
