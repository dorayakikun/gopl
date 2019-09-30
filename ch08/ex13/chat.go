package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	name     string
	receiver chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.receiver <- msg
			}

		case cli := <-entering:
			if len(clients) != 0 {
				var s strings.Builder
				s.WriteString("all members:")
				for c := range clients {
					s.WriteString("\n")
					s.WriteString(c.name)
				}
				cli.receiver <- s.String()
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.receiver)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	keep := make(chan struct{})
	go clientWriter(conn, ch)
	go keepConnection(conn, keep)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	cl := client{who, ch}
	entering <- cl

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		keep <- struct{}{}
	}

	leaving <- cl
	messages <- who + " has left"
	one := []byte{}

	// 2重でCloseすることを防止する
	_, err := conn.Read(one)
	if err != nil {
		return
	}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func keepConnection(conn net.Conn, keep <-chan struct{}) {
	for {
		select {
		case <-keep:
			continue
		case <-time.After(5 * time.Second):
			conn.Close()
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}