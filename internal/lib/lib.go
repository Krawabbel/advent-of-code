package lib

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Panicf(format string, a ...any) {
	err := fmt.Errorf(format, a...)
	panic(err)
}

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

func MustBeEqual[T comparable](t1, t2 T) {
	if !reflect.DeepEqual(t1, t2) {
		Panicf("%v != %v", t1, t2)
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

func MustPressEnter() {
	fmt.Println("Press 'Enter' to continue...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	Must(err)
}

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

func Abs[T constraints.Signed](t T) T {
	if t > 0 {
		return t
	}
	return -t
}

type Element struct{}

func MakeElement() Element {
	return Element{}
}

func (e *Element) IsElement() bool {
	return true
}
