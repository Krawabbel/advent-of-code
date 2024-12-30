package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

const (
	DIR_R = iota
	DIR_U
	DIR_L
	DIR_D
)

type Line struct {
	length, dir int
}

type Input struct {
	lines [][]Line
}

func preprocess(data []byte) Input {
	parts := lib.SplitLines(string(data))

	lines := make([][]Line, len(parts))
	for i, part := range parts {
		instrs := strings.Split(part, ",")
		lines[i] = make([]Line, len(instrs))

		for j, instr := range instrs {
			var dir int
			switch instr[0] {
			case 'R':
				dir = DIR_R
			case 'D':
				dir = DIR_D
			case 'L':
				dir = DIR_L
			case 'U':
				dir = DIR_U
			default:
				panic("unexpected direction")
			}

			length := lib.MustToInt(instr[1:])

			lines[i][j] = Line{dir: dir, length: length}
		}
	}

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

type Point struct {
	i, j int
}

func trace1(lines []Line) lib.Set[Point] {
	pos := lib.MakeSet[Point]()

	i, j := 0, 0
	for _, l := range lines {

		for range l.length {
			switch l.dir {
			case DIR_U:
				i++
			case DIR_D:
				i--
			case DIR_L:
				j--
			case DIR_R:
				j++
			}
			pos.Insert(Point{i, j})
		}
	}

	return pos
}

func part1(in Input) {

	lib.MustBeEqual(len(in.lines), 2)
	pos1 := trace1(in.lines[0])
	pos2 := trace1(in.lines[1])

	sol := -1
	for p := range pos1.Elements {
		if pos2.Contains(p) {
			dist := lib.Abs(p.i) + lib.Abs(p.j)
			if (dist > 0) && (sol < 0 || dist < sol) {
				sol = dist
			}
		}
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func trace2(lines []Line) map[Point]int {
	steps := make(map[Point]int)

	i, j := 0, 0
	s := 0
	for _, l := range lines {

		for range l.length {
			switch l.dir {
			case DIR_U:
				i++
			case DIR_D:
				i--
			case DIR_L:
				j--
			case DIR_R:
				j++
			}
			s++

			p := Point{i, j}
			if _, exists := steps[p]; !exists {
				steps[p] = s
			}
		}
	}

	return steps
}

func part2(in Input) {
	lib.MustBeEqual(len(in.lines), 2)
	steps1 := trace2(in.lines[0])
	steps2 := trace2(in.lines[1])

	sol := -1
	for p, s1 := range steps1 {
		if s2, contains := steps2[p]; contains {
			dist := s1 + s2
			if (dist > 0) && (sol < 0 || dist < sol) {
				sol = dist
			}
		}
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}
