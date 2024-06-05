package pkg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

func GetGitHubAPI(path string) (*http.Response, error) {
	token := os.Getenv("GITHUB_TOKEN")
	url := fmt.Sprintf("https://api.github.com%s", path)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	return client.Do(req)
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
