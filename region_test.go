package gocalc_test

import "gocalc"

type Interval [2]float64

var _ gocalc.Region = Interval{}

func (i Interval) Contains(p gocalc.Point) bool {
	if p.Len() != 1 {
		panic("Interval.Contains: P dim mismatch")
	}
	x := p.Map(0).ToFloat()
	return i[0] <= x && x <= i[1]
}

func (i Interval) Dim() int {
	return 1
}

func (i Interval) Sup(int) gocalc.Real {
	return SimpleReal(i[1])
}

func (i Interval) Inf(int) gocalc.Real {
	return SimpleReal(i[0])
}

func (i Interval) Corner() gocalc.Vector {
	return SimpleVector{i[0]}
}

func (i Interval) Lengths() []gocalc.Real {
	return []gocalc.Real{SimpleReal(i[1] - i[0])}
}

type Rect []Interval

var _ gocalc.Region = Rect{}

func (r Rect) Contains(p gocalc.Point) bool {
	if p.Len() != 2 {
		panic("Rect.Contains: P dim mismatch")
	}
	for i := 0; i < len(r); i++ {
		x := p.Map(i).ToFloat()
		if r[i][0] > x || x > r[i][1] {
			return false
		}
	}
	return true
}

func (r Rect) Dim() int {
	return len(r)
}

func (r Rect) Sup(i int) gocalc.Real {
	return SimpleReal(r[i][1])
}

func (r Rect) Inf(i int) gocalc.Real {
	return SimpleReal(r[i][0])
}

func (r Rect) Corner() gocalc.Vector {
	v := make(SimpleVector, len(r))
	for i := 0; i < len(r); i++ {
		v[i] = r[i][0]
	}
	return v
}

func (r Rect) Lengths() []gocalc.Real {
	lengths := make([]gocalc.Real, len(r))
	for i := 0; i < len(r); i++ {
		lengths[i] = SimpleReal(r[i][1] - r[i][0])
	}
	return lengths
}
