package gocalc

type Map[T1, T2 any] interface {
	Map(T1) T2
}

type RealValued interface {
	Region
	Map[Point, Real]
}

type VectorValued interface {
	Region
	Map[Point, Vector]
}
