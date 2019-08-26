package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {

	var r io.Reader
	r = os.Stdin
	dec := xml.NewDecoder(r)

	stack := []*Element{
		{}, // dummy root
	} // stack of element names

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		top := stack[len(stack)-1]
		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{Type: tok.Name, Attr: tok.Attr}
			top.Children = append(top.Children, el)
			stack = append(stack, el)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			top.Children = append(top.Children, CharData(tok.Copy()))
		}
	}
	for _, child := range stack[0].Children {
		_, ok := child.(*Element)
		if ok {
			printNode(os.Stdout, child, "")
		}
	}
}

func printNode(w io.Writer, node Node, prefix string) {
	switch n := node.(type) {
	case *Element:
		fmt.Fprintf(w, "%s<%s>\n", prefix, n.Type.Local)
		for _, child := range n.Children {
			printNode(w, child, "  " + prefix)
		}
		fmt.Fprintf(w, "%s</%s>\n", prefix, n.Type.Local)
	case CharData:
		fmt.Fprintf(w, "%s%s\n", prefix, string(n))
	default:
		panic("unexpected")
	}
}

func rootElement(dec *xml.Decoder) Node {
	tok, err := dec.Token()
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	return tok
}

func printChildren(depth int, children []Node) {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			fmt.Printf("%s%s\n", strings.Repeat("\t", depth), c.Type.Local)
			printChildren(depth+1, c.Children)
		case CharData:
			fmt.Printf("%s%s\n", strings.Repeat("\t", depth), c)
		}
	}
}
