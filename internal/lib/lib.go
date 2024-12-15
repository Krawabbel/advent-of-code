package lib

import (
	"fmt"
)

func Panicf(format string, a ...any) {
	err := fmt.Errorf(format, a...)
	panic(err)
}
