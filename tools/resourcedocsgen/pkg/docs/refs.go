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

// refChoosableTags maps each docs language to the tag it contributes to a
// <pulumi-choosable> `values` attribute.
var refChoosableTags = map[language.Language]string{
	language.CSharp: "csharp",
	language.Go:     "go",
	language.NodeJS: "javascript,typescript",
	language.Python: "python",
	language.YAML:   "yaml",
	language.Java:   "java",
	language.HCL:    "hcl",
}

type refLanguage struct {
	lang language.Language
	tag  string
}

// refLanguages pairs every docs language with its choosable tag. It is built by
// iterating language.All() so the "renders in language.All() order" guarantee is
// structural rather than a claim about a hand-maintained list, and so a language
// added without a tag fails at init instead of silently emitting values="".
var refLanguages = buildRefLanguages()

func buildRefLanguages() []refLanguage {
	var langs []refLanguage
	for lang := range language.All() {
		tag, ok := refChoosableTags[lang]
		if !ok {
			panic(fmt.Errorf("refChoosableTags missing entry for language %s", lang))
		}
		langs = append(langs, refLanguage{lang: lang, tag: tag})
	}
	return langs
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
// Use this for descriptions rendered once for all languages: resource,
// function and method comments, deprecation messages, nested type descriptions
// and the package description. Comments that are already rendered per-language
// (property and enum comments) go through [resolveRefsForLanguage] instead.
func (dctx *Context) resolveRefs(description string) string {
	return dctx.transformRefs(description, dctx.renderRef)
}

// resolveRefsForLanguage replaces every `{{% ref %}}` shortcode with the name
// produced by lang's DocLanguageHelper. Intended for comments that are
// already rendered per-language (property/enum/nested-type descriptions).
func (dctx *Context) resolveRefsForLanguage(description string, lang language.Language) string {
	return dctx.transformRefs(description, func(pkgRef schema.PackageReference, ref schema.DocRef) string {
		helper := dctx.getLanguageDocHelper(lang)
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
// each parsed ref to substitute, along with the package reference resolved once
// for the whole description. The guards short-circuit the parse/render
// round-trip for descriptions with no refs and for uninitialised contexts.
func (dctx *Context) transformRefs(
	description string, substitute func(schema.PackageReference, schema.DocRef) string,
) string {
	if description == "" || dctx.pkg == nil {
		return description
	}
	if !refShortcodeMarker.MatchString(description) {
		return description
	}
	pkgRef := dctx.pkg.Reference()
	rendered, err := pkgRef.InterpretPulumiRefs(description, func(ref schema.DocRef) (string, bool) {
		return substitute(pkgRef, ref), true
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
// <pulumi-choosable> elements. Languages that produced the same name share an
// element, so a ref that reads the same everywhere but Go costs two elements
// rather than seven. Both the elements and the languages within an element's
// `values` list are ordered by language.All().
func (dctx *Context) renderRef(pkgRef schema.PackageReference, ref schema.DocRef) string {
	type refGroup struct {
		name string
		tags []string
	}
	var groups []*refGroup
	byName := map[string]*refGroup{}
	for _, entry := range refLanguages {
		helper := dctx.getLanguageDocHelper(entry.lang)
		name, ok, err := helper.ResolveDocRef(pkgRef, schema.DocRef{}, ref)
		if err != nil {
			slog.Warn("resolving ref via language helper",
				"pkg", dctx.pkg.Name, "lang", entry.lang.String(), "ref", ref.Ref, "err", err)
		}
		if !ok || name == "" {
			name = defaultRefName(ref)
		}
		group, seen := byName[name]
		if !seen {
			group = &refGroup{name: name}
			byName[name] = group
			groups = append(groups, group)
		}
		group.tags = append(group.tags, entry.tag)
	}
	// Every language agrees, so there is nothing for the reader to choose
	// between: emit the name as plain text.
	if len(groups) == 1 {
		return groups[0].name
	}

	var b strings.Builder
	for _, group := range groups {
		// class="inline" is load-bearing: the registry theme sets
		// `pulumi-choosable { display: block }` in @layer base
		// (theme/src/scss/main.scss), so without the class every ref would
		// break its sentence onto its own line.
		fmt.Fprintf(&b, `<pulumi-choosable type="language" values=%q class="inline">%s</pulumi-choosable>`,
			strings.Join(group.tags, ","), group.name)
	}
	return b.String()
}

// defaultRefName mirrors the fallback rendering that schema.interpretPulumiRefs
// uses when the resolver returns false: a property name, the last segment of a
// type/function token, or the raw ref as a last resort. We prefer the last
// token segment over Type.String()/Function.Token so unresolved refs read as
// user-recognisable names instead of internal `pkg:module/name:Type` tokens.
//
// We mirror it here rather than returning false and letting upstream do it,
// because a false return would discard the per-language renderings the other
// helpers produced for the same ref. The tradeoff is that upstream's fallback
// never runs for registry docs: if it ever grows past plain text (rendering
// refs as links, say), this opts out until it is taught to do so here.
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
