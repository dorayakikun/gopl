package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(os.Stdout, url)
	}
}

func outline(w io.Writer, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(w, doc, startElement, endElement)

	return nil
}

func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

var depth int

func startElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Fprintf(w,"%*s<%s%s/>\n", depth*2, "", n.Data, attrstring(n.Attr))
		} else {
			fmt.Fprintf(w,"%*s<%s%s>\n", depth*2, "", n.Data, attrstring(n.Attr))
		}
		depth++
	}
}

func endElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Fprintf(w,"%*s</%s>\n", depth*2, "", n.Data)
	}
}

func attrstring(attr []html.Attribute) string {
	var sb strings.Builder
	for _, a := range attr {
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
	}
	return sb.String()
}