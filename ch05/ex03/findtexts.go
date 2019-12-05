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
	for _, text := range findTexts(nil, doc) {
		fmt.Println(text)
	}
}

func findTexts(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data != "script" && n.Data != "style") {
		if c := n.FirstChild; c != nil {
			if c.Type == html.TextNode {
				texts = append(texts, c.Data)
			}
		}
	}
	if c := n.FirstChild; c != nil {
		texts = findTexts(texts, c)
	}
	if s := n.NextSibling; s != nil {
		texts = findTexts(texts, s)
	}
	return texts
}
