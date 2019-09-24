package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type work struct {
	list  []string
	depth int
}

type unseen struct {
	link  string
	depth int
}

var depth = flag.Int("depth", int(^uint(0)>>1), "depth")

//!+
func main() {
	flag.Parse()

	worklist := make(chan work)
	unseenLinks := make(chan unseen)

	go func() { worklist <- work{flag.Args(), 0} }()

	for i := 0; i < 20; i++ {
		go func() {
			for u := range unseenLinks {
				if u.depth > *depth {
					continue
				}
				foundLinks := crawl(u.link)
				go func() { worklist <- work{foundLinks, u.depth + 1} }()
			}
		}()
	}

	seen := make(map[string]bool)
	for w := range worklist {
		for _, link := range w.list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- unseen{link, w.depth}
			}
		}
	}
}

//!-
