package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
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
				img.Set(px, py, newton(z))
			case 1:
				z := complex64(complex(x, y))
				img.Set(px, py, newtonC64(z))
			case 2:
				z := &ComplexBigFloat{Real: big.NewFloat(x), Imag: big.NewFloat(y)}
				img.Set(px, py, newtonbigfloat(z))
			case 3:
				z := &ComplexRat{
					Real: big.NewRat(int64(x), 1),
					Imag: big.NewRat(int64(y), 1),
				}
				img.Set(px, py, newtonRat(z))
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

// ===
// complex64
// ===
func fc64(x complex64) complex64 {
	return x*x*x*x - 1
}

func dfc64(x complex64) complex64 {
	return 4 * x * x * x
}
func newtonC64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z - fc64(z)/dfc64(z)

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

// ===
// ComplexBigFloat
// ===
type ComplexBigFloat struct {
	Real *big.Float
	Imag *big.Float
}

func (c *ComplexBigFloat) Add(other *ComplexBigFloat) *ComplexBigFloat {
	return &ComplexBigFloat{
		Real: new(big.Float).Add(
			c.Real,
			other.Real,
		),
		Imag: new(big.Float).Add(
			c.Imag,
			other.Imag,
		)}
}
func (c *ComplexBigFloat) Sub(other *ComplexBigFloat) *ComplexBigFloat {
	return &ComplexBigFloat{
		Real: new(big.Float).Sub(
			c.Real,
			other.Real,
		),
		Imag: new(big.Float).Sub(
			c.Imag,
			other.Imag,
		)}
}
func (c *ComplexBigFloat) Mul(other *ComplexBigFloat) *ComplexBigFloat {
	// (1+2i)(3+4i) = 1*3 + 1*4i + 2*3i +2*4i^2
	//              = 3 + 4i + 6i -8
	//              = 3-8+4i+6i
	return &ComplexBigFloat{
		Real: new(big.Float).Sub(
			new(big.Float).Mul(
				c.Real,
				other.Real,
			),
			new(big.Float).Mul(
				c.Imag,
				other.Imag,
			)),
		Imag: new(big.Float).Add(
			new(big.Float).Mul(
				c.Real,
				other.Imag,
			),
			new(big.Float).Mul(
				c.Imag,
				other.Real,
			))}
}
func (c *ComplexBigFloat) Div(other *ComplexBigFloat) *ComplexBigFloat {
	// x1 + y1i / x2 + y2i =(x1x2+y1y2)/x2^2+y^2 + (x2y1-x1y2)/x2^2+y^2i
	// real1*real2+imag1*imag2/denominator + real2*imag1-real1*imag2/denominator
	// 分母 = x2^2+y^2
	denominator := new(big.Float).Add(
		new(big.Float).Mul(other.Real, other.Real),
		new(big.Float).Mul(other.Imag, other.Imag),
	)
	fr := new(big.Float).Add(
		new(big.Float).Mul(c.Real, other.Real),
		new(big.Float).Mul(c.Imag, other.Imag),
	)
	fi := new(big.Float).Sub(
		new(big.Float).Mul(other.Real, c.Imag),
		new(big.Float).Mul(c.Real, other.Imag),
	)
	// 0または有限数でない場合は何もしない
	if (denominator.IsInf() || denominator == big.NewFloat(0)) ||
		(fr.IsInf() || fr == big.NewFloat(0)) ||
		(fi.IsInf() || fi == big.NewFloat(0)) {
		return c
	}
	return &ComplexBigFloat{
		Real: new(big.Float).Quo(
			fr,
			denominator,
		),
		Imag: new(big.Float).Quo(
			fi,
			denominator,
		),
	}
}
func (c ComplexBigFloat) Abs() float64 {
	// sqrt(real^2 + imag^2)
	real, _ := c.Real.Float64()
	imag, _ := c.Imag.Float64()
	return math.Hypot(real, imag)
}
func fbigfloat(x *ComplexBigFloat) *ComplexBigFloat {
	// 元計算では実部だけ引いているので、実部1虚部0として再現
	one := &ComplexBigFloat{
		Real: big.NewFloat(1),
		Imag: big.NewFloat(0),
	}
	return x.Mul(x).Mul(x).Mul(x).Sub(one)
}
func dfbigfloat(x *ComplexBigFloat) *ComplexBigFloat {
	four := &ComplexBigFloat{
		Real: big.NewFloat(4),
		Imag: big.NewFloat(0),
	}
	z := four.Mul(x).Mul(x).Mul(x)
	return z
}
func newtonbigfloat(z *ComplexBigFloat) color.Color {
	const iterations = 30
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z.Sub(fbigfloat(z).Div(dfbigfloat(z)))

		// z^4 - 1 = 0のとりうる解は、z = ±1 または z = ±iなので
		// それぞれの解とzの差が0近似値となったものを正解値とする

		// 1
		one := &ComplexBigFloat{Real: big.NewFloat(1), Imag: big.NewFloat(0)}
		// -1
		none := &ComplexBigFloat{Real: big.NewFloat(-1), Imag: big.NewFloat(0)}
		// i
		i := &ComplexBigFloat{Real: big.NewFloat(0), Imag: big.NewFloat(1)}
		// i
		ni := &ComplexBigFloat{Real: big.NewFloat(-0), Imag: big.NewFloat(-1)}
		if one.Sub(z).Abs() < epsilon {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if none.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if i.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if ni.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}

// ===
// ComplexRat
// ===
type ComplexRat struct {
	Real *big.Rat
	Imag *big.Rat
}

func (c *ComplexRat) Add(other *ComplexRat) *ComplexRat {
	return &ComplexRat{
		Real: new(big.Rat).Add(c.Real, other.Real),
		Imag: new(big.Rat).Add(c.Imag, other.Imag),
	}
}
func (c *ComplexRat) Sub(other *ComplexRat) *ComplexRat {
	return &ComplexRat{
		Real: new(big.Rat).Sub(c.Real, other.Real),
		Imag: new(big.Rat).Sub(c.Imag, other.Imag),
	}
}

func (c *ComplexRat) Mul(other *ComplexRat) *ComplexRat {
	// (1+2i)(3+4i) = 1*3 + 1*4i + 2*3i +2*4i^2
	//              = 3 + 4i + 6i -8
	//              = 3-8+4i+6i
	return &ComplexRat{
		Real: new(big.Rat).Sub(
			new(big.Rat).Mul(
				c.Real,
				other.Real,
			),
			new(big.Rat).Mul(
				c.Imag,
				other.Imag,
			)),
		Imag: new(big.Rat).Add(
			new(big.Rat).Mul(
				c.Real,
				other.Imag,
			),
			new(big.Rat).Mul(
				c.Imag,
				other.Real,
			))}
}
func (c *ComplexRat) Div(other *ComplexRat) *ComplexRat {
	// x1 + y1i / x2 + y2i =(x1x2+y1y2)/x2^2+y^2 + (x2y1-x1y2)/x2^2+y^2i
	// real1*real2+imag1*imag2/denominator + real2*imag1-real1*imag2/denominator
	// 分母 = x2^2+y^2
	denominator := new(big.Rat).Add(
		new(big.Rat).Mul(other.Real, other.Real),
		new(big.Rat).Mul(other.Imag, other.Imag),
	)

	v, _ := denominator.Float64()
	if v == 0 {
		return c
	}

	fr := new(big.Rat).Add(
		new(big.Rat).Mul(c.Real, other.Real),
		new(big.Rat).Mul(c.Imag, other.Imag),
	)
	fi := new(big.Rat).Sub(
		new(big.Rat).Mul(other.Real, c.Imag),
		new(big.Rat).Mul(c.Real, other.Imag),
	)
	return &ComplexRat{
		Real: new(big.Rat).Quo(
			fr,
			denominator,
		),
		Imag: new(big.Rat).Quo(
			fi,
			denominator,
		),
	}
}
func (c *ComplexRat) Abs() float64 {
	real, _ := c.Real.Float64()
	imag, _ := c.Imag.Float64()
	return math.Hypot(real, imag)
}

// z^4 - 1 = 0を求める関数
func fRat(x *ComplexRat) *ComplexRat {
	one := &ComplexRat{
		Real: big.NewRat(1, 1),
		Imag: big.NewRat(0, 1),
	}
	return x.Mul(x).Mul(x).Mul(x).Sub(one)
}

// 導関数
// https://ja.wolframalpha.com/input/?i=z%5E4+-1%E3%81%AE%E5%B0%8E%E9%96%A2%E6%95%B0
func dfRat(x *ComplexRat) *ComplexRat {
	four := &ComplexRat{
		Real: big.NewRat(4, 1),
		Imag: big.NewRat(0, 1),
	}
	return four.Mul(x).Mul(x).Mul(x)
}
func newtonRat(z *ComplexRat) color.Color {
	const iterations = 5
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z.Sub(fRat(z).Div(dfRat(z)))

		// 1
		one := &ComplexRat{Real: big.NewRat(1, 1), Imag: big.NewRat(0, 1)}
		// -1
		none := &ComplexRat{Real: big.NewRat(-1, 1), Imag: big.NewRat(0, 1)}
		// i
		i := &ComplexRat{Real: big.NewRat(0, 1), Imag: big.NewRat(1, 1)}
		// -i
		ni := &ComplexRat{Real: big.NewRat(0, 1), Imag: big.NewRat(-1, 1)}

		if one.Sub(z).Abs() < epsilon {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if none.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if i.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if ni.Sub(z).Abs() < epsilon {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}
