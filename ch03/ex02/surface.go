package main

// p236を参照

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var mode = flag.Int("mode", 0, "draw mode")

func main() {
	flag.Parse()
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			f := func() {
				for _, point := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
					if math.IsInf(point, 0) {
						return
					}
				}
				fmt.Printf("<polygon points='%g,%g %g,%g, %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
			f()
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y, *mode)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64, mode int) float64 {
	switch mode {
	case 0:
		r := math.Hypot(x, y)
		return math.Sin(r) / r
	case 1:
		r := math.Hypot(x, y)
		return math.Sin(-x) * math.Pow(1.5, -r)
	case 2:
		return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
	default:
		return math.Sin(x*y/10) / 10
	}
}
