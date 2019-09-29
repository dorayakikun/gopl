package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetch(url string) (filename string, n int64, err error) {
	responses := make(chan *http.Response)
	errors := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func(url string) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			errors <- err
			return
		}
		req.WithContext(ctx)
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			errors <- err
			return
		}
		responses <- resp
	}(url)
	go func(url string) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			errors <- err
			return
		}
		req.WithContext(ctx)
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			errors <- err
			return
		}
		responses <- resp
	}(url)
	go func(url string) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			errors <- err
			return
		}
		req.WithContext(ctx)
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			errors <- err
			return
		}
		responses <- resp
	}(url)

	var resp *http.Response
	select {
	case resp = <-responses:
		cancel()
	case err := <-errors:
		fmt.Printf("request failed: %s", err)
		cancel()
	}
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
