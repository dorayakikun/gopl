package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"
)

type context struct {
	cwd      string
	dataPort *net.TCPAddr
	conn     net.Conn
}

func detailed(file os.FileInfo) []byte {
	var buf bytes.Buffer

	stat := file.Sys().(*syscall.Stat_t)
	fmt.Fprint(&buf, file.Mode().String())
	fmt.Fprintf(&buf, " 1 %d %d ", stat.Uid, stat.Gid)
	fmt.Fprintf(&buf, fmt.Sprintf("%12d", file.Size()))
	fmt.Fprintf(&buf, file.ModTime().Format(" Jan _2 15:04 "))
	fmt.Fprintf(&buf, "%s\n", file.Name())
	return buf.Bytes()
}

func handleCommand(ctx *context, line string) {
	s := strings.Split(line, " ")

	if len(s) == 0 {
		log.Print("missing command")
		return
	}

	command := s[0]
	log.Printf("CMD: %q", line)
	switch command {
	case "USER":
		fmt.Fprintln(ctx.conn, "331 password required")
	case "PASS":
		fmt.Fprintln(ctx.conn, "230 logged in")
	case "CWD":
		cwd, err := os.Getwd()
		if err != nil {
			log.Print(err)
			fmt.Fprintln(ctx.conn, "451 local error in processing")
			return
		}

		os.Chdir(s[1])

		cwd, err = os.Getwd()
		if err != nil {
			log.Print(err)
			fmt.Fprintln(ctx.conn, "451 local error in processing")
			return
		}

		ctx.cwd = cwd

		fmt.Fprintf(ctx.conn,"220 directory changed to %s\n", cwd)
	case "PORT":
		addr := strings.Split(s[1], ",")

		if len(addr) != 6 {
			log.Printf("invalid address: %s", addr)
			return
		}

		p1, err := strconv.ParseUint(addr[4], 10, 16)
		if err != nil {
			fmt.Fprintln(ctx.conn, "451 local error in processing")
			return
		}
		p2, err := strconv.ParseUint(addr[5], 10, 16)
		if err != nil {
			log.Print(err)
			fmt.Fprintln(ctx.conn,"451 local error in processing")
			return
		}
		port := p1*256 + p2
		ip1, _ := strconv.ParseUint(addr[0], 10, 8)
		ip2, _ := strconv.ParseUint(addr[1], 10, 8)
		ip3, _ := strconv.ParseUint(addr[2], 10, 8)
		ip4, _ := strconv.ParseUint(addr[3], 10, 8)

		ctx.dataPort = &net.TCPAddr{IP: net.IPv4(uint8(ip1), uint8(ip2), uint8(ip3), uint8(ip4)), Port: int(port)}

		fmt.Fprintf(ctx.conn, "200 data port is now %d\n", port)
	case "LIST":
		fmt.Fprintln(ctx.conn, "150 file status ok")

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

		var dirname string
		path.Join(ctx.cwd, dirname)
		if len(s) > 1 {
			dirname = s[1]
		} else {
			dirname = "."
		}

		fi, err := os.Stat(path.Join(ctx.cwd, dirname))
		if err != nil {
			fmt.Fprintf(ctx.conn, "550 file not found: %s\n", dirname)
			return
		}
		var files []os.FileInfo
		if fi.IsDir() {
			files, err = ioutil.ReadDir(path.Join(ctx.cwd, dirname))
			if err != nil {
				// FIXME: 適切なコードとメッセージに置き換える
				fmt.Fprintln(ctx.conn, "501 invalid parameter or argument")
				return
			}
		} else {
			files = []os.FileInfo{fi}
		}
		for _, file := range files {
			c.Write(detailed(file))
		}
		fmt.Fprintln(ctx.conn, "226 closing data connection")
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
		fmt.Println("221 closing connection...")
		ctx.conn.Close()
	default:
		fmt.Fprintln(ctx.conn, "504 Command not implemented for that parameter.")
	}
}

func handleClient(conn net.Conn) {
	ctx := &context{conn: conn}
	scanner := bufio.NewScanner(ctx.conn)
	fmt.Fprintln(ctx.conn, "220 Welcome to this FTP server!")
	for scanner.Scan() {
		handleCommand(ctx, scanner.Text())
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
