package gocalc

type Map[T1, T2 any] interface {
	Map(T1) T2
}

type Sequence interface {
	Map[int, Real]
}

type Point interface {
	Element
	Finite
	Sequence
}

type Vector interface {
	Point
	Additive
	Scale(Real) Vector
	Inner(Vector) Real
	Basis() []Vector
}

type PV = Vector

func Distance(p, q Point) Real {
	if p.Len() != q.Len() {
		panic("Distance: points have different dimensions")
	}
	var v Real = p.Map(0)
	for i := 0; i < p.Len(); i++ {
		v = v.Add(p.Map(i).Mul(q.Map(i)).(Real)).(Real)
	}
	return Sqrt(v)
}
