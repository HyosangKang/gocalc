package gocalc_test

import (
	"gocalc"
	"testing"
)

func TestGradient(t *testing.T) {
	f := Double{
		Rect: []Interval{
			{0, 1},
			{0, 1},
		},
		Eval: func(x, y float64) float64 {
			return x*x + y*y
		},
	}
	p := SimplePoint{
		SimpleReal(0.5),
		SimpleReal(0.5),
	}
	t.Log(gocalc.Gradient(f.Map, p, SimpleReal(1e-6)))
}

type Multiple struct {
	Rect
	Eval func(...float64) []float64
}

var _ gocalc.VectorValued = Multiple{}

func (m Multiple) Map(p gocalc.Point) gocalc.Vector {
	if p.Len() != len(m.Rect) {
		panic("Multiple.Map: P dim mismatch")
	}
	x := make([]float64, p.Len())
	for i := range x {
		x[i] = p.Map(i).ToFloat()
	}
	y := m.Eval(x...)
	v := make(SimplePoint, len(y))
	for i := range y {
		v[i] = SimpleReal(y[i])
	}
	return v
}

func TestDifferential(t *testing.T) {
	f := Multiple{
		Rect: []Interval{
			{0, 1},
			{0, 1},
		},
		Eval: func(x ...float64) []float64 {
			return []float64{x[0]*x[0] + x[1]*x[1], x[0]*x[0] - x[1]*x[1]}
		},
	}
	p := SimplePoint{
		SimpleReal(0.5),
		SimpleReal(0.5),
	}
	diff := gocalc.Differential(f, p, SimpleReal(1e-6))
	v := SimplePoint{
		SimpleReal(1),
		SimpleReal(1),
	}
	t.Log(diff(v))
}
