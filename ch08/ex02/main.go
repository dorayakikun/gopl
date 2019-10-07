package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

type context struct {
	cwd      string
	dataPort *net.TCPAddr
	conn     net.Conn
}

func detailed(file os.FileInfo) []byte {
	var buf bytes.Buffer
	fmt.Fprint(&buf, file.Mode().String())
	fmt.Fprintf(&buf, " 1 %s %s ", "dorayakikun", "dorayakikun")
	fmt.Fprintf(&buf, fmt.Sprintf("%12d", file.Size()))
	fmt.Fprintf(&buf, file.ModTime().Format(" Jan _2 15:04 "))
	fmt.Fprintf(&buf, "%s\n", file.Name())
	return buf.Bytes()
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

		writer.Write([]byte("150 file status ok\n"))
		writer.Flush()

		var dirname string
		path.Join(ctx.cwd, dirname)
		if len(s) > 1 {
			dirname = s[1]
		} else {
			dirname = "."
		}

		fi, err := os.Stat(path.Join(ctx.cwd, dirname))
		if err != nil {
			writer.WriteString(fmt.Sprintf("550 file not found: %s\n", dirname))
			writer.Flush()
			return
		}
		var files []os.FileInfo
		if fi.IsDir() {
			files, err = ioutil.ReadDir(path.Join(ctx.cwd, dirname))
			if err != nil {
				// FIXME: 適切なコードとメッセージに置き換える
				writer.WriteString("501 invalid parameter or argument\n")
				writer.Flush()
				return
			}
		} else {
			files = []os.FileInfo{fi}
		}

		for _, file := range files {
			c.Write(detailed(file))
		}

		writer.WriteString("226 closing data connection\n")
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
