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

var vec = map[byte]Point{
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

type Controller interface {
	push(byte) int
}

type Robot struct {
	pos  Point
	ctrl Controller
	pad  map[byte]Point
	id   int
}

func NewRobot(id int, ctrl Controller, pad map[byte]Point) *Robot {
	return &Robot{
		id:   id,
		pos:  pad[PAD_A],
		ctrl: ctrl,
		pad:  pad,
	}
}

func isValidPosition(pad map[byte]Point, pos Point) bool {
	for _, ref := range pad {
		if ref == pos {
			return true
		}
	}
	return false
}

func (r *Robot) move(dir byte) int {
	next := PointAdd(r.pos, vec[dir])
	if isValidPosition(r.pad, next) {
		r.pos = next
	}
	return r.ctrl.push(dir)
}

var (
	updown    = []byte{PAD_UP, PAD_DOWN}
	leftright = []byte{PAD_LEFT, PAD_RIGHT}
)

func (r *Robot) moveTo(target Point) int {
	dist := PointSub(target, r.pos)
	moves := calcMoves(dist)

	// fmt.Println("\nMOVES", r.pos, dist, target, moves)

	dirorder := append(updown, leftright...)
	if r.pos.j == 0 {
		dirorder = append(leftright, updown...)
	}

	sum := 0
	for _, mvdir := range dirorder {
		mvcount := moves[mvdir]
		for range mvcount {
			sum += r.move(mvdir)
		}
	}
	return sum
}

func (r *Robot) press() int {
	return r.ctrl.push(PAD_A)
}

func (r *Robot) push(btn byte) int {
	val := r.moveTo(r.pad[btn]) + r.press()
	// fmt.Printf(" Robot %d pushes %s\n", r.id, string(btn))
	// fmt.Print(string(btn))
	return val
}

func calcMoves(d Point) map[byte]int {
	mv := make(map[byte]int)

	if d.i < 0 {
		mv[PAD_UP] -= d.i
	} else if d.i == 0 {
		// do nothing
	} else {
		mv[PAD_DOWN] += d.i
	}

	if d.j > 0 {
		mv[PAD_RIGHT] += d.j
	} else if d.j == 0 {
		// do nothing
	} else {
		mv[PAD_LEFT] -= d.j
	}

	return mv
}

type Operator struct{}

func (o *Operator) push(btn byte) int {
	fmt.Print(string(btn))
	return 1
}

func part1(in Input) {

	sol := 0

	for _, code := range in.codes {

		fmt.Printf("%s: ", string(code))

		var ctrl Controller = new(Operator)       // you
		ctrl = NewRobot(2, ctrl, POS_DIRECTIONAL) // -40 degrees
		ctrl = NewRobot(1, ctrl, POS_DIRECTIONAL) // high levels of radiation
		ctrl = NewRobot(0, ctrl, POS_NUMERIC)     // depressurized

		sum := 0
		for _, btn := range code {
			sum += ctrl.push(btn)
			// fmt.Println(" =>", string(btn))
			// fmt.Println(POS_NUMERIC[btn], r0.pos)
			// fmt.Print(" ")
		}

		tmp := strings.ReplaceAll(string(code), "A", "")
		fac := lib.MustToInt(tmp)
		sol += sum * fac

		fmt.Printf("\n %d * %d\n", sum, fac)
		// lib.MustPressEnter()
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
