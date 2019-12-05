package elementbytagname

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByTagName(t *testing.T) {
	data := []struct {
		s         string
		tagname   []string
		len       int
		explected []string
	}{
		{
			s:         `<div id="container"><p id="caption">caption</p></div>`,
			tagname:   []string{"p"},
			len:       2,
			explected: []string{"p", "p"},
		},
		{
			s:         `<div id="container"><a href="#"><img src="hoge.png" /></a></div>`,
			tagname:   []string{"img", "a"},
			len:       4,
			explected: []string{"a", "img", "img", "a"},
		},
	}

	for _, d := range data {
		n, _ := html.Parse(strings.NewReader(d.s))
		nodes := ElementByTagName(n, d.tagname...)
		if len(nodes) != d.len {
			t.Errorf("len is %d want %d\n", len(nodes), d.len)
			return
		}

		for i, node := range nodes {
			if node.Data != d.explected[i] {
				t.Errorf("node.Data is %s want %s\n", node.Data, d.explected[i])
				return
			}
		}
	}
}
