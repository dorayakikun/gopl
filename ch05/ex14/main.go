package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Text()
}

func run(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	var r, c int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return errors.New(fmt.Sprintf("%q want: [\"r\"  \"coloumn\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		r, c = a, b
		break
	}

	var sx, sy int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return errors.New(fmt.Sprintf("%q want: [\"sx\"  \"sy\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		sx, sy = a, b
	}

	var gx, gy int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return errors.New(fmt.Sprintf("%q want: [\"gx\"  \"gy\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		gx, gy = a, b
	}

	// TODO 座標と動けるマスを持てるstructを作成する
	maze := make([][]rune, r)
	for i := 0; i < r; i++ {
		maze[i] = make([]rune, c)

		scanner.Scan()
		cells := []rune(scanner.Text())

		for j, cell := range cells {
			if cell != '.' && cell != '#' {
				return errors.New(fmt.Sprintf("%q want '.' or '#'\n", cell))
			}
			maze[i][j] = cell
		}
	}

	return nil
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
