package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	x        int
	y        int
	cellType rune
	steps    int
}

func main() {
	steps, err := run(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(steps)
}

func run(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	var row, column int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return 0, errors.New(fmt.Sprintf("%q want: [\"r\"  \"coloumn\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		row, column = a, b
		break
	}

	var sy, sx int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return 0, errors.New(fmt.Sprintf("%q want: [\"sy\"  \"sx\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		sy, sx = a, b
		break
	}

	var gy, gx int
	for scanner.Scan() {
		rc := strings.Split(scanner.Text(), " ")

		if len(rc) != 2 {
			return 0, errors.New(fmt.Sprintf("%q want: [\"gy\"  \"gx\"]\n", rc))
		}
		a, err := strconv.Atoi(rc[0])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[0]))
		}
		b, err := strconv.Atoi(rc[1])
		if err != nil {
			return 0, errors.New(fmt.Sprintf("%q is not number\n", rc[1]))
		}
		gy, gx = a, b
		break
	}

	maze := make([][]*cell, row)
	for i := 0; i < row; i++ { // y
		maze[i] = make([]*cell, column)

		scanner.Scan()
		cells := []rune(scanner.Text())

		for j, c := range cells { // x
			if c != '.' && c != '#' {
				return 0, errors.New(fmt.Sprintf("%q want '.' or '#'\n", c))
			}
			maze[i][j] = &cell{x: j + 1, y: i + 1, cellType: c, steps: -1}
		}
	}

	// debug
	//for _, r := range maze {
	//	for _, c := range r {
	//		fmt.Print(string(c.cellType))
	//	}
	//	fmt.Println()
	//}

	start := maze[sy-1][sx-1]
	start.steps = 0
	queue := []*cell{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		var nexts []*cell

		// up
		uy, ux := current.y+1, current.x
		if 0 < uy && uy < row {
			cell := maze[uy-1][ux-1]
			if cell.steps < 0 && cell.cellType == '.' {
				nexts = append(nexts, cell)
			}
		}
		// down
		dy, dx := current.y-1, current.x
		if 0 < dy && dy < row {
			cell := maze[dy-1][dx-1]
			if cell.steps < 0 && cell.cellType == '.' {
				nexts = append(nexts, cell)
			}
		}
		// right
		ry, rx := current.y, current.x+1
		if 0 < rx && rx < column {
			cell := maze[ry-1][rx-1]
			if cell.steps < 0 && cell.cellType == '.' {
				nexts = append(nexts, cell)
			}
		}
		// left
		ly, lx := current.y, current.x-1
		if 0 < lx && lx < column {
			cell := maze[ly-1][lx-1]
			if cell.steps < 0 && cell.cellType == '.' {
				nexts = append(nexts, cell)
			}
		}

		for _, n := range nexts {
			n.steps = current.steps + 1
			queue = append(queue, n)
			if n.x == gx && n.y == gy {
				return n.steps, nil
				continue
			}
		}
	}

	return 0, errors.New("missing maze goal\n")
}
