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
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertDirEqual asserts that each file located under root is byte-for-byte identical
// with it's corresponding golden file. It also asserts on the directory structure itself.
//
// To update files under test, use:
//
//	go test ./... -update
//
// WARNING: Multiple AssertDirEqual should not run in the same go test. If you need to
// call AssertDirEqual more then once in a test, place each in a subtest:
//
//	t.Run("docs", func(t *testing.T) { util.AssertDirEqual(t, baseDocsOutDir) })
//	t.Run("tree", func(t *testing.T) { util.AssertDirEqual(t, basePackageTreeJSONOutDir) })
func AssertDirEqual(t *testing.T, root string, options ...AssertOption) {
	var opts assertOptions
	for _, o := range options {
		o(&opts)
	}
	var structure []string
	require.NoError(t, filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if path == root || err != nil {
			return err
		}

		pathName := strings.TrimLeft(strings.TrimPrefix(path, root), "/")
		require.NotEmpty(t, pathName, "internal error - pathName should not be empty")
		structure = append(structure, pathName)
		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		require.NoError(t, err, "could not read walked file")

		fileContent := string(content)
		for _, f := range opts.preCompareFileTransform {
			fileContent = f(t, fileContent)
		}

		t.Run(pathName, func(t *testing.T) { autogold.ExpectFile(t, autogold.Raw(fileContent)) })
		return nil
	}))
	t.Run("directory-structure", (func(t *testing.T) { autogold.ExpectFile(t, structure) }))
}

// AssertDirsEqual asserts that the file-tree at actual matches the file-tree at expected.
//
// If you just want to assert that a directory structure is unchanged or show updates, see
// [AssertDirEqual]. AssertDirsEqual is about showing that two already written out
// directories are equivalent.
func AssertDirsEqual(t *testing.T, expected, actual string, options ...AssertOption) {
	require.NotEqual(t, expected, actual, "cannot assert on the same directory")
	var opts assertOptions
	for _, o := range options {
		o(&opts)
	}
	var expectedStructure, actualStructure []string
	require.NoError(t, filepath.WalkDir(expected, func(path string, d fs.DirEntry, err error) error {
		if path == expected || err != nil {
			return err
		}

		pathName := strings.TrimLeft(strings.TrimPrefix(path, expected), "/")
		require.NotEmpty(t, pathName, "internal error - pathName should not be empty")
		expectedStructure = append(expectedStructure, pathName)
		if d.IsDir() {
			return nil
		}

		actualPath := filepath.Join(actual, pathName)
		actualContentBytes, err := os.ReadFile(actualPath)
		if errors.Is(err, os.ErrNotExist) {
			assert.Fail(t, "Missing expected file [%s/]%s", pathName, actual)
			return nil
		}
		require.NoError(t, err, "Could not open existing file %q", actualPath)

		expectedContentBytes, err := os.ReadFile(path)
		require.NoError(t, err, "walked file should exist")

		expectedContent := string(expectedContentBytes)
		actualContent := string(actualContentBytes)
		for _, f := range opts.preCompareFileTransform {
			expectedContent = f(t, expectedContent)
			actualContent = f(t, actualContent)
		}

		assert.Equalf(t, expectedContent, actualContent, "File %s doesn't match", pathName)

		return nil
	}))

	require.NoError(t, filepath.WalkDir(actual, func(path string, d fs.DirEntry, err error) error {
		if path == actual || err != nil {
			return err
		}

		pathName := strings.TrimLeft(strings.TrimPrefix(path, actual), "/")
		require.NotEmpty(t, pathName, "internal error - pathName should not be empty")
		actualStructure = append(actualStructure, pathName)
		return nil
	}))

	assert.ElementsMatch(t, expectedStructure, actualStructure,
		"Directory structure does not match")
}

type AssertOption func(*assertOptions)

type assertOptions struct {
	preCompareFileTransform []func(t *testing.T, fileContent string) (newFileContent string)
}

// Apply a regexp based transformation to all files before they are used in comparison..
func AssertOptionsPreCompareTransform(regex, with string) AssertOption {
	r, err := regexp.Compile(regex) // Pre-compile the regexp
	return func(o *assertOptions) {
		o.preCompareFileTransform = append(o.preCompareFileTransform, func(t *testing.T, src string) string {
			require.NoError(t, err)
			return r.ReplaceAllString(src, with)
		})
	}
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
