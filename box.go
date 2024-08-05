package gocalc

type Bounded interface {
	Dim() int
	Inf(int) Real
	Sup(int) Real
}

type Region interface {
	Set[Point]
	Bounded
}

type Box interface {
	Base() Vector
	Corners() <-chan Vector
}

// Mesh returns the mesh grid of a region.
func Mesh(b Box, nsub int) ([]Index, []Vector) {
	p := b.Base()
	f := p.Map(0)
	df := Rational(f, [2]int{1, nsub})
	var vds []Vector
	for q := range b.Corners() {
		vd := q.Add(p.AddInv()).(Vector).Scale(df).(Vector)
		vds = append(vds, vd)
	}
	var idx []Index
	var ps []Vector
	max := make([]int, len(vds))
	for i := range vds {
		max[i] = nsub
	}
	n := 0
	for v := range SubInc(vds, nsub) {
		idx = append(idx, ToIndex(n, max))
		ps = append(ps, v)
	}
	return idx, ps
}

func SubInc(vds []Vector, n int) <-chan Vector {
	if len(vds) < 1 {
		panic("SubInc: insufficienet vectors")
	}
	if n < 1 {
		panic("SubInc: Wrong number of increments")
	}
	ch := make(chan Vector)
	go func() {
		defer close(ch)
		p := vds[0].Zero().(Vector)
		for i := 0; i <= n; i++ {
			if len(vds) == 1 {
				ch <- p
			} else {
				for v := range SubInc(vds[1:], n) {
					ch <- p.Add(v).(Vector)
				}
			}
			p = p.Add(vds[0]).(Vector)
		}
	}()
	return ch
}
