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
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
)

func stringPtr(s string) *string {
	return &s
}

func TestConvertAPIPackageToPackageMeta(t *testing.T) {
	t.Parallel()
	createdAt := time.Date(2025, 4, 10, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		input   apiPackageMetadata
		want    pkg.PackageMeta
		wantErr bool
	}{
		{
			name: "provider package",
			input: apiPackageMetadata{
				Name:         "aws-test",
				Publisher:    "Pulumi",
				Description:  "AWS provider",
				LogoURL:      "https://example.com/logo.png",
				RepoURL:      "https://github.com/pulumi/pulumi-aws",
				Category:     "cloud",
				IsFeatured:   true,
				PackageTypes: []string{"native"},
				SchemaURL:    "https://example.com/schema.json",
				Version:      "v1.0.0",
				Title:        "AWS",
				CreatedAt:    createdAt,
			},
			want: pkg.PackageMeta{
				Name:          "aws-test",
				Publisher:     "Pulumi",
				Description:   "AWS provider",
				LogoURL:       "https://example.com/logo.png",
				RepoURL:       "https://github.com/pulumi/pulumi-aws",
				Category:      "cloud",
				Featured:      true,
				Native:        true,
				Component:     false,
				SchemaFileURL: "https://example.com/schema.json",
				Version:       "v1.0.0",
				Title:         "AWS",
				UpdatedOn:     createdAt.Unix(),
			},
		},
		{
			name: "component package",
			input: apiPackageMetadata{
				Name:         "test-component",
				PackageTypes: []string{"component"},
				CreatedAt:    createdAt,
			},
			want: pkg.PackageMeta{
				Name:      "test-component",
				Component: true,
				Native:    false,
				UpdatedOn: createdAt.Unix(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := convertAPIPackageToPackageMeta(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFileSystemProvider(t *testing.T) {
	t.Parallel()
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	packagesDir := filepath.Join(tmpDir, "themes", "default", "data", "registry", "packages")
	require.NoError(t, os.MkdirAll(packagesDir, fs.ModePerm))

	// Create test metadata files
	testPackages := []pkg.PackageMeta{
		{
			Name:      "aws-test",
			Publisher: "Pulumi",
			Version:   "v1.0.0",
			UpdatedOn: time.Now().Unix(),
			Native:    true,
			Component: false,
			Category:  "cloud",
			Featured:  true,
		},
		{
			Name:      "kubernetes-test",
			Publisher: "Pulumi",
			Version:   "v2.0.0",
			UpdatedOn: time.Now().Unix(),
			Native:    true,
			Component: false,
			Category:  "cloud",
			Featured:  true,
		},
	}

	for _, p := range testPackages {
		data, err := yaml.Marshal(p)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(packagesDir, p.Name+".yaml"), data, fs.ModePerm)
		require.NoError(t, err)
	}

	provider := NewFileSystemProvider(tmpDir)

	t.Run("GetPackageMetadata", func(t *testing.T) {
		t.Parallel()
		got, err := provider.GetPackageMetadata(context.Background(), "aws-test")
		require.NoError(t, err)
		assert.Equal(t, testPackages[0], got)

		_, err = provider.GetPackageMetadata(context.Background(), "non-existent-pkg-test")
		assert.Error(t, err)
	})

	t.Run("ListPackageMetadata", func(t *testing.T) {
		t.Parallel()
		got, err := provider.ListPackageMetadata(context.Background())
		require.NoError(t, err)
		assert.Len(t, got, len(testPackages))
		assert.ElementsMatch(t, testPackages, got)
	})

	t.Run("GetPackageMetadata errors", func(t *testing.T) {
		t.Parallel()
		t.Run("non-existent file", func(t *testing.T) {
			provider := NewFileSystemProvider(tmpDir)
			_, err := provider.GetPackageMetadata(context.Background(), "non-existent")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "reading the metadata file")
		})

		t.Run("invalid yaml", func(t *testing.T) {
			t.Parallel()

			invalidDir := t.TempDir()
			invalidPackagesDir := filepath.Join(invalidDir, "themes", "default", "data", "registry", "packages")
			require.NoError(t, os.MkdirAll(invalidPackagesDir, fs.ModePerm))

			err := os.WriteFile(filepath.Join(invalidPackagesDir, "invalid.yaml"), []byte("invalid yaml: ]["), fs.ModePerm)
			require.NoError(t, err)

			provider := NewFileSystemProvider(invalidDir)
			_, err = provider.GetPackageMetadata(context.Background(), "invalid")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "unmarshalling the metadata file")
		})
	})

	t.Run("ListPackageMetadata errors", func(t *testing.T) {
		t.Parallel()
		t.Run("non-existent directory", func(t *testing.T) {
			tmpDir := t.TempDir()
			provider := NewFileSystemProvider(tmpDir)
			_, err := provider.ListPackageMetadata(context.Background())
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "reading directory")
		})

		t.Run("ignores non-yaml files", func(t *testing.T) {
			tmpDir := t.TempDir()
			packagesDir := filepath.Join(tmpDir, "themes", "default", "data", "registry", "packages")
			require.NoError(t, os.MkdirAll(packagesDir, fs.ModePerm))

			// Create a valid yaml file
			validPkg := pkg.PackageMeta{Name: "valid", Publisher: "Pulumi"}
			validYAML, err := yaml.Marshal(validPkg)
			require.NoError(t, err)
			err = os.WriteFile(filepath.Join(packagesDir, "valid.yaml"), validYAML, fs.ModePerm)
			require.NoError(t, err)

			// Create an invalid yaml file
			err = os.WriteFile(filepath.Join(packagesDir, "readme.txt"), []byte("not valid yaml"), fs.ModePerm)
			require.NoError(t, err)

			provider := NewFileSystemProvider(tmpDir)
			packages, err := provider.ListPackageMetadata(context.Background())
			require.NoError(t, err)
			assert.Len(t, packages, 1)
			assert.Equal(t, "valid", packages[0].Name)
		})
	})
}

func TestRegistryAPIProvider(t *testing.T) {
	t.Parallel()
	t.Run("GetPackageMetadata", func(t *testing.T) {
		t.Parallel()
		t.Run("successfully returns package metadata", func(t *testing.T) {
			t.Parallel()
			// configure test server
			expectedPackage := apiPackageMetadata{
				Name:      "aws-test",
				Publisher: "Pulumi",
				Version:   "v1.0.0",
				CreatedAt: time.Now(),
			}
			responseBody := packageListResponse{
				Packages: []apiPackageMetadata{expectedPackage},
			}
			jsonResponse, err := json.Marshal(responseBody)
			require.NoError(t, err)

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/packages", r.URL.Path)
				assert.Equal(t, "name=aws-test", r.URL.RawQuery)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(jsonResponse)
				require.NoError(t, err)
			}))
			defer server.Close()

			// rreate client using test server
			provider := &registryAPIProvider{
				apiURL: server.URL,
				client: http.DefaultClient,
			}

			got, err := provider.GetPackageMetadata(context.Background(), "aws-test")
			require.NoError(t, err)
			assert.Equal(t, "aws-test", got.Name)
			assert.Equal(t, "Pulumi", got.Publisher)
		})

		t.Run("non existent package", func(t *testing.T) {
			t.Parallel()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}))
			defer server.Close()

			provider := &registryAPIProvider{
				apiURL: server.URL,
				client: http.DefaultClient,
			}

			_, err := provider.GetPackageMetadata(context.Background(), "non-existent-pkg")
			assert.Error(t, err)
		})

		t.Run("multiple packages found", func(t *testing.T) {
			// Setup response with multiple packages
			response := packageListResponse{
				Packages: []apiPackageMetadata{
					{Name: "aws-test", Publisher: "Pulumi"},
					{Name: "aws-test", Publisher: "Community"},
				},
			}
			jsonResponse, err := json.Marshal(response)
			require.NoError(t, err)

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(jsonResponse)
				require.NoError(t, err)
			}))
			defer server.Close()

			provider := &registryAPIProvider{
				apiURL: server.URL,
				client: http.DefaultClient,
			}

			_, err = provider.GetPackageMetadata(context.Background(), "aws-test")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "multiple packages found")
		})
	})

	t.Run("ListPackageMetadata", func(t *testing.T) {
		t.Parallel()
		t.Run("successfully lists packages with pagination", func(t *testing.T) {
			t.Parallel()
			// configure to test with pagination
			firstPage := struct {
				Packages          []apiPackageMetadata `json:"packages"`
				ContinuationToken *string              `json:"continuationToken"`
			}{
				Packages: []apiPackageMetadata{
					{Name: "aws-test", Publisher: "Pulumi"},
					{Name: "azure-test", Publisher: "Pulumi"},
				},
				ContinuationToken: stringPtr("page2"),
			}
			firstPageJSON, err := json.Marshal(firstPage)
			require.NoError(t, err)

			secondPage := struct {
				Packages          []apiPackageMetadata `json:"packages"`
				ContinuationToken *string              `json:"continuationToken"`
			}{
				Packages: []apiPackageMetadata{
					{Name: "gcp-test", Publisher: "Pulumi"},
				},
				ContinuationToken: nil,
			}
			secondPageJSON, err := json.Marshal(secondPage)
			require.NoError(t, err)

			// setup test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/packages", r.URL.Path)
				w.WriteHeader(http.StatusOK)

				// Return different responses based on continuation token
				if r.URL.Query().Get("continuationToken") == "page2" {
					_, err = w.Write(secondPageJSON)
					require.NoError(t, err)
				} else {
					_, err = w.Write(firstPageJSON)
					require.NoError(t, err)
				}
			}))
			defer server.Close()

			provider := &registryAPIProvider{
				apiURL: server.URL,
				client: http.DefaultClient,
			}

			got, err := provider.ListPackageMetadata(context.Background())
			require.NoError(t, err)
			assert.Len(t, got, 3)
			assert.Equal(t, []string{"aws-test", "azure-test", "gcp-test"}, []string{got[0].Name, got[1].Name, got[2].Name})
		})
		t.Run("server error", func(t *testing.T) {
			t.Parallel()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))
			defer server.Close()

			provider := &registryAPIProvider{
				apiURL: server.URL,
				client: http.DefaultClient,
			}

			_, err := provider.ListPackageMetadata(context.Background())
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "unexpected status code 500")
		})
	})
}
