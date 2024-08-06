package gocalc

// func Union(s, t FiniteSet[T]) FiniteSet {
// 	u := s.Empty()
// 	for x := s.Pop(); x != nil; x = s.Pop() {
// 		u.Insert(x)
// 	}
// 	for x := t.Pop(); x != nil; x = t.Pop() {
// 		u.Insert(x)
// 	}
// 	return u
// }

// func Intersection(s, t FiniteSet) FiniteSet {
// 	u := s.Empty()
// 	var set []Element
// 	for x := s.Pop(); x != nil; x = s.Pop() {
// 		set = append(set, x)
// 		if t.Contains(x) {
// 			u.Insert(x)
// 		}
// 	}
// 	for _, x := range set {
// 		s.Insert(x)
// 	}
// 	return u
// }

// func Subset(s, t FiniteSet) bool {
// 	isSubset := true
// 	var set []Element
// 	for x := s.Pop(); x != nil; x = s.Pop() {
// 		set = append(set, x)
// 		if !t.Contains(x) {
// 			isSubset = false
// 			break
// 		}
// 	}
// 	for _, x := range set {
// 		s.Insert(x)
// 	}
// 	return isSubset
// }

// func Combination(s FiniteSet, r int) <-chan FiniteSet {
// 	var set []Element
// 	for x := s.Pop(); x != nil; x = s.Pop() {
// 		isNew := true
// 		for _, y := range set {
// 			if y.Equals(x) {
// 				isNew = false
// 				break
// 			}
// 		}
// 		if isNew {
// 			set = append(set, x)
// 		}
// 	}
// 	ch := make(chan FiniteSet)
// 	go func() {
// 		defer close(ch)
// 		for idx := range CombInt(len(set), r) {
// 			new := s.Empty()
// 			for _, i := range idx {
// 				new.Insert(set[i])
// 			}
// 			ch <- new
// 		}
// 	}()
// 	return ch
// }

// // CombInt returns the set of all r-combinations of a int from 0 to n-1.
// func CombInt(n, r int) <-chan []int {
// 	ch := make(chan []int)
// 	go func() {
// 		defer close(ch)
// 		if n == r {
// 			var idx []int
// 			for i := 0; i < n; i++ {
// 				idx = append(idx, i)
// 			}
// 			ch <- idx
// 		} else if r == 0 {
// 			ch <- []int{}
// 		} else {
// 			for idx := range CombInt(n-1, r-1) {
// 				ch <- append(idx, n-1)
// 			}
// 			for idx := range CombInt(n-1, r) {
// 				ch <- idx
// 			}
// 		}
// 	}()
// 	return ch
// }
