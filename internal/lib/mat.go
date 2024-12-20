package lib

import (
	"fmt"
	"slices"
)

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

func PrintMatf[T any](format string, ts [][]T) {
	for _, row := range ts {
		for _, t := range row {
			fmt.Printf(format, t)
		}
		fmt.Println()
	}
}
