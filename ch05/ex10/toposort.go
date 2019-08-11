package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string] []string)[]string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(m map[string][]string)
	visitAll = func(m map[string][]string) {
		for k, _ := range m {
			if !seen[k] {
				seen[k] = true

				order = append(order, k)
			}
		}
	}

	visitAll(m)

	return order
}