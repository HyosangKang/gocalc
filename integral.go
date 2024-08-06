package gocalc

type VectorValued interface {
	Region
	Map[Point, Vector]
}

type Manifold interface {
	Locals() <-chan VectorValued
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
	var sum Real
	for prm := range f.Locals() {
		rank := prm.Corner().Len()
		mesh := Mesh(prm, nsub)
		for wdg := range f.Wedges() {
			for i, p := range mesh {
				nbhr := Neighbors(i, nsub, rank)
				if len(nbhr) != rank {
					continue
				}
				bdry := make([]Vector, rank)
				for i := 0; i < rank; i++ {
					bdry[i] = mesh[nbhr[i]].Add(p.AddInv()).(Vector)
				}
				vol := Volume(bdry, wdg.Wedge())
				sum = sum.Add(vol.Mul(wdg.Map(p)).(Real)).(Real)
			}
		}
	}
	return sum
}

func Volume(bdry []Vector, wdg []int) Real {
	mat := make([][]Real, len(bdry))
	for i := 0; i < len(bdry); i++ {
		mat[i] = make([]Real, len(bdry))
		for j := 0; j < len(bdry); j++ {
			mat[i][j] = bdry[i].Map(wdg[j]).(Real)
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
