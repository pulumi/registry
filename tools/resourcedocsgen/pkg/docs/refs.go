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
	"log/slog"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// refChoosableValues is the per-language `values` attribute used for
// <pulumi-choosable>. Iteration order matches language.All().
var refChoosableValues = [...]struct {
	lang   language.Language
	values string
}{
	{language.CSharp, "csharp"},
	{language.Go, "go"},
	{language.NodeJS, "javascript,typescript"},
	{language.Python, "python"},
	{language.YAML, "yaml"},
	{language.Java, "java"},
	{language.HCL, "hcl"},
}

func init() {
	// Guard against language.All() gaining an entry without a matching
	// choosable-values mapping, which would otherwise silently emit values="".
	seen := map[language.Language]bool{}
	for _, v := range refChoosableValues {
		seen[v.lang] = true
	}
	for lang := range language.All() {
		if !seen[lang] {
			panic(fmt.Errorf("refChoosableValues missing entry for language %s", lang))
		}
	}
}

// refShortcodeMarker is a permissive substring check for `{{% ref` (with
// optional space) so the guard doesn't miss `{{%ref` or `{{%  ref`.
var refShortcodeMarker = regexp.MustCompile(`\{\{%\s*ref[\s%]`)

// resolveRefs replaces every `{{% ref %}}` shortcode in description with a
// language-aware rendering. When every supported language's DocLanguageHelper
// produces the same name the shortcode is replaced with that name as plain
// text; otherwise the replacement is a sequence of inline <pulumi-choosable>
// elements so the registry frontend can show each language its own name.
//
// Scope: this only handles descriptions that flow through decomposeDocstring
// or processDescription (resources, functions, methods). Property comments,
// nested type descriptions, and enum value comments are single-language and
// go through [resolveRefsForLanguage] instead.
func (dctx *Context) resolveRefs(description string) string {
	return dctx.transformRefs(description, func(ref schema.DocRef) string {
		return dctx.renderRef(ref)
	})
}

// resolveRefsForLanguage replaces every `{{% ref %}}` shortcode with the name
// produced by lang's DocLanguageHelper. Intended for comments that are
// already rendered per-language (property/enum/nested-type descriptions).
func (dctx *Context) resolveRefsForLanguage(description string, lang language.Language) string {
	pkgRef := dctx.pkg.Reference()
	helper := dctx.getLanguageDocHelper(lang)
	return dctx.transformRefs(description, func(ref schema.DocRef) string {
		name, ok, err := helper.ResolveDocRef(pkgRef, schema.DocRef{}, ref)
		if err != nil {
			slog.Warn("resolving ref via language helper",
				"pkg", dctx.pkg.Name, "lang", lang.String(), "ref", ref.Ref, "err", err)
		}
		if !ok || name == "" {
			return defaultRefName(ref)
		}
		return name
	})
}

// transformRefs walks description with schema.InterpretPulumiRefs and hands
// each parsed ref to substitute. The guards short-circuit the parse/render
// round-trip for descriptions with no refs and for uninitialised contexts.
func (dctx *Context) transformRefs(description string, substitute func(schema.DocRef) string) string {
	if description == "" || dctx.pkg == nil {
		return description
	}
	if !refShortcodeMarker.MatchString(description) {
		return description
	}
	pkgRef := dctx.pkg.Reference()
	rendered, err := pkgRef.InterpretPulumiRefs(description, func(ref schema.DocRef) (string, bool) {
		return substitute(ref), true
	})
	if err != nil {
		slog.Warn("interpreting pulumi refs; leaving description unchanged",
			"pkg", dctx.pkg.Name, "err", err)
		return description
	}
	return rendered
}

// renderRef resolves ref via every language's DocLanguageHelper and returns
// either the shared name (when all languages agree) or a sequence of inline
// <pulumi-choosable> elements — one per language, in language.All() order.
func (dctx *Context) renderRef(ref schema.DocRef) string {
	pkgRef := dctx.pkg.Reference()
	names := make([]string, 0, len(refChoosableValues))
	var shared string
	first, allSame := true, true
	for _, entry := range refChoosableValues {
		helper := dctx.getLanguageDocHelper(entry.lang)
		name, ok, err := helper.ResolveDocRef(pkgRef, schema.DocRef{}, ref)
		if err != nil {
			slog.Warn("resolving ref via language helper",
				"pkg", dctx.pkg.Name, "lang", entry.lang.String(), "ref", ref.Ref, "err", err)
		}
		if !ok || name == "" {
			name = defaultRefName(ref)
		}
		names = append(names, name)
		switch {
		case first:
			shared, first = name, false
		case name != shared:
			allSame = false
		}
	}
	if allSame {
		return shared
	}

	var b strings.Builder
	for i, entry := range refChoosableValues {
		fmt.Fprintf(&b, `<pulumi-choosable type="language" values=%q>%s</pulumi-choosable>`,
			entry.values, names[i])
	}
	return b.String()
}

// defaultRefName mirrors the fallback rendering that schema.interpretPulumiRefs
// uses when the resolver returns false: a property name, the last segment of a
// type/function token, or the raw ref as a last resort. We prefer the last
// token segment over Type.String()/Function.Token so unresolved refs read as
// user-recognisable names instead of internal `pkg:module/name:Type` tokens.
func defaultRefName(ref schema.DocRef) string {
	if ref.Property != "" {
		return ref.Property
	}
	var tok string
	switch {
	case ref.Type != nil:
		tok = ref.Type.String()
	case ref.Function != nil:
		tok = ref.Function.Token
	default:
		return ref.Ref
	}
	if idx := strings.LastIndex(tok, ":"); idx != -1 && idx+1 < len(tok) {
		return tok[idx+1:]
	}
	return tok
}
