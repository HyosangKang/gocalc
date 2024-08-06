package gocalc_test

import (
	"gocalc"
)

type SimpleVector []float64

func (p SimpleVector) Equals(q any) bool {
	if qq, ok := q.(gocalc.Point); !ok {
		return false
	} else if len(p) != qq.Len() {
		return false
	} else {
		for i := 0; i < len(p); i++ {
			if !SimpleReal(p[i]).Equals(qq.Map(i)) {
				return false
			}
		}
		return true
	}
}

func (p SimpleVector) Map(i int) gocalc.Real {
	return SimpleReal(p[i])
}

func (p SimpleVector) Len() int {
	return len(p)
}

func (p SimpleVector) Add(r gocalc.Additive) gocalc.Additive {
	if pp, ok := r.(gocalc.Point); !ok {
		panic("SimplePoint.Add: input has no coordniate")
	} else {
		q := make(SimpleVector, len(p))
		for i := 0; i < len(p); i++ {
			q[i] = p[i] + pp.Map(i).ToFloat()
		}
		return q
	}
}

func (p SimpleVector) Zero() gocalc.Additive {
	return make(SimpleVector, len(p))
}

func (p SimpleVector) AddInv() gocalc.Additive {
	q := make(SimpleVector, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = -p[i]
	}
	return q
}

func (p SimpleVector) Scale(k gocalc.Real) gocalc.Vector {
	q := make(SimpleVector, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = k.ToFloat() * p[i]
	}
	return q
}

func (p SimpleVector) Inner(q gocalc.Vector) gocalc.Real {
	if qq, ok := q.(gocalc.Point); !ok {
		panic("SimplePoint.Inner: input has no coordniate")
	} else if len(p) != qq.Len() {
		panic("SimplePoint.Inner: dimension mismatch")
	} else {
		var s float64
		for i := 0; i < len(p); i++ {
			s += p[i] * qq.Map(i).ToFloat()
		}
		return SimpleReal(s)
	}
}

func (p SimpleVector) Basis() []gocalc.Vector {
	basis := make([]gocalc.Vector, p.Len())
	for i := 0; i < p.Len(); i++ {
		basis[i] = make(SimpleVector, p.Len())
		basis[i].(SimpleVector)[i] = 1
	}
	return basis
}
