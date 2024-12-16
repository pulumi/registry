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

import "fmt"

// The source of download URLs for use during metadata generation.
type downloadURLs interface {
	schemaURL() string
	indexURL() string
	installationConfigurationURL() string
}

var (
	_ downloadURLs = githubDownloadURLs{}
	_ downloadURLs = rawDownloadURLs{}
)

// The set of URLs that can be downloaded from GitHub.
type githubDownloadURLs struct {
	repoSlug repoSlug

	providerName string
	version      string
	schemaFile   string
}

func (s githubDownloadURLs) schemaURL() string {
	if s.schemaFile == "" {
		s.schemaFile = fmt.Sprintf("provider/cmd/pulumi-resource-%s/schema.json", s.providerName)
	}

	// we should be able to take the repo URL + the version + the schema url and
	// construct a file that we can download and read
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s",
		s.repoSlug, s.version, s.schemaFile)
}

func (s githubDownloadURLs) docsURL(name string) string {
	return "https://raw.githubusercontent.com/" + s.repoSlug.String() + "/" + s.version + "/docs/" + name
}

func (s githubDownloadURLs) indexURL() string {
	return s.docsURL("_index.md")
}

func (s githubDownloadURLs) installationConfigurationURL() string {
	return s.docsURL("installation-configuration.md")
}

type rawDownloadURLs struct {
	schemaFileURL                    string
	indexFileURL                     string
	installationConfigurationFileURL string
}

func (s rawDownloadURLs) schemaURL() string { return s.schemaFileURL }
func (s rawDownloadURLs) indexURL() string  { return s.indexFileURL }
func (s rawDownloadURLs) installationConfigurationURL() string {
	return s.installationConfigurationFileURL
}
