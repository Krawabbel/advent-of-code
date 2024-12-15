package main

import (
	"aoc2024/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	count map[int]int
}

func preprocess(data []byte) Input {
	ages := lib.MustToInts(strings.Split(string(data), ","))
	count := make(map[int]int)
	for _, age := range ages {
		count[age]++
	}
	return Input{count: count}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func step(last map[int]int) map[int]int {
	next := make(map[int]int)
	for age, count := range last {
		if age == 0 {
			next[8] += count
			next[6] += count
		} else {
			next[age-1] += count
		}
	}
	return next
}

func simulate(input Input, days int) int {
	count := lib.CloneMap(input.count)
	for range days {
		count = step(count)
	}
	sol := 0
	for _, c := range count {
		sol += c
	}
	return sol
}

func part1(input Input) {
	fmt.Println("SOLUTION TO PART 1:", simulate(input, 80))
}

func part2(input Input) {
	fmt.Println("SOLUTION TO PART 2:", simulate(input, 256))
}
