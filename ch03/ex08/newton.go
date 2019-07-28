package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
	"z/calc"
)

const epsilon = 1e-5

var mode = flag.Int("mode", 0, "0: complex128, 1: complex64, 2: big.Float, 3: big.Rat")

func main() {
	flag.Parse()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			switch *mode {
			case 0:
				z := complex(x, y)
				img.Set(px, py, calc.Newton(z))
			case 1:
				z := complex64(complex(x, y))
				img.Set(px, py, calc.NewtonC64(z))
			case 2:
				z := &calc.ComplexBigFloat{Real: big.NewFloat(x), Imag: big.NewFloat(y)}
				img.Set(px, py, calc.NewtonBigFloat(z))
			case 3:
				z := &calc.ComplexRat{
					Real: big.NewRat(int64(x), 1),
					Imag: big.NewRat(int64(y), 1),
				}
				img.Set(px, py, calc.NewtonRat(z))
			}
		}
	}
	// new image
	nimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r int
			var g int
			var b int
			var a int

			if px == 1023 {
				// 終端に到達しているなら、色の平均を求めない
				if py == 1023 {
					nimg.Set(px, py, img.RGBAAt(px, py))
					continue
				}
				// x終端に到達している場合は、直下のピクセルと比べて色の平均を取得する
				r += int(img.RGBAAt(px, py).R)
				r += int(img.RGBAAt(px, py+1).R)
				r /= 2

				g += int(img.RGBAAt(px, py).G)
				g += int(img.RGBAAt(px, py+1).G)
				g /= 2

				b += int(img.RGBAAt(px, py).B)
				b += int(img.RGBAAt(px, py+1).B)
				b /= 2

				a += int(img.RGBAAt(px, py).A)
				a += int(img.RGBAAt(px, py+1).A)
				a /= 2

				nimg.Set(px, py, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
				continue
			}

			// x終端に到達している場合は、右のピクセルと比べて色の平均を取得する
			if py == 1023 {
				r += int(img.RGBAAt(px, py).R)
				r += int(img.RGBAAt(px+1, py).R)
				r /= 2

				g += int(img.RGBAAt(px, py).G)
				g += int(img.RGBAAt(px+1, py).G)
				g /= 2

				b += int(img.RGBAAt(px, py).B)
				b += int(img.RGBAAt(px+1, py).B)
				b /= 2

				a += int(img.RGBAAt(px, py).A)
				a += int(img.RGBAAt(px+1, py).A)
				a /= 2

				nimg.Set(px, py, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
				continue
			}

			// それ以外は、右、右下、直下の3pxと自分自身の4pxで色の平均を割り出す
			r += int(img.RGBAAt(px, py).R)
			r += int(img.RGBAAt(px+1, py).R)
			r += int(img.RGBAAt(px, py+1).R)
			r += int(img.RGBAAt(px+1, py+1).R)
			r /= 4

			g += int(img.RGBAAt(px, py).G)
			g += int(img.RGBAAt(px+1, py).G)
			g += int(img.RGBAAt(px, py+1).G)
			g += int(img.RGBAAt(px+1, py+1).G)
			g /= 4

			b += int(img.RGBAAt(px, py).B)
			b += int(img.RGBAAt(px+1, py).B)
			b += int(img.RGBAAt(px, py+1).B)
			b += int(img.RGBAAt(px+1, py+1).B)
			b /= 4

			a += int(img.RGBAAt(px, py).A)
			a += int(img.RGBAAt(px+1, py).A)
			a += int(img.RGBAAt(px, py+1).A)
			a += int(img.RGBAAt(px+1, py+1).A)
			a /= 4

			nimg.Set(px, py, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
		}
	}
	err := png.Encode(os.Stdout, nimg)
	if err != nil {
		fmt.Printf("failed encoding: %s", err.Error())
	}
}
