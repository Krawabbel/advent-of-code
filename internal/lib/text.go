package lib

import (
	"fmt"
	"strings"
)

func SplitLines(txt string) []string {
	return strings.Split(txt, "\n")
}

func PrintFull[T any](t T) {
	fmt.Printf("%+v\n", t)
}
