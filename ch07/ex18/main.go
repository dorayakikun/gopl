package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element // stack of element names

	//el := rootElement(dec)
	//root, ok := el.(Element)
	//if !ok {
	//	fmt.Printf("%t %+v", ok, root)
	//	os.Exit(0)
	//}
	stack = append(stack, Element{})

  	var hoge Element
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
			el := Element{ Type: tok.Name, Attr: tok.Attr}
			top.Children = append(top.Children, el)
			stack = append(stack, el)
		case xml.EndElement:
			hoge = top
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			top.Children = append(top.Children, CharData(tok.Copy()))
		}
	}
	fmt.Println(hoge)
	//printChildren(1, root.Children)
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
