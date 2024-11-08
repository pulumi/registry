// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
)

func AssertDirEqual(t *testing.T, root string) {
	var structure []string
	require.NoError(t, filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		t.Logf("examining %s", path)
		if path == root {
			return nil
		}
		if err != nil {
			return err
		}

		pathName := strings.TrimLeft(strings.TrimPrefix(path, root), "/")
		require.NotEmpty(t, pathName, "internal error - pathName should not be empty")
		t.Logf("named path - %s", pathName)
		structure = append(structure, pathName)
		if d.IsDir() {
			return nil
		}

		t.Logf("reading file %s", pathName)
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		t.Run(pathName, func(t *testing.T) { autogold.ExpectFile(t, autogold.Raw(content)) })
		return nil
	}))
	t.Run("directory-structure", (func(t *testing.T) { autogold.ExpectFile(t, structure) }))
}

func WriteFile(t *testing.T, path, contents string) {
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o700))
	require.NoError(t, os.WriteFile(path, []byte(contents), 0o600))
}

func ReadFile(t *testing.T, path string) string {
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(b)
}
