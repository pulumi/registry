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
	"testing"

	"github.com/pulumi/registry/tools/resourcedocsgen/internal/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataBridgedProvider(t *testing.T) {
	t.Parallel()
	testMetadata(t, testMetadataArgs{
		repoSlug:   "pulumi/pulumi-random",
		version:    "v4.16.7",
		schemaFile: "provider/cmd/pulumi-resource-random/schema.json",
	})
}

func TestMetadataNativeProvider(t *testing.T) {
	t.Parallel()

	var metadataDir, pacakgeDocsDir string
	testMetadata(t, testMetadataArgs{
		repoSlug:   "pulumi/pulumi-command",
		version:    "v1.0.0",
		schemaFile: "provider/cmd/pulumi-resource-command/schema.json",
		assert: func(t *testing.T, metadata, pacakgeDocs string) {
			defaultAssert(t, metadata, pacakgeDocs)
			metadataDir = metadata
			pacakgeDocsDir = pacakgeDocs
		},
	})

	t.Run("test-remote-equivalence", func(t *testing.T) {
		t.Parallel()
		testMetadata(t, testMetadataArgs{
			repoSlug: "pulumi/pulumi-command",
			version:  "v1.0.0",
			schemaFileURL: "https://raw.githubusercontent.com/pulumi/pulumi-command/" +
				"v1.0.0/provider/cmd/pulumi-resource-command/schema.json",
			indexFileURL: "https://raw.githubusercontent.com/pulumi/pulumi-command/" +
				"v1.0.0/docs/_index.md",
			assert: func(t *testing.T, metadata, pacakgeDocs string) {
				util.AssertDirsEqual(t, metadataDir, metadata,
					// We fix the time stamp, since we expect URL based lookups to have stable time stamps.
					util.AssertOptionsPreCompareTransform("updated_on: [0-9]+", "updated_on: 1719590084"))
				util.AssertDirsEqual(t, pacakgeDocsDir, pacakgeDocs)
			},
		})
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
	repoSlug, version           string
	schemaFile                  string
	schemaFileURL, indexFileURL string

	// assert is a custom assert on a successfully run.
	//
	// If no assert is provided, then [defaultAssert] is used.
	assert func(t *testing.T, metadataDir, pacakgeDocsDir string)
	// Options to give to [defaultAssert].
	//
	// Cannot be provided with assert, since assert overrides the default assert.
	assertOptions []util.AssertOption

	// Assert that the command fails, and that the error contains errorContains.
	errorContains string
}

func defaultAssert(t *testing.T, metadataDir, pacakgeDocsDir string, opts ...util.AssertOption) {
	t.Run("metadata", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, metadataDir, opts...) })
	t.Run("index", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, pacakgeDocsDir, opts...) })
}

func testMetadata(t *testing.T, args testMetadataArgs) {
	cmd := PackageMetadataCmd()
	metadataDir := t.TempDir()
	pacakgeDocsDir := t.TempDir()
	cmd.SetArgs([]string{
		"--repoSlug", args.repoSlug,
		"--schemaFile", args.schemaFile,
		"--schemaFileURL", args.schemaFileURL,
		"--indexFileURL", args.indexFileURL,
		"--version", args.version,
		"--metadataDir", metadataDir,
		"--packageDocsDir", pacakgeDocsDir,
	})

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
		args.assert(t, metadataDir, pacakgeDocsDir)
		assert.Nil(t, args.assertOptions, "args.assertOptions are not used when args.assert is non-nil")
	} else {
		defaultAssert(t, metadataDir, pacakgeDocsDir, args.assertOptions...)
	}
}
