package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(s string) {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		*w += WordCounter(1)
	}
}

type RowCounter int

func (r *RowCounter) Write(s string) {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		*r += RowCounter(1)
	}
}

func main() {
	var w WordCounter
	w.Write("hello hello hello ")
	fmt.Printf("w:\t%d\n", w)

	var r RowCounter
	r.Write("Hello\nHi\nJa\nFoo")
	fmt.Printf("r:\t%d\n", r)
}
