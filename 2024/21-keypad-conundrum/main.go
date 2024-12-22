package main

import (
	"aoc/internal/lib"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	codes [][]byte
}

func preprocess(data []byte) Input {
	codes := bytes.Split(data, []byte{'\n'})

	return Input{codes: codes}
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

func PointSub(u, v Point) Point {
	return Point{u.i - v.i, u.j - v.j}
}

func PointAdd(u, v Point) Point {
	return Point{u.i + v.i, u.j + v.j}
}

const (
	PAD_UP    = '^'
	PAD_LEFT  = '<'
	PAD_DOWN  = 'v'
	PAD_RIGHT = '>'
	PAD_A     = 'A'
	PAD_0     = '0'
	PAD_1     = '1'
	PAD_2     = '2'
	PAD_3     = '3'
	PAD_4     = '4'
	PAD_5     = '5'
	PAD_6     = '6'
	PAD_7     = '7'
	PAD_8     = '8'
	PAD_9     = '9'
)

var POS_NUMERICAL = map[byte]Point{
	PAD_7: {0, 0},
	PAD_8: {0, 1},
	PAD_9: {0, 2},
	PAD_4: {1, 0},
	PAD_5: {1, 1},
	PAD_6: {1, 2},
	PAD_1: {2, 0},
	PAD_2: {2, 1},
	PAD_3: {2, 2},
	PAD_0: {3, 1},
	PAD_A: {3, 2},
}

var POS_DIRECTIONAL = map[byte]Point{
	PAD_UP:    {0, 1},
	PAD_A:     {0, 2},
	PAD_LEFT:  {1, 0},
	PAD_DOWN:  {1, 1},
	PAD_RIGHT: {1, 2},
}

var BUTTONS_DIRECTIONAL = map[Point]byte{}
var BUTTONS_NUMERICAL = map[Point]byte{}

func init() {

	for key, val := range POS_NUMERICAL {
		BUTTONS_NUMERICAL[val] = key
	}

	for key, val := range POS_DIRECTIONAL {
		BUTTONS_DIRECTIONAL[val] = key
	}
}

var dirs = map[byte]Point{
	PAD_UP:    {-1, 0},
	PAD_DOWN:  {+1, 0},
	PAD_LEFT:  {0, -1},
	PAD_RIGHT: {0, +1},
}

type Edge struct {
	begin, end byte
}

func calcMoves(dist Point) map[byte]int {
	mv := make(map[byte]int)

	if dist.i < 0 {
		mv[PAD_UP] -= dist.i
	} else if dist.i == 0 {
		// do nothing
	} else {
		mv[PAD_DOWN] += dist.i
	}

	if dist.j > 0 {
		mv[PAD_RIGHT] += dist.j
	} else if dist.j == 0 {
		// do nothing
	} else {
		mv[PAD_LEFT] -= dist.j
	}

	return mv
}

func explorePaths(pos0 Point, moves map[byte]int, positions map[byte]Point, buttons map[Point]byte) [][]byte {

	if len(moves) == 0 {
		// fmt.Println("END")
		return make([][]byte, 1)
	}

	paths := make([][]byte, 0)

loop:
	for dir, vec := range dirs {

		// fmt.Println("EXPLORE", string(dir))

		steps, found := moves[dir]
		if !found {
			// fmt.Println("CONTINUE: no move")
			continue loop
		}

		pos1 := pos0
		for range steps {
			pos1 = PointAdd(pos1, vec)
			if _, valid := buttons[pos1]; !valid {
				// fmt.Println("CONTINUE: invalid position")
				continue loop
			}
		}

		// fmt.Println("NEXT")

		delete(moves, dir)
		ps := explorePaths(pos1, moves, positions, buttons)
		moves[dir] = steps

		// fmt.Println("NEXT Paths", ps)

		prev := bytes.Repeat([]byte{dir}, steps)
		for _, next := range ps {
			full := append(prev, next...)
			paths = append(paths, full)
		}
	}

	return paths
}

func calcPaths(begin, end byte, positions map[byte]Point, buttons map[Point]byte) [][]byte {
	if begin == end {
		return make([][]byte, 1)
	}

	p0, valid := positions[begin]
	if !valid {
		return nil
	}

	p1, valid := positions[end]
	if !valid {
		return nil
	}

	dist := PointSub(p1, p0)

	moves := calcMoves(dist)

	// fmt.Printf("CALC : %s -> %s : %v\n", string(begin), string(end), dist)

	return explorePaths(p0, moves, positions, buttons)
}

func makeMap(positions map[byte]Point, buttons map[Point]byte) map[Edge][][]byte {

	dict := make(map[Edge][][]byte)

	for begin := range positions {
		for end := range positions {
			edge := Edge{begin: begin, end: end}

			paths := calcPaths(begin, end, positions, buttons)
			dict[edge] = paths

			// fmt.Printf("EDGE %s -> %s\n", string(begin), string(end))
			// for _, path := range paths {
			// 	fmt.Println(" * ", string(path))
			// }
		}
	}

	return dict
}

func buildSeq(code []byte, maps []map[Edge][][]byte) string {

	if len(maps) == 0 {
		return string(code)
	}

	seq := ""

	last := byte(PAD_A)
	for _, next := range code {
		best := "INITIAL"
		paths := maps[0][Edge{last, next}]
		for _, p := range paths {
			s := buildSeq(append(p, PAD_A), maps[1:])
			if len(s) < len(best) || best == "INITIAL" {
				best = s
			}
		}

		if best != "INITIAL" {
			seq += best
		}

		last = next
	}

	return seq
}

func part1(in Input) {

	mapDirpad := makeMap(POS_DIRECTIONAL, BUTTONS_DIRECTIONAL)

	mapNumpad := makeMap(POS_NUMERICAL, BUTTONS_NUMERICAL)

	// paths, found := mapNumpad[Edge{PAD_7, PAD_0}]
	// lib.MustBeTrue(found)
	// for _, p := range paths {
	// 	fmt.Println("===>", string(p))
	// }

	// lib.MustPressEnter()

	sol := 0
	for _, code := range in.codes {
		seq := buildSeq(code, []map[Edge][][]byte{mapNumpad, mapDirpad, mapDirpad})
		fmt.Printf("CODE %s: %s\n", string(code), seq)

		numericPart := lib.MustToInt(strings.ReplaceAll(string(code), "A", ""))
		cost := len(seq)
		sol += cost * numericPart

		// // *** SIMULATION / VALIDATION ***
		// simRobots := make([]Point, len(padsBruteForce))
		// for i, pad := range padsBruteForce {
		// 	simRobots[i] = pad[PAD_A]
		// }
		// have := sim(seq, simRobots, padsBruteForce)
		// lib.MustBeEqual(string(code), have)
		// lib.MustPressEnter()
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}

// var padsBruteForce = []map[byte]Point{POS_DIRECTIONAL, POS_DIRECTIONAL, POS_NUMERICAL}

// func isValidPosition(pad map[byte]Point, pos Point) bool {
// 	for _, p := range pad {
// 		if p == pos {
// 			return true
// 		}
// 	}
// 	return false
// }

// func sim(seq string, robots []Point, pads []map[byte]Point) string {
// 	if seq == "" {
// 		return ""
// 	}
// 	return apply(seq[0], robots, pads) + sim(seq[1:], robots, pads)
// }

// func apply(button byte, robots []Point, pads []map[byte]Point) string {

// 	if len(robots) == 0 {
// 		return string(button)
// 	}

// 	if button == PAD_A {
// 		for b, p := range pads[0] {
// 			if p == robots[0] {
// 				return apply(b, robots[1:], pads[1:])
// 			}
// 		}
// 		panic("got here")
// 	}

// 	robots[0] = PointAdd(robots[0], dirs[button])

// 	if !isValidPosition(pads[0], robots[0]) {
// 		panic("invalid position")
// 	}

// 	return ""
// }
