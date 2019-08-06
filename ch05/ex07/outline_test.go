package main

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestForEachNode(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			s: `<p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul>`,
			expected: `<html>
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
`,
		},
		{
			s: "",
			expected: `<html>
  <head/>
  <body/>
</html>
`,
		},
	}

	for _, d := range data {
		n, _ := html.Parse(strings.NewReader(d.s))
		var buf bytes.Buffer
		forEachNode(&buf, n, startElement, endElement)
		if buf.String() != d.expected {
			t.Errorf("\n%q \nwant:\n %q", buf.String(), d.expected)
		}
	}
}