package calc

import (
	"image/color"
	"math"
	"math/big"
)

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
	zero := big.NewRat(0, 10)

	p, q := c.Real, c.Imag
	p, q = p.Abs(p), q.Abs(q)
	if p.Cmp(q) < 0 {
		p, q = q, p
	}
	if p.Cmp(zero) == 0 {
		return 0
	}
	// TODO math.Hypot(p, q)のシュミレートをどうするか

	fp, _ := p.Float64()
	fq, _ := q.Float64()
	return math.Hypot(fp, fq)
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
func NewtonRat(z *ComplexRat) color.Color {
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