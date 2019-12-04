package sexpr

import (
	"fmt"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Value string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct {
}

type EndList struct {
}

type Decoder2 struct {
	l *lexer
}

func (d *Decoder2) Token() (Token, error) {
	switch d.l.token {
	case scanner.Ident:
		s := d.l.text()
		return Symbol{s}, nil
	case scanner.String:
		s, _ := strconv.Unquote(d.l.text()) // NOTE: ignoring errors
		d.l.next()
		return String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(d.l.text()) // NOTE: ignoring errors
		d.l.next()
		return Int{i}, nil
	case '(':
		d.l.next()
		return StartList{}, nil
	case ')':
		d.l.next()
		return EndList{}, nil
	}
	return nil, fmt.Errorf("unexpected token %q", d.l.text())
}
