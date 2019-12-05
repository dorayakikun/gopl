package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Value string
}

func (s Symbol) String() string {
	return s.Value
}

type String struct {
	Value string
}

func (s String) String() string {
	return fmt.Sprintf("%q", s.Value)
}

type Int struct {
	Value int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type StartList struct {
}

func (s StartList) String() string {
	return "("
}

type EndList struct {
}

func (e EndList) String() string {
	return ")"
}

type Decoder2 struct {
	l *lexer
}

func (d *Decoder2) Token() (Token, error) {
	switch d.l.token {
	case scanner.Ident:
		s := d.l.text()
		d.l.next()
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
	case scanner.EOF:
		return nil, io.EOF
	}
	return nil, fmt.Errorf("unexpected token %q", d.l.text())
}
