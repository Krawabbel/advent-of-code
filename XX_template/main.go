package main

import (
	"aoc2024/internal/util"
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	util.Must(err)

	input := string(blob)

	part1(input)

	part2(input)
}

func part2(input string) {

	sol := 0

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part1(input string) {

	sol := 0

	fmt.Println("SOLUTION TO PART 2:", sol)
}
