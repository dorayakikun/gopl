package main

import "fmt"

var a = make(chan int)
var b = make(chan int)
var c = make(chan int)
var d = make(chan int)
var e = make(chan int)
var f = make(chan int)
var g = make(chan int)
var h = make(chan int)

var done = make(chan struct{})

func main() {
	go func() {
		for i := 0; i < 10000; i++ {
			a <- i
		}
		done <- struct{}{}
	}()
	go func() {
		for {
			num := <-a
			fmt.Printf("a: %d\n", num)
			b <- num
		}
	}()
	go func() {
		for {
			num := <-b
			fmt.Printf("b: %d\n", num)
			c <- num
		}
	}()
	go func() {
		for {
			num := <-c
			fmt.Printf("c: %d\n", num)
			d <- num
		}
	}()
	go func() {
		for {
			num := <-d
			fmt.Printf("d: %d\n", num)
			e <- num
		}
	}()
	go func() {
		for {
			num := <-e
			fmt.Printf("e: %d\n", num)
			f <- num
		}
	}()
	go func() {
		for {
			num := <-f
			fmt.Printf("f: %d\n", num)
			g <- num
		}
	}()
	go func() {
		for {
			num := <-g
			fmt.Printf("g: %d\n", num)
			h <- num
		}
	}()
	go func() {
		for {
			num := <-h
			fmt.Printf("h: %d\n", num)
		}
	}()

	<-done
}
