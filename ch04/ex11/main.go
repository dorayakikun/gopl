package main

import (
	"fmt"
	"github.com/dorayakikun/go-study/ch04/ex11/github"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	if token == "" {
		fmt.Println("missing Personal Access Token")
		os.Exit(2)
	}

	// TODO -mがなければ任意のエディタを開く
	// go run main.go create <OWNER> <REPO> [-m <MESSAGE>] [-l <LABELS>]
	//body := "bodyです"
	//labels := []string{"E-easy", "I-ICE"}
	//err := github.CreateIssue("dorayakikun", "go-study", token, &github.CreateIssueParameter{"aaaaaaaaaaaaaaa", &body, &labels})
	//if err != nil {
	//	fmt.Printf("errr: %v", err.Error())
	//}

	// go run main.go get <OWNER> <REPO> <ISSUE_NUMBER>
	err := github.GetIssue("dorayakikun", "go-study", "8", token)
	if err != nil {
		fmt.Printf("errr: %v", err.Error())
	}

	// TODO -mがなければ任意のエディタを開く
	// go run main.go edit <OWNER> <REPO> <ISSUE_NUMBER> [-m <MESSAGE>] [-l <LABELS>]
	//body := "bodyです part2"
	//err := github.PatchIssue("dorayakikun", "go-study", "8", token, &github.EditIssueParameter{nil, &body, nil, nil})
	//if err != nil {
	//	fmt.Printf("errr: %v", err.Error())
	//}

	// go run main.go close <OWNER> <REPO> <ISSUE_NUMBER>
	//state := "closed"
	//err := github.PatchIssue("dorayakikun", "go-study", "8", token, &github.EditIssueParameter{nil, nil, &state, nil})
	//if err != nil {
	//	fmt.Printf("errr: %v", err.Error())
	//}
}
