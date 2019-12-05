package github

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"time"
)

type Milestone struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Creator   *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type MilestonesResult struct {
	Milestones []*Milestone
}

func FetchMilestones(owner string, repo string) (MilestonesResult, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/milestones", owner, repo)
	fmt.Printf("GET %s\n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return MilestonesResult{}, errors.Wrap(err, "fetch milestones failed")
	}

	if resp.StatusCode >= 400 {
		return MilestonesResult{}, fmt.Errorf("bad request milestones (%d)", resp.StatusCode)
	}

	var ret MilestonesResult
	if err := json.NewDecoder(resp.Body).Decode(&(ret.Milestones)); err != nil {
		return MilestonesResult{}, errors.Wrap(err, "decode milestones failed")
	}
	return ret, nil
}

func WriteMilestones(w http.ResponseWriter, result *MilestonesResult) error {
	t := template.Must(template.New("issues").Parse(`
		<h1>Milestones</h1>
		<table>
          <tr>
			<th>#</th>
			<th>state</th>
			<th>creator</th>
			<th>title</th>
		  </tr>
		  {{range .Milestones}}
          <tr>
		    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		    <td>{{.State}}</td>
		    <td><a href='{{.Creator.HTMLURL}}'><img alt="{{.Creator.Login}}" src="{{.Creator.AvatarURL}}" style="width 64; height: 64;" /></a></td>
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
