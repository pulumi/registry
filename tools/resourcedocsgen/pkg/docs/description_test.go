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
	"os"
	"path/filepath"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/require"
)

func TestProcessDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		prefix string
	}{
		{"lambda-description"},
		{"scaleway-k8s-cluster-description"}, // Repro: https://github.com/pulumi/registry/issues/4202
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.prefix, func(t *testing.T) {
			t.Parallel()

			input := readFile(t, filepath.Join("testdata", tt.prefix+".md"))

			testPackageSpec := newTestPackageSpec()
			schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.ValidationOptions{
				AllowDanglingReferences: true,
			})
			require.NoError(t, err, "importing spec")

			dctx := NewContext("test", schemaPkg)
			actual := dctx.processDescription(input, dctx.getSupportedSnippetLanguages(false, nil)).description

			autogold.ExpectFile(t, autogold.Raw(actual))
		})
	}
}

func TestDecomposeDocstringDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		prefix string
	}{
		{"lambda-description"},                 // renders code choosers
		{"certificate-validation-description"}, // renders legacy shortcode examples
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.prefix, func(t *testing.T) {
			t.Parallel()

			input := readFile(t, filepath.Join("testdata", tt.prefix+".md"))

			testPackageSpec := newTestPackageSpec()
			schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.ValidationOptions{
				AllowDanglingReferences: true,
			})
			require.NoError(t, err, "importing spec")

			dctx := NewContext("test", schemaPkg)
			actual := dctx.decomposeDocstring(input, dctx.getSupportedSnippetLanguages(false, nil)).description

			autogold.ExpectFile(t, autogold.Raw(actual))
		})
	}
}

func readFile(t *testing.T, filepath string) string {
	inputBytes, err := os.ReadFile(filepath)
	require.NoError(t, err)
	return string(inputBytes)
}
