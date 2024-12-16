package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result int
	args   []int
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := string(blob)

	lines := strings.Split(input, "\n")

	eqs := make([]equation, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, ": ")
		result, err := strconv.Atoi(parts[0])
		lib.Must(err)
		strArgs := strings.Split(parts[1], " ")
		args := make([]int, len(strArgs))
		for j, a := range strArgs {
			arg, err := strconv.Atoi(a)
			lib.Must(err)
			args[j] = arg
		}
		eqs[i] = equation{result: result, args: args}
	}

	part1(eqs)

	part2(eqs)
}

func solvePart1(res, tmp int, args []int) bool {
	if len(args) == 0 {
		return tmp == res
	}

	if solvePart1(res, tmp+args[0], args[1:]) {
		return true
	}

	return solvePart1(res, tmp*args[0], args[1:])
}

func part1(eqs []equation) {

	sol := 0

	for _, eq := range eqs {
		if solvePart1(eq.result, 0, eq.args) {
			sol += eq.result
		}
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func solvePart2(res, tmp int, args []int) bool {
	if len(args) == 0 {
		return tmp == res
	}

	if solvePart2(res, tmp+args[0], args[1:]) {
		return true
	}

	if solvePart2(res, tmp*args[0], args[1:]) {
		return true
	}

	val, err := strconv.Atoi(fmt.Sprintf("%d%d", tmp, args[0]))
	lib.Must(err)
	return solvePart2(res, val, args[1:])
}

func part2(eqs []equation) {

	sol := 0

	for _, eq := range eqs {
		if solvePart2(eq.result, 0, eq.args) {
			sol += eq.result
		}
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}
