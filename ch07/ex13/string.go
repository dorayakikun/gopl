package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return fmt.Sprintf("Var: %s", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("literal: %g", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("unary: x: %s op: %s", u.x, string(u.op))
}

func (b binary) String() string {
	return fmt.Sprintf("binary: x: %s y: %s op: %s", b.x, b.y, string(b.op))
}

func (c call) String() string {
	b := strings.Builder{}

	b.WriteString("call:\n")
	b.WriteString(fmt.Sprintf("\tfn: %s\n", c.fn))

	if len(c.args) > 0 {
		b.WriteString("\targs: [\n")
		for _, a := range c.args {
			b.WriteString(fmt.Sprintf("\t\t%s\n", a))
		}
		b.WriteString("\t]\n")
	}
	return b.String()
}