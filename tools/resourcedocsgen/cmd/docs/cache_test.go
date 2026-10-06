// Copyright 2026, Pulumi Corporation.
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

package docs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/pulumi/registry/tools/resourcedocsgen/internal/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const cacheTestSchemaFormat = `{
  "name": "cachetest",
  "version": "9.9.9",
  "resources": {
    "cachetest:index:Thing": {
      "description": %q,
      "properties": {"name": {"type": "string"}},
      "inputProperties": {"name": {"type": "string"}}
    }
  }
}`

// schemaServer serves a cachetest schema whose resource description the test can change,
// and counts the requests it receives.
type schemaServer struct {
	*httptest.Server
	mu          sync.Mutex
	description string
	requests    atomic.Int32
}

func newSchemaServer(t *testing.T, description string) *schemaServer {
	s := &schemaServer{description: description}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		s.requests.Add(1)
		s.mu.Lock()
		defer s.mu.Unlock()
		fmt.Fprintf(w, cacheTestSchemaFormat, s.description)
	}))
	t.Cleanup(s.Close)
	return s
}

func (s *schemaServer) setDescription(description string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.description = description
}

// runCacheTestGeneration runs `docs registry cachetest` against a package YAML pointing
// at schemaURL and returns the generated resource page.
func runCacheTestGeneration(t *testing.T, registryDir, docsOutDir, navOutDir, schemaURL string) string {
	// The invalid publisher makes the registry API lookup fail before any request,
	// so the schema can only come from schema_file_url.
	util.WriteFile(t,
		filepath.Join(registryDir, "themes", "default", "data", "registry", "packages", "cachetest.yaml"),
		`name: cachetest
title: cachetest
publisher: "contains-invalid-chars!!!"
repo_url: https://github.com/example/pulumi-cachetest
version: v9.9.9
schema_file_url: `+schemaURL+"\n")

	cmd := resourceDocsFromRegistryCmd(http.DefaultClient)
	cmd.SetArgs([]string{
		"cachetest",
		"--baseDocsOutDir", docsOutDir,
		"--basePackageTreeJSONOutDir", navOutDir,
		"--registryDir", registryDir,
	})
	require.NoError(t, cmd.Execute())
	return util.ReadFile(t, filepath.Join(docsOutDir, "cachetest", "api-docs", "thing", "_index.md"))
}

// A schema_file_url that points at a branch can change while the package YAML doesn't.
// The second run must pick up the new schema rather than reuse the cached output.
func TestCacheMutableSchemaURL(t *testing.T) {
	t.Parallel()
	server := newSchemaServer(t, "Description A.")
	schemaURL := server.URL + "/my-branch/schema.json"
	registryDir, docsOutDir, navOutDir := t.TempDir(), t.TempDir(), t.TempDir()

	page := runCacheTestGeneration(t, registryDir, docsOutDir, navOutDir, schemaURL)
	require.Contains(t, page, "Description A.")

	server.setDescription("Description B.")
	page = runCacheTestGeneration(t, registryDir, docsOutDir, navOutDir, schemaURL)
	assert.True(t, strings.Contains(page, "Description B."), "second run reused output generated from the old schema")
	assert.Equal(t, int32(2), server.requests.Load())

	// With the schema unchanged, a third run downloads it again but skips generation.
	page = runCacheTestGeneration(t, registryDir, docsOutDir, navOutDir, schemaURL)
	assert.True(t, strings.Contains(page, "Description B."))
	assert.Equal(t, int32(3), server.requests.Load())
}

// A schema_file_url pinned to the package's version is immutable, so a rerun with
// unchanged YAML must not touch the network at all.
func TestCachePinnedSchemaURLSkipsFetch(t *testing.T) {
	t.Parallel()
	server := newSchemaServer(t, "Description A.")
	schemaURL := server.URL + "/v9.9.9/schema.json"
	registryDir, docsOutDir, navOutDir := t.TempDir(), t.TempDir(), t.TempDir()

	runCacheTestGeneration(t, registryDir, docsOutDir, navOutDir, schemaURL)
	require.Equal(t, int32(1), server.requests.Load())

	runCacheTestGeneration(t, registryDir, docsOutDir, navOutDir, schemaURL)
	assert.Equal(t, int32(1), server.requests.Load(), "pinned schema was fetched again")
}

func TestIsPinnedSchemaURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, url, version string
		want               bool
	}{
		{
			"github tag", "https://raw.githubusercontent.com/pulumi/pulumi-aiven/v6.60.0/provider/cmd/pulumi-resource-aiven/schema.json", //nolint:lll
			"v6.60.0", true,
		},
		{
			"tag without v", "https://raw.githubusercontent.com/example/pulumi-x/1.2.3/schema.json",
			"v1.2.3", true,
		},
		{
			"opentofu", "https://djoiyj6oj2oxz.cloudfront.net/schemas/registry.opentofu.org/airbytehq/airbyte/1.3.0/schema.json",
			"v1.3.0", true,
		},
		{
			"commit sha", "https://raw.githubusercontent.com/pulumi/pulumi-hcl/86982a86c75cd97920e38080c230997bbdc20348/registry/schema.json", //nolint:lll
			"v0.18.4", true,
		},
		{
			"branch", "https://raw.githubusercontent.com/example/pulumi-x/my-branch/provider/cmd/pulumi-resource-x/schema.json",
			"v9.9.9", false,
		},
		{
			"other version in path", "https://raw.githubusercontent.com/example/pulumi-x/v1.0.0/schema.json",
			"v9.9.9", false,
		},
		{"s3 without version", "https://example-bucket.s3.amazonaws.com/x/schema.json", "v9.9.9", false},
		{"empty version", "https://example.com/v/schema.json", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, isPinnedSchemaURL(tt.url, tt.version))
		})
	}
}
