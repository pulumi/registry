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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const defaultPackageCategory = pkg.PackageCategoryCloud

var featuredPackages = []string{
	"aws",
	"azure-native",
	"gcp",
	"kubernetes",
}

func PackageMetadataCmd() *cobra.Command {
	var repoSlug optional[repoSlug, *repoSlug]
	var providerName string
	var categoryStr string
	var component bool
	var publisher string
	var schemaFile string
	var title string
	var version string
	var metadataDir string
	var packageDocsDir string
	var schemaFileURL string
	var indexFileURL string
	var installationConfigurationURL string

	cmd := &cobra.Command{
		Use:   "metadata <args>",
		Short: "Generate package metadata from Pulumi schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			if repoSlug.isSet && schemaFile == "" && providerName == "" {
				providerName = strings.Replace(repoSlug.v.name, "pulumi-", "", -1)
			}

			publishedDate := time.Now()

			if schemaFileURL == "" {
				if !repoSlug.isSet {
					return errors.New("repoSlug is required unless schemaFileURL is passed")
				}
				schemaFileURL = inferSchemaFileURL(providerName, schemaFile, repoSlug.String(), version)
				// try and get the version release data using the github releases API
				//
				// We only attempt to infer the date when schemaFileURL is missing, implying we are
				// sourcing this from a GH repo.
				if pd, ok, err := inferPublishDate(repoSlug.v, version); err != nil {
					return errors.WithMessage(err, "failed to infer publish date")
				} else if ok {
					publishedDate = pd
				}
			} else {
				if schemaFile != "" {
					glog.Warning("schemaFile was ignored - schemaFileURL takes precedence")
				}
			}
			mainSpec, err := readRemoteSchemaFile(schemaFileURL, repoSlug.v)
			if err != nil {
				return errors.WithMessage(err, "unable to read remote schema file")
			}

			// Unify the version passed in (if any) with the version on the schema (if any).
			switch {
			// If we don't have any version, then error.
			case mainSpec.Version == "" && version == "":
				return fmt.Errorf("No version available for package %s", providerName)
			// Version has been specified, so use it.
			//
			// This can override mainSpec.Version.
			case version != "":
				mainSpec.Version = version
			// Infer the version from the schema.
			default:
				contract.Assertf(mainSpec.Version != "", "impossible - version was checked earlier")
				version = mainSpec.Version
			}

			// Unify the repo passed in (if any) with the version in the schema (if any).
			switch {
			case mainSpec.Repository == "" && !repoSlug.isSet:
				return fmt.Errorf("No repository available for package %s", providerName)
			case repoSlug.isSet:
				mainSpec.Repository = "https://github.com/" + repoSlug.String()
			default:
				r, err := url.Parse(mainSpec.Repository)
				if err != nil {
					return errors.Wrapf(err, "parsing repo url %q from schema", mainSpec.Repository)
				}

				if r.Hostname() != "github.com" {
					return fmt.Errorf("mainSpec.Repository host must be from GH, found %q", r.Hostname())
				}

				// r.Path starts with a "/", and may end with a "/". For example:
				//
				// If the URL is "https://github.com/example/pulumi-example", then the path would be
				// "/example/pulumi-example". repoSlug.Set expects a slug of the form "{owner}/{repo}",
				// so we need to strip the leading "/". If there is a trailing "/", we need to strip
				// that too.
				if err := repoSlug.Set(strings.Trim(r.Path, "/")); err != nil {
					return errors.Wrapf(err, "parsing repo slug from %q", mainSpec.Repository)
				}
			}

			status := pkg.PackageStatusGA
			if strings.HasPrefix(version, "v0.") || strings.HasPrefix(version, "0.") {
				status = pkg.PackageStatusPublicPreview
			}

			category, err := getPackageCategory(mainSpec, categoryStr)
			if err != nil {
				return errors.Wrap(err, "getting category")
			}

			// If the title was not overridden, then try to determine
			// the title from the schema.
			if title == "" {
				// If the schema for this package does not have the
				// displayName, then use its package name.
				if mainSpec.DisplayName == "" {
					title = mainSpec.Name
					// Eventually all of Pulumi's own packages will have the displayName
					// set in their schema but for the time being until they are updated
					// with that info, let's lookup the proper title from the lookup map.
					if v, ok := pkg.TitleLookup[mainSpec.Name]; ok {
						title = v
					}
				} else {
					title = mainSpec.DisplayName
				}
			}

			component = component /* CLI flag */ || isComponent(mainSpec.Keywords)
			native := !component && (mainSpec.Attribution == "" || isNative(mainSpec.Keywords))

			// if there's a publisher then we need to use that immediately if there is no
			// publisher on cmd, then try and use packageSpec if there's no publisher or
			// packageSpec publisher, then assume repo owner is the publisher otherwise error
			var publisherName string
			switch {
			case publisher != "" /* CLI flag */ :
				publisherName = publisher
			case mainSpec.Publisher != "":
				publisherName = mainSpec.Publisher
			default:
				contract.Assertf(repoSlug.v.owner != "", "repoSlug.owner is non-empty by construction")
				publisherName = cases.Title(language.Und, cases.NoLower).String(repoSlug.v.owner)
			}

			pm := pkg.PackageMeta{
				Name:        mainSpec.Name,
				Description: mainSpec.Description,
				LogoURL:     mainSpec.LogoURL,
				Publisher:   publisherName,
				Title:       title,

				RepoURL:       mainSpec.Repository,
				SchemaFileURL: schemaFileURL,
				PackageStatus: status,
				UpdatedOn:     publishedDate.Unix(),
				Version:       version,

				Category:  category,
				Component: component,
				Featured:  isFeaturedPackage(mainSpec.Name),
				Native:    native,
			}
			packageMetaBytes, err := yaml.Marshal(pm)
			if err != nil {
				return errors.Wrap(err, "generating package metadata")
			}

			if err := pkg.EmitFile(metadataDir, mainSpec.Name+".yaml", append([]byte(
				`# WARNING: this file was generated by resourcedocsgen
# Do not edit by hand unless you're certain you know what you are doing!
`), packageMetaBytes...)); err != nil {
				return errors.Wrap(err, "writing metadata file")
			}

			if packageDocsDir == "" {
				// if the user hasn't specified an packageDocsDir, we will default to
				// the path within the registry folder.
				packageDocsDir = "themes/default/content/registry/packages/" + mainSpec.Name
			}

			requiredFiles := []struct {
				name, url string
			}{
				{"_index.md", indexFileURL},
				{"installation-configuration.md", installationConfigurationURL},
			}

			for _, requiredFile := range requiredFiles {
				docsURLOrDefault := func(specified, name string) string {
					if specified != "" {
						return specified
					}
					return "https://raw.githubusercontent.com/" + repoSlug.String() + "/" + version + "/docs/" + name
				}

				url := docsURLOrDefault(requiredFile.url, requiredFile.name)

				content, err := readRemoteFile(url, repoSlug.v.owner)
				if err != nil {
					return err
				}
				if len(content) == 0 {
					continue
				}

				// Normalize end of line representation
				content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

				if rest, ok := bytes.CutPrefix(bytes.TrimLeft(content, "\n\t\r "), []byte("---\n")); ok {
					content = append([]byte(`---
# WARNING: this file was fetched from `+url+`
# Do not edit by hand unless you're certain you know what you are doing!
`), rest...)
				} else {
					return fmt.Errorf(`expected file %s to start with YAML front-matter ("---\n"), found leading %q`,
						url, strings.Split(string(content), "\n")[0])
				}

				if err := pkg.EmitFile(packageDocsDir, requiredFile.name, content); err != nil {
					return errors.Wrap(err, fmt.Sprintf("writing %s file", requiredFile))
				}
			}

			return nil
		},
	}

	cmd.Flags().Var(&repoSlug, "repoSlug", "The repository slug e.g. pulumi/pulumi-provider")
	cmd.Flags().StringVar(&providerName, "providerName", "", "The name of the provider e.g. aws, aws-native. "+
		"Required when there is no schemaFile flag specified.")
	cmd.Flags().StringVar(&schemaFileURL, "schemaFileURL", "",
		`The URL from which the schema can be retrieved.

schemaFileURL takes precedence over schemaFile.`)
	cmd.Flags().StringVar(&indexFileURL, "indexFileURL", "",
		`The URL from which the docs/_index.md file can be retrieved.`)
	cmd.Flags().StringVar(&installationConfigurationURL, "installationConfigurationFileURL", "",
		`The URL from which the docs/installation-configuration.md file can be retrieved.`)
	cmd.Flags().StringVarP(&schemaFile, "schemaFile", "s", "",
		"Relative path to the schema.json file from the root of the repository. If no schemaFile is specified,"+
			" then providerName is required so the schemaFile path can "+
			"be inferred to be provider/cmd/pulumi-resource-<providerName>/schema.json")
	cmd.Flags().StringVar(&version, "version", "", "The version of the package")
	cmd.Flags().StringVar(&categoryStr, "category", "", fmt.Sprintf("The category for the package. Value must "+
		"match one of the keys in the map: %v", pkg.CategoryNameMap))
	cmd.Flags().StringVar(&publisher, "publisher", "", "The publisher's display name to be shown in the package. "+
		"This will default to Pulumi")
	cmd.Flags().StringVar(&title, "title", "",
		"The display name of the package. If omitted, the name of the package will be used")
	cmd.Flags().BoolVar(&component, "component", false,
		"Whether or not this package is a component and not a provider")
	cmd.Flags().StringVar(&metadataDir, "metadataDir", "themes/default/data/registry/packages",
		"The location to save the metadata - this will default to the folder "+
			"structure that the registry expects (themes/default/data/registry/packages)")
	cmd.Flags().StringVar(&packageDocsDir, "packageDocsDir", "",
		"The location to save the package docs - this will default to the folder "+
			"structure that the registry expects (themes/default/content/registry/packages)")

	return cmd
}

func readRemoteFile(url, repoOwner string) ([]byte, error) {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("downloading remote file from %s: %w", url, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// For pulumi repos, we have hard coded top-level config files in the registry.
		// To avoid overwriting them prematurely while we migrate, we default to returning nil, which will allow the
		// registry to fall back on top-level config files already in existence since we won't write empty content.
		if repoOwner == "pulumi" && resp.StatusCode == 404 {
			return nil, nil
		}
		// For third-level providers, send an error if files could not be found.
		return nil, fmt.Errorf("finding remote file at %s: %s", url, resp.Status)
	}

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading contents of remote file")
	}

	return contents, nil
}

func getPackageCategory(mainSpec *pschema.PackageSpec, categoryOverrideStr string) (pkg.PackageCategory, error) {
	var category pkg.PackageCategory
	var err error

	// If a category override was passed-in, use that instead of what's in the schema.
	if categoryOverrideStr != "" {
		glog.V(2).Infof("Using category override name %s\n", categoryOverrideStr)
		n, ok := pkg.CategoryNameMap[categoryOverrideStr]
		if !ok {
			return "", fmt.Errorf("invalid override for category name %s", categoryOverrideStr)
		}
		category = n
	} else if c, ok := pkg.CategoryLookup[mainSpec.Name]; ok {
		glog.V(2).Infoln("Using the category for this package from the lookup map")
		// TODO: This condition can be removed when all packages under the `pulumi` org
		// have a proper category tag in their schema.
		category = c
	}

	if category != "" {
		return category, nil
	}

	glog.V(2).Infoln("Looking-up category from the keywords in the schema")
	category, err = getCategoryFromKeywords(mainSpec.Keywords)
	if err != nil {
		return "", errors.Wrap(err, "getting the category from keywords")
	}

	return category, nil
}

// getCategoryFromKeywords searches for a tag in the provided keywords slice
// with a prefix of category/. Returns the converted category type if such a tag
// is found. Otherwise, returns PackageCategoryCloud always as the default.
func getCategoryFromKeywords(keywords []string) (pkg.PackageCategory, error) {
	categoryTag := getTagWithPrefixFromKeywords(keywords, "category/")
	if categoryTag == nil {
		return defaultPackageCategory, nil
	}

	categoryName := strings.Replace(*categoryTag, "category/", "", -1)
	category, ok := pkg.CategoryNameMap[categoryName]
	if !ok {
		return defaultPackageCategory, fmt.Errorf("invalid category tag %s", *categoryTag)
	}

	return category, nil
}

func isComponent(keywords []string) bool {
	return getTagFromKeywords(keywords, "kind/component") != nil
}

func isFeaturedPackage(pkgName string) bool {
	return slices.Contains(featuredPackages, pkgName)
}

func isNative(keywords []string) bool {
	return getTagFromKeywords(keywords, "kind/native") != nil
}

func getTagWithPrefixFromKeywords(keywords []string, tagPrefix string) *string {
	for _, k := range keywords {
		if strings.HasPrefix(k, tagPrefix) {
			return &k
		}
	}

	glog.V(2).Infof("A tag with the prefix %q was not found in the package's keywords", tagPrefix)
	return nil
}

func getTagFromKeywords(keywords []string, tag string) *string {
	for _, k := range keywords {
		if k == tag {
			return &k
		}
	}

	glog.V(2).Infof("The tag %q was not found in the package's keywords", tag)
	return nil
}

func getGitHubTags(repoSlug string) ([]pkg.GitHubTag, error) {
	path := fmt.Sprintf("/repos/%s/tags", repoSlug)
	tagsResp, err := pkg.GetGitHubAPI(path)
	if err != nil {
		return nil, fmt.Errorf("getting tags info for %s: %w", repoSlug, err)
	}
	defer tagsResp.Body.Close()

	if tagsResp.StatusCode != 200 {
		respBody, err := io.ReadAll(tagsResp.Body)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("getting tags info for %s: %s", repoSlug, tagsResp.Status))
		}

		return nil, fmt.Errorf("getting tags info for %s: %s", repoSlug, string(respBody))
	}

	var tags []pkg.GitHubTag
	err = json.NewDecoder(tagsResp.Body).Decode(&tags)
	if err != nil {
		return nil, fmt.Errorf("constructing tags information for %s: %w", repoSlug, err)
	}

	return tags, nil
}

// optional makes a pflag.Value optional
type optional[V any, T interface {
	*V
	pflag.Value
}] struct {
	v     V
	isSet bool
}

func (s *optional[U, T]) String() string { return T(&s.v).String() }

func (s *optional[U, T]) Set(input string) error {
	if input == "" {
		return nil
	}
	s.isSet = true
	return T(&s.v).Set(input)
}

func (s *optional[U, T]) Type() string { return T(&s.v).Type() }

type repoSlug struct{ name, owner string }

func (s repoSlug) String() string { return s.owner + "/" + s.name }

func (s *repoSlug) Set(input string) error {
	if strings.Contains(input, "https") || strings.Contains(input, "github.com") {
		return fmt.Errorf("Expected repoSlug to be in the format of `{owner}/{repo}`"+
			" but got %q", input)
	}

	githubSlugParts := strings.Split(input, "/")

	if len(githubSlugParts) != 2 {
		return fmt.Errorf("Expected repoSlug to be in the format of `{owner}/{repo}`"+
			" but got %q", input)
	}

	owner := strings.TrimSpace(githubSlugParts[0])
	name := strings.TrimSpace(githubSlugParts[1])

	switch {
	case owner == "" && name == "":
		return fmt.Errorf("Expected non-empty {owner} and non-empty {name} in {owner}/{name}, but got %q", input)
	case owner == "":
		return fmt.Errorf("Expected non-empty {owner} in {owner}/%s, but got %q", name, input)
	case name == "":
		return fmt.Errorf("Expected non-empty {name} in %s/{name}, but got %q", owner, input)
	}

	s.owner = owner
	s.name = name
	return nil
}

func (s repoSlug) Type() string { return "repo slug" }

func inferSchemaFileURL(providerName, schemaFile, repo, version string) string {
	if schemaFile == "" {
		schemaFile = fmt.Sprintf("provider/cmd/pulumi-resource-%s/schema.json", providerName)
	}

	// we should be able to take the repo URL + the version + the schema url and
	// construct a file that we can download and read
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s",
		repo, version, schemaFile)
}

func readRemoteSchemaFile(schemaFileURL string, repo repoSlug) (*pschema.PackageSpec, error) {
	schema, err := readRemoteFile(schemaFileURL, repo.owner)
	if err != nil {
		return nil, err
	}

	if schema == nil {
		return nil, fmt.Errorf("unable to get contents of schema file at %q", schemaFileURL)
	}

	// The source schema can be in YAML format. If that's the case
	// convert it to JSON first.
	if strings.HasSuffix(schemaFileURL, ".yaml") {
		schema, err = yaml.YAMLToJSON(schema)
		if err != nil {
			return nil, errors.Wrap(err, "reading YAML schema")
		}
	}

	spec := &pschema.PackageSpec{}
	if err := json.Unmarshal(schema, spec); err != nil {
		return nil, errors.Wrap(err, "unmarshalling schema into a PackageSpec")
	}
	return spec, nil
}

func inferPublishDate(repo repoSlug, version string) (time.Time, bool, error) {
	// try and get the version release data using the github releases API
	tags, err := getGitHubTags(repo.String())
	if err != nil {
		return time.Time{}, false, errors.Wrap(err, "github tags")
	}

	var commitDetails string
	for _, tag := range tags {
		if tag.Name == version {
			commitDetails = tag.Commit.URL
			break
		}
	}

	if commitDetails == "" {
		return time.Time{}, false, nil
	}

	var commit pkg.GitHubCommit
	// now let's make a request to the specific commit to get the date
	commitResp, err := http.Get(commitDetails) //nolint:gosec
	if err != nil {
		return time.Time{}, false, fmt.Errorf("getting release info for %s: %w", repo, err)
	}

	defer commitResp.Body.Close()
	err = json.NewDecoder(commitResp.Body).Decode(&commit)
	if err != nil {
		return time.Time{}, false, fmt.Errorf("constructing commit information for %s: %w", repo, err)
	}

	return commit.Commit.Author.Date, true, nil
}
