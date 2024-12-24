package lib

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

type Vector[T Number] struct {
	coords []T
}

func Vec[T Number](ts ...T) Vector[T] {
	return Vector[T]{coords: ts}
}

func (v Vector[T]) Slice() []T {
	return v.coords
}

func (v Vector[T]) Size() int {
	return len(v.coords)
}

func (v Vector[T]) Get(i int) T {
	return v.coords[i]
}

func (v Vector[T]) String() string {
	return "VEC [" + v.Hash("%v", ", ") + "]"
}

func (v Vector[T]) Hash(format, sep string) string {
	f := func(t T) string { return fmt.Sprintf(format, t) }
	strs := Transform(v.Slice(), f)
	return strings.Join(strs, sep)
}

func VecAdd[T Number](u, v Vector[T]) Vector[T] {
	MustBeEqual(u.Size(), v.Size())
	w := make([]T, u.Size())
	for i := range w {
		w[i] = u.coords[i] + v.coords[i]
	}
	return Vector[T]{coords: w}
}

func VecSub[T Number](u, v Vector[T]) Vector[T] {
	MustBeEqual(u.Size(), v.Size())
	w := make([]T, u.Size())
	for i := range w {
		w[i] = u.coords[i] - v.coords[i]
	}
	return Vector[T]{coords: w}
}

func VecScalarMult[T Number](u T, v Vector[T]) Vector[T] {
	f := func(t T) T { return u * t }
	w := Transform(v.Slice(), f)
	return Vector[T]{coords: w}
}

func VecInnerProd[T Number](u, v Vector[T]) T {
	MustBeEqual(u.Size(), v.Size())
	var w T = 0
	for i := range u.Size() {
		w += u.coords[i] * v.coords[i]
	}
	return w
}
