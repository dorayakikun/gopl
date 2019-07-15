package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var mode = flag.Int("mode", 0, "0: SHA256, 1: SHA384, 2: SHA512")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 2 {
		fmt.Printf("len(args) is %d want 2\n", len(args))
		os.Exit(2)
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		switch *mode {
		case 0:
			fmt.Printf("%q is %x\t(SHA256)\n", s, sha256.Sum256([]byte(s)))
		case 1:
			fmt.Printf("%q is %x\t(SHA384)\n", s, sha512.Sum384([]byte(s)))
		case 2:
			fmt.Printf("%q is %x\t(SHA512)\n", s, sha512.Sum512([]byte(s)))
		default:
			fmt.Printf("mode is %d want 0-2\n", *mode)
			os.Exit(2)
		}
	}
}
