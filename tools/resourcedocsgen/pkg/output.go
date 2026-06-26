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

// neutralizeHugoShortcodes disarms Hugo shortcode delimiters inside fenced code
// blocks, where Hugo would resolve them before Markdown rendering. Inline code
// spans and indented code blocks are outside the current generator surface.
func neutralizeHugoShortcodes(contents []byte) []byte {
	lines := strings.Split(string(contents), "\n")
	inFence := false
	var fenceChar byte
	var fenceLen int
	for i, line := range lines {
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
		contents = neutralizeHugoShortcodes(contents)
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
