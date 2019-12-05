package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 100, "number of goroutines")
	flag.Parse()
	var r, w chan struct{}

	start := make(chan struct{})
	w = start

	fmt.Println("start prepare")
	for i := 0; i < n; i++ {
		r, w = w, make(chan struct{})
		go func(r, w chan struct{}) {
			<-r
			w <- struct{}{}
		}(r, w)
	}
	fmt.Println("end prepare")
	t := time.Now()
	start <- struct{}{}
	<-w
	fmt.Println(n, time.Now().Sub(t))
}
