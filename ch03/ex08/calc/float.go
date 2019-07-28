package calc

import (
	"image/color"
	"math/big"
)

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
	if denominator.IsInf() || denominator.Cmp(new(big.Float)) == 0 {
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
func (c ComplexBigFloat) Abs() *big.Float {
	zero := new(big.Float)

	if c.Real.IsInf() || c.Imag.IsInf() {
		return new(big.Float).SetInf(true)
	}

	p, q := new(big.Float).Abs(c.Real), new(big.Float).Abs(c.Imag)
	if p.Cmp(q) < 0 {
		p, q = q, p
	}

	if p.Cmp(zero) == 0 {
		return zero
	}

	q = new(big.Float).Quo(q, p)
	return new(big.Float).
		Mul(p, new(big.Float).
			Sqrt(
				new(big.Float).Add(
					big.NewFloat(1),
					new(big.Float).Mul(q, q),
				),
			),
		)
}
func fBigfloat(x *ComplexBigFloat) *ComplexBigFloat {
	// 元計算では実部だけ引いているので、実部1虚部0として再現
	one := &ComplexBigFloat{
		Real: big.NewFloat(1),
		Imag: big.NewFloat(0),
	}
	return x.Mul(x).Mul(x).Mul(x).Sub(one)
}
func dfBigfloat(x *ComplexBigFloat) *ComplexBigFloat {
	four := &ComplexBigFloat{
		Real: big.NewFloat(4),
		Imag: big.NewFloat(0),
	}
	z := four.Mul(x).Mul(x).Mul(x)
	return z
}
func NewtonBigFloat(z *ComplexBigFloat) color.Color {
	const iterations = 30
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z.Sub(fBigfloat(z).Div(dfBigfloat(z)))

		// z^4 - 1 = 0のとりうる解は、z = ±1 または z = ±iなので
		// それぞれの解とzの差が0近似値となったものを正解値とする

		// 1
		one := &ComplexBigFloat{Real: big.NewFloat(1), Imag: big.NewFloat(0)}
		// -1
		none := &ComplexBigFloat{Real: big.NewFloat(-1), Imag: big.NewFloat(0)}
		// i
		i := &ComplexBigFloat{Real: big.NewFloat(0), Imag: big.NewFloat(1)}
		// -i
		ni := &ComplexBigFloat{Real: big.NewFloat(0), Imag: big.NewFloat(-1)}
		if one.Sub(z).Abs().Cmp(big.NewFloat(epsilon)) < 0 {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if none.Sub(z).Abs().Cmp(big.NewFloat(epsilon)) < 0 {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if i.Sub(z).Abs().Cmp(big.NewFloat(epsilon)) < 0 {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if ni.Sub(z).Abs().Cmp(big.NewFloat(epsilon)) < 0 {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}
