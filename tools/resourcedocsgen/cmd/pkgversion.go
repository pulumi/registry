package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/spf13/cobra"

	"io/ioutil"
	"net/http"
	"strings"
)

func CheckVersion() *cobra.Command {

	var repoSlug string
	cmd := &cobra.Command{
		Use:   "pkgversion",
		Short: "Check a Pulumi package version",
		Long: `Print the most recent version of a Pulumi package. If the most recent version of a Pulumi package is the same as the version published in the Pulumi Registry, print nothing.

The most recent version of a Pulumi package is taken to be the most recent release published in GitHub.

The version in the registry is defined in the YAML file at:

    https://raw.githubusercontent.com/pulumi/registry/master/themes/default/data/registry/packages/${PKG#pulumi/pulumi-}.yaml`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if strings.Contains(repoSlug, "https") || strings.Contains(repoSlug, "github.com") {
				return errors.New(fmt.Sprintf("Expected repoSlug to be in the format of `owner/repo`"+
					" but got %q", repoSlug))
			}

			repoName := ""
			githubSlugParts := strings.Split(repoSlug, "/")
			if len(githubSlugParts) > 0 {
				repoName = githubSlugParts[1]
			}

			version, err := getLatestVersion(repoSlug)
			if err != nil {
				return err
			}

			pkgName := strings.TrimPrefix(repoName, "pulumi-")
			pkgMetadata := fmt.Sprintf("https://raw.githubusercontent.com/pulumi/registry/master/themes/default/data/registry/packages/%s.yaml", pkgName)
			regVersion, err := getRegistryVersion(pkgMetadata)
			if err != nil {
				return err
			}

			// print version tag if there's a difference, and not if there isn't
			// we assume that the published latest version from the provider repo is the desired one, so any difference
			// between versions should indicate an update to the registry version
			if version != regVersion {
				fmt.Println(version)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&repoSlug, "repoSlug", "", "The repository slug e.g. pulumi/pulumi-provider")
	cmd.MarkFlagRequired("repoSlug")
	return cmd
}

func getLatestVersion(repoSlug string) (string, error) {
	path := fmt.Sprintf("/repos/%s/releases/latest", repoSlug)
	resp, err := pkg.GetGitHubAPI(path)

	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("getting latest version from https://api.github.com%s", path))
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.Wrap(err, fmt.Sprintf("Could not find a release at https://api.github.com%s", path))
	}

	var tag pkg.GitHubTag
	err = json.NewDecoder(resp.Body).Decode(&tag)

	if err != nil {
		return "", errors.Wrap(err, "failure reading contents of latest tag")
	}

	return tag.Name, nil
}

func getRegistryVersion(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("getting latest version from %s", url))
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.Wrap(err, "file not found")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failure reading contents of remote file")
	}

	var meta pkg.PackageMeta
	err = yaml.Unmarshal(contents, &meta)
	if err != nil {
		return "", errors.Wrap(err, "error unmarshalling yaml file")
	}

	return meta.Version, nil
}
