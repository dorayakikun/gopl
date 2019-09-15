package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn) {
	tz := os.Getenv("TZ")
	defer c.Close()
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%s\n", tz))
		if err != nil {
			return // e.g., client disconnected
		}
		_, err = io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var port = flag.String("port", "", "port number")

	flag.Parse()

	if *port == "" {
		log.Fatal("port flag is empty")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
