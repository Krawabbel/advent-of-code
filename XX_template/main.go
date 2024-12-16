package main

import (
	"aoc/internal/lib"
	"bytes"
	"fmt"
	"os"
)

type Input struct {
	grid [][]byte
}

func preprocess(data []byte) Input {
	grid := bytes.Split(data, []byte{'\n'})

	return Input{grid: grid}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func part1(input Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(input Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
