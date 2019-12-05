package outline

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el := forEachNode(c, id, pre, post)
		if el != nil {
			return el
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

func attrstring(attr []html.Attribute) string {
	var sb strings.Builder
	for _, a := range attr {
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
	}
	return sb.String()
}
