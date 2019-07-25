package github

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const IssuesURL = "https://api.github.com/search/issues"

// TODO POST /repos/:owner/:repo/issues
func CreateIssue(owner string, repo string) (error){
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo), nil)
	if err != nil {
		return errors.Wrap(err, "create req failed")
	}

	// cf. https://developer.github.com/v3/#current-version
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "req send failed")
	}

	if resp.StatusCode >= 400 {
		return errors.Errorf("create issue failed status code(%d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	fmt.Printf("%v", resp.Body)

	return nil
}
// TODO GET GET /repos/:owner/:repo/issues/:issue_number

// TODO PATCH /repos/:owner/:repo/issues/:issue_number