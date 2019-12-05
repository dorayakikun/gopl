package outline

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByID(t *testing.T) {
	s := `<div id="container"><p id="caption">caption</p></div>`
	n, _ := html.Parse(strings.NewReader(s))
	el := ElementByID(n, "caption")

	if el.Type != html.ElementNode {
		t.Errorf("el.Type is %v want %v", el.Type, html.ElementNode)
	}
	if el.Data != "p" {
		t.Errorf("el.Data is %s want %s", el.Data, "p")
	}
}
