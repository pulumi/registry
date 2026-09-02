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
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// resourceRef is a `{{% ref %}}` shortcode pointing at the test package's
// module-level Resource.
const resourceRef = `{{% ref #/resources/prov:module%2Fresource:Resource %}}`

// stubDocLanguageHelper answers ResolveDocRef with a canned result. The
// embedded interface is nil: ref rendering only ever calls ResolveDocRef, so
// any other call is a bug we want to fail loudly.
type stubDocLanguageHelper struct {
	codegen.DocLanguageHelper

	name string
	ok   bool
}

func (s stubDocLanguageHelper) ResolveDocRef(
	pkg schema.PackageReference, selfRef, ref schema.DocRef,
) (string, bool, error) {
	return s.name, s.ok, nil
}

// newRefTestContext returns a Context over the test package whose language
// helpers resolve every ref to names[lang].
func newRefTestContext(t *testing.T, names map[language.Language]string) *Context {
	t.Helper()

	schemaPkg, err := schema.ImportSpec(newTestPackageSpec(), nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)
	for lang := range language.All() {
		name, ok := names[lang]
		dctx.docHelpers[lang] = stubDocLanguageHelper{name: name, ok: ok}
	}
	return dctx
}

func TestRefLanguagesMatchesLanguageAll(t *testing.T) {
	t.Parallel()

	var want []language.Language
	for lang := range language.All() {
		want = append(want, lang)
	}

	require.Len(t, refLanguages, len(want))
	for i, entry := range refLanguages {
		assert.Equal(t, want[i], entry.lang, "refLanguages must follow language.All() order")
		assert.NotEmpty(t, entry.tag, "language %s has no choosable tag", entry.lang)
	}
}

func TestResolveRefsAllLanguagesAgree(t *testing.T) {
	t.Parallel()

	dctx := newRefTestContext(t, map[language.Language]string{
		language.CSharp: "Resource",
		language.Go:     "Resource",
		language.NodeJS: "Resource",
		language.Python: "Resource",
		language.YAML:   "Resource",
		language.Java:   "Resource",
		language.HCL:    "Resource",
	})

	got := dctx.resolveRefs("See also " + resourceRef + " for details.")
	assert.Equal(t, "See also Resource for details.", got)
}

func TestResolveRefsGroupsLanguagesByName(t *testing.T) {
	t.Parallel()

	// C#, TypeScript, YAML, Java and HCL agree, so they share one element;
	// Go and Python each get their own.
	dctx := newRefTestContext(t, map[language.Language]string{
		language.CSharp: "Resource",
		language.Go:     "module.Resource",
		language.NodeJS: "Resource",
		language.Python: "_module.Resource",
		language.YAML:   "Resource",
		language.Java:   "Resource",
		language.HCL:    "Resource",
	})

	got := dctx.resolveRefs("See also " + resourceRef + " for details.")
	assert.Equal(t, `See also `+
		`<pulumi-choosable type="language" values="csharp,javascript,typescript,yaml,java,hcl" `+
		`class="inline">Resource</pulumi-choosable>`+
		`<pulumi-choosable type="language" values="go" class="inline">module.Resource</pulumi-choosable>`+
		`<pulumi-choosable type="language" values="python" class="inline">_module.Resource</pulumi-choosable>`+
		` for details.`, got)
}

func TestResolveRefsFallsBackWhenNoHelperResolves(t *testing.T) {
	t.Parallel()

	// Every helper declines, so every language falls back to the same default
	// name — which means the reader has nothing to choose between.
	dctx := newRefTestContext(t, map[language.Language]string{})

	got := dctx.resolveRefs("See also " + resourceRef + " for details.")
	assert.NotContains(t, got, "{{%", "the shortcode should not survive")
	assert.NotContains(t, got, "pulumi-choosable", "identical fallbacks should render as plain text")
	assert.Contains(t, got, "Resource")
}

func TestResolveRefsForLanguage(t *testing.T) {
	t.Parallel()

	dctx := newRefTestContext(t, map[language.Language]string{
		language.Go:     "module.Resource",
		language.Python: "_module.Resource",
	})

	assert.Equal(t, "See also module.Resource for details.",
		dctx.resolveRefsForLanguage("See also "+resourceRef+" for details.", language.Go))
	assert.Equal(t, "See also _module.Resource for details.",
		dctx.resolveRefsForLanguage("See also "+resourceRef+" for details.", language.Python))
}

func TestResolveRefsLeavesRefFreeDescriptionsAlone(t *testing.T) {
	t.Parallel()

	dctx := newRefTestContext(t, map[language.Language]string{})

	const description = "Nothing to resolve here.\n\n```go\nfmt.Println(\"{{ not a ref }}\")\n```\n"
	assert.Equal(t, description, dctx.resolveRefs(description))
	assert.Empty(t, dctx.resolveRefs(""))
}

func TestResolveRefsInImportDetails(t *testing.T) {
	t.Parallel()

	dctx := newRefTestContext(t, map[language.Language]string{
		language.CSharp: "Resource",
		language.Go:     "module.Resource",
		language.NodeJS: "Resource",
		language.Python: "Resource",
		language.YAML:   "Resource",
		language.Java:   "Resource",
		language.HCL:    "Resource",
	})

	info := dctx.processDescription(
		"Intro paragraph.\n\n## Import\n\nImport an existing "+resourceRef+" like so.\n",
		dctx.getSupportedSnippetLanguages(false, nil))

	assert.NotContains(t, info.importDetails, "{{%", "refs in the Import section should be resolved")
	assert.Contains(t, info.importDetails, `<pulumi-choosable type="language" values="go" `+
		`class="inline">module.Resource</pulumi-choosable>`)
}
