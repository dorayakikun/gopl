package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const BaseURL = "https://api.github.com"

// POST /repos/:owner/:repo/issues
func PostIssue(owner string, repo string, token string, issue *PostIssueParameter) error {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", BaseURL, owner, repo)

	buf, err := json.Marshal(issue)
	if err != nil {
		return errors.Wrap(err, "json marshal failed")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(buf))
	if err != nil {
		return errors.Wrap(err, "create req failed")
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "req send failed")
	}

	if resp.StatusCode >= 400 {
		return errors.Errorf("post issue failed status code(%d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return errors.Wrap(err, "decode failed")
	}

	buf, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		return errors.Wrap(err, "marshal indent failed")
	}
	fmt.Printf("post issue succeed:\n%s", string(buf))

	return nil
}

// GET GET /repos/:owner/:repo/issues/:issue_number
func GetIssue(owner string, repo string, issueNumber string, token string) error {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", BaseURL, owner, repo, issueNumber)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "create req failed")
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "req send failed")
	}

	if resp.StatusCode >= 400 {
		return errors.Errorf("get issue failed status code(%d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return errors.Wrap(err, "decode failed")
	}

	buf, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return errors.Wrap(err, "marshal indent failed")
	}
	fmt.Println(string(buf))

	return nil
}

// PATCH /repos/:owner/:repo/issues/:issue_number
func PatchIssue(owner string, repo string, issueNumber string, token string, issue *PatchIssueParameter) error {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", BaseURL, owner, repo, issueNumber)

	buf, err := json.Marshal(issue)
	if err != nil {
		return errors.Wrap(err, "json marshal failed")
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(buf))
	if err != nil {
		return errors.Wrap(err, "create req failed")
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "req send failed")
	}

	if resp.StatusCode >= 400 {
		return errors.Errorf("patch issue failed status code(%d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return errors.Wrap(err, "decode failed")
	}

	buf, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		return errors.Wrap(err, "marshal indent failed")
	}
	fmt.Printf("patch issue succeed:\n%s", string(buf))

	return nil
}
