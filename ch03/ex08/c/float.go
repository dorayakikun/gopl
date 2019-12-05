package c

import (
	"image/color"
	"math/big"
)

type ComplexBigFloat struct {
	Real *big.Float
	Imag *big.Float
}

func (lhs *ComplexBigFloat) Add(rhs *ComplexBigFloat) *ComplexBigFloat {
	return &ComplexBigFloat{
		Real: new(big.Float).Add(
			lhs.Real,
			rhs.Real,
		),
		Imag: new(big.Float).Add(
			lhs.Imag,
			rhs.Imag,
		)}
}
func (lhs *ComplexBigFloat) Sub(rhs *ComplexBigFloat) *ComplexBigFloat {
	return &ComplexBigFloat{
		Real: new(big.Float).Sub(
			lhs.Real,
			rhs.Real,
		),
		Imag: new(big.Float).Sub(
			lhs.Imag,
			rhs.Imag,
		)}
}
func (lhs *ComplexBigFloat) Mul(rhs *ComplexBigFloat) *ComplexBigFloat {
	/**
	    (a + bi) + (c + di)
	  = (ac - bd) + i(bc + ad)
	*/
	a := lhs.Real
	b := lhs.Imag
	c := rhs.Real
	d := rhs.Imag

	return &ComplexBigFloat{
		Real: new(big.Float).Sub(
			new(big.Float).Mul(a, c),
			new(big.Float).Mul(b, d),
		),
		Imag: new(big.Float).Add(
			new(big.Float).Mul(b, c),
			new(big.Float).Mul(a, d),
		)}
}
func (lhs *ComplexBigFloat) Div(rhs *ComplexBigFloat) *ComplexBigFloat {
	/**
	    (a + bi)/(c + di)
	  = (ac + bd) / c^2 + di^2 + i(bc-ad) / c^2 + di^2
	*/
	a := lhs.Real
	b := lhs.Imag
	c := rhs.Real
	d := rhs.Imag

	denominator := new(big.Float).Add(
		new(big.Float).Mul(c, c),
		new(big.Float).Mul(d, d),
	)
	fr := new(big.Float).Add(
		new(big.Float).Mul(a, c),
		new(big.Float).Mul(b, d),
	)
	fi := new(big.Float).Sub(
		new(big.Float).Mul(b, c),
		new(big.Float).Mul(a, d),
	)
	// 0または有限数でない場合は何もしない
	if denominator.IsInf() || denominator.Cmp(new(big.Float)) == 0 {
		return lhs
	}
	return &ComplexBigFloat{
		Real: new(big.Float).Quo(fr, denominator),
		Imag: new(big.Float).Quo(fi, denominator),
	}
}
func (lhs ComplexBigFloat) Abs() *big.Float {
	zero := new(big.Float)

	if lhs.Real.IsInf() || lhs.Imag.IsInf() {
		return new(big.Float).SetInf(true)
	}

	p, q := new(big.Float).Abs(lhs.Real), new(big.Float).Abs(lhs.Imag)
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
