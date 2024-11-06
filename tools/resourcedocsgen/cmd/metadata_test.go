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
	"os"
	"path/filepath"
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest //PackageMetadataCmd relies on global state.
func TestMetadataBridgedProvider(t *testing.T) {
	testMetadata(t, testMetadataArgs{
		providerName: "random",
		version:      "v4.16.7",
		schemaFile:   "provider/cmd/pulumi-resource-random/schema.json",
		checkIndex:   true,
	})
}

//nolint:paralleltest //PackageMetadataCmd relies on global state.
func TestMetadataNativeProvider(t *testing.T) {
	testMetadata(t, testMetadataArgs{
		providerName: "command",
		version:      "v1.0.0",
		schemaFile:   "provider/cmd/pulumi-resource-command/schema.json",
	})
}

//nolint:paralleltest // PackageMetadataCmd relies on global state.
func TestMetadataComponentProvider(t *testing.T) {
	testMetadata(t, testMetadataArgs{
		providerName: "aws-apigateway",
		version:      "v2.6.1",
		schemaFile:   "schema.yaml",
	})
}

type testMetadataArgs struct {
	providerName, version string
	schemaFile            string
	checkIndex            bool
}

func testMetadata(t *testing.T, args testMetadataArgs) {
	cmd := PackageMetadataCmd()
	metadataDir := t.TempDir()
	pacakgeDocsDir := t.TempDir()
	cmd.SetArgs([]string{
		"--repoSlug", "pulumi/pulumi-" + args.providerName,
		"--schemaFile", args.schemaFile,
		"--version", args.version,
		"--metadataDir", metadataDir,
		"--packageDocsDir", pacakgeDocsDir,
	})
	require.NoError(t, cmd.Execute())
	t.Run(args.providerName+".yaml", func(t *testing.T) {
		t.Parallel()
		autogold.ExpectFile(t, autogold.Raw(
			readFile(t, filepath.Join(metadataDir, args.providerName+".yaml")),
		))
	})
	if args.checkIndex {
		t.Run("_index.md", func(t *testing.T) {
			t.Parallel()
			autogold.ExpectFile(t, autogold.Raw(
				readFile(t, filepath.Join(pacakgeDocsDir, "_index.md")),
			))
		})
	}
}

func readFile(t *testing.T, path string) string {
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(b)
}
