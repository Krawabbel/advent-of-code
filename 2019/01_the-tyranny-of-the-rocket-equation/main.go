package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
)

type Input struct {
	mass []int
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))

	mass := make([]int, len(lines))
	for i, l := range lines {
		mass[i] = lib.MustToInt(l)
	}

	return Input{mass: mass}
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
	sol := 0
	for _, mass := range in.mass {
		sol += mass/3 - 2
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	for _, mass := range in.mass {
		fuel := mass/3 - 2
		for fuel > 0 {
			sol += fuel
			fuel = fuel/3 - 2
		}
	}
	fmt.Println("SOLUTION TO PART 2:", sol)
}
