package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
)

type Input struct {
	lines []string
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))

	return Input{lines: lines}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func part1(in Input) {
	fmt.Printf("%+v\n", in)
	sol := 0
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
