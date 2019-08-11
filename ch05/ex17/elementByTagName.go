package elementbytagname

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	return forEachNode(doc, name, startElement, endElement)
}

func forEachNode(n *html.Node, name []string, pre, post func(n *html.Node, name []string) bool) []*html.Node {
	var nodes []*html.Node
	if pre != nil {
		if pre(n, name) {
			nodes = append(nodes, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el := forEachNode(c, name, pre, post)
		if el != nil {
			nodes = append(nodes, el...)
		}
	}

	if post != nil {
		if post(n, name) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func startElement(n *html.Node, name []string) bool {
	if n.Type == html.ElementNode {
		for _, v := range name {
			if n.Data == v {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, name []string) bool {
	if n.Type == html.ElementNode {
		for _, v := range name {
			if n.Data == v {
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