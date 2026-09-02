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
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

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

// resolveRefs walks description with schema.InterpretPulumiRefs and replaces
// every `{{% ref %}}` shortcode with a language-aware rendering produced by
// [renderRef]. selfRef identifies the entity being documented (used to render
// self-property refs unqualified); pass the zero schema.DocRef{} if there is
// no such context.
func (dctx *Context) resolveRefs(description string, selfRef schema.DocRef) string {
	if description == "" || dctx.pkg == nil {
		return description
	}
	// Skip the parse/render round-trip when there are no refs — otherwise
	// goldmark normalizes trailing whitespace and other formatting even for
	// docstrings we'd otherwise leave alone.
	if !strings.Contains(description, "{{% ref") {
		return description
	}
	pkgRef := dctx.pkg.Reference()
	rendered, err := pkgRef.InterpretPulumiRefs(description, func(ref schema.DocRef) (string, bool) {
		return dctx.renderRef(pkgRef, selfRef, ref), true
	})
	if err != nil {
		// Malformed refs (e.g. unknown token): leave the description untouched
		// so the failure surfaces in the rendered docs.
		return description
	}
	return rendered
}

// renderRef asks each language's DocLanguageHelper to resolve ref and returns
// the appropriate replacement text: the shared name when every language agrees,
// otherwise a sequence of inline <pulumi-choosable> elements — one per
// language — so the registry frontend can show each language its own name.
func (dctx *Context) renderRef(pkgRef schema.PackageReference, selfRef, ref schema.DocRef) string {
	perLang := map[language.Language]string{}
	allSame := true
	var shared string
	for lang := range language.All() {
		helper := dctx.getLanguageDocHelper(lang)
		name, ok, err := helper.ResolveDocRef(pkgRef, selfRef, ref)
		if err != nil || !ok {
			// Helper couldn't resolve this ref — fall back to schema's default
			// rendering (property name, type token, or ref string) for this
			// language.
			name = defaultRefName(ref)
		}
		perLang[lang] = name
		if shared == "" {
			shared = name
		} else if name != shared {
			allSame = false
		}
	}
	if allSame {
		return shared
	}

	var b strings.Builder
	for lang := range language.All() {
		fmt.Fprintf(&b, `<pulumi-choosable type="language" values=%q>%s</pulumi-choosable>`,
			refChoosableValues[lang], perLang[lang])
	}
	return b.String()
}

// defaultRefName mirrors the fallback rendering that schema.interpretPulumiRefs
// uses when the resolver callback returns false.
func defaultRefName(ref schema.DocRef) string {
	switch {
	case ref.Property != "":
		return ref.Property
	case ref.Type != nil:
		return ref.Type.String()
	case ref.Function != nil:
		return ref.Function.Token
	default:
		return ref.Ref
	}
}
