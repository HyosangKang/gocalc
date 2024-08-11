package gocalc

type T interface {
	Rank() int
	Map[Vector, T]
}

type t struct {
	rank int
	tsr  map[int]func(r) r
}

func (t t) Rank() int {
	return t.rank
}

func (t t) Map(v Vector) T {
	return nil
}

type r float64

var _ T = r(0)

func (r r) Rank() int {
	return 0
}

func (r r) Map(v Vector) T {
	return r
}
