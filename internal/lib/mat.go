package lib

import "slices"

func Mat[T any](m, n int, t T) [][]T {
	mat := make([][]T, m)
	for i := range mat {
		mat[i] = slices.Repeat([]T{t}, n)
	}
	return mat
}

func InMatBounds[T any](mat [][]T, i, j int) bool {
	return i >= 0 && i < len(mat) && j >= 0 && j < len(mat[i])
}
