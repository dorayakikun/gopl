package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for _, url := range os.Args[1:] {
		select {
		case b := <-ch:
			fmt.Println(b)
		case <-time.After(1000 * time.Millisecond):
			fmt.Printf("timed out %s\n", url)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	out, err := os.Create(strconv.Itoa(start.Nanosecond()) + ".html")
	if err != nil {
		ch <- fmt.Sprintf("failed create file %v", err)
	}
	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
