package gocalc_test

import (
	"gocalc"
	"testing"
)

type SimpleReal float64

var _ gocalc.Real = SimpleReal(0)

const epsilon = 1e-6

func (f SimpleReal) Equals(x any) bool {
	if xfloat, ok := x.(SimpleReal); ok {
		return f-xfloat > -epsilon && f-xfloat < epsilon
	}
	return false
}

func (f SimpleReal) ToFloat() float64 {
	return float64(f)
}

func (f SimpleReal) Zero() gocalc.Additive {
	return SimpleReal(0)
}

func (f SimpleReal) Add(g gocalc.Additive) gocalc.Additive {
	return SimpleReal(float64(f) + g.(SimpleReal).ToFloat())
}

func (f SimpleReal) AddInv() gocalc.Additive {
	return SimpleReal(-f)
}

func (f SimpleReal) One() gocalc.Multiplicative {
	return SimpleReal(1)
}

func (f SimpleReal) Mul(g gocalc.Multiplicative) gocalc.Multiplicative {
	return SimpleReal(float64(f) * g.(SimpleReal).ToFloat())
}

func (f SimpleReal) MulInv() gocalc.Multiplicative {
	if f.Equals(f.Zero()) {
		return nil
	}
	return SimpleReal(1 / f)
}

func (f SimpleReal) Integer(n int) gocalc.Real {
	return SimpleReal(float64(n))
}

func (f SimpleReal) GreaterThan(g gocalc.Real) bool {
	return f.ToFloat() > g.ToFloat()
}

func TestSqrt(t *testing.T) {
	a := SimpleReal(2)
	t.Log(gocalc.Sqrt(a))
}

func TestRational(t *testing.T) {
	f := SimpleReal(0)
	r := [2]int{1, 2}
	t.Log(gocalc.Rational(f, r))
}
