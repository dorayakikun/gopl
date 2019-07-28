package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
	"z/github"
)

func main() {
	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	if token == "" {
		fmt.Println("missing Personal Access Token")
		fmt.Println("export GITHUB_PERSONAL_ACCESS_TOKEN=<YOUR ACCESS TOKEN>")
		os.Exit(2)
	}

	if len(os.Args) == 1 {
		showUsage()
		return
	}

	create := flag.NewFlagSet("create", flag.ExitOnError)
	// create message
	cm := create.String("m", "", "message")
	// create labels
	cl := create.String("l", "","labels")

	get := flag.NewFlagSet("get", flag.ExitOnError)

	edit := flag.NewFlagSet("edit", flag.ExitOnError)
	// edit message
	em := edit.String("m", "", "message")
	// edit labels
	el := edit.String( "l", "", "labels")

	close := flag.NewFlagSet("close", flag.ExitOnError)

	switch os.Args[1] {
	case "create":
		create.Parse(os.Args[2:])
	case "get":
		get.Parse(os.Args[2:])
	case "edit":
		edit.Parse(os.Args[2:])
	case "close":
		close.Parse(os.Args[2:])
	}

	if create.Parsed() {
		// go run main.go create [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO>
		if len(create.Args()) != 2 {
			fmt.Printf("missing args\nusage: go run main.go create [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO>\n")
			os.Exit(2)
		}

		var title string
		var body string
		if *cm == "" {
			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "vi"
			}

			t := time.Now()
			fname := fmt.Sprintf("/tmp/go-study-tmp%d", t.Unix())
			cmd := exec.Command(editor, fname)

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			defer os.Remove(fname)
			if err != nil {
				fmt.Printf("open editor failed")
			}

			buf, err := ioutil.ReadFile(fname)
			if err != nil {
				fmt.Printf("open file failed\n")
				os.Exit(1)
			}
			lines := strings.Split(string(buf), "\n")

			if len(lines) > 3 {
				title = lines[0]
				body = strings.Join(lines[2:], "\n")
			} else {
				title = lines[0]
			}

		} else {
			title = *cm
		}

		var bodyPtr *string
		if body == "" {
			bodyPtr = nil
		} else {
			bodyPtr = &body
		}

		var clPtr *[]string
		if *cl == "" {
			clPtr = nil
		} else {
			cl := strings.Split(*cl, ",")
			clPtr = &cl
		}

		err := github.PostIssue(create.Arg(0), create.Arg(1), token, &github.PostIssueParameter{title, bodyPtr, clPtr})
		if err != nil {
			fmt.Printf("errr: %v", err.Error())
		}
		return
	}

	if get.Parsed() {
		// go run main.go get <OWNER> <REPO> <ISSUE_NUMBER>
		if len(get.Args()) != 3 {
			fmt.Printf("missing args\nusage: go run main.go get <OWNER> <REPO> <ISSUE_NUMBER>\n")
			os.Exit(2)
		}
		err := github.GetIssue(get.Arg(0), get.Arg(1), get.Arg(2), token)
		if err != nil {
			fmt.Printf("errr: %v", err.Error())
		}
		return
	}
	if edit.Parsed() {
		// go run main.go edit [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO> <ISSUE_NUMBER>
		if len(edit.Args()) != 3 {
			fmt.Printf("missing args\nusage: go run main.go edit [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO> <ISSUE_NUMBER>\n")
			os.Exit(2)
		}
		var title string
		var body string
		if *em == "" {
			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "vi"
			}

			t := time.Now()
			fname := fmt.Sprintf("/tmp/go-study-tmp%d", t.Unix())
			cmd := exec.Command(editor, fname)

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			defer os.Remove(fname)
			if err != nil {
				fmt.Printf("open editor failed")
			}

			buf, err := ioutil.ReadFile(fname)
			if err != nil {
				fmt.Printf("open file failed\n")
				os.Exit(1)
			}
			lines := strings.Split(string(buf), "\n")

			if len(lines) > 3 {
				title = lines[0]
				body = strings.Join(lines[2:], "\n")
			} else {
				title = lines[0]
			}

		} else {
			title = *em
		}

		var bodyPtr *string
		if body == "" {
			bodyPtr = nil
		} else {
			bodyPtr = &body
		}

		var elPtr *[]string
		if *el == "" {
			elPtr = nil
		} else {
			el := strings.Split(*el, ",")
			elPtr = &el
		}

		err := github.PatchIssue(edit.Arg(0), edit.Arg(1), edit.Arg(2), token, &github.PatchIssueParameter{&title, bodyPtr, nil, elPtr})
		if err != nil {
			fmt.Printf("errr: %v", err.Error())
		}
	}

	if close.Parsed() {
		// go run main.go close <OWNER> <REPO> <ISSUE_NUMBER>
		if len(close.Args()) != 3 {
			fmt.Printf("missing args\nusage: go run main.go close <OWNER> <REPO> <ISSUE_NUMBER>\n")
			os.Exit(2)
		}
		state := "closed"
		err := github.PatchIssue(close.Arg(0), close.Arg(1), close.Arg(2), token, &github.PatchIssueParameter{nil, nil, &state, nil})
		if err != nil {
			fmt.Printf("errr: %v", err.Error())
		}
		return
	}
}

func showUsage() {
	fmt.Println("usage: go run main.go create [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO>")
	fmt.Println("       go run main.go get <OWNER> <REPO> <ISSUE_NUMBER>")
	fmt.Println("       go run main.go edit [-m <MESSAGE>] [-l <LABELS>] <OWNER> <REPO> <ISSUE_NUMBER>")
	fmt.Println("       go run main.go close <OWNER> <REPO> <ISSUE_NUMBER>")
}