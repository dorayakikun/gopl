package main

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//f, err := os.Open("foo.xml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal()
	}
}

func run(r io.Reader, w io.Writer) error {
	dec := xml.NewDecoder(r)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.Wrap(err, "xml decode failed")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			for _, a := range tok.Attr {
				if a.Name.Local == "id" {
					stack = append(stack, "#"+a.Value)
				}
				if a.Name.Local == "class" {
					stack = append(stack, "."+a.Value)
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Fprintf(w, "%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
	return nil
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
