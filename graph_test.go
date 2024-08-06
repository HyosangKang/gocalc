package gocalc_test

import (
	"gocalc"
	"math"
	"testing"
)

func TestToInt(t *testing.T) {
	idx := []int{0, 1, 2}
	max := 2
	n := gocalc.ToInt(idx, max)
	jdx := gocalc.ToIdx(n, max, len(idx))
	t.Log(n)
	t.Log(jdx)
}

type Single struct {
	Interval
	Eval func(float64) float64
	Prec func(float64, float64) float64
}

var (
	s                = Single{}
	_ gocalc.Grapher = s
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

func TestGraph1(t *testing.T) {
	s := Single{
		Interval: Interval{-2 * math.Pi, 2 * math.Pi},
		Eval: func(x float64) float64 {
			return math.Sin(x)
		},
		Prec: func(d, c float64) float64 {
			return d
		},
	}

	gocalc.Graph(s, gocalc.GraphOption{
		Nsub: 100,
		Xmin: -7, Xmax: 7,
		Ymin: -1, Ymax: 1,
		Width: 600, Height: 600,
		Filename: "test_graph_1.png",
	})
}

func TestGraph2(t *testing.T) {
	single := Single{
		Interval: Interval{-2, 2},
		Eval: func(x float64) float64 {
			if -1e-6 < x && x < 1e-6 {
				return 0
			}
			return 1 / x
		},
		Prec: func(d, c float64) float64 {
			if -1e-6 < c && c < 1e-6 {
				return d * 1e12
			}
			return d / (math.Abs(c) * (d + math.Abs(c)))
		},
	}

	gocalc.Graph(single, gocalc.GraphOption{
		Nsub: 100,
		Xmin: -2, Xmax: 2,
		Ymin: -10, Ymax: 10,
		Width: 600, Height: 600,
		Filename: "test_graph_2.png",
	})
}

func TestGraph3(t *testing.T) {
	double := Double{
		Rect: []Interval{
			{-2 * math.Pi, 2 * math.Pi},
			{-2 * math.Pi, 2 * math.Pi},
		},
		Eval: func(x, y float64) float64 {
			return x*x + y*y
		},
	}
	gocalc.Graph(double, gocalc.GraphOption{
		Nsub: 20,
		Xmin: -10, Xmax: 10,
		Ymin: -10, Ymax: 80,
		Width: 600, Height: 600,
		Filename: "test_graph_3.png",
	})
}
