package main

import (
	"aoc2024/internal/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	fields []int
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	util.Must(err)

	input := string(blob)

	strLines := strings.Split(input, "\n")

	lines := make([]line, len(strLines))
	for i, l := range strLines {
		strFields := strings.Split(l, " ")

		fields := make([]int, len(strFields))
		for j, f := range strFields {
			field, err := strconv.Atoi(f)
			util.Must(err)
			fields[j] = field
		}
		lines[i] = line{fields: fields}
	}

	part1(lines)

	part2(lines)
}

func part1(lines []line) {

	sol := 0

	for i, l := range lines {
		fmt.Println(i, l)
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(eqs []line) {
	sol := 0

	fmt.Println("SOLUTION TO PART 2:", sol)
}
