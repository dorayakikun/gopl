package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type StringReader struct {
	s string
	i int
	p int
}

func NewReader(sd string, p int) *StringReader {
	return &StringReader{sd, 0, p }
}

func (d *StringReader) Read(b []byte) (n int, err error) {
	if d.i >= len(d.s) {
		err = io.EOF
		return
	}

	end := d.i + d.p
	if end >= len(d.s) {
		n = copy(b, []byte(d.s[d.i:]))
	} else {
		n = copy(b, []byte(d.s[d.i:end]))
	}
	d.i += n
	return
}

func main() {
	for _, domstring := range os.Args[1:] {
		outline(domstring)
	}
}

func outline(domstring string) error {
	// 10byteずつ読み込むように
	doc, err := html.Parse(NewReader(domstring, 10))
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
