package main

import (
	"aoc2024/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	lines []string
}

func preprocess(data []byte) Input {
	lines := strings.Split(string(data), "\n")
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

func part1(input Input) {
	sol := 0
	for _, l := range input.lines {
		first, last := -1, -1
		for _, r := range l {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				last = int(r - '0')
				if first < 0 {
					first = last
				}
			default:
				// ignore
			}
		}
		sol += 10*first + last
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func match(str, next string) bool {
	return len(str) >= len(next) && str[:len(next)] == next
}

func part2(input Input) {
	sol := 0
	for _, l := range input.lines {
		first, last := -1, -1
		for i := 0; i < len(l); i++ {
			digit := -1
			b := l[i]
			switch b {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				digit = int(b - '0')
			default:
				switch {
				case match(l[i:], "zero"):
					digit = 0
				case match(l[i:], "one"):
					digit = 1
				case match(l[i:], "two"):
					digit = 2
				case match(l[i:], "three"):
					digit = 3
				case match(l[i:], "four"):
					digit = 4
				case match(l[i:], "five"):
					digit = 5
				case match(l[i:], "six"):
					digit = 6
				case match(l[i:], "seven"):
					digit = 7
				case match(l[i:], "eight"):
					digit = 8
				case match(l[i:], "nine"):
					digit = 9
				}
			}
			if digit > -1 {
				last = digit
				if first < 0 {
					first = digit
				}
			}
		}
		sol += 10*first + last
	}
	fmt.Println("SOLUTION TO PART 2:", sol)
}
