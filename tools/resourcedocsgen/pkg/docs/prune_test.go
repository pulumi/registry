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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const emptyModulesPackage = "prunepkg"

// newEmptyModulesPackageSpec returns a package spec exercising the three kinds of module the pruning pass has to tell
// apart:
//
//   - "shared" holds nothing but an object type used by resources elsewhere, the shape awsx's `awsx` module has. It
//     has no page of its own to generate, so it is pruned.
//   - "tree" holds nothing itself, but its child "tree/v1" has a resource, so it is kept as a parent.
//   - "funcs" has only a function — content, so it is kept.
func newEmptyModulesPackageSpec() schema.PackageSpec {
	return schema.PackageSpec{
		Name:        emptyModulesPackage,
		Version:     "0.0.1",
		Description: "A fake provider package used for testing empty module pruning.",
		Types: map[string]schema.ComplexTypeSpec{
			"prunepkg:shared:SharedBucket": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A bucket type shared across modules.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"bucketName": {
							Description: "The name of the bucket.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
					},
				},
			},
		},
		Resources: map[string]schema.ResourceSpec{
			"prunepkg:index:RootResource": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A package-level resource.",
					Type:        "object",
				},
				InputProperties: map[string]schema.PropertySpec{
					"bucket": {
						Description: "A bucket, whose type lives in the shared module.",
						TypeSpec:    schema.TypeSpec{Ref: "#/types/prunepkg:shared:SharedBucket"},
					},
				},
			},
			"prunepkg:tree/v1:LeafResource": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A resource in a nested module.",
					Type:        "object",
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			"prunepkg:funcs:getThing": {
				Description: "A function in a function-only module.",
			},
		},
	}
}

func newEmptyModulesContext(t *testing.T) *Context {
	t.Helper()

	spec := newEmptyModulesPackageSpec()
	schemaPkg, err := schema.ImportSpec(spec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	return NewContext("test", schemaPkg)
}

func TestPruneEmptyModules(t *testing.T) {
	t.Parallel()

	dctx := newEmptyModulesContext(t)
	modules := dctx.modules()

	names := make([]string, 0, len(modules))
	for name := range modules {
		names = append(names, name)
	}
	sort.Strings(names)

	// The root module is always kept, even though it would be pruned on a resource count of its own.
	assert.Equal(t, []string{"", "funcs", "tree", "tree/v1"}, names)

	rootChildren := make([]string, 0, len(modules[""].children))
	for _, child := range modules[""].children {
		rootChildren = append(rootChildren, child.mod)
	}
	sort.Strings(rootChildren)
	assert.Equal(t, []string{"funcs", "tree"}, rootChildren,
		"the pruned module must not be left dangling in its parent's children")
}

func TestPruneEmptyModulesGeneratePackage(t *testing.T) {
	t.Parallel()

	dctx := newEmptyModulesContext(t)
	files, err := dctx.GeneratePackage()
	require.NoError(t, err, "generating package")

	paths := make([]string, 0, len(files))
	for p := range files {
		paths = append(paths, p)
	}
	sort.Strings(paths)

	assert.NotContains(t, paths, "shared/_index.md", "a types-only module must not get a page")
	assert.Contains(t, paths, "_index.md")
	assert.Contains(t, paths, "tree/_index.md")
	assert.Contains(t, paths, "tree/v1/_index.md")
	assert.Contains(t, paths, "funcs/_index.md")

	rootIndex := string(files["_index.md"])
	assert.NotContains(t, rootIndex, `title="shared"`, "a pruned module must not be listed in its parent's modules")
	assert.Contains(t, rootIndex, `title="tree"`)
	assert.Contains(t, rootIndex, `title="funcs"`)

	// The main regression risk: the pruned module's types are still rendered inline on the pages of the resources
	// that reference them.
	resourcePage := string(files["rootresource/_index.md"])
	require.NotEmpty(t, resourcePage, "expected a page for RootResource")
	assert.Contains(t, resourcePage, "SharedBucket", "supporting types from a pruned module must still be documented")
}

func TestPruneEmptyModulesPackageTree(t *testing.T) {
	t.Parallel()

	dctx := newEmptyModulesContext(t)
	tree, err := dctx.GeneratePackageTree()
	require.NoError(t, err, "generating package tree")

	modules := map[string][]PackageTreeItem{}
	for _, item := range tree {
		if item.Type == entryTypeModule {
			modules[item.Name] = item.Children
		}
	}

	assert.NotContains(t, modules, "shared", "a pruned module must not appear in the nav")
	assert.Contains(t, modules, "funcs")
	require.Contains(t, modules, "tree", "a module whose only content is a non-empty child must be kept")
	require.Len(t, modules["tree"], 1)
	assert.Equal(t, "v1", modules["tree"][0].Name)
}

// A package whose modules are all pruned still gets its own top-level page, titled as a package rather than as a
// module.
func TestPruneEmptyModulesKeepsRootPage(t *testing.T) {
	t.Parallel()

	spec := newEmptyModulesPackageSpec()
	delete(spec.Resources, "prunepkg:tree/v1:LeafResource")
	delete(spec.Functions, "prunepkg:funcs:getThing")

	schemaPkg, err := schema.ImportSpec(spec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)
	files, err := dctx.GeneratePackage()
	require.NoError(t, err, "generating package")

	rootIndex := string(files["_index.md"])
	require.NotEmpty(t, rootIndex, "expected a top-level page")
	assert.Contains(t, rootIndex, `title_tag: "prunepkg Package"`)
	assert.NotContains(t, rootIndex, `<h2 id="modules">`, "no modules survived, so there is no modules section")
	assert.Contains(t, rootIndex, `title="RootResource"`, "the package's own resources are still listed")
}
