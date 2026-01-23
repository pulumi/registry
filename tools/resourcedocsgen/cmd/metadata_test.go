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

package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/registry/tools/resourcedocsgen/internal/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataBridgedProvider(t *testing.T) {
	t.Parallel()
	testMetadata(t, testMetadataArgs{
		repoSlug:   "pulumi/pulumi-random",
		publisher:  "Pulumi",
		version:    "v4.16.7",
		schemaFile: "provider/cmd/pulumi-resource-random/schema.json",
	})
}

func TestMetadataNativeProvider(t *testing.T) {
	t.Parallel()

	var metadataDir, packageDocsDir string
	testMetadata(t, testMetadataArgs{
		repoSlug:   "pulumi/pulumi-command",
		version:    "v1.0.0",
		schemaFile: "provider/cmd/pulumi-resource-command/schema.json",
		assert: func(t *testing.T, metadata, packageDocs string) {
			defaultAssert(t, metadata, packageDocs)
			metadataDir = metadata
			packageDocsDir = packageDocs
		},
	})

	t.Run("test-remote-equivalence", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/schema.json":
				b, err := os.ReadFile("./testdata/command-v1.0.0.json")
				require.NoError(t, err)
				var schema schema.PackageSpec
				require.NoError(t, json.Unmarshal(b, &schema))
				schema.Version = "v1.0.0"
				bytes, err := json.Marshal(schema)
				require.NoError(t, err)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			// We need to serve a docs file, since from-url based commands require a doc file.
			case "/docs/_index.md":
				_, err := w.Write([]byte(`---
---`))
				require.NoError(t, err)
			default:
				require.Fail(t, "Unknown URL path %q", r.URL.Path)
			}
		}))
		defer server.Close()

		testMetadata(t, testMetadataArgs{
			providerName:  "command",
			schemaFileURL: server.URL + "/schema.json",
			indexFileURL:  server.URL + "/docs/_index.md",
			assert: func(t *testing.T, metadata, packageDocs string) {
				require.NoError(t, os.Remove(filepath.Join(packageDocs, "_index.md")))
				util.AssertDirsEqual(t, metadataDir, metadata,
					// We fix the time stamp, since we expect URL based lookups to have stable time stamps.
					util.AssertOptionsPreCompareTransform("updated_on: -?[0-9]{5,}", "updated_on: *********"),
					//nolint:lll
					util.AssertOptionsPreCompareTransform("https://raw.githubusercontent.com/pulumi/pulumi-command/v1.0.0/provider/cmd/pulumi-resource-command", server.URL),
				)
				util.AssertDirsEqual(t, packageDocsDir, packageDocs)
			},
		})
	})
}

func TestRepoURLFromRemoteSchema(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/schema.json":
			schema := &schema.Package{
				Name:       "test",
				Version:    ref(semver.MustParse("1.0.0")),
				Repository: "https://github.com/pulumi/pulumi-test",
				Publisher:  "Some Publisher",
				Provider:   &schema.Resource{},
			}
			bytes, err := schema.MarshalJSON()
			require.NoError(t, err)
			_, err = w.Write(bytes)
			require.NoError(t, err)
		case "/docs/_index.md":
			_, err := w.Write([]byte(`---
layout: package
---

# some docs`))
			require.NoError(t, err)
		default:
			assert.Failf(t, "unknown path %s", r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer server.Close()
	testMetadata(t, testMetadataArgs{
		providerName:  "test",
		schemaFileURL: server.URL + "/schema.json",
		indexFileURL:  server.URL + "/docs/_index.md",
		assertOptions: []util.AssertOption{
			util.AssertOptionsPreCompareTransform("[0-9]{5}", "*****"),
		},
	})
}

func TestMissingMetadataFile(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/schema.json":
			schema := &schema.Package{
				Name:       "test",
				Version:    ref(semver.MustParse("1.0.0")),
				Repository: "https://github.com/pulumi/pulumi-test",
				Publisher:  "example",
				Provider:   &schema.Resource{},
			}
			bytes, err := schema.MarshalJSON()
			require.NoError(t, err)
			_, err = w.Write(bytes)
			require.NoError(t, err)
		case "/docs/_index.md":
			_, err := w.Write([]byte(`---
layout: package
---

# some docs`))
			require.NoError(t, err)
		default:
			assert.Failf(t, "unknown path %s", r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer server.Close()
	testMetadata(t, testMetadataArgs{
		providerName:  "test",
		schemaFileURL: server.URL + "/schema.json",
		indexFileURL:  server.URL + "/docs/_index.md",
		assertOptions: []util.AssertOption{
			util.AssertOptionsPreCompareTransform("[0-9]{5}", "*****"),
		},
	})
}

func TestMissingAllDocsFiles(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/schema.json":
			schema := &schema.Package{
				Name:       "test",
				Version:    ref(semver.MustParse("1.0.0")),
				Repository: "https://github.com/example/pulumi-does-not-exist",
				Publisher:  "example",
				Provider:   &schema.Resource{},
			}
			bytes, err := schema.MarshalJSON()
			require.NoError(t, err)
			_, err = w.Write(bytes)
			require.NoError(t, err)
		default:
			assert.Failf(t, "unknown path %s", r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer server.Close()
	testMetadata(t, testMetadataArgs{
		providerName:  "test",
		schemaFileURL: server.URL + "/schema.json",
		errorContains: `missing URL for required file "_index.md"`,
	})
}

func TestMetadataComponentProvider(t *testing.T) {
	t.Parallel()
	testMetadata(t, testMetadataArgs{
		repoSlug:   "pulumi/pulumi-aws-apigateway",
		version:    "v2.6.1",
		schemaFile: "schema.yaml",
	})
}

type testMetadataArgs struct {
	providerName                             string // shared args
	repoSlug, version, schemaFile, publisher string // from-github args
	schemaFileURL, indexFileURL              string // from-url args
	assert                                   func(t *testing.T, metadataDir, pacakgeDocsDir string)
	assertOptions                            []util.AssertOption
	errorContains                            string
}

func defaultAssert(t *testing.T, metadataDir, packageDocsDir string, opts ...util.AssertOption) {
	t.Run("metadata", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, metadataDir, opts...) })
	t.Run("index", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, packageDocsDir, opts...) })
}

func testMetadata(t *testing.T, args testMetadataArgs) {
	t.Helper()
	cmd := PackageMetadataCmd()
	metadataDir := t.TempDir()
	packageDocsDir := t.TempDir()

	var cliArgs []string
	if args.schemaFileURL != "" || args.indexFileURL != "" {
		assert.Empty(t, args.repoSlug)
		assert.Empty(t, args.schemaFile)
		assert.Empty(t, args.version)
		cliArgs = []string{
			"from-urls",
			"--schemaFileURL", args.schemaFileURL,
			"--indexFileURL", args.indexFileURL,
		}
	} else {
		assert.Empty(t, args.schemaFileURL)
		assert.Empty(t, args.indexFileURL)
		cliArgs = []string{
			"from-github",
			"--repoSlug", args.repoSlug,
			"--schemaFile", args.schemaFile,
			"--version", args.version,
			"--publisher", args.publisher,
		}
	}

	cmd.SetArgs(append(cliArgs,
		"--providerName", args.providerName,
		"--metadataDir", metadataDir,
		"--packageDocsDir", packageDocsDir,
	))

	// Capture output into the test
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	defer func() {
		if stdout.Len() > 0 {
			t.Log(stdout.String())
		}
		if stderr.Len() > 0 {
			t.Log(stderr.String())
		}
	}()

	if args.errorContains != "" {
		assert.ErrorContains(t, cmd.Execute(), args.errorContains)
		assert.Nil(t, args.assert,
			"args.assert will not be called when args.errorContains is non-empty, so it should not be set")
		return
	}
	require.NoError(t, cmd.Execute())
	if args.assert != nil {
		args.assert(t, metadataDir, packageDocsDir)
		assert.Nil(t, args.assertOptions, "args.assertOptions are not used when args.assert is non-nil")
	} else {
		defaultAssert(t, metadataDir, packageDocsDir, args.assertOptions...)
	}
}

func ref[T any](t T) *T { return &t }

func TestComputeEditURLFromGitHubUserContentURL(t *testing.T) {
	t.Parallel()
	inputBaseURL := "https://raw.githubusercontent.com/"
	expectedBaseURL := "https://github.com/"

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "pulumi-random main branch",
			input:    inputBaseURL + "pulumi/pulumi-random/v4.16.7/docs/_index.md",
			expected: expectedBaseURL + "pulumi/pulumi-random/blob/v4.16.7/docs/_index.md",
		},
		{
			name:     "non-pulumi repo custom branch",
			input:    inputBaseURL + "not-pulumi/random-repo/mybranch/docs/_index.md",
			expected: expectedBaseURL + "not-pulumi/random-repo/blob/mybranch/docs/_index.md",
		},
		{
			name:     "long file path",
			input:    inputBaseURL + "not-pulumi/random-repo/mybranch/docs/subdir_1/subdir_2/subdir_3/_index.md",
			expected: expectedBaseURL + "not-pulumi/random-repo/blob/mybranch/docs/subdir_1/subdir_2/subdir_3/_index.md",
		},
		{
			name:     "short file path",
			input:    inputBaseURL + "not-pulumi/random-repo/branch/_index.md",
			expected: expectedBaseURL + "not-pulumi/random-repo/blob/branch/_index.md",
		},
		{
			name:     "invalid URL",
			input:    inputBaseURL + "incomplete/url",
			expected: "",
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, computeEditURLFromGitHubUserContentURL(tt.input))
		})
	}
}
