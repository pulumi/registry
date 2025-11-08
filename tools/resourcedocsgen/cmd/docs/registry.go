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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"

	"github.com/golang/glog"

	"github.com/pkg/errors"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/publishers"
	concpool "github.com/sourcegraph/conc/pool"
)

func getRepoSlug(repoURL string) (string, error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", errors.Wrapf(err, "parsing repo url %s", repoURL)
	}

	return u.Path, nil
}

func writeSchemaFile(baseSchemasOutDir, packageName string, schemaBytes []byte) error {
	// Create the package-specific directory if it doesn't exist
	packageDir := filepath.Join(baseSchemasOutDir, packageName)
	if err := os.MkdirAll(packageDir, 0o755); err != nil {
		return errors.Wrapf(err, "creating directory %s", packageDir)
	}

	// Write the schema.json file
	schemaFilePath := filepath.Join(packageDir, "schema.json")
	glog.Infof("Writing schema file to %s", schemaFilePath)
	if err := os.WriteFile(schemaFilePath, schemaBytes, 0o600); err != nil {
		return errors.Wrapf(err, "writing schema file to %s", schemaFilePath)
	}

	return nil
}

func genResourceDocsForPackageFromRegistryMetadata(
	metadata pkg.PackageMeta, docsOutDir, packageTreeJSONOutDir, schemasOutDir string,
) error {
	glog.Infoln("Generating docs for", metadata.Name)

	schemaFileURL, err := getSchemaFileURL(metadata)
	if err != nil {
		return fmt.Errorf("failed to get schema_file_url: %w", err)
	}

	glog.Infoln("Reading remote schema file from registry")
	schemaBytes, err := getSchemaFromRegistry(metadata, schemaFileURL)
	if err != nil {
		if errors.Is(err, ErrPackageNotFound) {
			glog.Infoln(err)
		} else {
			glog.Warningf("Error getting schema from registry: %s", err)
		}

		glog.Infoln("Falling back to reading remote schema file from VCS")
		schemaBytes, err = getSchemaFromVCS(metadata, schemaFileURL)
		if err != nil {
			return fmt.Errorf("getting schema from VCS for %q: %w", metadata.Name, err)
		}
	}

	// Write the schema.json file if schemasOutDir is provided
	if schemasOutDir != "" {
		if err := writeSchemaFile(schemasOutDir, metadata.Name, schemaBytes); err != nil {
			return fmt.Errorf("writing schema file for %q: %w", metadata.Name, err)
		}
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

var ErrPackageNotFound = errors.New("package not found")

func getSchemaFromRegistry(metadata pkg.PackageMeta, schemaURL string) ([]byte, error) {
	backendURL := os.Getenv("PULUMI_BACKEND_URL")
	if backendURL == "" {
		backendURL = "https://api.pulumi.com/api"
	}

	// Check if the schema URL is from OpenTofu registry
	source := "pulumi"
	if schemaURL != "" {
		parsedURL, err := url.Parse(schemaURL)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing schema URL %s", schemaURL)
		}
		if strings.HasPrefix(parsedURL.Path, "/schemas/registry.opentofu.org/") {
			source = "opentofu"
		}
	}

	publisher := publishers.GetPublisherName(metadata.Publisher)
	if publisher == "" {
		return nil, errors.Errorf("publisher %q not found, please add it to publisher-names.json", metadata.Publisher)
	}

	version, err := semver.Parse(strings.TrimPrefix(metadata.Version, "v"))
	if err != nil {
		return nil, errors.Wrapf(err, "parsing version %q", metadata.Version)
	}

	apiURL := fmt.Sprintf("%s/preview/registry/packages/%s/%s/%s/versions/%s",
		backendURL, source, publisher, metadata.Name, version)

	//nolint:gosec // We're constructing the URL based on a predefined pattern.
	metadataResp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving package metadata for %q", apiURL)
	}

	defer contract.IgnoreClose(metadataResp.Body)

	if metadataResp.StatusCode == 404 {
		return nil, fmt.Errorf("%w in registry: %q", ErrPackageNotFound, apiURL)
	}
	if metadataResp.StatusCode >= 400 {
		return nil, errors.Errorf("failed to retrieve package metadata for %q: %s", apiURL, metadataResp.Status)
	}

	var response struct {
		SchemaURL string `json:"schemaURL"`
	}

	if err := json.NewDecoder(metadataResp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	schemaResp, err := http.Get(response.SchemaURL)
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving schema file from %q", response.SchemaURL)
	}
	defer contract.IgnoreClose(schemaResp.Body)

	if schemaResp.StatusCode >= 400 {
		return nil, errors.Errorf("failed to retrieve schema file from %q: %s", response.SchemaURL, schemaResp.Status)
	}

	return io.ReadAll(schemaResp.Body)
}

func getSchemaFromVCS(metadata pkg.PackageMeta, schemaFileURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", schemaFileURL, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "creating request for %q", schemaFileURL)
	}

	pkg.AddGitHubAuthHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "reading schema file from VCS %s", schemaFileURL)
	}

	defer contract.IgnoreClose(resp.Body)

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		// ignore error, we'll just return the status code in that case
		if err != nil {
			return nil, fmt.Errorf("failed to get schema file from VCS %s: %s", schemaFileURL, resp.Status)
		}
		return nil, fmt.Errorf("failed to get schema file from VCS %s: %s\n%s", schemaFileURL, resp.Status, string(bodyBytes))
	}

	schemaBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading response body from %s", schemaFileURL)
	}

	return schemaBytes, nil
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

func getRegistryPackagesPath(repoPath string) string {
	return filepath.Join(repoPath, "themes", "default", "data", "registry", "packages")
}

func genResourceDocsForAllRegistryPackages(
	registryRepoPath, baseDocsOutDir, basePackageTreeJSONOutDir, baseSchemasOutDir string,
) error {
	registryPackagesPath := getRegistryPackagesPath(registryRepoPath)
	metadataFiles, err := os.ReadDir(registryPackagesPath)
	if err != nil {
		return errors.Wrap(err, "reading the registry packages dir")
	}

	pool := concpool.New().WithErrors().WithMaxGoroutines(runtime.NumCPU())
	for _, f := range metadataFiles {
		f := f
		pool.Go(func() error {
			glog.Infof("=== starting %s ===\n", f.Name())
			glog.Infoln("Processing metadata file")
			metadataFilePath := filepath.Join(registryPackagesPath, f.Name())

			b, err := os.ReadFile(metadataFilePath)
			if err != nil {
				return errors.Wrapf(err, "reading the metadata file %s", metadataFilePath)
			}

			var metadata pkg.PackageMeta
			if err := yaml.Unmarshal(b, &metadata); err != nil {
				return errors.Wrapf(err, "unmarshalling the metadata file %s", metadataFilePath)
			}

			docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
			err = genResourceDocsForPackageFromRegistryMetadata(
				metadata, docsOutDir, basePackageTreeJSONOutDir, baseSchemasOutDir,
			)
			if err != nil {
				return errors.Wrapf(err, "generating resource docs using metadata file info %s", f.Name())
			}

			glog.Infof("=== completed %s ===", f.Name())
			return nil
		})
	}

	return pool.Wait()
}

func resourceDocsFromRegistryCmd() *cobra.Command {
	var baseDocsOutDir string
	var basePackageTreeJSONOutDir string
	var baseSchemasOutDir string
	var registryDir string

	cmd := &cobra.Command{
		Use:   "registry [pkgName]",
		Short: "Generate resource docs for a package from the registry",
		Long: "Generate resource docs for all packages in the registry or specific packages. " +
			"Pass a package name in the registry as an optional arg to generate docs only for that package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			registryDir, err := filepath.Abs(registryDir)
			if err != nil {
				return errors.Wrap(err, "finding the cwd")
			}

			if len(args) > 0 {
				glog.Infoln("Generating docs for a single package:", args[0])
				registryPackagesPath := getRegistryPackagesPath(registryDir)
				pkgName := args[0]
				metadataFilePath := filepath.Join(registryPackagesPath, pkgName+".yaml")
				b, err := os.ReadFile(metadataFilePath)
				if err != nil {
					return errors.Wrapf(err, "reading the metadata file %s", metadataFilePath)
				}

				var metadata pkg.PackageMeta
				if err := yaml.Unmarshal(b, &metadata); err != nil {
					return errors.Wrapf(err, "unmarshalling the metadata file %s", metadataFilePath)
				}

				docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")

				err = genResourceDocsForPackageFromRegistryMetadata(
					metadata, docsOutDir, basePackageTreeJSONOutDir, baseSchemasOutDir,
				)
				if err != nil {
					return errors.Wrapf(err, "generating docs for package %q from registry metadata", pkgName)
				}
			} else {
				glog.Infoln("Generating docs for all packages in the registry...")
				err := genResourceDocsForAllRegistryPackages(
					registryDir, baseDocsOutDir, basePackageTreeJSONOutDir, baseSchemasOutDir,
				)
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
	cmd.Flags().StringVar(&baseSchemasOutDir, "baseSchemasOutDir",
		"",
		"The directory path to write the schema.json files to (optional)")
	cmd.Flags().StringVar(&registryDir, "registryDir",
		".",
		"The root of the pulumi/registry directory")

	return cmd
}
