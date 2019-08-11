package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		fmt.Printf("p is %v", p)
	}()
	panic(0)
}
