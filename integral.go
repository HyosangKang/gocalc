package gocalc

type Parametric interface {
	Region
	Map[Point, Vector]
}

type Manifold interface {
	Locals() <-chan Parametric
}

type Wedge interface {
	Map[Point, Real]
	Wedge() []int
}

type Form interface {
	Manifold
	Wedges() <-chan Wedge
}

func Integral(f Form, nsub int) Real {
	var sum []Real
	for prm := range f.Locals() {
		s := prm.Corner().Map(0).Zero().(Real)
		rank := prm.Corner().Len()
		mesh := Mesh(prm, nsub)
		for wdg := range f.Wedges() {
			for i, p := range mesh {
				pp := prm.Map(p)
				nbhr := Neighbors(i, nsub, rank)
				if len(nbhr) != rank {
					continue
				}
				bdry := make([]Vector, rank)
				for j := 0; j < rank; j++ {
					qq := prm.Map(mesh[nbhr[j]])
					bdry[j] = qq.Add(pp.AddInv()).(Vector)
				}
				vol := Volume(bdry, wdg.Wedge())
				s = s.Add(vol.Mul(wdg.Map(pp)).(Real)).(Real)
			}
		}
		sum = append(sum, s)
	}
	v := sum[0]
	for i := 1; i < len(sum); i++ {
		v = v.Add(sum[i]).(Real)
	}
	return v
}

func Volume(bdry []Vector, wdg []int) Real {
	mat := make([][]Real, len(bdry))
	for i := 0; i < len(bdry); i++ {
		mat[i] = make([]Real, len(bdry))
		for j := 0; j < len(bdry); j++ {
			mat[i][j] = bdry[i].Map(wdg[j])
		}
	}
	return Det(mat)
}

func Det(mat [][]Real) Real {
	r, c := len(mat), len(mat[0])
	if r != c {
		panic("Det: matrix is not square")
	}
	if r == 1 {
		return mat[0][0]
	}
	det := mat[0][0].Zero().(Real)
	for i := 0; i < r; i++ {
		var sub [][]Real
		for j := 0; j < r; j++ {
			if j == i {
				continue
			}
			row := make([]Real, r-1)
			copy(row, mat[j][1:])
			sub = append(sub, row)
		}
		sign := mat[0][0].One().(Real)
		if i%2 == 1 {
			sign = sign.AddInv().(Real)
		}
		det = det.Add(Det(sub).Mul(sign).Mul(mat[i][0]).(Real)).(Real)
	}
	return det
}
