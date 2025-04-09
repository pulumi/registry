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
	"time"

	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"

	"github.com/golang/glog"

	"github.com/pkg/errors"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	concpool "github.com/sourcegraph/conc/pool"
)

// PackageMetadataProvider is an interface for providers that retrieve package metadata
// from either the filesystem directory or the Pulumi Registry API.
type PackageMetadataProvider interface {
	// GetPackageMetadata returns metadata for a specific package
	GetPackageMetadata(pkgName string) (*pkg.PackageMeta, error)
	// ListPackageMetadata returns metadata for all packages
	ListPackageMetadata() ([]*pkg.PackageMeta, error)
}

// FileSystemProvider implements PackageMetadataProvider using the local yaml data files
// in the pulumi/registry repository.
type fileSystemProvider struct {
	registryDir string
}

// RegistryAPIProvider implements PackageMetadataProvider using the Pulumi API
// to retrieve package metadata.
type registryAPIProvider struct {
	apiURL string
}

// PackageMetadata represents the API response structure for package metadata
// from the Pulumi Registry API.
// TODO: import this from the pulumi-service if possible.
type PackageMetadata struct {
	Name          string    `json:"name"`
	Publisher     string    `json:"publisher"`
	Source        string    `json:"source"`
	Version       string    `json:"version"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	LogoURL       string    `json:"logoUrl,omitempty"`
	RepoURL       string    `json:"repoUrl,omitempty"`
	Category      string    `json:"category,omitempty"`
	IsFeatured    bool      `json:"isFeatured"`
	PackageTypes  []string  `json:"packageTypes,omitempty"`
	PackageStatus string    `json:"packageStatus"`
	SchemaURL     string    `json:"schemaURL"`
	CreatedAt     time.Time `json:"createdAt"`
}

// PackageListResponse represents the API response structure for package lists
type PackageListResponse struct {
	Packages []PackageMetadata `json:"packages"`
}

// NewFileSystemProvider creates a new PackageMetadataProvider that reads from the filesystem
func NewFileSystemProvider(registryDir string) PackageMetadataProvider {
	return &fileSystemProvider{
		registryDir: registryDir,
	}
}

// NewAPIProvider creates a new PackageMetadataProvider that reads from the Pulumi API
func NewAPIProvider(apiURL string) PackageMetadataProvider {
	return &registryAPIProvider{
		apiURL: apiURL,
	}
}

// contains checks if a string is in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getRepoSlug(repoURL string) (string, error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", errors.Wrapf(err, "parsing repo url %s", repoURL)
	}

	return u.Path, nil
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

	resp, err := http.Get(schemaFileURL) //nolint:gosec
	if err != nil {
		return errors.Wrapf(err, "reading schema file from VCS %s", schemaFileURL)
	}

	defer contract.IgnoreClose(resp.Body)

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

func getRegistryPackagesPath(repoPath string) string {
	return filepath.Join(repoPath, "themes", "default", "data", "registry", "packages")
}

func genResourceDocsForAllRegistryPackages(
	provider PackageMetadataProvider,
	baseDocsOutDir, basePackageTreeJSONOutDir string,
) error {
	metadataList, err := provider.ListPackageMetadata()
	if err != nil {
		return errors.Wrap(err, "listing package metadata")
	}

	pool := concpool.New().WithErrors().WithMaxGoroutines(runtime.NumCPU())
	for _, metadata := range metadataList {
		metadata := metadata
		pool.Go(func() error {
			glog.Infof("=== starting %s ===\n", metadata.Name)
			docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
			err = genResourceDocsForPackageFromRegistryMetadata(*metadata, docsOutDir, basePackageTreeJSONOutDir)
			if err != nil {
				return errors.Wrapf(err, "generating resource docs using metadata file info %s", metadata.Name)
			}

			glog.Infof("=== completed %s ===", metadata.Name)
			return nil
		})
	}
	return pool.Wait()
}

func convertAPIPackageToPackageMeta(apiPkg PackageMetadata) (*pkg.PackageMeta, error) {
	return &pkg.PackageMeta{
		Name:          apiPkg.Name,
		Publisher:     apiPkg.Publisher,
		Description:   apiPkg.Description,
		LogoURL:       apiPkg.LogoURL,
		RepoURL:       apiPkg.RepoURL,
		Category:      pkg.PackageCategory(apiPkg.Category),
		Featured:      apiPkg.IsFeatured,
		Native:        contains(apiPkg.PackageTypes, "native"),
		Component:     contains(apiPkg.PackageTypes, "component"),
		PackageStatus: pkg.PackageStatus(apiPkg.PackageStatus),
		SchemaFileURL: apiPkg.SchemaURL,
		Version:       apiPkg.Version,
		Title:         apiPkg.Title,
		UpdatedOn:     apiPkg.CreatedAt.Unix(),
	}, nil
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
			var provider PackageMetadataProvider
			if useAPI {
				provider = NewAPIProvider(apiURL)
			} else {
				provider = NewFileSystemProvider(registryDir)
			}

			if len(args) > 0 {
				glog.Infoln("Generating docs for a single package:", args[0])
				metadata, err := provider.GetPackageMetadata(args[0])
				if err != nil {
					return errors.Wrapf(err, "getting metadata for package %q", args[0])
				}

				docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
				err = genResourceDocsForPackageFromRegistryMetadata(*metadata, docsOutDir, basePackageTreeJSONOutDir)
				if err != nil {
					return errors.Wrapf(err, "generating docs for package %q from registry metadata", args[0])
				}
			} else {
				glog.Infoln("Generating docs for all packages in the registry...")
				err := genResourceDocsForAllRegistryPackages(provider, baseDocsOutDir, basePackageTreeJSONOutDir)
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

// GetPackageMetadata implements PackageMetadataProvider for fileSystemProvider
func (p *fileSystemProvider) GetPackageMetadata(pkgName string) (*pkg.PackageMeta, error) {
	metadataFilePath := filepath.Join(getRegistryPackagesPath(p.registryDir), pkgName+".yaml")
	b, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading the metadata file %s", metadataFilePath)
	}

	var metadata pkg.PackageMeta
	if err := yaml.Unmarshal(b, &metadata); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling the metadata file %s", metadataFilePath)
	}

	return &metadata, nil
}

// ListPackageMetadata implements PackageMetadataProvider for fileSystemProvider
func (p *fileSystemProvider) ListPackageMetadata() ([]*pkg.PackageMeta, error) {
	registryPackagesPath := getRegistryPackagesPath(p.registryDir)
	files, err := os.ReadDir(registryPackagesPath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading directory %s", registryPackagesPath)
	}

	// Count YAML files to pre-allocate the slice mostly to appease the linter.
	var metadataCount int
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			metadataCount++
		}
	}

	metadataList := make([]*pkg.PackageMeta, 0, metadataCount)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		metadata, err := p.GetPackageMetadata(strings.TrimSuffix(file.Name(), ".yaml"))
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}

// GetPackageMetadata implements PackageMetadataProvider for registryAPIProvider
func (p *registryAPIProvider) GetPackageMetadata(pkgName string) (*pkg.PackageMeta, error) {
	resp, err := http.Get(fmt.Sprintf("%s/packages?name=%s", p.apiURL, pkgName))
	if err != nil {
		return nil, errors.Wrapf(err, "fetching package metadata from API for %s", pkgName)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code %d when fetching package metadata", resp.StatusCode)
	}

	var response PackageListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "decoding API response")
	}

	if len(response.Packages) == 0 {
		return nil, errors.Errorf("no package found with name %s", pkgName)
	}

	if len(response.Packages) > 1 {
		return nil, errors.Errorf("multiple packages found with name %s", pkgName)
	}

	return convertAPIPackageToPackageMeta(response.Packages[0])
}

// ListPackageMetadata implements PackageMetadataProvider for registryAPIProvider
func (p *registryAPIProvider) ListPackageMetadata() ([]*pkg.PackageMeta, error) {
	resp, err := http.Get(p.apiURL + "/packages")
	if err != nil {
		return nil, errors.Wrap(err, "fetching package list from API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code %d when fetching package list", resp.StatusCode)
	}

	var response PackageListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "decoding API response")
	}

	metadataList := make([]*pkg.PackageMeta, 0, len(response.Packages))
	for _, apiPkg := range response.Packages {
		metadata, err := convertAPIPackageToPackageMeta(apiPkg)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}
