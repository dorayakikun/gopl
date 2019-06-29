package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) == 0 {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames != nil {
			filenames[input.Text()] = append(filenames[input.Text()], f.Name())
		}
	}
}