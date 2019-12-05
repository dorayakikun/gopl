package main

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	children, err := run(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	for _, child := range children {
		_, ok := child.(*Element)
		if ok {
			printNode(os.Stdout, child, "")
		}
	}
}

func run(r io.Reader, w io.Writer) ([]Node, error) {
	dec := xml.NewDecoder(r)

	stack := []*Element{
		{}, // dummy root
	} // stack of element names

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "decode failed")
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
	return stack[0].Children, nil
}

func printNode(w io.Writer, node Node, prefix string) {
	switch n := node.(type) {
	case *Element:
		fmt.Fprintf(w, "%s<%s>\n", prefix, n.Type.Local)
		for _, child := range n.Children {
			printNode(w, child, "  "+prefix)
		}
		fmt.Fprintf(w, "%s</%s>\n", prefix, n.Type.Local)
	case CharData:
		fmt.Fprintf(w, "%s%s\n", prefix, string(n))
	default:
		panic("unexpected")
	}
}
