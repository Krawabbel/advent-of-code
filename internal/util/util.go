package util

import "strconv"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustBeTrue(b bool) {
	if !b {
		panic("false is not true")
	}
}

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
