package main

import (
	"fmt"
	"time"
)

var counter int

func main() {

	ping := make(chan struct{})
	pong := make(chan struct{})

	go sendPing(ping, pong)
	go sendPong(ping, pong)

	ping <- struct{}{}

	for {
		time.Sleep(time.Duration(time.Second))
		fmt.Printf("counter: %d\n", counter)
	}
}

func sendPing(ping <-chan struct{}, pong chan<- struct{}) {
	for {
		<-ping
		fmt.Println("ping")
		pong <- struct{}{}
	}
}

func sendPong(ping chan<- struct{}, pong <-chan struct{}) {
	for {
		<-pong
		fmt.Println("pong")
		counter++
		ping <- struct{}{}
	}
}
