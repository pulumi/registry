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

package docs

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/registry/tools/resourcedocsgen/internal/tests/util"
	"github.com/stretchr/testify/require"
)

func TestGenerateDocsSinglePackage(t *testing.T) {
	t.Parallel()
	const partialMetadata = `category: Utility
component: false
description: A Pulumi package to safely use randomness in Pulumi programs.
featured: false
logo_url: ""
name: random
native: false
package_status: ga
publisher: Pulumi
repo_url: https://github.com/pulumi/pulumi-random
title: random
updated_on: 1729058626
version: v4.18.2`

	t.Run("registry", func(t *testing.T) {
		t.Parallel()
		registryDir := t.TempDir()

		// Use a wrong schema_file_path to ensure the schema can only be loaded via the Registry API.
		util.WriteFile(t,
			filepath.Join(registryDir, "themes", "default", "data", "registry", "packages", "random.yaml"),
			partialMetadata+`
schema_file_path: does/not/exist/schema.json`)
		basePackageTreeJSONOutDir := t.TempDir()
		baseDocsOutDir := t.TempDir()

		cmd := resourceDocsFromRegistryCmd()
		cmd.SetArgs([]string{
			"random", /* pkgName */
			"--baseDocsOutDir", baseDocsOutDir,
			"--basePackageTreeJSONOutDir", basePackageTreeJSONOutDir,
			"--registryDir", registryDir,
		})
		require.NoError(t, cmd.Execute())
		t.Run("docs", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, baseDocsOutDir) })
		t.Run("tree", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, basePackageTreeJSONOutDir) })
	})

	t.Run("vcs", func(t *testing.T) {
		t.Parallel()
		registryDir := t.TempDir()

		// Use a publisher name containing invalid characters to ensure the schema can only be loaded via the VCS
		util.WriteFile(t,
			filepath.Join(registryDir, "themes", "default", "data", "registry", "packages", "random.yaml"),
			partialMetadata+`
schema_file_path: provider/cmd/pulumi-resource-random/schema.json
publisher: "contains-invalid-chars!!!"
`)
		basePackageTreeJSONOutDir := t.TempDir()
		baseDocsOutDir := t.TempDir()

		cmd := resourceDocsFromRegistryCmd()
		cmd.SetArgs([]string{
			"random", /* pkgName */
			"--baseDocsOutDir", baseDocsOutDir,
			"--basePackageTreeJSONOutDir", basePackageTreeJSONOutDir,
			"--registryDir", registryDir,
		})
		require.NoError(t, cmd.Execute())
		t.Run("docs", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, baseDocsOutDir) })
		t.Run("tree", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, basePackageTreeJSONOutDir) })
		t.Run("check-schema-file-url", func(t *testing.T) {
			t.Parallel()
			registryDirV2 := t.TempDir()
			util.WriteFile(t,
				filepath.Join(registryDirV2, "themes", "default", "data", "registry", "packages", "random.yaml"),
				partialMetadata+`
publisher: "contains-invalid-chars!!!"
schema_file_url: https://raw.githubusercontent.com/pulumi/pulumi-random/v4.18.2/provider/cmd/pulumi-resource-random/schema.json`) //nolint:lll

			basePackageTreeJSONOutDirV2 := t.TempDir()
			baseDocsOutDirV2 := t.TempDir()

			cmd := resourceDocsFromRegistryCmd()
			cmd.SetArgs([]string{
				"random", /* pkgName */
				"--baseDocsOutDir", baseDocsOutDirV2,
				"--basePackageTreeJSONOutDir", basePackageTreeJSONOutDirV2,
				"--registryDir", registryDirV2,
			})
			require.NoError(t, cmd.Execute())
			util.AssertDirsEqual(t, baseDocsOutDir, baseDocsOutDirV2)
			util.AssertDirsEqual(t, basePackageTreeJSONOutDir, basePackageTreeJSONOutDirV2)
		})
	})
}

func TestGenerateDocsAllPackage(t *testing.T) {
	t.Parallel()

	packageMetadata := map[string]string{
		"random": `category: Utility
component: false
description: A Pulumi package to safely use randomness in Pulumi programs.
featured: false
logo_url: ""
name: random
native: false
package_status: ga
publisher: Pulumi
repo_url: https://github.com/pulumi/pulumi-random
schema_file_path: provider/cmd/pulumi-resource-random/schema.json
title: random
updated_on: 1729058626
version: v4.18.2`,
		"aws-apigateway": `category: Cloud
component: true
description: Pulumi Amazon Web Services (AWS) API Gateway Components.
featured: false
logo_url: https://raw.githubusercontent.com/pulumi/pulumi-aws-apigateway/main/assets/logo.png
name: aws-apigateway
native: false
package_status: ga
publisher: Pulumi
repo_url: https://github.com/pulumi/pulumi-aws-apigateway
schema_file_path: schema.yaml
title: AWS API Gateway
updated_on: 1727876192
version: v2.6.2`,
	}

	t.Run("registry", func(t *testing.T) {
		t.Parallel()
		registryDir := t.TempDir()
		for pkgName, partialMetadata := range packageMetadata {
			// Use a wrong schema_file_path to ensure the schema can only be loaded via the Registry API.
			util.WriteFile(t,
				filepath.Join(registryDir, "themes", "default", "data", "registry", "packages", pkgName+".yaml"),
				partialMetadata+`
schema_file_path: does/not/exist/schema.json`)
		}

		basePackageTreeJSONOutDir := t.TempDir()
		baseDocsOutDir := t.TempDir()

		cmd := resourceDocsFromRegistryCmd()
		cmd.SetArgs([]string{
			"--baseDocsOutDir", baseDocsOutDir,
			"--basePackageTreeJSONOutDir", basePackageTreeJSONOutDir,
			"--registryDir", registryDir,
		})
		require.NoError(t, cmd.Execute())
		t.Run("docs", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, baseDocsOutDir) })
		t.Run("tree", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, basePackageTreeJSONOutDir) })
	})

	t.Run("vcs", func(t *testing.T) {
		t.Parallel()
		registryDir := t.TempDir()
		for pkgName, partialMetadata := range packageMetadata {
			// Use a publisher name containing invalid characters to ensure the schema can only be loaded via the VCS
			util.WriteFile(t,
				filepath.Join(registryDir, "themes", "default", "data", "registry", "packages", pkgName+".yaml"),
				partialMetadata+`
publisher: "contains-invalid-chars!!!"`)
		}

		basePackageTreeJSONOutDir := t.TempDir()
		baseDocsOutDir := t.TempDir()

		cmd := resourceDocsFromRegistryCmd()
		cmd.SetArgs([]string{
			"--baseDocsOutDir", baseDocsOutDir,
			"--basePackageTreeJSONOutDir", basePackageTreeJSONOutDir,
			"--registryDir", registryDir,
		})
		require.NoError(t, cmd.Execute())
		t.Run("docs", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, baseDocsOutDir) })
		t.Run("tree", func(t *testing.T) { t.Parallel(); util.AssertDirEqual(t, basePackageTreeJSONOutDir) })
	})
}
