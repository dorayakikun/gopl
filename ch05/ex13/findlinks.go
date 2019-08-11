package main

import (
	"fmt"
	"io"
	"log"
	"os"
	neturl "net/url"
	"net/http"
	"path/filepath"

	"z/links"
)

func breadthFirst(f func(domain string, item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			u, err := neturl.Parse(item)
			if err != nil {
				log.Print(err)
			}
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(u.Hostname(), item)...)
			}
		}
	}
}

func crawl(domain string, url string) []string {
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
		if u.Hostname() != domain {
			continue
		}

		resp, err := http.Get(l)
		defer resp.Body.Close()
		if err != nil {
			log.Fatalf("fetch failed %v\n", err)
			continue
		}

		if u.Path == "" {
			continue
		}

		p := domain + u.Path

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
				continue
			}
		}

		f, err := os.Create(p)
		defer f.Close()
		if err != nil {
			log.Fatalf("copy failed %v\n", err)
			continue
		}

		io.Copy(f, resp.Body)
	}

	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
