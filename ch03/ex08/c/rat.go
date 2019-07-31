package c

import (
	"image/color"
	"math/big"
)

type ComplexRat struct {
	Real *big.Rat
	Imag *big.Rat
}

func (lhs *ComplexRat) Add(rhs *ComplexRat) *ComplexRat {
	return &ComplexRat{
		Real: new(big.Rat).Add(lhs.Real, rhs.Real),
		Imag: new(big.Rat).Add(lhs.Imag, rhs.Imag),
	}
}
func (lhs *ComplexRat) Sub(rhs *ComplexRat) *ComplexRat {
	return &ComplexRat{
		Real: new(big.Rat).Sub(lhs.Real, rhs.Real),
		Imag: new(big.Rat).Sub(lhs.Imag, rhs.Imag),
	}
}

func (lhs *ComplexRat) Mul(rhs *ComplexRat) *ComplexRat {
	/**
		(a + bi)(lhs + di)
	  = (ac - bd) + i(bc + ad)
	*/
	a := lhs.Real
	b := lhs.Imag
	c := rhs.Real
	d := rhs.Imag
	return &ComplexRat{
		// ac - bd
		Real: new(big.Rat).Sub(
			new(big.Rat).Mul(a, c),
			new(big.Rat).Mul(b, d)),
		// bc -ad
		Imag: new(big.Rat).Add(
			new(big.Rat).Mul(b, c),
			new(big.Rat).Mul(a, d))}
}
func (lhs *ComplexRat) Div(rhs *ComplexRat) *ComplexRat {
	/**
	    (a + bi)/(c + di)
	  = (ac + bd) / c^2 + di^2 + i(bc-ad) / c^2 + di^2
	*/

	a := lhs.Real
	b := lhs.Imag
	c := rhs.Real
	d := rhs.Imag

	// c^2 + di^2
	denominator := new(big.Rat).Add(
		new(big.Rat).Mul(c, c),
		new(big.Rat).Mul(d, d),
	)

	zero := new(big.Rat).SetFloat64(0)
	if denominator.Cmp(zero) == 0 {
		return &ComplexRat{zero, zero}
	}

	// ac + bd
	fr := new(big.Rat).Add(
		new(big.Rat).Mul(a, c),
		new(big.Rat).Mul(b, d),
	)
	// bc - ad
	fi := new(big.Rat).Sub(
		new(big.Rat).Mul(b, c),
		new(big.Rat).Mul(a, d),
	)
	return &ComplexRat{
		Real: new(big.Rat).Quo(fr, denominator),
		Imag: new(big.Rat).Quo(fi, denominator),
	}
}
func (lhs *ComplexRat) Abs() *big.Rat {
	p := lhs.Real
	q := lhs.Imag

	// Absは1としか比較しないはずなので、平方根は不要
	return new(big.Rat).Add(
		new(big.Rat).Mul(p, p),
		new(big.Rat).Mul(q, q),
	)
}

// z^4 - 1 = 0を求める関数
func fRat(x *ComplexRat) *ComplexRat {
	one := &ComplexRat{
		Real: big.NewRat(1, 1),
		Imag: new(big.Rat),
	}
	return x.Mul(x).Mul(x).Mul(x).Sub(one)
}

// 導関数
// https://ja.wolframalpha.com/input/?i=z%5E4+-1%E3%81%AE%E5%B0%8E%E9%96%A2%E6%95%B0
func dfRat(x *ComplexRat) *ComplexRat {
	four := &ComplexRat{
		Real: big.NewRat(4, 1),
		Imag: new(big.Rat),
	}
	return four.Mul(x).Mul(x).Mul(x)
}
func NewtonRat(z *ComplexRat) color.Color {
	const iterations = 5
	const contrast = 15

	epsilonRat := new(big.Rat).Mul(big.NewRat(1, 10000), big.NewRat(1, 10000))
	// 1
	one := &ComplexRat{Real: new(big.Rat).SetFloat64(1), Imag: new(big.Rat).SetFloat64(0)}
	// -1
	negaOne := &ComplexRat{Real: new(big.Rat).SetFloat64(-1), Imag: new(big.Rat).SetFloat64(0)}
	// i
	i := &ComplexRat{Real: new(big.Rat).SetFloat64(0), Imag: new(big.Rat).SetFloat64(1)}
	// -i
	negaI := &ComplexRat{Real: new(big.Rat).SetFloat64(0), Imag: new(big.Rat).SetFloat64(-1)}
	for n := uint8(0); n < iterations; n++ {
		// ニュートン法のアルゴリズムは右記を参照. https://algorithm.joho.info/mathematics/newton-method-program/
		// 漸化式 αn - f(αn) / df(αn)
		z = z.Sub(fRat(z).Div(dfRat(z)))
		// fmt.Printf("z: %s\n", z.Abs().FloatString(10))
		if one.Sub(z).Abs().Cmp(epsilonRat) < 0 {
			return color.RGBA{R: contrast * n, G: 0, B: 0, A: 255}
		} else if negaOne.Sub(z).Abs().Cmp(epsilonRat) < 0 {
			return color.RGBA{R: 0, G: 0, B: contrast * n, A: 255}
		} else if i.Sub(z).Abs().Cmp(epsilonRat) < 0 {
			return color.RGBA{R: 0, G: contrast * n, B: 0, A: 255}
		} else if negaI.Sub(z).Abs().Cmp(epsilonRat) < 0 {
			return color.RGBA{R: 0, G: contrast * n, B: contrast * n, A: 255}
		}
	}
	return color.Black
}
