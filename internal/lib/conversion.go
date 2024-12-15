package lib

import "strconv"

func MustToInt(s string) int {
	i, err := strconv.Atoi(s)
	Must(err)
	return i
}

func MustToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i] = MustToInt(s)
	}
	return ints
}
