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
	"os"
	"path/filepath"
	"strings"

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

// fileSystemProvider implements PackageMetadataProvider using the local yaml data files
// in the pulumi/registry repository.
type fileSystemProvider struct {
	registryDir string
}

// NewFileSystemProvider creates a new PackageMetadataProvider that reads from the filesystem
func NewFileSystemProvider(registryDir string) PackageMetadataProvider {
	return &fileSystemProvider{
		registryDir: registryDir,
	}
}

func getRegistryPackagesPath(repoPath string) string {
	return filepath.Join(repoPath, "themes", "default", "data", "registry", "packages")
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
