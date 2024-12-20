package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	winning, personal []int
}

type Input struct {
	cards []Card
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))
	cards := make([]Card, len(lines))
	for i, line := range lines {
		next := strings.ReplaceAll(line, "  ", " ")
		for len(next) < len(line) {
			line = next
			next = strings.ReplaceAll(line, "  ", " ")
		}
		line = next
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := lib.MustToInts(strings.Split(parts[0], " "))
		personal := lib.MustToInts(strings.Split(parts[1], " "))
		cards[i] = Card{winning: winning, personal: personal}
	}
	return Input{cards: cards}
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
	for _, card := range in.cards {
		n := 0
		for _, p := range card.personal {
			for _, w := range card.winning {
				if p == w {
					n++
				}
			}
		}
		if n > 0 {
			sol += (1 << (n - 1))
		}
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	counts := make([]int, len(in.cards))

	for i, card := range in.cards {
		counts[i]++

		m := 0
		for _, p := range card.personal {
			for _, w := range card.winning {
				if p == w {
					m++
				}
			}
		}

		// fmt.Printf("%dx %d -> %d -> %dx %d..%d\n", counts[i], i, m, counts[i], i+1, i+m)

		for j := 0; j < m; j++ {
			if i+j+1 < len(counts) {
				counts[i+j+1] += counts[i]
			}
		}

	}

	sol := 0
	for _, c := range counts {
		sol += c
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}
