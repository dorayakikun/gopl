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

	// sin(r) / r の最小と最大はそれぞれ -0.22, 1.0の近似値なのでmin, maxは計算せずに定数を置く
	// https://www.wolframalpha.com/input/?i=min+max+sin((x%5E2%2By%5E2)%5E(1%2F2))+%2F+(x%5E2%2By%5E2)%5E(1%2F2),+0+%3C%3D+x+%3C%3D+30,+0+%3C%3D+y+%3C%3D+30
	minZ = -0.22
	maxZ = 1.0
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
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			f := func() {
				for _, point := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
					if math.IsInf(point, 0) {
						return
					}
				}
				z := (az + bz + cz + dz) / 4
				// zを0〜1に正規化
				nz := (z-minZ)/maxZ - minZ
				red := int(nz * 255)
				blue := 255 - red
				fmt.Printf("<polygon points='%g,%g %g,%g, %g,%g %g,%g' fill='rgb(%d, 0, %d)' />\n", ax, ay, bx, by, cx, cy, dx, dy, red, blue)
			}
			f()
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
