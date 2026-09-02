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
	"fmt"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// refShortcodeRegex matches `{{% ref DEST %}}` shortcodes. The destination is
// captured non-greedily so a stray `%` in the middle doesn't consume beyond the
// intended closing delimiter.
var refShortcodeRegex = regexp.MustCompile(`\{\{%\s*ref\s+(\S.*?)\s*%\}\}`)

// refChoosableValues maps each supported language to the `values` attribute
// used by <pulumi-choosable>. Kept in sync with the mapping in markupBlock.
var refChoosableValues = map[language.Language]string{
	language.NodeJS: "javascript,typescript",
	language.Python: "python",
	language.Go:     "go",
	language.CSharp: "csharp",
	language.Java:   "java",
	language.YAML:   "yaml",
	language.HCL:    "hcl",
}

// resolveRefs replaces every `{{% ref DEST %}}` shortcode in description with a
// language-aware rendering. If every language's DocLanguageHelper produces the
// same name, the shortcode is replaced with that name as plain text; otherwise
// the replacement is a sequence of inline <pulumi-choosable> elements so the
// registry frontend can show each language its own name.
//
// selfRef identifies the entity being documented (used to render self-property
// refs unqualified). Pass the zero schema.DocRef{} if there is no such context.
func (dctx *Context) resolveRefs(description string, selfRef schema.DocRef) string {
	if description == "" || !strings.Contains(description, "{{% ref") {
		return description
	}
	return refShortcodeRegex.ReplaceAllStringFunc(description, func(match string) string {
		m := refShortcodeRegex.FindStringSubmatch(match)
		if len(m) != 2 {
			return match
		}
		return dctx.renderRef(m[1], selfRef)
	})
}

// renderRef resolves a single ref destination for every supported language and
// returns the appropriate replacement markup (plain text or wrapped in
// per-language <pulumi-choosable> elements).
func (dctx *Context) renderRef(dest string, selfRef schema.DocRef) string {
	if dctx.pkg == nil {
		return dest
	}
	pkgRef := dctx.pkg.Reference()

	// Interpret the ref as a standalone doc snippet once per language, so each
	// DocLanguageHelper can resolve it (or fall back) independently.
	snippet := "{{% ref " + dest + " %}}"
	perLang := map[language.Language]string{}
	var fallback string
	for lang := range language.All() {
		helper := dctx.getLanguageDocHelper(lang)
		rendered, err := pkgRef.InterpretPulumiRefs(snippet, func(ref schema.DocRef) (string, bool) {
			name, ok, err := helper.ResolveDocRef(pkgRef, selfRef, ref)
			if err != nil || !ok {
				return "", false
			}
			return name, true
		})
		if err != nil {
			// Malformed ref: leave the shortcode as-is so the source is
			// preserved and the failure is visible in the rendered docs.
			return snippet
		}
		perLang[lang] = rendered
		fallback = rendered
	}

	allSame := true
	for _, v := range perLang {
		if v != fallback {
			allSame = false
			break
		}
	}
	if allSame {
		return fallback
	}

	var b strings.Builder
	for lang := range language.All() {
		name, ok := perLang[lang]
		if !ok {
			continue
		}
		fmt.Fprintf(&b, `<pulumi-choosable type="language" values=%q>%s</pulumi-choosable>`,
			refChoosableValues[lang], name)
	}
	return b.String()
}
