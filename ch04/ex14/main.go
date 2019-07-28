package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"z/github"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		urlPart := strings.Split(r.URL.Path, "/")

		if len(urlPart) != 3 {
			fmt.Fprintln(w, "usage: http://localhost:8000/:owner/repo")
			return
		}

		owner := urlPart[1]
		repo := urlPart[2]

		// issues
		issues, err := github.FetchIssues(owner, repo)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		if err := github.WriteIssues(w, &issues); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}

		// milestones
		milestones, err := github.FetchMilestones(owner, repo)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		if err := github.WriteMilestones(w, &milestones); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}

		// users
		users, err := github.FetchUsers()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := github.WriteUsers(w, &users); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
