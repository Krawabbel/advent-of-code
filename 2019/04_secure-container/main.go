package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	a, b int
}

func preprocess(data []byte) Input {
	str := strings.Trim(string(data), "\n")
	parts := strings.Split(str, "-")
	a := lib.MustToInt(parts[0])
	b := lib.MustToInt(parts[1])
	return Input{a: a, b: b}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func to6Digits(n int) []int {
	out := make([]int, 6)

}

func explore(a, b []int, l int, same bool) int {
	n := len(a)

	if n == 0 {
		if same {
			return 0
		}
		return 1
	}

	for j := l + 1; j < 10; j++ {
		sum += explore(a, b, n-1, j, same)
	}
	return sum
}

func part1(in Input) {
	sol := explore(in.a, in.b, 6, 0, false)
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
