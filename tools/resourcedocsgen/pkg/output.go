// Copyright 2024, Pulumi Corporation.
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
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// hugoShortcodeDelim matches Hugo's shortcode delimiters `{{%` and `{{<`.
var hugoShortcodeDelim = regexp.MustCompile(`\{\{([%<])`)

// wrappingShortcode matches a line that is a registry fence-wrapping block
// shortcode delimiter — chooser / choosable / example(s) — opening or closing.
//
// These shortcodes wrap fenced code examples from the OUTSIDE and never
// legitimately appear inside a code example, so they must never be neutralized.
// A malformed upstream doc (e.g. registry.opentofu.org-mirrored provider docs
// with a stray, unpaired code fence) can desync the fence tracker below and make
// it believe such a line sits inside a fence. Neutralizing the delimiter there
// aborts the Hugo build — the paired closing tag is still recognized — and trips
// the shortcode-delimiter lint. Anchoring on these lines both preserves the
// shortcode and re-syncs the tracker (see neutralizeHugoShortcodes).
//
// The match is anchored to the start of the (whitespace-trimmed) line so a
// delimiter appearing mid-line inside genuine example code is still neutralized.
var wrappingShortcode = regexp.MustCompile(
	`^\s*\{\{[%<]\s*/?(choosable-inline|choosable|chooser|examples|example)\b`)

// neutralizeHugoShortcodes disarms Hugo shortcode delimiters inside fenced code
// blocks, where Hugo would resolve them before Markdown rendering. Inline code
// spans and indented code blocks are outside the current generator surface.
//
// Registry fence-wrapping shortcodes (chooser/choosable/example(s)) are never
// neutralized: a line that is one of them is, by definition, outside a code
// example, so encountering one re-syncs the fence tracker. This keeps a
// malformed upstream fence from corrupting real shortcodes further down the page
// while still disarming genuine code-example delimiters that follow.
func neutralizeHugoShortcodes(contents []byte) []byte {
	lines := strings.Split(string(contents), "\n")
	inFence := false
	var fenceChar byte
	var fenceLen int
	for i, line := range lines {
		if wrappingShortcode.MatchString(line) {
			inFence = false
			continue
		}
		if c, n, rest := fenceRun(line); n > 0 {
			switch {
			case !inFence:
				inFence, fenceChar, fenceLen = true, c, n
				continue
			case c == fenceChar && n >= fenceLen && strings.TrimSpace(rest) == "":
				inFence = false
				continue
			}
		}
		if inFence {
			lines[i] = hugoShortcodeDelim.ReplaceAllString(line, "{{ $1")
		}
	}
	return []byte(strings.Join(lines, "\n"))
}

// fenceRun returns the fence char, run length, and trailing text if the line
// starts with a CommonMark fence marker.
func fenceRun(line string) (byte, int, string) {
	indent := 0
	for indent < len(line) && line[indent] == ' ' {
		indent++
	}
	if indent > 3 {
		return 0, 0, ""
	}
	line = line[indent:]
	if len(line) < 3 || (line[0] != '`' && line[0] != '~') {
		return 0, 0, ""
	}
	c := line[0]
	n := 0
	for n < len(line) && line[n] == c {
		n++
	}
	if n < 3 {
		return 0, 0, ""
	}
	return c, n, line[n:]
}

// EmitFile writes the file with the provided contents in the output
// directory outDir.
func EmitFile(outDir, relPath string, contents []byte) error {
	// Do not emit a file if there are no contents to write.
	if contents == nil {
		return nil
	}
	if strings.HasSuffix(relPath, ".md") {
		contents = NormalizeDocs(contents)
	}
	p := path.Join(outDir, relPath)
	if err := os.MkdirAll(path.Dir(p), 0o700); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}
