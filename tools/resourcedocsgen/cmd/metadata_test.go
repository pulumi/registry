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
	"testing"

	"github.com/pulumi/registry/tools/resourcedocsgen/internal/tests/util"
	"github.com/stretchr/testify/require"
)

func TestMetadataBridgedProvider(t *testing.T) {
	t.Parallel()
	testMetadata(t, testMetadataArgs{
		providerName: "random",
		version:      "v4.16.7",
		schemaFile:   "provider/cmd/pulumi-resource-random/schema.json",
	})
}

func TestMetadataNativeProvider(t *testing.T) {
	t.Parallel()

	var metadataDir, pacakgeDocsDir string
	testMetadata(t, testMetadataArgs{
		providerName: "command",
		version:      "v1.0.0",
		schemaFile:   "provider/cmd/pulumi-resource-command/schema.json",
		assert: func(t *testing.T, metadata, pacakgeDocs string) {
			defaultAssert(t, metadata, pacakgeDocs)
			metadataDir = metadata
			pacakgeDocsDir = pacakgeDocs
		},
	})

	t.Run("test-remote-equivalence", func(t *testing.T) {
		t.Parallel()
		testMetadata(t, testMetadataArgs{
			providerName: "command",
			version:      "v1.0.0",
			schemaFileURL: "https://raw.githubusercontent.com/pulumi/pulumi-command/" +
				"v1.0.0/provider/cmd/pulumi-resource-command/schema.json",
			indexFileURL: "https://raw.githubusercontent.com/pulumi/pulumi-command/" +
				"v1.0.0/docs/_index.md",
			assert: func(t *testing.T, metadata, pacakgeDocs string) {
				util.AssertDirsEqual(t, metadataDir, metadata,
					// We fix the time stamp, since we expect URL based lookups to have stable time stamps.
					util.OptionAssertDirsEqualReplace("updated_on: [0-9]+", "updated_on: 1719590084"))
				util.AssertDirsEqual(t, pacakgeDocsDir, pacakgeDocs)
			},
		})
	})
}

func TestMetadataComponentProvider(t *testing.T) {
	t.Parallel()
	testMetadata(t, testMetadataArgs{
		providerName: "aws-apigateway",
		version:      "v2.6.1",
		schemaFile:   "schema.yaml",
	})
}

type testMetadataArgs struct {
	providerName, version       string
	schemaFile                  string
	schemaFileURL, indexFileURL string
	assert                      func(t *testing.T, metadataDir, pacakgeDocsDir string)
}

func defaultAssert(t *testing.T, metadataDir, pacakgeDocsDir string) {
	t.Run("metadata", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, metadataDir) })
	t.Run("index", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, pacakgeDocsDir) })
}

func testMetadata(t *testing.T, args testMetadataArgs) {
	cmd := PackageMetadataCmd()
	metadataDir := t.TempDir()
	pacakgeDocsDir := t.TempDir()
	cmd.SetArgs([]string{
		"--repoSlug", "pulumi/pulumi-" + args.providerName,
		"--schemaFile", args.schemaFile,
		"--schemaFileURL", args.schemaFileURL,
		"--indexFileURL", args.indexFileURL,
		"--version", args.version,
		"--metadataDir", metadataDir,
		"--packageDocsDir", pacakgeDocsDir,
	})
	require.NoError(t, cmd.Execute())
	if args.assert != nil {
		args.assert(t, metadataDir, pacakgeDocsDir)
	} else {
		defaultAssert(t, metadataDir, pacakgeDocsDir)
	}
}
