package c

import (
	"image/color"
	"math/cmplx"
)

const epsilon = 1e-5

func Newton(z complex128) color.Color {
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
