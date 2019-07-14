package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("len(args) is %d want 3", len(args))
		os.Exit(2)
	}
	a := args[1]
	b := args[2]
	// a sha256
	as := sha256.Sum256([]byte(a))
	// b sha256
	bs := sha256.Sum256([]byte(b))

	var count int
	for i := range as {
		// xorでa, bのアンマッチをあぶり出し
		count += int(pc[as[i]^bs[i]])
	}
	fmt.Printf("(%q, %q) diff is %d", a, b, count)
}
