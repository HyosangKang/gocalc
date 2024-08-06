package gocalc_test

import (
	"gocalc"
	"math"
)

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

type Single struct {
	Interval
	Eval func(float64) float64
	Prec func(float64, float64) float64
}

var (
	s                = Single{}
	_ gocalc.Grapher = s
	_ gocalc.Cont    = s
)

func (s Single) Eps(d gocalc.Real, x gocalc.Point) gocalc.Real {
	return SimpleReal(s.Prec(d.ToFloat(), x.Map(0).ToFloat()))
}

func (s Single) Map(p gocalc.Point) gocalc.Real {
	if p.Len() != 1 {
		panic("SimpleSigle.Map: P dim mismatch")
	}
	x := p.Map(0).ToFloat()
	return SimpleReal(s.Eval(x))
}

func (s Single) Lift(p gocalc.Point) gocalc.Point {
	x := p.Map(0).ToFloat()
	return &SimpleVector{
		x,
		s.Eval(x),
	}
}

func (s Single) Project(p gocalc.Point) gocalc.Point {
	return p
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

type Double struct {
	Rect
	Eval func(float64, float64) float64
}

var (
	d                = Double{}
	_ gocalc.Grapher = d
)

func (d Double) Map(p gocalc.Point) gocalc.Real {
	if p.Len() != 2 {
		panic("Double.Map: P dimemsion mismatch")
	}
	x, y := p.Map(0).ToFloat(), p.Map(1).ToFloat()
	return SimpleReal(d.Eval(x, y))
}

func (d Double) Lift(p gocalc.Point) gocalc.Point {
	x, y := p.Map(0).ToFloat(), p.Map(1).ToFloat()
	return SimpleVector{x, y, d.Eval(x, y)}
}

func (d Double) Project(p gocalc.Point) gocalc.Point {
	t1, t2 := -1.0, -1.0
	x, y, z := p.Map(0).ToFloat(), p.Map(1).ToFloat(), p.Map(2).ToFloat()
	x, y = x*math.Cos(t1)-y*math.Sin(t1), x*math.Sin(t1)+y*math.Cos(t1)
	y, z = y*math.Cos(t2)-z*math.Sin(t2), y*math.Sin(t2)+z*math.Cos(t2)
	return SimpleVector{x, y}
}
