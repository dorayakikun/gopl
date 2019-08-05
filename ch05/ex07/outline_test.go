package main

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
	"testing"
)

func TestForEachNode(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul>`
	expected := `<html>
  <head/>
  <body>
    <p>
    </p>
    <ul>
      <li>
        <a href="foo">
        </a>
      </li>
      <li>
        <a href="/bar/baz">
        </a>
      </li>
    </ul>
  </body>
</html>
`
	n, _ := html.Parse(strings.NewReader(s))

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("create pipe failed: %s\n", err.Error())
	}
	defer r.Close()

	stdout := os.Stdout
	os.Stdout = w

	forEachNode(n, startElement, endElement)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != expected {
		t.Errorf("\n%q \nwant:\n %q", buf.String(), expected)
	}
}