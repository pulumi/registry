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
	"context"
	"testing"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stubSchemaLoader is a schema.Loader that always returns the same package.
// Used by tests that bind a package whose Dependencies list references one
// external package.
type stubSchemaLoader struct {
	pkg *schema.Package
}

var _ schema.Loader = (*stubSchemaLoader)(nil)

func (l *stubSchemaLoader) LoadPackage(_ string, _ *semver.Version) (*schema.Package, error) {
	return l.pkg, nil
}

func (l *stubSchemaLoader) LoadPackageV2(_ context.Context, _ *schema.PackageDescriptor) (*schema.Package, error) {
	return l.pkg, nil
}

// TestResolveRefsExternal exercises `{{% ref %}}` shortcodes whose destination
// points at an external package via the `/pkg/version/schema.json#/...` prefix.
// External refs are resolved through the current package's Dependencies list,
// with the package loaded by the schema loader.
func TestResolveRefsExternal(t *testing.T) {
	t.Parallel()

	// Build an "extdep" package separately. Doc refs in the current package
	// can resolve into it via the Dependencies list, mirroring how `$ref`
	// resolves external schema types.
	extSpec := schema.PackageSpec{
		Name:    "extdep",
		Version: "1.0.0",
		Resources: map[string]schema.ResourceSpec{
			"extdep:mod:ExtResource": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						"extProp": {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			"extdep:mod:extFunction": {},
		},
		Types: map[string]schema.ComplexTypeSpec{
			"extdep:mod:ExtType": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"extField": {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
			},
		},
	}
	extPkg, err := schema.ImportSpec(extSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{})
	require.NoError(t, err)

	loader := &stubSchemaLoader{pkg: extPkg}
	extVersion := semver.MustParse("1.0.0")
	spec := schema.PackageSpec{
		Name:         "test",
		Version:      "1.0.0",
		Dependencies: []schema.PackageDescriptor{{Name: "extdep", Version: &extVersion}},
	}
	pkg, diags, err := schema.BindSpec(spec, loader, schema.ValidationOptions{})
	require.NoError(t, err)
	require.False(t, diags.HasErrors(), "diags: %v", diags)

	dctx := NewContext("test", pkg)

	// choose builds a run of inline <pulumi-choosable> elements, one per
	// supported language in language.All() order (matches refs.renderRef).
	choose := func(csharp, gogo, ts, py, yaml, java, hcl string) string {
		return `<pulumi-choosable type="language" values="csharp">` + csharp + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="go">` + gogo + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="javascript,typescript">` + ts + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="python">` + py + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="yaml">` + yaml + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="java">` + java + `</pulumi-choosable>` +
			`<pulumi-choosable type="language" values="hcl">` + hcl + `</pulumi-choosable>`
	}

	t.Run("ResourceRef", func(t *testing.T) {
		t.Parallel()
		got := dctx.resolveRefs(schema.DocRef{},
			"See {{% ref /extdep/v1.0.0/schema.json#/resources/extdep:mod:ExtResource %}} for details.")
		want := "See " + choose(
			"ExtResource", "mod.ExtResource", "ExtResource", "_mod.ExtResource",
			"extdep:mod:ExtResource", "ExtResource", "ExtResource",
		) + " for details."
		assert.Equal(t, want, got)
	})

	t.Run("TypeRef", func(t *testing.T) {
		t.Parallel()
		got := dctx.resolveRefs(schema.DocRef{},
			"Uses {{% ref /extdep/v1.0.0/schema.json#/types/extdep:mod:ExtType %}} as input.")
		want := "Uses " + choose(
			"ExtType", "mod.ExtType", "ExtType", "ExtTypeArgs",
			"ExtType", "ExtType", "ExtType",
		) + " as input."
		assert.Equal(t, want, got)
	})

	t.Run("FunctionRef", func(t *testing.T) {
		t.Parallel()
		got := dctx.resolveRefs(schema.DocRef{},
			"Call {{% ref /extdep/v1.0.0/schema.json#/functions/extdep:mod:extFunction %}} to invoke.")
		want := "Call " + choose(
			"extFunction", "ExtFunction", "extFunction", "ext_function",
			"extdep:mod:extFunction", "extFunction", "extFunction",
		) + " to invoke."
		assert.Equal(t, want, got)
	})

	t.Run("ExternalPropertyRef", func(t *testing.T) {
		t.Parallel()
		got := dctx.resolveRefs(schema.DocRef{},
			"See {{% ref /extdep/v1.0.0/schema.json#/resources/extdep:mod:ExtResource/properties/extProp %}}.")
		want := "See " + choose(
			"extProp", "mod.ExtResource.ExtProp", "ExtResource.extProp", "_mod.ExtResource.ext_prop",
			"extProp", "extProp", "extProp",
		) + "."
		assert.Equal(t, want, got)
	})

	t.Run("MissingExternalRefLeavesShortcode", func(t *testing.T) {
		t.Parallel()
		// A destination the loader can't satisfy should leave the shortcode
		// intact and log a warning rather than mangling the description.
		in := "Broken {{% ref /extdep/v1.0.0/schema.json#/resources/extdep:mod:DoesNotExist %}}."
		got := dctx.resolveRefs(schema.DocRef{}, in)
		assert.Equal(t, in, got)
	})
}
