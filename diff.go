package gocalc

type Continuous interface {
	Delta(Point, Real) Real
}

func Partial(f func(Point) Real, p, u Vector, delta Real) Real {
	x := p.Add(u.Scale(delta)).(Vector)
	return f(x).Add(f(p).AddInv()).(Real).Mul(delta.MulInv()).(Real)
}

func Gradient(f func(Point) Real, p Vector, delta Real) Vector {
	var v Vector = p.Zero().(Vector)
	for n := range p.Basis() {
		d := Partial(f, p, n, delta)
		v = v.Add(n.Scale(d)).(Vector)
	}
	return v
}

func Differential(f Map[Point, Vector], p Vector, delta Real) func(Vector) Vector {
	return func(v Vector) Vector {
		var w Vector = p.Zero().(Vector)
		var i int
		for n := range p.Basis() {
			ff := func(q Point) Real {
				return f.Map(q).Map(i).(Real)
			}
			d := Gradient(ff, p, delta)
			w = w.Add(n.Scale(d.Inner(v))).(Vector)
			i++
		}
		return w
	}
}
