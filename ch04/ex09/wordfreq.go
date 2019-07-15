package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	freqwords := make(map[string]int)

	for scanner.Scan() {
		freqwords[scanner.Text()]++
	}
	fmt.Print("word\tcount\t\n")
	for w := range freqwords {
		fmt.Printf("%s\t%d\n", w, freqwords[w])
	}
}
