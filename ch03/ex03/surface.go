package main

import (
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

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	fmt.Println("<defs><linearGradient id='gradient'>" +
		"<stop offset='0' stop-color='red' />" +
		"<stop offset='1' stop-color='blue' />" +
		"</linearGradient></defs>")

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
				fmt.Printf("<polygon points='%g,%g %g,%g, %g,%g %g,%g' fill='url(#gradient)' />\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
			f()
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
