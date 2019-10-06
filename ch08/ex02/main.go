package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type context struct {
	cwd      string
	dataPort *net.TCPAddr
	conn     net.Conn
}

func handleCommand(line string, ctx *context) {
	writer := bufio.NewWriter(ctx.conn)

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
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			return
		}

		os.Chdir(s[1])

		cwd, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx.cwd = cwd

		writer.WriteString(fmt.Sprintf("220 directory changed to %s\n", cwd))
		writer.Flush()
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

		ip1, _ := strconv.ParseUint(addr[0], 10, 8)
		ip2, _ := strconv.ParseUint(addr[0], 10, 8)
		ip3, _ := strconv.ParseUint(addr[0], 10, 8)
		ip4, _ := strconv.ParseUint(addr[0], 10, 8)

		ctx.dataPort = &net.TCPAddr{IP: net.IPv4(uint8(ip1), uint8(ip2), uint8(ip3), uint8(ip4)), Port: int(port)}

		writer.WriteString(fmt.Sprintf("200 data port is now %d\n", port))
		writer.Flush()
	case "LIST":
		log.Printf("%+v\n", ctx)

		src := *ctx.conn.LocalAddr().(*net.TCPAddr)
		src.Port = src.Port - 1

		var dest net.TCPAddr
		if ctx.dataPort != nil {
			dest = *ctx.dataPort
		} else {
			dest = *ctx.conn.RemoteAddr().(*net.TCPAddr)
		}

		c, err := net.DialTCP("tcp", &src, &dest)
		defer c.Close()
		if err != nil {
			log.Fatal(err)
		}

		c.Write([]byte("125 data connection already open\n"))

		var dirname string
		if len(s) > 1 {
			dirname = s[1]
		} else {
			dirname = "."
		}

		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			log.Print(err)
			c.Write([]byte("501 invalid parameter or argument\n"))
			return
		}

		for file := range files {
			c.Write([]byte(fmt.Sprintf("%s\n", file)))
		}

		c.Write([]byte(fmt.Sprintf("226 closing data connection\n")))
		writer.Flush()

	case "RETR":
		log.Printf("%+v\n", ctx)

		src := *ctx.conn.LocalAddr().(*net.TCPAddr)
		src.Port = src.Port - 1

		var dest net.TCPAddr
		if ctx.dataPort != nil {
			dest = *ctx.dataPort
		} else {
			dest = *ctx.conn.RemoteAddr().(*net.TCPAddr)
		}

		c, err := net.DialTCP("tcp", &src, &dest)
		defer c.Close()
		if err != nil {
			log.Fatal(err)
		}

	case "QUIT":
		writer.WriteString("221 closing connection...\n")
		ctx.conn.Close()
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	ctx := &context{}
	ctx.conn = conn

	writer.WriteString("220 Welcome to this FTP server!\n")
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Print(err)
			}
		}
		handleCommand(strings.Trim(line, "\r\n"), ctx)
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
