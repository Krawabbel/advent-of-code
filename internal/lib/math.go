package lib

import "golang.org/x/exp/constraints"

func Abs[T constraints.Float | constraints.Integer](t T) T {
	if t > 0 {
		return t
	}
	return -t
}
