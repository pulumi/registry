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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"

	"github.com/golang/glog"

	"github.com/pkg/errors"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/registry/svc"
	concpool "github.com/sourcegraph/conc/pool"
)

func getRepoSlug(repoURL string) (string, error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", errors.Wrapf(err, "parsing repo url %s", repoURL)
	}

	return u.Path, nil
}

func addGitHubAuthHeaders(req *http.Request) {
	// Check if the request is for GitHub domains
	host := req.URL.Host
	if host == "github.com" || host == "api.github.com" || host == "raw.githubusercontent.com" {
		// Add GitHub token from environment variable if available
		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			req.Header.Add("Authorization", "Bearer "+token)
		}
	}
}

func genResourceDocsForPackageFromRegistryMetadata(
	metadata pkg.PackageMeta, docsOutDir, packageTreeJSONOutDir string,
) error {
	glog.Infoln("Generating docs for", metadata.Name)

	schemaFileURL, err := getSchemaFileURL(metadata)
	if err != nil {
		return fmt.Errorf("failed to get schema_file_url: %w", err)
	}
	glog.Infoln("Reading remote schema file from VCS")

	req, err := http.NewRequest("GET", schemaFileURL, nil)
	if err != nil {
		return errors.Wrapf(err, "creating request for %s", schemaFileURL)
	}

	addGitHubAuthHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "reading schema file from VCS %s", schemaFileURL)
	}

	defer contract.IgnoreClose(resp.Body)

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		// ignore error, we'll just return the status code in that case
		if err != nil {
			return fmt.Errorf("failed to get schema file from VCS %s: %s", schemaFileURL, resp.Status)
		}
		return fmt.Errorf("failed to get schema file from VCS %s: %s\n%s", schemaFileURL, resp.Status, string(bodyBytes))
	}

	schemaBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "reading response body from %s", schemaFileURL)
	}

	var mainSpec pschema.PackageSpec

	switch {
	// The source schema can be in YAML format.
	case strings.HasSuffix(schemaFileURL, ".yaml"):
		if err := yaml.Unmarshal(schemaBytes, &mainSpec); err != nil {
			return errors.Wrap(err, "unmarshalling YAML schema into PackageSpec")
		}
	// If we don't have another format, assume JSON.
	default:
		if err := json.Unmarshal(schemaBytes, &mainSpec); err != nil {
			return errors.Wrap(err, "unmarshalling JSON schema into PackageSpec")
		}
	}

	pulPkg, genctx, err := getPulumiPackageFromSchema(docsOutDir, mainSpec)
	if err != nil {
		return errors.Wrap(err, "generating package from schema file")
	}

	glog.Infoln("Running docs generator...")
	if err := generateDocsFromSchema(docsOutDir, genctx); err != nil {
		return errors.Wrap(err, "generating docs from schema")
	}

	glog.Infoln("Generating the package tree JSON file...")
	if err := generatePackageTree(packageTreeJSONOutDir, pulPkg.Name, genctx); err != nil {
		return errors.Wrap(err, "generating package tree")
	}

	return nil
}

func getSchemaFileURL(metadata pkg.PackageMeta) (string, error) {
	if metadata.SchemaFileURL != "" {
		return metadata.SchemaFileURL, nil
	}

	// We don't have an explicit SchemaFileURL, so migrate from SchemaFilePath.

	if metadata.RepoURL == "" {
		return "", errors.Errorf("metadata for package %q does not contain the repo_url", metadata.Name)
	}

	schemaFilePath := fmt.Sprintf(defaultSchemaFilePathFormat, metadata.Name)
	if p := metadata.SchemaFilePath; p != "" { //nolint:staticcheck
		schemaFilePath = p
	}

	// Make sure the schema file path does not have a leading slash.
	// We'll add in the URL format below. It's easier to read that way.
	schemaFilePath = strings.TrimPrefix(schemaFilePath, "/")

	repoSlug, err := getRepoSlug(metadata.RepoURL)
	if err != nil {
		return "", errors.WithMessage(err, "could not get repo slug")
	}

	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repoSlug, metadata.Version, schemaFilePath), nil
}

func genResourceDocsForAllRegistryPackages(
	ctx context.Context,
	provider svc.PackageMetadataProvider,
	baseDocsOutDir, basePackageTreeJSONOutDir string,
) error {
	metadataList, err := provider.ListPackageMetadata(ctx)
	if err != nil {
		return errors.Wrap(err, "listing package metadata")
	}

	pool := concpool.New().WithErrors().WithMaxGoroutines(runtime.NumCPU())
	for _, metadata := range metadataList {
		pool.Go(func() error {
			glog.Infof("=== starting %s ===\n", metadata.Name)
			docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
			err = genResourceDocsForPackageFromRegistryMetadata(metadata, docsOutDir, basePackageTreeJSONOutDir)
			if err != nil {
				return errors.Wrapf(err, "generating resource docs using metadata file info %s", metadata.Name)
			}

			glog.Infof("=== completed %s ===", metadata.Name)
			return nil
		})
	}
	return pool.Wait()
}

func resourceDocsFromRegistryCmd() *cobra.Command {
	var baseDocsOutDir string
	var basePackageTreeJSONOutDir string
	var registryDir string
	var useAPI bool
	var apiURL string

	cmd := &cobra.Command{
		Use:   "registry [pkgName]",
		Short: "Generate resource docs for a package from the registry",
		Long: "Generate resource docs for all packages in the registry or specific packages. " +
			"Pass a package name in the registry as an optional arg to generate docs only for that package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var provider svc.PackageMetadataProvider
			if useAPI {
				provider = svc.NewAPIProvider(apiURL)
			} else {
				provider = svc.NewFileSystemProvider(registryDir)
			}

			if len(args) > 0 {
				glog.Infoln("Generating docs for a single package:", args[0])
				metadata, err := provider.GetPackageMetadata(ctx, args[0])
				if err != nil {
					return errors.Wrapf(err, "getting metadata for package %q", args[0])
				}

				docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
				err = genResourceDocsForPackageFromRegistryMetadata(metadata, docsOutDir, basePackageTreeJSONOutDir)
				if err != nil {
					return errors.Wrapf(err, "generating docs for package %q from registry metadata", args[0])
				}
			} else {
				glog.Infoln("Generating docs for all packages in the registry...")
				err := genResourceDocsForAllRegistryPackages(ctx, provider, baseDocsOutDir, basePackageTreeJSONOutDir)
				if err != nil {
					return errors.Wrap(err, "generating docs for all packages from registry metadata")
				}
			}

			glog.Infoln("Done!")
			return nil
		},
	}

	cmd.Flags().StringVar(&baseDocsOutDir, "baseDocsOutDir",
		"../../content/registry/packages",
		"The directory path to where the docs will be written to")
	cmd.Flags().StringVar(&basePackageTreeJSONOutDir, "basePackageTreeJSONOutDir",
		"../../static/registry/packages/navs",
		"The directory path to write the package tree JSON file to")
	cmd.Flags().StringVar(&registryDir, "registryDir",
		".",
		"The root of the pulumi/registry directory")
	cmd.Flags().BoolVar(&useAPI, "use-api", false, "Use the Pulumi Registry API instead of local files")
	cmd.Flags().StringVar(&apiURL, "api-url",
		"https://api.pulumi.com/api/preview/registry",
		"URL of the Pulumi Registry API")

	return cmd
}
