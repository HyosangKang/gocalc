package gocalc

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Index []int

// ToInt converts an index to an integer with the maximum indices at each dimension
func (idx Index) ToInt(max []int) int {
	n := idx[len(idx)-1]
	for i := len(idx) - 2; i >= 0; i-- {
		n = n*(max[i]+1) + idx[i]
	}
	return n
}

// ToIndex converts an integer to an index with the maximum indices at each dimension
func ToIndex(i int, max []int) Index {
	idx := make(Index, len(max))
	for j := 0; j < len(max); j++ {
		idx[j] = i % (max[j] + 1)
		i /= (max[j] + 1)
	}
	return idx
}

func (idx Index) Neighbors(max []int) <-chan Index {
	ch := make(chan Index)
	go func() {
		defer close(ch)
		for i := 0; i < len(idx); i++ {
			if idx[i] < max[i] {
				jdx := make(Index, len(idx))
				copy(jdx, idx)
				jdx[i]++
				ch <- jdx
			}
		}
	}()
	return ch
}

type Cont interface {
	Eps(Real, Point) Real
}

type Grapher interface {
	Box
	Lift(Point) Point
	Project(Point) Point
}

func Graph(gpr Grapher, opt GraphOption) {
	img := image.NewPaletted(
		image.Rect(0, 0, opt.Width, opt.Height),
		[]color.Color{
			color.White,
			color.Black,
		})
	idxx, ps := Mesh(gpr, opt.Nsub)
	max := make([]int, gpr.Base().Len())
	for i := 0; i < gpr.Base().Len(); i++ {
		max[i] = opt.Nsub
	}
	for i, p := range ps {
		half := Rational(p.Map(0), [2]int{1, 2})
		idx := idxx[i]
		for jdx := range idx.Neighbors(max) {
			q := ps[jdx.ToInt(max)]
			pp, qq := gpr.Project(gpr.Lift(p)), gpr.Project(gpr.Lift(q))
			if cgpr, ok := gpr.(Cont); ok {
				d := Distance(p, q)
				e := Distance(pp, qq)
				m := p.Add(q).(Vector).Scale(half).(Vector)
				if e.GreaterThan(cgpr.Eps(d, m)) {
					opt.DrawLine(img, pp, qq)
				}
			} else {
				opt.DrawLine(img, pp, qq)
			}
		}
	}
	fp, _ := os.Create(opt.Filename)
	defer fp.Close()
	png.Encode(fp, img)
}

func Contour(gpr Grapher, opt GraphOption) {
}

type GraphOption struct {
	Nsub                   int
	Xmin, Xmax, Ymin, Ymax float64
	Width, Height          int
	Levels                 []float64 // for Contour only
	Filename               string
}

func (opt GraphOption) DrawLine(img *image.Paletted, p, q Point) {
	x0, y0 := p.Map(0), p.Map(1)
	x1, y1 := q.Map(0), q.Map(1)
	i0, j0 := opt.Pixel(x0.ToFloat(), y0.ToFloat())
	i1, j1 := opt.Pixel(x1.ToFloat(), y1.ToFloat())

	DrawPoint := func(i, j int) {
		for l := -1; l <= 1; l++ {
			img.Set(i+l, j, color.Black)
			img.Set(i, j+l, color.Black)
		}
	}

	n := Max(Abs(i0-i1), Abs(j0-j1))
	if n == 0 {
		DrawPoint(i0, j0)
	} else {
		is := Linspace(i0, i1, n)
		js := Linspace(j0, j1, n)
		for k := 0; k <= int(n); k++ {
			i, j := is[k], js[k]
			DrawPoint(i, j)
		}
	}
}

func (opt GraphOption) Pixel(x, y float64) (int, int) {
	dw := (opt.Xmax - opt.Xmin) / float64(opt.Width)
	dh := (opt.Ymax - opt.Ymin) / float64(opt.Height)
	i := int((x - opt.Xmin) / dw)
	j := int((opt.Ymax - y) / dh)
	return i, j
}

func Linspace[T int | float64](a, b T, n int) []T {
	if n <= 0 {
		return nil
	}
	af, bf := float64(a), float64(b)
	var h float64 = (bf - af) / float64(n)
	var arr []T
	for i := 0; i <= n; i++ {
		arr = append(arr, T(af+float64(i)*h))
	}
	return arr
}

func Max[T int | float64](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func Abs[T int | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
