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
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"

	"github.com/golang/glog"

	"github.com/pkg/errors"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
)

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
	if metadata.RepoURL == "" {
		return errors.Errorf("metadata for package %q does not contain the repo_url", metadata.Name)
	}

	schemaFilePath := fmt.Sprintf(defaultSchemaFilePathFormat, metadata.Name)
	if metadata.SchemaFilePath != "" {
		schemaFilePath = metadata.SchemaFilePath
	}

	// Make sure the schema file path does not have a leading slash.
	// We'll add in the URL format below. It's easier to read that way.
	schemaFilePath = strings.TrimPrefix(schemaFilePath, "/")

	repoSlug, err := getRepoSlug(metadata.RepoURL)
	if err != nil {
		return errors.WithMessage(err, "could not get repo slug")
	}

	glog.Infoln("Reading remote schema file from VCS")
	// TODO: Support raw URLs for other VCS too.
	schemaFileURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repoSlug, metadata.Version, schemaFilePath)
	resp, err := http.Get(schemaFileURL) //nolint:gosec
	if err != nil {
		return errors.Wrapf(err, "reading schema file from VCS %s", schemaFileURL)
	}

	defer contract.IgnoreClose(resp.Body)

	schemaBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "reading response body from %s", schemaFileURL)
	}

	// The source schema can be in YAML format. If that's the case
	// convert it to JSON first.
	if strings.HasSuffix(schemaFilePath, ".yaml") {
		schemaBytes, err = yaml.YAMLToJSON(schemaBytes)
		if err != nil {
			return errors.Wrap(err, "reading YAML schema")
		}
	}

	mainSpec = &pschema.PackageSpec{}
	if err := json.Unmarshal(schemaBytes, mainSpec); err != nil {
		return errors.Wrap(err, "unmarshalling schema into a PackageSpec")
	}

	pulPkg, err := getPulumiPackageFromSchema(docsOutDir)
	if err != nil {
		return errors.Wrap(err, "generating package from schema file")
	}

	glog.Infoln("Running docs generator...")
	if err := generateDocsFromSchema(docsOutDir, pulPkg); err != nil {
		return errors.Wrap(err, "generating docs from schema")
	}

	glog.Infoln("Generating the package tree JSON file...")
	if err := generatePackageTree(packageTreeJSONOutDir, pulPkg.Name); err != nil {
		return errors.Wrap(err, "generating package tree")
	}

	return nil
}

func getRegistryPackagesPath(repoPath string) string {
	return filepath.Join(repoPath, "themes", "default", "data", "registry", "packages")
}

const maxConcurrency = 1

func genResourceDocsForAllRegistryPackages(registryRepoPath, baseDocsOutDir, basePackageTreeJSONOutDir string) error {
	registryPackagesPath := getRegistryPackagesPath(registryRepoPath)
	metadataFiles, err := os.ReadDir(registryPackagesPath)
	if err != nil {
		return errors.Wrap(err, "reading the registry packages dir")
	}

	var m sync.Mutex
	var errs []error
	addError := func(err error) {
		m.Lock()
		defer m.Unlock()
		errs = append(errs, err)
	}

	// Concurrency is limited via a checkout-checkin system.
	//
	// Each time a process starts, it checks out a hallpass from the monitor (blocking
	// if none are available). When the process completes, it returns it's hallpass to
	// the monitor.
	//
	// Because there are a fixed number of hallpasses (maxConcurrency), and each
	// process needs a hallpass to run, we limit effective concurrency to the number
	// of available hallpasses.
	var hallpass struct{}
	monitor := make(chan struct{}, maxConcurrency)
	for i := 0; i < maxConcurrency; i++ {
		monitor <- hallpass
	}

	var w sync.WaitGroup
	w.Add(len(metadataFiles))
	for _, f := range metadataFiles {
		go func(f os.DirEntry) {
			defer func(hallpass struct{}) {
				w.Done()
				monitor <- hallpass
			}(<-monitor)
			glog.Infof("=== starting %s ===\n", f.Name())
			glog.Infoln("Processing metadata file")
			metadataFilePath := filepath.Join(registryPackagesPath, f.Name())

			b, err := os.ReadFile(metadataFilePath)
			if err != nil {
				addError(errors.Wrapf(err, "reading the metadata file %s", metadataFilePath))
				return
			}

			var metadata pkg.PackageMeta
			if err := yaml.Unmarshal(b, &metadata); err != nil {
				addError(errors.Wrapf(err, "unmarshalling the metadata file %s", metadataFilePath))
				return
			}

			docsOutDir := filepath.Join(baseDocsOutDir, metadata.Name, "api-docs")
			err = genResourceDocsForPackageFromRegistryMetadata(metadata, docsOutDir, basePackageTreeJSONOutDir)
			if err != nil {
				addError(errors.Wrapf(err, "generating resource docs using metadata file info %s", f.Name()))
				return
			}

			glog.Infof("=== completed %s ===", f.Name())
		}(f)
	}

	w.Wait()
	contract.Assertf(len(monitor) == maxConcurrency,
		"Every process should have completed, so every flag should be checked in")

	return stderrors.Join(errs...)
}

func resourceDocsFromRegistryCmd() *cobra.Command {
	var baseDocsOutDir string
	var basePackageTreeJSONOutDir string

	cmd := &cobra.Command{
		Use:   "registry [pkgName]",
		Short: "Generate resource docs for a package from the registry",
		Long: "Generate resource docs for all packages in the registry or specific packages. " +
			"Pass a package name in the registry as an optional arg to generate docs only for that package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			registryDir, err := os.Getwd()
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

				err = genResourceDocsForPackageFromRegistryMetadata(metadata, docsOutDir, basePackageTreeJSONOutDir)
				if err != nil {
					return errors.Wrapf(err, "generating docs for package %q from registry metadata", pkgName)
				}
			} else {
				glog.Infoln("Generating docs for all packages in the registry...")
				err := genResourceDocsForAllRegistryPackages(registryDir, baseDocsOutDir, basePackageTreeJSONOutDir)
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

	return cmd
}
