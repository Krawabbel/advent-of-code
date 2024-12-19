package lib

import (
	"fmt"
)

func Panicf(format string, a ...any) {
	err := fmt.Errorf(format, a...)
	panic(err)
}

func PrintSlice[T any](ts []T) {
	for _, t := range ts {
		fmt.Printf("%+v\n", t)
	}
}
