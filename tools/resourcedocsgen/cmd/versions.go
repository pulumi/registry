// Copyright 2026, Pulumi Corporation.
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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/blang/semver"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/publishers"
)

// discoveredVersion mirrors one entry of registry-mirror-discover's
// `--format json-array` output so scripts/versioned-docs/discover-versions.sh
// can consume either tool interchangeably. Version carries no leading "v".
type discoveredVersion struct {
	Version   string `json:"version"`
	SchemaURL string `json:"schema_url"`
}

type skippedVersion struct {
	Version string `json:"version"`
	Reason  string `json:"reason"`
}

type versionsReport struct {
	Package  string              `json:"package"`
	Versions []discoveredVersion `json:"versions"`
	Skipped  []skippedVersion    `json:"skipped"`
}

// VersionsCmd discovers the historical versions of a registry package for
// versioned docs generation. Unlike the private registry-mirror-discover tool
// (which covers blessed first-party providers), this works for any package
// with a GitHub repo_url: it enumerates git tags, picks the latest
// `--minor-count` minors of the last `--major-count` majors, and derives a
// version-pinned schema URL for each by substituting the version into the
// package's schema_file_url. Unresolvable versions are reported in `skipped`
// rather than failing the run (best-effort by design).
func VersionsCmd() *cobra.Command {
	var registryDir string
	var pkgName string
	var majorCount int
	var minorCount int
	var format string

	cmd := &cobra.Command{
		Use:   "versions",
		Short: "Discover a registry package's historical versions and schema URLs",
		RunE: func(cmd *cobra.Command, args []string) error {
			metadataPath := filepath.Join(getRegistryPackagesPath(registryDir), pkgName+".yaml")
			metadataBytes, err := os.ReadFile(metadataPath)
			if err != nil {
				return errors.Wrapf(err, "reading package metadata %s", metadataPath)
			}
			var metadata pkg.PackageMeta
			if err := yaml.Unmarshal(metadataBytes, &metadata); err != nil {
				return errors.Wrapf(err, "unmarshalling package metadata %s", metadataPath)
			}

			report, err := discoverVersions(metadata, majorCount, minorCount)
			if err != nil {
				return err
			}

			var out any = report
			if format == "json-array" {
				out = report.Versions
			} else if format != "json" {
				return fmt.Errorf("unknown --format %q (expected json or json-array)", format)
			}

			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(out)
		},
	}

	cmd.Flags().StringVar(&registryDir, "registryDir", ".", "The root of the registry repo")
	cmd.Flags().StringVar(&pkgName, "package", "", "The registry package name (matches the metadata YAML filename)")
	cmd.Flags().IntVar(&majorCount, "major-count", 3, "Number of most recent major versions to include")
	cmd.Flags().IntVar(&minorCount, "minor-count", 1, "Number of most recent minor versions to include per major")
	cmd.Flags().StringVar(&format, "format", "json",
		"Output format: json (versions + skipped report) or json-array (registry-mirror-discover compatible)")
	if err := cmd.MarkFlagRequired("package"); err != nil {
		panic(err)
	}

	return cmd
}

func getRegistryPackagesPath(registryDir string) string {
	return filepath.Join(registryDir, "themes", "default", "data", "registry", "packages")
}

func discoverVersions(metadata pkg.PackageMeta, majorCount, minorCount int) (*versionsReport, error) {
	repoSlug, err := repoSlugFromURL(metadata.RepoURL)
	if err != nil {
		return nil, err
	}

	tags, err := getAllGitHubTags(repoSlug)
	if err != nil {
		return nil, err
	}

	versions := selectVersions(tagNames(tags), majorCount, minorCount)
	if len(versions) == 0 {
		return nil, errors.Errorf("no semver release tags found for %s", repoSlug)
	}

	report := &versionsReport{
		Package:  metadata.Name,
		Versions: []discoveredVersion{},
		Skipped:  []skippedVersion{},
	}
	for _, v := range versions {
		schemaURL, err := resolveSchemaURL(metadata, v)
		if err != nil {
			report.Skipped = append(report.Skipped, skippedVersion{Version: v.String(), Reason: err.Error()})
			continue
		}
		report.Versions = append(report.Versions, discoveredVersion{Version: v.String(), SchemaURL: schemaURL})
	}
	return report, nil
}

func repoSlugFromURL(repoURL string) (string, error) {
	if repoURL == "" {
		return "", errors.New("package metadata does not contain a repo_url")
	}
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", errors.Wrapf(err, "parsing repo url %s", repoURL)
	}
	slug := strings.Trim(u.Path, "/")
	if strings.Count(slug, "/") != 1 {
		return "", errors.Errorf("expected repo_url path to be owner/repo, got %q", u.Path)
	}
	return slug, nil
}

// getAllGitHubTags pages through the tags API. Providers with long histories
// (e.g. aws) have hundreds of tags; a page cap bounds API usage while easily
// covering the last few majors, which are always the most recent tags.
func getAllGitHubTags(repoSlug string) ([]pkg.GitHubTag, error) {
	const perPage = 100
	const maxPages = 10

	var all []pkg.GitHubTag
	for page := 1; page <= maxPages; page++ {
		path := fmt.Sprintf("/repos/%s/tags?per_page=%d&page=%d", repoSlug, perPage, page)
		resp, err := pkg.GetGitHubAPI(path)
		if err != nil {
			return nil, errors.Wrapf(err, "getting tags for %s", repoSlug)
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			return nil, errors.Errorf("getting tags for %s: %s", repoSlug, resp.Status)
		}
		var tags []pkg.GitHubTag
		err = json.NewDecoder(resp.Body).Decode(&tags)
		resp.Body.Close()
		if err != nil {
			return nil, errors.Wrapf(err, "decoding tags for %s", repoSlug)
		}
		all = append(all, tags...)
		if len(tags) < perPage {
			break
		}
	}
	return all, nil
}

func tagNames(tags []pkg.GitHubTag) []string {
	names := make([]string, 0, len(tags))
	for _, t := range tags {
		names = append(names, t.Name)
	}
	return names
}

var releaseTagRe = regexp.MustCompile(`^v?\d+\.\d+\.\d+$`)

// selectVersions picks the latest `minorCount` minor lines of the last
// `majorCount` majors from a set of git tag names, taking the highest patch of
// each minor line. Pre-releases and non-semver tags are ignored. The result is
// sorted newest-first, matching registry-mirror-discover.
func selectVersions(tagNames []string, majorCount, minorCount int) []semver.Version {
	byMajor := map[uint64]map[uint64]semver.Version{}
	for _, name := range tagNames {
		if !releaseTagRe.MatchString(name) {
			continue
		}
		v, err := semver.Parse(strings.TrimPrefix(name, "v"))
		if err != nil || len(v.Pre) > 0 || len(v.Build) > 0 {
			continue
		}
		minors, ok := byMajor[v.Major]
		if !ok {
			minors = map[uint64]semver.Version{}
			byMajor[v.Major] = minors
		}
		if best, ok := minors[v.Minor]; !ok || v.GT(best) {
			minors[v.Minor] = v
		}
	}

	majors := make([]uint64, 0, len(byMajor))
	for m := range byMajor {
		majors = append(majors, m)
	}
	sort.Slice(majors, func(i, j int) bool { return majors[i] > majors[j] })
	if len(majors) > majorCount {
		majors = majors[:majorCount]
	}

	var out []semver.Version
	for _, m := range majors {
		minors := make([]uint64, 0, len(byMajor[m]))
		for minor := range byMajor[m] {
			minors = append(minors, minor)
		}
		sort.Slice(minors, func(i, j int) bool { return minors[i] > minors[j] })
		if len(minors) > minorCount {
			minors = minors[:minorCount]
		}
		for _, minor := range minors {
			out = append(out, byMajor[m][minor])
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GT(out[j]) })
	return out
}

// substituteVersionInURL rewrites a version-pinned schema URL to point at a
// different version. It replaces the path segment matching the package's
// current version — with or without the leading "v", covering both
// raw.githubusercontent.com tag paths (/v7.35.0/) and CDN paths for bridged
// OpenTofu providers (/0.14.3/). Returns "" when no segment matches.
func substituteVersionInURL(schemaFileURL, currentVersion string, target semver.Version) string {
	u, err := url.Parse(schemaFileURL)
	if err != nil || u.Path == "" {
		return ""
	}
	bare := strings.TrimPrefix(currentVersion, "v")
	segments := strings.Split(u.Path, "/")
	replaced := false
	for i, s := range segments {
		switch s {
		case "v" + bare:
			segments[i] = "v" + target.String()
			replaced = true
		case bare:
			segments[i] = target.String()
			replaced = true
		}
	}
	if !replaced {
		return ""
	}
	u.Path = strings.Join(segments, "/")
	return u.String()
}

// resolveSchemaURL finds a fetchable, version-pinned schema URL for the given
// version: first by substituting the version into schema_file_url (or the
// default schema path convention), then falling back to the Pulumi Registry
// API. The returned URL is validated with a request before being accepted.
func resolveSchemaURL(metadata pkg.PackageMeta, target semver.Version) (string, error) {
	var candidates []string

	if metadata.SchemaFileURL != "" {
		if substituted := substituteVersionInURL(metadata.SchemaFileURL, metadata.Version, target); substituted != "" {
			candidates = append(candidates, substituted)
		}
	}

	if len(candidates) == 0 && metadata.RepoURL != "" {
		// Same convention as getSchemaFileURL in cmd/docs/registry.go.
		schemaFilePath := fmt.Sprintf("provider/cmd/pulumi-resource-%s/schema.json", metadata.Name)
		if p := metadata.SchemaFilePath; p != "" { //nolint:staticcheck
			schemaFilePath = strings.TrimPrefix(p, "/")
		}
		if slug, err := repoSlugFromURL(metadata.RepoURL); err == nil {
			candidates = append(candidates,
				fmt.Sprintf("https://raw.githubusercontent.com/%s/v%s/%s", slug, target.String(), schemaFilePath))
		}
	}

	var lastErr error
	for _, candidate := range candidates {
		err := checkURLFetchable(candidate)
		if err == nil {
			return candidate, nil
		}
		lastErr = err
	}

	// Fall back to the Pulumi Registry API, which serves schemas for published
	// package versions (same endpoint used by getSchemaFromRegistry in
	// cmd/docs/registry.go).
	if registryURL, err := schemaURLFromRegistryAPI(metadata, target); err == nil {
		return registryURL, nil
	} else if lastErr == nil {
		lastErr = err
	} else {
		lastErr = errors.Wrapf(lastErr, "registry API fallback also failed: %v", err)
	}

	return "", errors.Wrapf(lastErr, "no fetchable schema URL for v%s", target.String())
}

func checkURLFetchable(rawURL string) error {
	req, err := http.NewRequest("HEAD", rawURL, nil)
	if err != nil {
		return errors.Wrapf(err, "creating request for %s", rawURL)
	}
	pkg.AddGitHubAuthHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "checking %s", rawURL)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return errors.Errorf("checking %s: %s", rawURL, resp.Status)
	}
	return nil
}

func schemaURLFromRegistryAPI(metadata pkg.PackageMeta, target semver.Version) (string, error) {
	backendURL := os.Getenv("PULUMI_BACKEND_URL")
	if backendURL == "" {
		backendURL = "https://api.pulumi.com/api"
	}

	source := "pulumi"
	if metadata.SchemaFileURL != "" {
		if parsed, err := url.Parse(metadata.SchemaFileURL); err == nil &&
			strings.HasPrefix(parsed.Path, "/schemas/registry.opentofu.org/") {
			source = "opentofu"
		}
	}

	publisher := publishers.GetPublisherName(metadata.Publisher)
	if publisher == "" {
		return "", errors.Errorf("publisher %q not found in publisher-names.json", metadata.Publisher)
	}

	apiURL := fmt.Sprintf("%s/registry/packages/%s/%s/%s/versions/%s",
		backendURL, source, publisher, metadata.Name, target.String())
	resp, err := http.Get(apiURL) //nolint:gosec // URL constructed from a predefined pattern.
	if err != nil {
		return "", errors.Wrapf(err, "querying registry API %s", apiURL)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return "", errors.Errorf("querying registry API %s: %s", apiURL, resp.Status)
	}

	var response struct {
		SchemaURL string `json:"schemaURL"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", errors.Wrap(err, "unmarshalling registry API response")
	}
	if response.SchemaURL == "" {
		return "", errors.Errorf("registry API returned no schemaURL for %s v%s", metadata.Name, target.String())
	}
	return response.SchemaURL, nil
}
