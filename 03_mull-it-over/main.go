package main

import (
	"aoc2024/internal/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const rePart1Expr = `mul\(([0-9]{1,3}),([0-9]{1,3})\)`

var rePart1 = regexp.MustCompile(rePart1Expr)

const rePart2Expr = `(mul)\(([0-9]{1,3}),([0-9]{1,3})\)|(do\(\))|(don't\(\))`

var rePart2 = regexp.MustCompile(rePart2Expr)

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	util.Must(err)

	input := string(blob)

	part1(input)

	part2(input)
}

func part1(input string) {

	matches := rePart1.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, m := range matches {
		a, err := strconv.Atoi(m[1])
		util.Must(err)

		b, err := strconv.Atoi(m[2])
		util.Must(err)

		sum += a * b
	}

	fmt.Println("SOLUTION PART 1:", sum)
}

func part2(input string) {
	matches := rePart2.FindAllStringSubmatch(input, -1)

	do := true
	sum := 0
	for _, m := range matches {

		switch {
		case m[1] == "mul":
			if do {

				a, err := strconv.Atoi(m[2])
				util.Must(err)

				b, err := strconv.Atoi(m[3])
				util.Must(err)

				sum += a * b
			}
		case m[4] == "do()":
			do = true
		case m[5] == "don't()":
			do = false
		}

	}
	fmt.Println("SOLUTION PART 2:", sum)
}
