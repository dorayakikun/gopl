package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		log.Fatal("missing clock")
	}

	for _, arg := range args[1:] {
		s := strings.Split(arg, "=")
		if len(s) != 2 {
			log.Fatalf("iligal format :%s e.g. NewYork=localhost:8080", arg)
		}
		go func(city, address string) {
			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		}(s[0], s[1])
	}
	for {
		if false {
			break
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
