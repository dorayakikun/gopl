package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	tcp, ok := c.(*net.TCPConn)
	if !ok {
		log.Fatalln("c is not TCPConn")
	}
	input := bufio.NewScanner(tcp)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(tcp, input.Text(), 1*time.Second, &wg)
	}
	go func() {
		wg.Wait()
		println("reverb2 done")
		tcp.CloseWrite()
	}()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
