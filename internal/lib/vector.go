package lib

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

type Vector[T Number] struct {
	coords []T
}

func Vec[T Number](ts ...T) Vector[T] {
	return Vector[T]{coords: ts}
}

func (v *Vector[T]) Slice() []T {
	return v.coords
}

func (v *Vector[T]) Size() int {
	return len(v.coords)
}

func (v *Vector[T]) Get(i int) T {
	return v.coords[i]
}

func VecAdd[T Number](u, v Vector[T]) Vector[T] {
	MustBeEqual(u.Size(), v.Size())
	w := make([]T, u.Size())
	for i := range w {
		w[i] = u.coords[i] + v.coords[i]
	}
	return Vector[T]{coords: w}
}
