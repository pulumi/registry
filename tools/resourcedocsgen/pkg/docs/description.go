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

package docs

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const (
	beginCodeBlock = "<!--Start PulumiCodeChooser -->"
	endCodeBlock   = "<!--End PulumiCodeChooser -->"
)

type codeLocation struct {
	open  int
	close int
}

func getCodeSection(doc string) []codeLocation {
	var fences []codeLocation

	startIndex := 0
	for {
		open := strings.Index(doc[startIndex:], beginCodeBlock)
		if open == -1 {
			break
		}
		var fence codeLocation
		fence.open = startIndex + open
		startIndex += open + len(beginCodeBlock)
		closing := strings.Index(doc[startIndex:], endCodeBlock)

		contract.Assertf(closing != -1, "this should never happen: "+
			"there should be equal amounts of opening and closing code block markers")

		fence.close = startIndex + closing

		startIndex += closing + len(endCodeBlock)
		fences = append(fences, fence)
	}
	return fences
}

func markupBlock(block, supportedSnippetLanguages string) string {
	languages := []struct{ tag, choosable string }{
		{"typescript", "<div>\n<pulumi-choosable type=\"language\" values=\"javascript,typescript\">\n\n"},
		{"python", "<div>\n<pulumi-choosable type=\"language\" values=\"python\">\n\n"},
		{"go", "<div>\n<pulumi-choosable type=\"language\" values=\"go\">\n\n"},
		{"csharp", "<div>\n<pulumi-choosable type=\"language\" values=\"csharp\">\n\n"},
		{"java", "<div>\n<pulumi-choosable type=\"language\" values=\"java\">\n\n"},
		{"yaml", "<div>\n<pulumi-choosable type=\"language\" values=\"yaml\">\n\n"},
		{"hcl", "<div>\n<pulumi-choosable type=\"language\" values=\"hcl\">\n\n"},
	}
	const (
		chooserStartFmt = "<div>\n<pulumi-chooser type=\"language\" options=\"%s\"></pulumi-chooser>\n</div>\n"
		choosableEnd    = "</pulumi-choosable>\n</div>\n"
	)

	var markedUpBlock strings.Builder
	// first, append the start chooser
	markedUpBlock.WriteString(fmt.Sprintf(chooserStartFmt, supportedSnippetLanguages))

	for _, lang := range languages {
		// Add language specific open choosable
		markedUpBlock.WriteString(lang.choosable)
		// find our language - because we have no guarantee of order from our input, we need to find
		// both code fences and then append the content in the order that docsgen expects.
		start := strings.Index(block, "```"+lang.tag)
		if start == -1 {
			markedUpBlock.WriteString("```\n")
			markedUpBlock.WriteString(defaultMissingExampleSnippetPlaceholder)
			markedUpBlock.WriteString("\n```\n")
		} else {
			// find end index - this is the next code fence.
			endLangBlock := start + len("```"+lang.tag) + strings.Index(block[start+len("```"+lang.tag):], "```")
			// append code to block, and include code fences
			fence := block[start : endLangBlock+len("```")]
			markedUpBlock.WriteString(dedentContinuationLines(fence, len(lineIndent(block, start))))
			markedUpBlock.WriteRune('\n')
		}
		// add closing choosable
		markedUpBlock.WriteString(choosableEnd)
	}
	return markedUpBlock.String()
}

func (dctx *Context) processDescription(description, supportedSnippetLanguages string) docInfo {
	importDetails := ""
	parts := strings.Split(description, "\n\n## Import")
	if len(parts) > 1 {
		importDetails = parts[1]
		description = parts[0]
	}

	codeBlocks := getCodeSection(description)

	startIndex := 0
	var markedUpDescription string
	for _, block := range codeBlocks {
		// append text
		markedUpDescription += description[startIndex:block.open]
		codeBlock := description[block.open:block.close]
		markedUp := markupBlock(codeBlock, supportedSnippetLanguages)
		markedUpDescription += indentContinuationLines(markedUp, lineIndent(description, block.open))
		startIndex = block.close + len(endCodeBlock)
	}
	// append remainder of description, if any
	markedUpDescription += description[startIndex:]

	return docInfo{
		description:   markedUpDescription,
		importDetails: importDetails,
	}
}

func lineIndent(s string, offset int) string {
	prefix := s[strings.LastIndexByte(s[:offset], '\n')+1 : offset]
	if strings.TrimLeft(prefix, " \t") != "" {
		return ""
	}
	return prefix
}

func indentContinuationLines(s string, prefix string) string {
	if prefix == "" {
		return s
	}
	lines := strings.Split(s, "\n")
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			lines[i] = prefix + lines[i]
		}
	}
	return strings.Join(lines, "\n")
}

func dedentContinuationLines(s string, n int) string {
	if n == 0 {
		return s
	}
	lines := strings.Split(s, "\n")
	for i := 1; i < len(lines); i++ {
		width := len(lines[i]) - len(strings.TrimLeft(lines[i], " \t"))
		lines[i] = lines[i][min(width, n):]
	}
	return strings.Join(lines, "\n")
}
