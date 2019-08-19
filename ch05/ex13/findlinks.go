package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"z/links"
)

func breadthFirst(f func(domains []string, item string) []string, worklist []string) {
	var domains []string
	for _, w := range worklist {
		u, err := neturl.Parse(w)
		if err != nil {
			log.Print(err)
		}
		domains = append(domains, u.Hostname())
	}

	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(domains, item)...)
			}
		}
	}
}

func crawl(domains []string, url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, l := range list {
		fmt.Println(l)
		u, err := neturl.Parse(l)
		if err != nil {
			log.Print(err)
			continue
		}

		var samedomain bool
		for _, d := range domains {
			if d == u.Hostname() {
				samedomain = true
			}
		}

		if !samedomain {
			continue
		}

		copyFile(l, u)
	}

	return list
}

func copyFile(link string, url *neturl.URL) {
	resp, err := http.Get(link)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("fetch failed %v\n", err)
		return
	}

	if url.Path == "" {
		return
	}

	p := url.Hostname() + url.Path

	// 明示的にファイル名を指定していないなら、index.htmlを作成する
	if p[len(p)-1] == '/' {
		p += "index.html"
	} else if filepath.Ext(p) == "" {
		p += "/index.html"
	}

	if _, err = os.Stat(filepath.Dir(p)); err != nil {
		err = os.MkdirAll(filepath.Dir(p), 0777)
		if err != nil {
			log.Fatalf("create dir failed %v\n", err)
			return
		}
	}

	f, err := os.Create(p)
	defer f.Close()
	if err != nil {
		log.Fatalf("create file failed %v\n", err)
		return
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalf("copy failed %v\n", err)
		return
	}
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
