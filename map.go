package gocalc

type Map[T1, T2 any] interface {
	Map(T1) T2
}
