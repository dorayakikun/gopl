package main

import (
	"fmt"
	"os"
	"time"
	"z/github"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go <KEYWORD>")
		os.Exit(2)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Printf("search failed: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("total count:%d\n", result.TotalCount)

	dateRanges := make(map[string][]*github.Issue)
	for _, item := range result.Items {
		duration := time.Since(item.CreatedAt).Hours()
		if duration < 744 {
			dateRanges["past month"] = append(dateRanges["past month"], item)
		} else if duration < 8760 {
			dateRanges["past year"] = append(dateRanges["past year"], item)
		} else {
			dateRanges["more old"] = append(dateRanges["more old"], item)
		}
	}

	for k, v := range dateRanges {
		fmt.Printf("[%s]\n", k)
		for _, item := range v {
			fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.RFC3339))
		}
	}
}
