package main

import (
	"aoc2024/internal/lib"
	"fmt"
	"os"
	"strings"
)

type antenna struct {
	x, y int
	freq rune
}

func isAntenna(r rune) bool {
	isLowercaseLetter := r >= 'a' && r <= 'z'
	isUppercaseLetter := r >= 'A' && r <= 'Z'
	isDigit := r >= '0' && r <= '9'
	return isLowercaseLetter || isUppercaseLetter || isDigit
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := string(blob)

	antennas := make([]antenna, 0)

	lines := strings.Split(input, "\n")

	m := len(lines)
	n := len(lines[0])

	for i, l := range lines {
		for j, r := range l {
			switch {
			case r == '.':
				// ignore
			case isAntenna(r):
				a := antenna{x: i, y: j, freq: r}
				antennas = append(antennas, a)
			default:
				fmt.Println("ignored", string(r))
			}
		}
	}

	part1(m, n, antennas)

	part2(m, n, antennas)
}

func inBounds(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func hash(x, y int) string {
	return fmt.Sprintf("%d+%d", x, y)
}

func part1(m, n int, as []antenna) {

	nodes := map[string]struct{}{}

	for i, a1 := range as {
		for j := i + 1; j < len(as); j++ {
			a2 := as[j]

			if a1.freq != a2.freq {
				continue
			}

			dx := a2.x - a1.x
			dy := a2.y - a1.y

			n1x := a1.x + 2*dx
			n1y := a1.y + 2*dy
			if inBounds(n1x, n1y, m, n) {
				nodes[hash(n1x, n1y)] = struct{}{}
			}

			n2x := a2.x - 2*dx
			n2y := a2.y - 2*dy
			if inBounds(n2x, n2y, m, n) {
				nodes[hash(n2x, n2y)] = struct{}{}
			}
		}
	}

	fmt.Println("SOLUTION TO PART 1:", len(nodes))
}

func part2(m, n int, as []antenna) {

	nodes := map[string]struct{}{}

	for i, a1 := range as {
		for j := i + 1; j < len(as); j++ {
			a2 := as[j]

			if a1.freq != a2.freq {
				continue
			}

			dx := a2.x - a1.x
			dy := a2.y - a1.y

			for n1x, n1y := a1.x, a1.y; inBounds(n1x, n1y, m, n); n1x, n1y = n1x+dx, n1y+dy {
				nodes[hash(n1x, n1y)] = struct{}{}
			}

			for n2x, n2y := a1.x, a1.y; inBounds(n2x, n2y, m, n); n2x, n2y = n2x-dx, n2y-dy {
				nodes[hash(n2x, n2y)] = struct{}{}
			}

		}
	}

	fmt.Println("SOLUTION TO PART 2:", len(nodes))
}
