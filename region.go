package gocalc

type Set[T Element] interface {
	Contains(T) bool
}

type Region interface {
	Set[Point]
	Corner() Vector
	Lengths() []Real
}

func Mesh(reg Region, nsub int) []Vector {
	base := reg.Corner()
	var bdry []Vector
	lengths := reg.Lengths()
	basis := base.Basis()
	if len(lengths) != len(basis) {
		panic("Mesh: lengths and basis have different dimensions")
	}
	for i, b := range basis {
		bb := b.Scale(lengths[i].Mul(Rational(base.Map(0), [2]int{1, nsub})).(Real))
		bdry = append(bdry, bb)
	}
	var mesh []Vector
	for v := range SubInc(bdry, nsub) {
		mesh = append(mesh, base.Add(v).(Vector))
	}
	return mesh
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
