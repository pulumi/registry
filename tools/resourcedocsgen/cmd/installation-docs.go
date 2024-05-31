package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/docs/tools/resourcedocsgen/pkg"
	"github.com/ryboe/q"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
)

func packageInstallationDocsCmd() *cobra.Command {
	var repoSlug string
	var packageDocsDir string
	var version string
	cmd := &cobra.Command{
		Use:   "installation-docs <args>",
		Short: "Generate installation docs from required docs files",
		RunE: func(cmd *cobra.Command, args []string) error {

			if strings.Contains(repoSlug, "https") || strings.Contains(repoSlug, "github.com") {
				return errors.New(fmt.Sprintf("Expected repoSlug to be in the format of `owner/repo`"+
					" but got %q", repoSlug))
			}

			if packageDocsDir == "" {
				// if the user hasn't specified a packageDocsDir, default to
				// the path within the registry folder.
				packageDocsDir = fmt.Sprintf("themes/default/content/registry/packages/%s", mainSpec.Name)
			}

			requiredFiles := []string{
				"_index.md",
				"installation-configuration.md",
			}
			for _, requiredFile := range requiredFiles {
				//requiredFilePath := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/docs/%s",
				//	repoSlug, version, requiredFile)
				//details, err := readRemoteFile(requiredFilePath)
				requiredFilePath := fmt.Sprintf("%s",
					requiredFile)
				details, err := os.ReadFile(requiredFilePath)
				q.Q(string(details), err)
				if err != nil {
					return err
				}

				if err := pkg.EmitFile(packageDocsDir, requiredFile, details); err != nil {
					panic("wtf?")
					return errors.Wrap(err, fmt.Sprintf("writing %s file", requiredFile))
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&version, "version", "", "The version of the package")
	cmd.Flags().StringVar(&repoSlug, "repoSlug", "", "The repository slug e.g. pulumi/pulumi-provider")
	cmd.Flags().StringVar(&packageDocsDir, "packageDocsDir", "", "The location to save the package docs - this will default to the folder "+
		"structure that the registry expects (themes/default/content/registry/packages)")

	cmd.MarkFlagRequired("repoSlug")
	cmd.MarkFlagRequired("version")

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
		return nil, errors.New(fmt.Sprintf("finding remote file at %s: %s", url, resp.Status))
	}

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading contents of remote file")
	}

	return contents, nil
}
