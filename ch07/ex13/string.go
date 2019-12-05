package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return fmt.Sprintf("%s", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x, string(b.op), b.y)
}

func (c call) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s(", c.fn))

	if len(c.args) > 0 {
		b.WriteString(fmt.Sprintf("%s", c.args[0]))
		for _, a := range c.args[1:] {
			b.WriteString(", ")
			b.WriteString(fmt.Sprintf("%s", a))
		}
	}
	b.WriteString(")")
	return b.String()
}
