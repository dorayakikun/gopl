package github

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
)

type User struct {
	ID        int
	Login     string
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}
type UsersResult struct {
	Users []*User
}

const url = "https://api.github.com/users"

func FetchUsers() (UsersResult, error) {
	fmt.Printf("GET %s\n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return UsersResult{}, errors.Wrap(err, "fetch users failed")
	}

	if resp.StatusCode >= 400 {
		return UsersResult{}, fmt.Errorf("bad request users (%d)", resp.StatusCode)
	}

	var ret UsersResult
	if err := json.NewDecoder(resp.Body).Decode(&(ret.Users)); err != nil {
		return UsersResult{}, errors.Wrap(err, "decode user failed")
	}
	return ret, nil
}

func WriteUsers(w http.ResponseWriter, result *UsersResult) error {
	t := template.Must(template.New("users").Parse(`
		<h1>Users</h1>
		<table>
		  <tr>
		    <th>name</th>
			<th>avatar</th>
		  </tr>
		  {{range .Users}}
		  <tr>
		    <td>{{.Login}}</td>
		    <td><a href="{{.HTMLURL}}"><img alt="{{.Login}}" src="{{.AvatarURL}}" style="width: 64;height: 64;"></img></a></td>
		  <tr>
		  {{end}}
		<table>
	`))

	if err := t.Execute(w, result); err != nil {
		return errors.Wrap(err, "write users failed")
	}
	return nil
}
