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

type Args struct {
	last, next byte
	level      int
}

var cache = map[Args]int{}

func buildEdge(last, next byte, maps []map[Edge][][]byte) int {
	args := Args{last, next, len(maps)}
	if _, found := cache[args]; !found {
		cache[args] = buildEdgeImpl(last, next, maps)
	}
	return cache[args]
}

func buildEdgeImpl(last, next byte, maps []map[Edge][][]byte) int {

	best := -1

	paths := maps[0][Edge{last, next}]
	for _, p := range paths {
		s := buildRemainder(PAD_A, append(p, PAD_A), maps[1:])
		if s < best || best == -1 {
			best = s
		}
	}

	lib.MustBeTrue(best >= 0)

	return best
}

func buildRemainder(last byte, code []byte, maps []map[Edge][][]byte) int {

	if len(maps) == 0 {
		return len(code)
	}

	if len(code) == 0 {
		return 0
	}
	next := code[0]

	return buildEdge(last, next, maps) + buildRemainder(next, code[1:], maps)
}

func part1(in Input) {
	fmt.Println("SOLUTION TO PART 1:", solve(in.codes, 2))
}

func part2(in Input) {
	fmt.Println("SOLUTION TO PART 2:", solve(in.codes, 25))
}

func solve(codes [][]byte, n int) int {

	mapDirpad := makeMap(POS_DIRECTIONAL, BUTTONS_DIRECTIONAL)

	mapNumpad := makeMap(POS_NUMERICAL, BUTTONS_NUMERICAL)

	// paths, found := mapNumpad[Edge{PAD_7, PAD_0}]
	// lib.MustBeTrue(found)
	// for _, p := range paths {
	// 	fmt.Println("===>", string(p))
	// }

	// lib.MustPressEnter()

	maps := []map[Edge][][]byte{mapNumpad}
	for range n {
		maps = append(maps, mapDirpad)
	}

	sol := 0
	for _, code := range codes {
		cost := buildRemainder(PAD_A, code, maps)
		fmt.Printf("CODE %s: %d\n", string(code), cost)

		numericPart := lib.MustToInt(strings.ReplaceAll(string(code), "A", ""))
		sol += cost * numericPart
	}

	return sol
}
