package gocalc

type Sequence interface {
	Map[int, Real]
}

type FiniteSequence interface {
	Finite
	Sequence
}
