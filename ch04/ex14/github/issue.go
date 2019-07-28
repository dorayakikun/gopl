package github

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"time"
)

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssuesResult struct {
	Issues []*Issue
}

func FetchIssues(owner string, repo string) (IssuesResult, error){
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	fmt.Printf("GET %s\n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return IssuesResult{}, errors.Wrap(err, "fetch issues failed")
	}

	if resp.StatusCode >= 400 {
		return IssuesResult{}, fmt.Errorf("bad request issues (%d)", resp.StatusCode)
	}

	var ret IssuesResult
	if err := json.NewDecoder(resp.Body).Decode(&(ret.Issues)); err != nil {
		return IssuesResult{}, errors.Wrap(err, "decode issues failed")
	}
	return ret, nil
}

func WriteIssues(w http.ResponseWriter, result *IssuesResult) error {
	t := template.Must(template.New("issues").Parse(`
		<h1>Issues</h1>
		<table>
		  <tr>
			<th>#</th>
			<th>state</th>
			<th>user</th>
			<th>title</th>
		  </tr>
          {{range .Issues}}
          <tr>
  		    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  		    <td>{{.State}}</td>
		    <td><a href='{{.User.HTMLURL}}'><img alt="{{.User.Login}}" src="{{.User.AvatarURL}}" style="width: 64; height: 64;" /></a></td>
  		    <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
          </tr>
		{{end}}
		</table>
	`))

	if err := t.Execute(w, result); err != nil {
		return errors.Wrap(err, "write users failed")
	}
	return nil
}