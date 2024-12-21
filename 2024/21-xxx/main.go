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

var POS_NUMERIC = map[byte]Point{
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

var dirs = map[byte]Point{
	PAD_UP:    {-1, 0},
	PAD_DOWN:  {+1, 0},
	PAD_LEFT:  {0, -1},
	PAD_RIGHT: {0, +1},
}

var POS_DIRECTIONAL = map[byte]Point{
	PAD_UP:    {0, 1},
	PAD_A:     {0, 2},
	PAD_LEFT:  {1, 0},
	PAD_DOWN:  {1, 1},
	PAD_RIGHT: {1, 2},
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

var (
	ErrInvalidPosition   = fmt.Errorf("invalid position")
	ErrBadDirection      = fmt.Errorf("bad direction")
	ErrExplorationFailed = fmt.Errorf("exploration failed")
)

// **********************************************************************

func solve(begin, button byte, pads []map[byte]Point) (string, error) {
	if len(pads) == 0 {
		return string(button), nil
	}

	pad := pads[0]
	p0, isValidPadPosition := pad[begin]
	if !isValidPadPosition {
		return "", ErrInvalidPosition
	}

	p1, found := pad[button]
	lib.MustBeTrue(found)

	dist := PointSub(p1, p0)
	moves := calcMoves(dist)

	return explore(PAD_A, moves, pads[1:])
}

func explore(begin byte, moves map[byte]int, pads []map[byte]Point) (string, error) {

	if len(moves) == 0 {
		return solve(begin, PAD_A, pads)
	}

	seq := ""
	best := -1

	for b := range dirs {

		s, err := move(begin, b, moves, pads)
		if err != nil {
			// fmt.Println("DEBUG: ERR: ", err)
			continue
		}
		// fmt.Println("DEBUG: NO ERR")

		c := len(s)
		isFirstValidSolution := best < 0
		isBetterSolution := c < best
		if isFirstValidSolution || isBetterSolution {
			best = c
			seq = s
		}
	}

	if best < 0 {
		return "", ErrExplorationFailed
	}

	return seq, nil
}

func move(begin, button byte, moves map[byte]int, pads []map[byte]Point) (string, error) {

	steps, found := moves[button]
	if !found || steps == 0 {
		return "", ErrBadDirection
	}

	seq := ""

	state := begin
	for range steps {
		s, err := solve(state, button, pads)
		if err != nil {
			return "", err
		}
		seq += s
		state = button
	}

	delete(moves, button)
	s, err := explore(state, moves, pads)
	moves[button] = steps

	return seq + s, err

}

func part1(in Input) {

	sol := 0

	pads := []map[byte]Point{
		POS_NUMERIC,
		POS_DIRECTIONAL,
		POS_DIRECTIONAL,
	}

	for _, code := range in.codes {

		fmt.Printf("%s: ", string(code))

		seq := ""
		state := byte(PAD_A)

		for _, button := range code {
			s, err := solve(state, button, pads)
			lib.Must(err)
			seq += s
			state = button
		}

		numericPart := lib.MustToInt(strings.ReplaceAll(string(code), "A", ""))
		cost := len(seq)
		sol += cost * numericPart

		fmt.Printf("%s (%d * %d)\n", seq, cost, numericPart)
		// lib.MustPressEnter()
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
