package main

import (
	"crawl/links"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
)

func breadthFirst(seed []string) {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	originaldomains := make(map[string]bool)

	for _, v := range seed {
		originaldomains[v] = true
	}

	go func() {
		worklist <- seed
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(originaldomains, link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(originaldomains map[string]bool, url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	var ret []string
	for _, l := range list {
		log.Print(l)
		u, err := neturl.Parse(l)
		if err != nil {
			log.Print(err)
			continue
		}
		// 最初に指定されたドメインと一致しないなら以降の処理をSKIP
		if _, ok := originaldomains["https://"+u.Hostname()]; !ok {
			continue
		}
		ret = append(ret, l)
		copyFile(l, u)
	}

	return ret
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
	breadthFirst(os.Args[1:])
}
