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
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateCLIPackage(t *testing.T) {
	t.Parallel()

	testPackageSpec := newTestPackageSpec()

	// Add a deprecated resource to the test spec.
	testPackageSpec.Resources["prov:module/deprecatedResource:DeprecatedResource"] = schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "This resource is deprecated.",
		},
		DeprecationMessage: "Use Resource instead.",
		InputProperties: map[string]schema.PropertySpec{
			"prop": {
				Description: "An input property.",
				TypeSpec:    schema.TypeSpec{Type: "string"},
			},
		},
	}

	schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)
	bundle, err := dctx.GenerateCLIPackage()
	require.NoError(t, err, "generating CLI package")

	t.Run("PackageMetadata", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, providerPackage, bundle.Package)
		assert.Equal(t, "0.0.1", bundle.PackageVersion)
	})

	t.Run("ResourceEntryHasMetadata", func(t *testing.T) {
		t.Parallel()
		// Find a resource entry — the test package has resources in module and module2.
		var found bool
		for key, entry := range bundle.Resources {
			if strings.Contains(key, "resource") && !strings.Contains(key, "deprecated") {
				found = true
				assert.NotEmpty(t, entry.Title, "resource %s should have a title", key)
				assert.NotEmpty(t, entry.Description, "resource %s should have a description", key)
				assert.NotEmpty(t, entry.Content, "resource %s should have content", key)
				// Title should NOT appear as a markdown heading in the content.
				assert.False(t, strings.HasPrefix(entry.Content, "# "),
					"content for %s should not start with a markdown title heading", key)
				break
			}
		}
		assert.True(t, found, "should have found at least one resource entry")
	})

	t.Run("FunctionEntryHasMetadata", func(t *testing.T) {
		t.Parallel()
		var found bool
		for key, entry := range bundle.Functions {
			found = true
			assert.NotEmpty(t, entry.Title, "function %s should have a title", key)
			assert.NotEmpty(t, entry.Description, "function %s should have a description", key)
			assert.NotEmpty(t, entry.Content, "function %s should have content", key)
			assert.False(t, strings.HasPrefix(entry.Content, "# "),
				"content for %s should not start with a markdown title heading", key)
			break
		}
		assert.True(t, found, "should have found at least one function entry")
	})

	t.Run("DeprecatedResourceHasMetadata", func(t *testing.T) {
		t.Parallel()
		var found bool
		for key, entry := range bundle.Resources {
			if strings.Contains(key, "deprecated") {
				found = true
				assert.True(t, entry.Deprecated, "deprecated resource %s should have Deprecated=true", key)
				assert.Equal(t, "Use Resource instead.", entry.DeprecationMessage)
				// Content should NOT contain the deprecation blockquote.
				assert.NotContains(t, entry.Content, "> **Deprecated:**",
					"content for %s should not contain deprecation blockquote", key)
				break
			}
		}
		assert.True(t, found, "should have found the deprecated resource entry")
	})

	t.Run("OverviewIsPopulated", func(t *testing.T) {
		t.Parallel()
		assert.NotEmpty(t, bundle.Overview)
		// Should contain module list with markdown bullet syntax.
		assert.Contains(t, bundle.Overview, "## Modules")
		assert.Contains(t, bundle.Overview, "- [")
		// Should contain package details.
		assert.Contains(t, bundle.Overview, "## Package Details")
	})

	t.Run("ContentHasNoWindowsLineEndings", func(t *testing.T) {
		t.Parallel()
		for key, entry := range bundle.Resources {
			assert.NotContains(t, entry.Content, "\r\n",
				"resource %s content should not contain CRLF", key)
		}
		for key, entry := range bundle.Functions {
			assert.NotContains(t, entry.Content, "\r\n",
				"function %s content should not contain CRLF", key)
		}
	})
}

func TestNormalizeWhitespace(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"strips CRLF", "line1\r\nline2", "line1\nline2"},
		{"converts tabs", "col1\tcol2", "col1    col2"},
		{"trims whitespace", "  hello  ", "hello"},
		{"combined", "\r\n\thello\r\n", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, normalizeWhitespace(tt.input))
		})
	}
}
