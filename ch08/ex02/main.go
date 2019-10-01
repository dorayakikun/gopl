package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
	}
}