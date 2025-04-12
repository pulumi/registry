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

package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
)

// PackageMetadataProvider is an interface for providers that retrieve package metadata
// from either the filesystem directory or the Pulumi Registry API.
type PackageMetadataProvider interface {
	// GetPackageMetadata returns metadata for a specific package
	GetPackageMetadata(ctx context.Context, pkgName string) (pkg.PackageMeta, error)
	// ListPackageMetadata returns metadata for all packages
	ListPackageMetadata(ctx context.Context) ([]pkg.PackageMeta, error)
}

// HTTPClient is an interface for making HTTP requests
//
//go:generate mockgen -destination=mock_httpclient_test.go -package=svc github.com/pulumi/registry/tools/resourcedocsgen/pkg/registry/svc HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// fileSystemProvider implements PackageMetadataProvider using the local yaml data files
// in the pulumi/registry repository.
type fileSystemProvider struct {
	registryDir string
}

// registryAPIProvider implements PackageMetadataProvider using the Pulumi API
// to retrieve package metadata.
type registryAPIProvider struct {
	apiURL string
	client HTTPClient
}

// PackageMetadata represents the API response structure for package metadata
// from the Pulumi Registry API.
// TODO: will eventually be available in the "github.com/pulumi/pulumi/sdk/v3/go/common/apitype" package.
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
		client: http.DefaultClient,
	}
}

func getRegistryPackagesPath(repoPath string) string {
	return filepath.Join(repoPath, "themes", "default", "data", "registry", "packages")
}

func convertAPIPackageToPackageMeta(apiPkg PackageMetadata) pkg.PackageMeta {
	return pkg.PackageMeta{
		Name:          apiPkg.Name,
		Publisher:     apiPkg.Publisher,
		Description:   apiPkg.Description,
		LogoURL:       apiPkg.LogoURL,
		RepoURL:       apiPkg.RepoURL,
		Category:      pkg.PackageCategory(apiPkg.Category),
		Featured:      apiPkg.IsFeatured,
		Native:        slices.Contains(apiPkg.PackageTypes, "native"),
		Component:     slices.Contains(apiPkg.PackageTypes, "component"),
		PackageStatus: pkg.PackageStatus(apiPkg.PackageStatus),
		SchemaFileURL: apiPkg.SchemaURL,
		Version:       apiPkg.Version,
		Title:         apiPkg.Title,
		UpdatedOn:     apiPkg.CreatedAt.Unix(),
	}
}

// GetPackageMetadata implements PackageMetadataProvider for fileSystemProvider
func (p *fileSystemProvider) GetPackageMetadata(ctx context.Context, pkgName string) (pkg.PackageMeta, error) {
	metadataFilePath := filepath.Join(getRegistryPackagesPath(p.registryDir), pkgName+".yaml")
	b, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return pkg.PackageMeta{}, errors.Wrapf(err, "reading the metadata file %s", metadataFilePath)
	}

	var metadata pkg.PackageMeta
	if err := yaml.Unmarshal(b, &metadata); err != nil {
		return pkg.PackageMeta{}, errors.Wrapf(err, "unmarshalling the metadata file %s", metadataFilePath)
	}

	return metadata, nil
}

// ListPackageMetadata implements PackageMetadataProvider for fileSystemProvider
func (p *fileSystemProvider) ListPackageMetadata(ctx context.Context) ([]pkg.PackageMeta, error) {
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

	metadataList := make([]pkg.PackageMeta, 0, metadataCount)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		metadata, err := p.GetPackageMetadata(ctx, strings.TrimSuffix(file.Name(), ".yaml"))
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}

// GetPackageMetadata implements PackageMetadataProvider for registryAPIProvider
func (p *registryAPIProvider) GetPackageMetadata(ctx context.Context, pkgName string) (pkg.PackageMeta, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/packages?name=%s", p.apiURL, pkgName), nil)
	if err != nil {
		return pkg.PackageMeta{}, errors.Wrapf(err, "creating request for package %s", pkgName)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return pkg.PackageMeta{}, errors.Wrapf(err, "fetching package metadata from API for %s", pkgName)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pkg.PackageMeta{}, errors.Errorf("unexpected status code %d when fetching package metadata", resp.StatusCode)
	}

	var response PackageListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return pkg.PackageMeta{}, errors.Wrap(err, "decoding API response")
	}

	switch len(response.Packages) {
	case 0:
		return pkg.PackageMeta{}, errors.Errorf("no package found with name %s", pkgName)
	case 1:
		metadata, err := convertAPIPackageToPackageMeta(response.Packages[0])
		if err != nil {
			return pkg.PackageMeta{}, err
		}
		return *metadata, nil
	default:
		return pkg.PackageMeta{}, errors.Errorf("multiple packages found with name %s", pkgName)
	}
}

// ListPackageMetadata implements PackageMetadataProvider for registryAPIProvider
func (p *registryAPIProvider) ListPackageMetadata(ctx context.Context) ([]pkg.PackageMeta, error) {
	var allPackages []pkg.PackageMeta
	// Maximum allowed by the API (must be less than 500). Request up to 499 to
	// account for pagination with minimum number of requests.
	const limit = 499
	continuationToken := ""

	for {
		url := fmt.Sprintf("%s/packages?limit=%d", p.apiURL, limit)
		if continuationToken != "" {
			url = fmt.Sprintf("%s&continuationToken=%s", url, continuationToken)
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "creating request for package list")
		}

		resp, err := p.client.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "fetching package list from API")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("unexpected status code %d when fetching package list", resp.StatusCode)
		}

		var response struct {
			Packages          []PackageMetadata `json:"packages"`
			ContinuationToken *string           `json:"continuationToken"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, errors.Wrap(err, "decoding API response")
		}

		for _, apiPkg := range response.Packages {
			metadata, err := convertAPIPackageToPackageMeta(apiPkg)
			if err != nil {
				return nil, err
			}
			allPackages = append(allPackages, *metadata)
		}

		// If there's no continuation token, we've reached the end
		if response.ContinuationToken == nil {
			break
		}

		continuationToken = *response.ContinuationToken
	}

	return allPackages, nil
}
