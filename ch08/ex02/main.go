package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func handleCommand(writer net.Conn, line string) {
	s := strings.Split(line, " ")
	command := s[0]
	switch command {
	case "USER":
		writer.Write([]byte("331 Password Required\n"))
	case "PASS":
		writer.Write([]byte("230 Logged in\n"))
	}
}

func handleClient(conn net.Conn) {

	reader := bufio.NewReader(conn)

	conn.Write([]byte("220 Welcome to this FTP server!\n"))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Print(err)
			}
		}
		handleCommand(conn, line)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleClient(conn)
	}
}
