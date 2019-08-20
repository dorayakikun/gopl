package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"z/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("任意の式を入力してください")

	scanner.Scan()
	s := scanner.Text()
	expr, err := eval.Parse(s)

	if err != nil {
		log.Fatal(err)
	}

	vars := map[eval.Var]bool{}
	err = expr.Check(vars)
	if err != nil {
		log.Fatal(err)
	}

	env := make(map[eval.Var]float64)
	for v := range vars {
		fmt.Printf("%s の値を入力してください\n", v)
		for scanner.Scan() {
			s = scanner.Text()
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Printf("数値を入力してください 入力された値: %s\n", s)
				continue
			}
			env[v] = f
			break
		}
	}

	fmt.Printf("計算結果: %g\n", expr.Eval(env))
}
