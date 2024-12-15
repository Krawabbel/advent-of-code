package lib

func CloneSlice1D[T any](ts []T) []T {
	cp := make([]T, len(ts))
	copy(cp, ts)
	return cp
}

func CloneSlice2D[T any](mat [][]T) [][]T {
	cp := make([][]T, len(mat))
	for i, vec := range mat {
		cp[i] = CloneSlice1D(vec)
	}
	return cp
}

func CloneMap[K comparable, T any](m map[K]T) map[K]T {
	cp := make(map[K]T)
	for key, val := range m {
		cp[key] = val
	}
	return cp
}
