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

package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

func GetGitHubAPI(path string) (*http.Response, error) {
	token := os.Getenv("GITHUB_TOKEN")
	url := "https://api.github.com" + path

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	return client.Do(req)
}

// AddGitHubAuthHeaders adds GitHub authentication headers to the request if the request is for GitHub domains
// and the GITHUB_TOKEN environment variable is set.
func AddGitHubAuthHeaders(req *http.Request) {
	// Check if the request is for GitHub domains
	host := req.URL.Host
	if host == "github.com" || host == "api.github.com" || host == "raw.githubusercontent.com" {
		// Add GitHub token from environment variable if available
		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			req.Header.Add("Authorization", "Bearer "+token)
		}
	}
}

type GitHubTag struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID string `json:"node_id"`
}

type GitHubCommit struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}
