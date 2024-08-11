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

type Element interface {
	Equals(any) bool
}

type Real interface {
	Element
	Additive
	Multiplicative
	GreaterThan(Real) bool
	ToFloat() float64
}

func Integer(f Real, n int) Real {
	if n == 0 {
		return f.Zero().(Real)
	}
	o := f.One().(Real)
	if n < 0 {
		o = o.AddInv().(Real)
		n = -n
	}
	z := f.Zero().(Real)
	for i := 0; i < n; i++ {
		z = z.Add(o).(Real)
	}
	return z
}

func Rational(f Real, r [2]int) Real {
	if r[1] == 0 {
		return nil
	}
	if r[0] == 0 {
		return f.Zero().(Real)
	}
	a, b := Integer(f, r[0]), Integer(f, r[1])
	return a.Mul(b.MulInv()).(Real)
}

func Sqrt(a Real) Real {
	x := a.One().(Real)
	half := Rational(x, [2]int{1, 2})
	for !a.Equals(x.Mul(x).(Real)) {
		x = x.Add(a.Mul(x.MulInv()).(Real)).(Real)
		x = x.Mul(half).(Real)
	}
	return x
}
