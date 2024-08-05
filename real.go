package gocalc

type Additive interface {
	Add(Additive) Additive
	AddInv() Additive
	Zero() Additive
}

type Multiplicative interface {
	Mul(Multiplicative) Multiplicative
	MulInv() Multiplicative
	One() Multiplicative
}

type Real interface {
	Element
	Additive
	Multiplicative
	GreaterThan(Real) bool
	ToFloat() float64
}

// Rational returns r[0]/r[1] in Real system of f.
func Rational(f Real, r [2]int) Real {
	if r[1] == 0 {
		return nil
	}
	if r[0] == 0 {
		return f.Zero().(Real)
	}
	var rr [2]Real
	for i := 0; i < 2; i++ {
		n := r[i]
		var o Real
		if r[i] > 0 {
			o = f.One().(Real)
		} else {
			o = f.One().(Real).AddInv().(Real)
			n = -n
		}
		t := f.Zero().(Real)
		for i := 0; i < n; i++ {
			t = t.Add(o).(Real)
		}
		rr[i] = t
	}
	return rr[0].Mul(rr[1].MulInv()).(Real)

}

func Sqrt(a Real) Real {
	x := a.One().(Real)
	for !a.Equals(x.Mul(x).(Real)) {
		x = x.Add(a.Mul(x.MulInv()).(Real)).(Real).Mul(Rational(x, [2]int{1, 2})).(Real)
	}
	return x
}
