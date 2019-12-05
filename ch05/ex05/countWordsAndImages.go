package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing url")
		fmt.Println("usage: go run countWordsAndImage.go <URL>")
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Printf("count words and images failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("words: %d, images: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	counts := make(map[string]int)

	counts = freqCount(counts, n)

	words = counts["words"]
	images = counts["images"]
	return
}

func freqCount(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode && (n.Data != "script" || n.Data == "style") {
		if c := n.FirstChild; c != nil {
			if c.Type == html.TextNode {
				counts["words"] += len(strings.Split(c.Data, " "))
			}
		}
	}
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		counts["images"]++
	}
	if c := n.FirstChild; c != nil {
		counts = freqCount(counts, c)
	}
	if s := n.NextSibling; s != nil {
		counts = freqCount(counts, s)
	}
	return counts
}
