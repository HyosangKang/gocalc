package gocalc

type Manifold interface {
	Locals() <-chan Parametric
}

type Tensor Map[[]int, Real]

type Form interface {
	Dim() int
	Rank() int
	Map[Vector, Tensor]
}

func Integrate(f Form, m Manifold) Real {
	return nil
}

type Parametric interface {
	Box
	Map[Point, Point]
}

func Volume(m Manifold) float64 {
	return 0
}

func Geodesic(m Manifold, p Vector) <-chan Vector {
	return nil
}
