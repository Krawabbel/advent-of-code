package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := string(blob)

	part1(input)

	part2(input)
}

func makeXMAS(fwd bool) string {
	if fwd {
		return "XMAS"
	}
	return "SAMX"
}

func find_part1(text string, fwd bool, dist int) [][]int {

	xmas := makeXMAS(fwd)

	matches := [][]int{}

	for i := range text {

		idxs := make([]int, 4)
		found := true

		for j := range 4 {

			pos := i + j*dist

			if pos < 0 || pos >= len(text) || text[pos] != xmas[j] {
				found = false
				break
			}

			idxs[j] = pos

		}

		if found {
			matches = append(matches, idxs)
		}

	}

	return matches
}

func part1(input string) {

	lines := strings.Split(input, "\n")
	cols := len(lines[0])
	text := strings.Join(lines, "*")

	matches := make([][]int, 0)
	matches = append(matches, find_part1(text, true, 1)...)
	matches = append(matches, find_part1(text, false, 1)...)
	matches = append(matches, find_part1(text, true, cols)...)
	matches = append(matches, find_part1(text, false, cols)...)
	matches = append(matches, find_part1(text, true, cols+1)...)
	matches = append(matches, find_part1(text, false, cols+1)...)
	matches = append(matches, find_part1(text, true, cols+2)...)
	matches = append(matches, find_part1(text, false, cols+2)...)

	paint(text, matches)

	sol := len(matches)

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func substr(text string, idxs ...int) string {
	str := ""
	for _, idx := range idxs {
		if idx >= len(text) {
			return ""
		}
		str += string(text[idx])
	}
	return str
}

func paint(text string, matches [][]int) {
	tmp := text
	tmp = strings.ReplaceAll(tmp, "X", ".")
	tmp = strings.ReplaceAll(tmp, "M", ".")
	tmp = strings.ReplaceAll(tmp, "A", ".")
	tmp = strings.ReplaceAll(tmp, "S", ".")
	for _, match := range matches {
		for _, idx := range match {
			tmp = tmp[:idx] + string(text[idx]) + tmp[idx+1:]
		}
	}
	tmp = strings.ReplaceAll(tmp, "*", "\n")
	fmt.Println(tmp)
}

func part2(input string) {

	lines := strings.Split(input, "\n")
	cols := len(lines[0])
	text := strings.Join(lines, "*")

	matches := make([][]int, 0)

	for i := range text {
		idxs := []int{i, i + 2, i + cols + 2, i + 2*cols + 2, i + 2*cols + 4}
		switch substr(text, idxs...) {
		case "MSAMS", "SMASM", "MMASS", "SSAMM":
			matches = append(matches, idxs)
		}
	}

	paint(text, matches)

	sol := len(matches)

	fmt.Println("SOLUTION TO PART 2:", sol)
}
