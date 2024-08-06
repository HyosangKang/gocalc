package gocalc_test

import (
	"gocalc"
	"math"
	"testing"
)

func TestDet(t *testing.T) {
	mat := [][]gocalc.Real{
		{SimpleReal(1), SimpleReal(2)},
		{SimpleReal(3), SimpleReal(4)},
	}
	t.Log(gocalc.Det(mat))
	mat = [][]gocalc.Real{
		{SimpleReal(1), SimpleReal(2), SimpleReal(3)},
		{SimpleReal(4), SimpleReal(5), SimpleReal(6)},
		{SimpleReal(7), SimpleReal(8), SimpleReal(9)},
	}
	t.Log(gocalc.Det(mat))
}

type Surface struct {
	Rect
	Local func(...float64) []float64
}

var _ gocalc.Manifold = Surface{}

func (s Surface) Locals() <-chan gocalc.Parametric {
	ch := make(chan gocalc.Parametric)
	go func() {
		defer close(ch)
		ch <- s
	}()
	return ch
}

func (s Surface) Map(p gocalc.Point) gocalc.Vector {
	if p.Len() != len(s.Rect) {
		panic("Surface.Map: P dim mismatch")
	}
	x := make([]float64, p.Len())
	for i := range x {
		x[i] = p.Map(i).ToFloat()
	}
	y := s.Local(x...)
	v := make(SimpleVector, len(y))
	copy(v, y)
	return v
}

type SimpleWedge struct {
	Eval func(...float64) float64
	W    []int
}

var _ gocalc.Wedge = SimpleWedge{}

func (w SimpleWedge) Map(p gocalc.Point) gocalc.Real {
	x := make([]float64, p.Len())
	for i := range x {
		x[i] = p.Map(i).ToFloat()
	}
	return SimpleReal(w.Eval(x...))
}

func (w SimpleWedge) Wedge() []int {
	return w.W
}

type SimpleForm struct {
	gocalc.Manifold
	Ws []SimpleWedge
}

var _ gocalc.Form = SimpleForm{}

func (s SimpleForm) Wedges() <-chan gocalc.Wedge {
	ch := make(chan gocalc.Wedge)
	go func() {
		defer close(ch)
		for _, w := range s.Ws {
			ch <- w
		}
	}()
	return ch
}

func TestIntegral(t *testing.T) {
	f := SimpleForm{
		Manifold: Surface{
			Rect: []Interval{
				{0, 2 * math.Pi},
				{0, math.Pi},
			},
			Local: func(x ...float64) []float64 {
				return []float64{
					math.Sin(x[1]) * math.Cos(x[0]),
					math.Sin(x[1]) * math.Sin(x[0]),
					math.Cos(x[1]),
				}
			},
		},
		Ws: []SimpleWedge{
			{
				Eval: func(x ...float64) float64 {
					return x[2] / (x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
				},
				W: []int{0, 1},
			},
			{
				Eval: func(x ...float64) float64 {
					return x[0] / (x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
				},
				W: []int{1, 2},
			},
			{
				Eval: func(x ...float64) float64 {
					return x[1] / (x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
				},
				W: []int{2, 0},
			},
		},
	}
	t.Log(gocalc.Integral(f, 100))
}
