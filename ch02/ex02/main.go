package main

import (
	"bufio"
	"flag"
	"fmt"
	tempconv "github.com/dorayakikun/go-study/ch02/ex01"
	"github.com/dorayakikun/go-study/ch02/ex02/lengthconv"
	"github.com/dorayakikun/go-study/ch02/ex02/weightconv"
	"os"
	"strconv"
)

var mode = flag.String("mode", "temperature", "Type of unit to convert")

func main() {
	flag.Parse()
	fmt.Println(*mode)
	if len(flag.Args()) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printWithUnit(input.Text(), *mode)
		}
	}
	for _, arg := range flag.Args() {
		printWithUnit(arg, *mode)
	}
}

func printWithUnit(s string, mode string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	switch mode {
	case "temperature":
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CToF(c))
	case "weight":
		p := weightconv.Pound(v)
		k := weightconv.Kilogram(v)
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
	case "length":
		f := lengthconv.Feet(v)
		m := lengthconv.Meters(v)
		fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	}
}
