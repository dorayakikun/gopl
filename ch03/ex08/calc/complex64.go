package calc

import (
	"image/color"
	"math/cmplx"
)

func FC64(x complex64) complex64 {
	return x*x*x*x - 1
}

func DFC64(x complex64) complex64 {
	return 4 * x * x * x
}
func NewtonC64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z - FC64(z)/DFC64(z)

		// z^4 - 1 = 0のとりうる解は、z = ±1 または z = ±iなので
		// それぞれの解とzの差が0近似値となったものを正解値とする
		if cmplx.Abs(complex128(1-z)) < epsilon {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if cmplx.Abs(complex128(-1-z)) < epsilon {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if cmplx.Abs(complex128(1-z)) < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if cmplx.Abs(complex128(-1i-z)) < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}