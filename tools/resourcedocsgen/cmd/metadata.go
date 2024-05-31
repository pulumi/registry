package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/pulumi/docs/tools/resourcedocsgen/pkg"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const defaultPackageCategory = pkg.PackageCategoryCloud

var mainSpec *pschema.PackageSpec

var featuredPackages = []string{
	"aws",
	"azure-native",
	"gcp",
	"kubernetes",
}

func PackageMetadataCmd() *cobra.Command {
	var repoSlug string
	var providerName string
	var categoryStr string
	var component bool
	var publisher string
	var schemaFile string
	var title string
	var version string
	var metadataDir string
	var packageDocsDir string

	cmd := &cobra.Command{
		Use:   "metadata <args>",
		Short: "Generate package metadata from Pulumi schema",
		RunE: func(cmd *cobra.Command, args []string) error {

			if strings.Contains(repoSlug, "https") || strings.Contains(repoSlug, "github.com") {
				return errors.New(fmt.Sprintf("Expected repoSlug to be in the format of `owner/repo`"+
					" but got %q", repoSlug))
			}

			repoName := ""
			repoOwner := ""
			githubSlugParts := strings.Split(repoSlug, "/")
			if len(githubSlugParts) > 0 {
				repoOwner = githubSlugParts[0]
				repoName = githubSlugParts[1]
			} else {
				return errors.New(fmt.Sprintf("Expected repoSlug to be in the format of `owner/repo`"+
					" but got %q", repoSlug))
			}

			if schemaFile == "" && providerName == "" {
				providerName = strings.Replace(repoName, "pulumi-", "", -1)
			}

			if schemaFile == "" {
				schemaFile = fmt.Sprintf("provider/cmd/pulumi-resource-%s/schema.json", providerName)
			}

			// we should be able to take the repo URL + the version + the schema url and
			// construct a file that we can download and read
			schemaFilePath := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s",
				repoSlug, version, schemaFile)

			schema, err := readRemoteFile(schemaFilePath)
			if err != nil {
				return err
			}

			if schema == nil {
				return fmt.Errorf("unable to get contents of schemaFile %q", schemaFilePath)
			}

			// The source schema can be in YAML format. If that's the case
			// convert it to JSON first.
			if strings.HasSuffix(schemaFile, ".yaml") {
				schema, err = yaml.YAMLToJSON(schema)
				if err != nil {
					return errors.Wrap(err, "reading YAML schema")
				}
			}

			// try and get the version release data using the github releases API
			tags, err := getGitHubTags(repoSlug)
			if err != nil {
				return errors.Wrap(err, "github tags")
			}

			var commitDetails string
			for _, tag := range tags {
				if tag.Name == version {
					commitDetails = tag.Commit.URL
					break
				}
			}

			publishedDate := time.Now()
			if commitDetails != "" {
				var commit pkg.GitHubCommit
				// now let's make a request to the specific commit to get the date
				commitResp, err := http.Get(commitDetails)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("getting release info for %s", repoSlug))
				}

				defer commitResp.Body.Close()
				err = json.NewDecoder(commitResp.Body).Decode(&commit)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("constructing commit information for %s", repoSlug))
				}

				publishedDate = commit.Commit.Author.Date
			}

			mainSpec = &pschema.PackageSpec{}
			if err := json.Unmarshal(schema, mainSpec); err != nil {
				return errors.Wrap(err, "unmarshalling schema into a PackageSpec")
			}
			mainSpec.Version = version

			if mainSpec.Repository == "" {
				// we already know the repo slug so we can reconstruct the repository name using that
				mainSpec.Repository = fmt.Sprintf("https://github.com/%s", repoSlug)
			}

			status := pkg.PackageStatusGA
			if strings.HasPrefix(version, "v0.") {
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

			native := mainSpec.Attribution == ""
			// If native is false, check if the schema has the "kind/native" tag in the Keywords
			// array.
			if !native {
				native = isNative(mainSpec.Keywords)
			}

			if !component {
				component = isComponent(mainSpec.Keywords)
			}

			if native && component {
				native = false
			}

			// if there's a publisher then we need to use that immediately
			// if there is no publisher on cmd, then try and use packageSpec
			// if there's no publisher or packageSpec publisher, then assume repo owner is the publisher
			// otherwise error
			publisherName := ""
			if publisher != "" {
				publisherName = publisher
			} else if publisher == "" && mainSpec.Publisher != "" {
				publisherName = mainSpec.Publisher
			} else if publisher == "" && repoOwner != "" {
				publisherName = cases.Title(language.Und, cases.NoLower).String(repoOwner)
			} else {
				return errors.New("unable to determine package publisher")
			}

			cleanSchemaFilePath := func(s string) string {
				s = strings.ReplaceAll(s, "../", "")
				s = strings.ReplaceAll(s, fmt.Sprintf("pulumi-%s", mainSpec.Name), "")
				return s
			}

			pm := pkg.PackageMeta{
				Name:        mainSpec.Name,
				Description: mainSpec.Description,
				LogoURL:     mainSpec.LogoURL,
				Publisher:   publisherName,
				Title:       title,

				RepoURL:        mainSpec.Repository,
				SchemaFilePath: cleanSchemaFilePath(schemaFile),

				PackageStatus: status,
				UpdatedOn:     publishedDate.Unix(),
				Version:       version,

				Category:  category,
				Component: component,
				Featured:  isFeaturedPackage(mainSpec.Name),
				Native:    native,
			}
			b, err := yaml.Marshal(pm)
			if err != nil {
				return errors.Wrap(err, "generating package metadata")
			}

			if metadataDir == "" {
				// if the user hasn't specified an metadataDir, we will default to
				// the path within the registry folder.
				metadataDir = "themes/default/data/registry/packages"
			}
			metadataFileName := fmt.Sprintf("%s.yaml", mainSpec.Name)
			if err := pkg.EmitFile(metadataDir, metadataFileName, b); err != nil {
				return errors.Wrap(err, "writing metadata file")
			}

			if packageDocsDir == "" {
				// if the user hasn't specified an packageDocsDir, we will default to
				// the path within the registry folder.
				packageDocsDir = fmt.Sprintf("themes/default/content/registry/packages/%s", mainSpec.Name)
			}

			requiredFiles := []string{
				"_index.md",
				"installation-configuration.md",
			}
			for _, requiredFile := range requiredFiles {
				requiredFilePath := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/docs/%s",
					repoSlug, version, requiredFile)
				details, err := readRemoteFile(requiredFilePath)
				if err != nil {
					return err
				}

				if err := pkg.EmitFile(packageDocsDir, requiredFile, details); err != nil {
					return errors.Wrap(err, fmt.Sprintf("writing %s file", requiredFile))
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&repoSlug, "repoSlug", "", "The repository slug e.g. pulumi/pulumi-provider")
	cmd.Flags().StringVar(&providerName, "providerName", "", "The name of the provider e.g. aws, aws-native. "+
		"Required when there is no schemaFile flag specified.")
	cmd.Flags().StringVarP(&schemaFile, "schemaFile", "s", "", "Relative path to the schema.json file from "+
		"the root of the repository. If no schemaFile is specified, then providerName is required so the schemaFile path can "+
		"be inferred to be provider/cmd/pulumi-resource-<providerName>/schema.json")
	cmd.Flags().StringVar(&version, "version", "", "The version of the package")
	cmd.Flags().StringVar(&categoryStr, "category", "", fmt.Sprintf("The category for the package. Value must "+
		"match one of the keys in the map: %v", pkg.CategoryNameMap))
	cmd.Flags().StringVar(&publisher, "publisher", "", "The publisher's display name to be shown in the package. "+
		"This will default to Pulumi")
	cmd.Flags().StringVar(&title, "title", "", "The display name of the package. If omitted, the name of the "+
		"package will be used")
	cmd.Flags().BoolVar(&component, "component", false, "Whether or not this package is a component and not a provider")
	cmd.Flags().StringVar(&metadataDir, "metadataDir", "", "The location to save the metadata - this will default to the folder "+
		"structure that the registry expects (themes/default/data/registry/packages)")
	cmd.Flags().StringVar(&packageDocsDir, "packageDocsDir", "", "The location to save the package docs - this will default to the folder "+
		"structure that the registry expects (themes/default/content/registry/packages)")

	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("repoSlug")

	return cmd
}

func readRemoteFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("downloading remote file from %s", url))
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// this ultimately checks that the file exists and has content
		return nil, nil
	}

	contents, err := ioutil.ReadAll(resp.Body)
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
		if n, ok := pkg.CategoryNameMap[categoryOverrideStr]; !ok {
			return "", errors.New(fmt.Sprintf("invalid override for category name %s", categoryOverrideStr))
		} else {
			category = n
		}
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
	var category pkg.PackageCategory
	if n, ok := pkg.CategoryNameMap[categoryName]; !ok {
		return defaultPackageCategory, errors.New(fmt.Sprintf("invalid category tag %s", *categoryTag))
	} else {
		category = n
	}

	return category, nil
}

func isComponent(keywords []string) bool {
	return getTagFromKeywords(keywords, "kind/component") != nil
}

func isFeaturedPackage(str string) bool {
	for _, v := range featuredPackages {
		if v == str {
			return true
		}
	}
	return false
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
		return nil, errors.Wrap(err, fmt.Sprintf("getting tags info for %s", repoSlug))
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
		return nil, errors.Wrap(err, fmt.Sprintf("constructing tags information for %s", repoSlug))
	}

	return tags, nil
}
