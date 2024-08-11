package gocalc

type Map[T1, T2 any] interface {
	Map(T1) T2
}

type Finite interface {
	Len() int
}

type Point interface {
	Element
	Finite
	Map[int, Real]
}

type Vector interface {
	Point
	Additive
	Scale(Real) Vector
	Inner(Vector) Real
	Basis() []Vector
}

func Distance(p, q Vector) Real {
	if p.Len() != q.Len() {
		panic("Distance: points have different dimensions")
	}
	v := p.Add(q.AddInv()).(Vector)
	return Sqrt(v.Inner(v))
}
