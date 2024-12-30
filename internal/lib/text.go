package lib

import (
	"fmt"
	"strings"
)

func SplitLines(txt string) []string {
	return strings.Split(strings.ReplaceAll(txt, "\r\n", "\n"), "\n")
}

func PrintFull[T any](t T) {
	fmt.Printf("%+v\n", t)
}
