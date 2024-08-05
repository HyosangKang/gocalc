package gocalc_test

import (
	"gocalc"
)

type SimplePoint []SimpleReal

var _ gocalc.Vector = SimplePoint{}

func (p SimplePoint) Equals(q any) bool {
	if qq, ok := q.(gocalc.Point); !ok {
		return false
	} else if len(p) != qq.Len() {
		return false
	} else {
		for i := 0; i < len(p); i++ {
			if !p[i].Equals(qq.Map(i)) {
				return false
			}
		}
		return true
	}
}

func (p SimplePoint) Map(i int) gocalc.Real {
	return p[i]
}

func (p SimplePoint) Len() int {
	return len(p)
}

func (p SimplePoint) Add(r gocalc.Additive) gocalc.Additive {
	if pp, ok := r.(gocalc.Point); !ok {
		panic("SimplePoint.Add: input has no coordniate")
	} else {
		q := make(SimplePoint, len(p))
		for i := 0; i < len(p); i++ {
			q[i] = SimpleReal(p[i].ToFloat() + pp.Map(i).ToFloat())
		}
		return q
	}
}

func (p SimplePoint) Zero() gocalc.Additive {
	return make(SimplePoint, len(p))
}

func (p SimplePoint) AddInv() gocalc.Additive {
	q := make(SimplePoint, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = p.Map(i).AddInv().(SimpleReal)
	}
	return q
}

func (p SimplePoint) Scale(k gocalc.Real) gocalc.Vector {
	q := make(SimplePoint, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = p.Map(i).Mul(k).(SimpleReal)
	}
	return q
}

func (p SimplePoint) Inner(q gocalc.Vector) gocalc.Real {
	if qq, ok := q.(gocalc.Point); !ok {
		panic("SimplePoint.Inner: input has no coordniate")
	} else if len(p) != qq.Len() {
		panic("SimplePoint.Inner: dimension mismatch")
	} else {
		var s SimpleReal
		for i := 0; i < len(p); i++ {
			s = s.Add(p[i].Mul(qq.Map(i)).(gocalc.Real)).(SimpleReal)
		}
		return s
	}
}

func (p SimplePoint) Basis() <-chan gocalc.Vector {
	ch := make(chan gocalc.Vector)
	go func() {
		defer close(ch)
		for i := 0; i < len(p); i++ {
			q := make(SimplePoint, len(p))
			q[i] = SimpleReal(1)
			ch <- q
		}
	}()
	return ch
}
