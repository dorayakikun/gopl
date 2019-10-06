package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleCommand(conn net.Conn, line string, cd chan<- string) {
	writer := bufio.NewWriter(conn)

	s := strings.Split(line, " ")

	if len(s) == 0 {
		log.Print("missing command")
		return
	}

	command := s[0]
	log.Printf("CMD: %q", line)
	switch command {
	case "USER":
		writer.WriteString("331 password required\n")
		writer.Flush()
	case "PASS":
		writer.WriteString("230 logged in\n")
		writer.Flush()
	case "CWD":
		cd <- s[1]
	case "PORT":
		addr := strings.Split(s[1], ",")

		if len(addr) != 6 {
			log.Printf("invalid address: %s", addr)
			return
		}

		p1, err := strconv.ParseUint(addr[4], 10, 16)
		if err != nil {
			log.Print(err)
			writer.WriteString("451 local error in processing\n")
			writer.Flush()
			return
		}
		p2, err := strconv.ParseUint(addr[5], 10, 16)
		if err != nil {
			log.Print(err)
			writer.WriteString("451 local error in processing\n")
			writer.Flush()
			return
		}
		port := p1<<8 | p2
		fmt.Printf("port: %d\n", port)
		writer.WriteString(fmt.Sprintf("200 data port is now %d\n", port))
		writer.Flush()
	case "LIST":
		writer.WriteString(fmt.Sprintf("125 data connection already open\n"))
		writer.Flush()


	case "QUIT":
		writer.WriteString("221 closing connection...\n")
		// TODO: 終了をchannelを通して行うようにする
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}

	// change directory
	cd := make(chan string)
	go func() {
		for {
			dir := <-cd
			fmt.Printf("cwd: %s\n", cwd)
			err = os.Chdir(dir)
			if err != nil {
				log.Print(err)
				return
			}
			cwd, err = os.Getwd()
			if err != nil {
				log.Print(err)
				return
			}
			writer.WriteString(fmt.Sprintf("220 directory changed to %s\n", cwd))
			writer.Flush()
		}
	}()

	writer.WriteString("220 Welcome to this FTP server!\n")
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Print(err)
			}
		}
		handleCommand(conn, strings.Trim(line, "\r\n"), cd)
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
		go handleClient(conn)
	}
}
