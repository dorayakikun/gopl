package main

import (
	"fmt"
	"os"
)

var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

type CircularError struct {
	path []string
}

func (e *CircularError) Error() string {
	return fmt.Sprintf("found circular path: %q", e.path)
}

func main() {
	courses, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool, stack []string) error

	visitAll = func(items map[string]bool, stack []string) error {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				stack = append(stack, item)
				err := visitAll(m[item], stack)
				if err != nil {
					return err
				}
				order = append(order, item)
				// 探査を終えたタイミングでstackを初期化
				stack = nil
			} else {
				for _, s := range stack {
					// 現在探索中のstackにitemsがいたらエラーとする
					if item == s {
						return &CircularError{
							path: append(append([]string(nil), stack...), item),
						}
					}
				}
			}
		}
		return nil
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	var stack []string
	err := visitAll(keys, stack)
	if err != nil {
		return order, err
	}
	return order, nil
}
