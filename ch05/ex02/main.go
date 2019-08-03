package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	fmt.Println("tag\tcount")
	for k, v := range countElementTypes(elements, doc) {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func countElementTypes(elements map[string]int, n * html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	if c := n.FirstChild; c != nil {
		countElementTypes(elements, c)
	}
	if s := n.NextSibling; s != nil {
		countElementTypes(elements, s)
	}
	return elements
}