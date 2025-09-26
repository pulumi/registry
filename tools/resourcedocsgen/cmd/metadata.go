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
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/spf13/cobra"
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
	cmd := &cobra.Command{
		Use:   "metadata <args>",
		Short: "Generate package metadata from Pulumi schema",
	}

	metadataDir := cmd.PersistentFlags().String("metadataDir", "themes/default/data/registry/packages",
		"The location to save the metadata - this will default to the folder "+
			"structure that the registry expects (themes/default/data/registry/packages)")
	packageDocsDir := cmd.PersistentFlags().String("packageDocsDir", "",
		"The location to save the package docs - this will default to the folder "+
			"structure that the registry expects (themes/default/content/registry/packages/<provider-name>)")

	cmd.AddCommand(
		packageMetadataFromGitHubCmd(metadataDir, packageDocsDir),
		packageMetadataFromURLsCmd(metadataDir, packageDocsDir),
	)

	return cmd
}

func packageMetadataFromURLsCmd(metadataDir, packageDocsDir *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "from-urls <args>",
		Short:        "Generate package metadata from a Pulumi schema",
		SilenceUsage: true,
	}

	// Required Flags

	providerName := cmd.Flags().String("providerName", "", "The name of the provider e.g. aws, aws-native. "+
		"Required when there is no schemaFile flag specified.")
	schemaFileURL := cmd.Flags().String("schemaFileURL", "", `The URL from which the schema can be retrieved.`)
	indexFileURL := cmd.Flags().String("indexFileURL", "",
		`The URL from which the docs/_index.md file can be retrieved.`)

	contract.AssertNoErrorf(stderrors.Join(
		cmd.MarkFlagRequired("providerName"),
		cmd.MarkFlagRequired("schemaFileURL"),
		cmd.MarkFlagRequired("indexFileURL"),
	), "impossible")

	// Optional Flags
	installationConfigurationURL := cmd.Flags().String("installationConfigurationFileURL", "",
		`The URL from which the docs/installation-configuration.md file can be retrieved.`)

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		var errs []error
		if *indexFileURL == "" {
			errs = append(errs, stderrors.New(`missing URL for required file "_index.md"`))
		}

		return stderrors.Join(errs...)
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// By the time this function runs, input *string values have been set by
		// cobra.
		//
		// They are guaranteed to be non-nil.

		providerName := *providerName
		schemaFileURL := *schemaFileURL
		indexFileURL := *indexFileURL
		installationConfigurationURL := *installationConfigurationURL
		metadataDir := *metadataDir
		packageDocsDir := *packageDocsDir

		// Get the schema.

		mainSpec, err := readRemoteSchemaFile(schemaFileURL, "")
		if err != nil {
			return errors.WithMessage(err, "unable to read remote schema file")
		}

		err = writePackageMetadata(mainSpec, providerName, schemaFileURL, metadataDir, time.Now(), nil)
		if err != nil {
			return err
		}

		// Write docs files

		if packageDocsDir == "" {
			// if the user hasn't specified a packageDocsDir, we will default
			// to the path within the registry folder.
			packageDocsDir = "themes/default/content/registry/packages/" + mainSpec.Name
		}

		var errs []error
		index, err := readDocsFile(indexFileURL)
		if err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, pkg.EmitFile(packageDocsDir, "_index.md", index))
		}

		if installationConfigurationURL != "" {
			installationConfiguration, err := readDocsFile(installationConfigurationURL)
			if err != nil {
				errs = append(errs, err)
			} else {
				errs = append(errs, pkg.EmitFile(packageDocsDir, "installation-configuration.md", installationConfiguration))
			}
		}

		return stderrors.Join(errs...)
	}

	return cmd
}

func packageMetadataFromGitHubCmd(metadataDir, packageDocsDir *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "from-github <args>",
		Short:        "Generate package metadata from Pulumi schema",
		SilenceUsage: true,
	}

	// Required Flags

	var repoSlug repoSlug
	var version string
	cmd.Flags().Var(&repoSlug, "repoSlug", "The repository slug e.g. pulumi/pulumi-provider")
	cmd.Flags().StringVar(&version, "version", "", "The version of the package")

	contract.AssertNoErrorf(stderrors.Join(
		cmd.MarkFlagRequired("repoSlug"),
		cmd.MarkFlagRequired("version"),
	), "impossible")

	// Optional Flags

	var providerName string
	var schemaFile string
	var publisher string

	cmd.Flags().StringVar(&providerName, "providerName", "", "The name of the provider e.g. aws, aws-native. "+
		"Required when there is no schemaFile flag specified.")
	cmd.Flags().StringVarP(&schemaFile, "schemaFile", "s", "",
		"Relative path to the schema.json file from the root of the repository. If no schemaFile is specified,"+
			" then providerName is required so the schemaFile path can "+
			"be inferred to be provider/cmd/pulumi-resource-<providerName>/schema.json")
	cmd.Flags().StringVar(&publisher, "publisher", "", "The publisher's display name to be shown in the package. "+
		"This will default to Pulumi")

	// Apply default values that depend on other inputs.
	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		if providerName == "" {
			providerName = strings.ReplaceAll(repoSlug.name, "pulumi-", "")
		}
		if schemaFile == "" {
			schemaFile = fmt.Sprintf("provider/cmd/pulumi-resource-%s/schema.json", providerName)
		}
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		metadataDir := *metadataDir
		packageDocsDir := *packageDocsDir

		publishedDate := time.Now()
		// try and get the version release data using the github releases API
		if pd, ok, err := inferPublishDate(repoSlug, version); err != nil {
			return errors.WithMessage(err, "failed to infer publish date")
		} else if ok {
			publishedDate = pd
		}

		// we should be able to take the repo URL + the version + the schema url and
		// construct a file that we can download and read
		schemaURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s",
			repoSlug, version, schemaFile)

		mainSpec, err := readRemoteSchemaFile(schemaURL, repoSlug.owner)
		if err != nil {
			return errors.WithMessage(err, "unable to read remote schema file")
		}

		contract.Assertf(version != "", "version is a required field")
		mainSpec.Version = version

		mainSpec.Repository = "https://github.com/" + repoSlug.String()

		// If the schema for this package does not have the displayName, and it
		// does have a pre-baked title, use that title.
		if mainSpec.DisplayName == "" {
			// Eventually all of Pulumi's own packages should have the
			// displayName set in their schema but for the time being until
			// they are updated with that info, let's lookup the proper title
			// from the lookup map.
			mainSpec.DisplayName = pkg.TitleLookup[mainSpec.Name]
		}

		// Apply the publisher CLI option to the schema, if specified.
		if publisher != "" {
			mainSpec.Publisher = publisher
		}

		err = writePackageMetadata(mainSpec, providerName, schemaURL, metadataDir, publishedDate, &repoSlug)
		if err != nil {
			return errors.Wrap(err, "generating package metadata")
		}

		if packageDocsDir == "" {
			// if the user hasn't specified a packageDocsDir, we will default to
			// the path within the registry folder.
			packageDocsDir = "themes/default/content/registry/packages/" + mainSpec.Name
		}

		remoteFiles := []struct {
			name     string
			required bool
		}{
			{"_index.md", true},
			{"installation-configuration.md", false},
		}

		for _, remoteFile := range remoteFiles {
			url := "https://raw.githubusercontent.com/" + repoSlug.String() + "/" + mainSpec.Version + "/docs/" + remoteFile.name
			content, err := readRemoteFile(url, repoSlug.owner)
			if err != nil {
				return err
			}
			if len(content) == 0 {
				continue
			}

			// Normalize end of line representation
			content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

			frontmatter := []byte(`---
# WARNING: this file was fetched from ` + url + `
# Do not edit by hand unless you're certain you know what you are doing!
`)
			// Only add edit_url for pulumi GitHub repos
			if strings.HasPrefix(mainSpec.Repository, "https://github.com/pulumi/") {
				editURL := mainSpec.Repository + "/blob/" + mainSpec.Version + "/docs/" + remoteFile.name
				frontmatter = append(frontmatter, []byte(`edit_url: `+editURL+"\n")...)
			}

			if rest, ok := bytes.CutPrefix(bytes.TrimLeft(content, "\n\t\r "), []byte("---\n")); ok {
				content = append(frontmatter, rest...)
			} else {
				return fmt.Errorf(`expected file %s to start with YAML front-matter ("---\n"), found leading %q`,
					url, strings.Split(string(content), "\n")[0])
			}

			if err := pkg.EmitFile(packageDocsDir, remoteFile.name, content); err != nil {
				return errors.Wrapf(err, "writing %s file", remoteFile.name)
			}
		}

		return nil
	}

	return cmd
}

func writePackageMetadata(
	spec *schema.PackageSpec, providerName, schemaURL, metadataDir string, updatedOn time.Time, repoSlug *repoSlug,
) error {
	// Validate the schema for usage.
	//
	// A schema is valid if:
	// - spec.Name is non-empty (and it matches providerName)
	// - spec.Version is non-empty, and valid semver
	// - spec.Keywords is valid
	// - spec.Publisher is non-empty

	var validationErrors []error

	if spec.Name != providerName && !legacyNameRuleException(spec.Name, providerName) {
		validationErrors = append(validationErrors, fmt.Errorf(
			"--providerName doesn't match the schema name: %q != %q", providerName, spec.Name))
	}
	contract.Assertf(providerName != "", "providerName is required, and should never be empty")

	if spec.Version == "" {
		validationErrors = append(validationErrors, fmt.Errorf("schema for %q has no version", spec.Name))
	} else if _, err := semver.Parse(strings.TrimPrefix(spec.Version, "v")); err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("version %q is not valid semver: %w", spec.Version, err))
	}

	category, err := getPackageCategory(spec)
	if err != nil {
		validationErrors = append(validationErrors, errors.Wrap(err, "getting category"))
	}

	if spec.Publisher == "" {
		if repoSlug != nil && legacyPublisherRuleException(*repoSlug) {
			spec.Publisher = getLegacyPublisher(*repoSlug)
		} else {
			validationErrors = append(validationErrors, stderrors.New("publisher is a required field on the schema"))
		}
	}

	// Check if the schema failed to validate.
	if err := stderrors.Join(validationErrors...); err != nil {
		return err
	}

	// The schema is now considered *valid* to publish.

	title := spec.Name
	if spec.DisplayName != "" {
		title = spec.DisplayName
	}

	component := isComponent(spec.Keywords)
	native := !component && (spec.Attribution == "" || isNative(spec.Keywords))
	return emitPackageMetadata(pkg.PackageMeta{
		Name:        spec.Name,
		Description: spec.Description,
		LogoURL:     spec.LogoURL,
		Publisher:   spec.Publisher,
		Title:       title,

		RepoURL:       spec.Repository,
		SchemaFileURL: schemaURL,
		PackageStatus: getPackageStatus(spec.Version),
		UpdatedOn:     updatedOn.Unix(),
		Version:       spec.Version,

		Category:  category,
		Component: component,
		Featured:  isFeaturedPackage(spec.Name),
		Native:    native,
	}, metadataDir)
}

// legacyNameRuleException prevents the registry from rejecting previously acceptable
// names, even if they don't agree.
//
// New providers should not be added to the legacyNameRuleException list. Instead, they
// should comply with the registry's preferred naming scheme.
func legacyNameRuleException(schemaName, providerName string) bool {
	// The only exception we have is runpod:
	//
	// The repo name is "runpod/pulumi-runpod-native" but the name of the provider is
	// "runpod" (not the inferred "runpod-native").
	return schemaName == "runpod" && providerName == "runpod-native"
}

// legacyPublisherExceptions is a set of repository slugs that are exceptions to the
// rule of requiring a publisher field in the schema.
var legacyPublisherExceptions = codegen.NewStringSet(
	"genesiscloud/pulumi-genesiscloud",
	"nuage-studio/pulumi-nuage",
	"wttech/pulumi-aem",
	"pulumi/pulumi-tls-self-signed-cert",
)

// legacyPublisherRuleException checks if the repository slug is an exception to the
// standard publisher naming rules. This allows certain repositories to maintain
// backward compatibility with the previous behavior of not requiring a publisher field
// in the schema.
//
// For existing providers, we allow the publisher field to be empty and derive it from
// the repository owner. New providers should not be added to this exception list and
// should explicitly specify their publisher in the schema.
func legacyPublisherRuleException(repoSlug repoSlug) bool {
	return legacyPublisherExceptions.Has(repoSlug.String())
}

// getLegacyPublisher returns the publisher name for repositories that are exceptions
// to the standard publisher naming rules. It converts the repository owner to title case
// to derive a publisher name when one is not explicitly specified in the schema.
// This maintains backward compatibility with previously established publisher names.
func getLegacyPublisher(repoSlug repoSlug) string {
	contract.Assertf(repoSlug.owner != "", "repoSlug.owner is non-empty by construction")
	return cases.Title(language.Und, cases.NoLower).String(repoSlug.owner)
}

func readRemoteFile(url, repoOwner string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "creating request for %q", url)
	}

	pkg.AddGitHubAuthHeaders(req)

	resp, err := http.DefaultClient.Do(req)
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

func getPackageCategory(mainSpec *schema.PackageSpec) (pkg.PackageCategory, error) {
	var category pkg.PackageCategory
	var err error

	if c, ok := pkg.CategoryLookup[mainSpec.Name]; ok {
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

	categoryName := strings.ReplaceAll(*categoryTag, "category/", "")
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

type repoSlug struct{ name, owner string }

func (s repoSlug) String() string { return s.owner + "/" + s.name }

func (s *repoSlug) Set(input string) error {
	if strings.Contains(input, "https") || strings.Contains(input, "github.com") {
		return fmt.Errorf("expected repoSlug to be in the format of `{owner}/{repo}`"+
			" but got %q", input)
	}

	githubSlugParts := strings.Split(input, "/")

	if len(githubSlugParts) != 2 {
		return fmt.Errorf("expected repoSlug to be in the format of `{owner}/{repo}`"+
			" but got %q", input)
	}

	owner := strings.TrimSpace(githubSlugParts[0])
	name := strings.TrimSpace(githubSlugParts[1])

	switch {
	case owner == "" && name == "":
		return fmt.Errorf("expected non-empty {owner} and non-empty {name} in {owner}/{name}, but got %q", input)
	case owner == "":
		return fmt.Errorf("expected non-empty {owner} in {owner}/%s, but got %q", name, input)
	case name == "":
		return fmt.Errorf("expected non-empty {name} in %s/{name}, but got %q", owner, input)
	}

	s.owner = owner
	s.name = name
	return nil
}

func (s repoSlug) Type() string { return "repo slug" }

func readRemoteSchemaFile(schemaFileURL, repoOwner string) (*schema.PackageSpec, error) {
	schemaBytes, err := readRemoteFile(schemaFileURL, repoOwner)
	if err != nil {
		return nil, err
	}

	if schemaBytes == nil {
		return nil, fmt.Errorf("unable to get contents of schema file at %q", schemaFileURL)
	}

	// The source schema can be in YAML format. If that's the case
	// convert it to JSON first.
	if strings.HasSuffix(schemaFileURL, ".yaml") {
		schemaBytes, err = yaml.YAMLToJSON(schemaBytes)
		if err != nil {
			return nil, errors.Wrap(err, "reading YAML schema")
		}
	}

	spec := &schema.PackageSpec{}
	if err := json.Unmarshal(schemaBytes, spec); err != nil {
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
	req, err := http.NewRequest("GET", commitDetails, nil)
	if err != nil {
		return time.Time{}, false, errors.Wrapf(err, "creating request for %q", commitDetails)
	}

	pkg.AddGitHubAuthHeaders(req)
	commitResp, err := http.DefaultClient.Do(req)
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

// readDocsFile is a specialized version of [readRemoteFile] for docs files.
//
// Docs files are expected to be markdown and have a YAML frontmatter.
func readDocsFile(url string) ([]byte, error) {
	content, err := readRemoteFile(url, "")
	if err != nil {
		return nil, err
	}

	// Normalize end of line representation
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

	rest, ok := bytes.CutPrefix(bytes.TrimLeft(content, "\n\t\r "), []byte("---\n"))
	if !ok {
		return nil, fmt.Errorf(`expected file %s to start with YAML front-matter ("---\n"), found leading %q`,
			url, strings.Split(string(content), "\n")[0])
	}
	return append([]byte(`---
# WARNING: this file was fetched from `+url+`
# Do not edit by hand unless you're certain you know what you are doing!
`), rest...), nil
}

func emitPackageMetadata(pm pkg.PackageMeta, metadataDir string) error {
	packageMetaBytes, err := yaml.Marshal(pm)
	if err != nil {
		return errors.Wrap(err, "generating package metadata")
	}

	if err := pkg.EmitFile(metadataDir, pm.Name+".yaml", append([]byte(
		`# WARNING: this file was generated by resourcedocsgen
# Do not edit by hand unless you're certain you know what you are doing!
`), packageMetaBytes...)); err != nil {
		return errors.Wrap(err, "writing metadata file")
	}
	return nil
}

func getPackageStatus(version string) pkg.PackageStatus {
	if strings.HasPrefix(version, "v0.") || strings.HasPrefix(version, "0.") {
		return pkg.PackageStatusPublicPreview
	}
	return pkg.PackageStatusGA
}
