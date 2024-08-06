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
