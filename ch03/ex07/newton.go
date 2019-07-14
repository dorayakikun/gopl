package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const epsilon = 1e-5

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
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
	png.Encode(os.Stdout, nimg)
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z - f(z)/df(z)

		// z^4 - 1 = 0のとりうる解は、z = ±1 または z = ±iなので
		// それぞれの解とzの差が0近似値となったものを正解値とする
		if cmplx.Abs(1-z) < epsilon {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if cmplx.Abs(-1-z) < epsilon {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if cmplx.Abs(1i-z) < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if cmplx.Abs(-1i-z) < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}

// z^4 - 1 = 0を求める関数
func f(x complex128) complex128 {
	return x*x*x*x - 1
}

// 導関数
// https://ja.wolframalpha.com/input/?i=z%5E4+-1%E3%81%AE%E5%B0%8E%E9%96%A2%E6%95%B0
func df(x complex128) complex128 {
	return 4 * x * x * x
}
