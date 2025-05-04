// Copyright 2025, Pulumi Corporation.
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
	"slices"
	"time"

	"github.com/pkg/errors"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
)

// NewAPIProvider creates a new PackageMetadataProvider that reads from the Pulumi API
func NewAPIProvider(apiURL string) PackageMetadataProvider {
	return &registryAPIProvider{
		apiURL: apiURL,
		client: http.DefaultClient,
	}
}

// registryAPIProvider implements PackageMetadataProvider using the Pulumi API
// to retrieve package metadata.
type registryAPIProvider struct {
	apiURL string
	client *http.Client
}

// apiPackageMetadata represents the API response structure for package metadata
// from the Pulumi Registry API.
type apiPackageMetadata struct {
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

// packageListResponse represents the API response structure for package lists
type packageListResponse struct {
	Packages []apiPackageMetadata `json:"packages"`
}

func convertAPIPackageToPackageMeta(apiPkg apiPackageMetadata) pkg.PackageMeta {
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

	var response packageListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return pkg.PackageMeta{}, errors.Wrap(err, "decoding API response")
	}

	switch len(response.Packages) {
	case 0:
		return pkg.PackageMeta{}, errors.Errorf("no package found with name %s", pkgName)
	case 1:
		metadata := convertAPIPackageToPackageMeta(response.Packages[0])
		return metadata, nil
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
			Packages          []apiPackageMetadata `json:"packages"`
			ContinuationToken *string              `json:"continuationToken"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, errors.Wrap(err, "decoding API response")
		}

		for _, apiPkg := range response.Packages {
			metadata := convertAPIPackageToPackageMeta(apiPkg)
			allPackages = append(allPackages, metadata)
		}

		// If there's no continuation token, we've reached the end
		if response.ContinuationToken == nil {
			break
		}

		continuationToken = *response.ContinuationToken
	}

	return allPackages, nil
}
