package main

import (
	"aoc2024/internal/lib"
	"bytes"
	"fmt"
	"os"
)

type Vector struct {
	i, j int
}

func Vec(i, j int) Vector {
	return Vector{i: i, j: j}
}

func VecAdd(u, v Vector) Vector {
	return Vector{i: u.i + v.i, j: u.j + v.j}
}

func VecMultElementwise(u Vector, v Vector) Vector {
	return Vector{i: u.i * v.i, j: u.j * v.j}
}

type Input struct {
	walls, boxes  map[Vector]struct{}
	robot         Vector
	moves         []byte
	width, height int
}

func preprocess(data []byte) Input {
	parts := bytes.Split(data, []byte{'\n', '\n'})

	walls := make(map[Vector]struct{})
	boxes := make(map[Vector]struct{})
	var robot Vector

	grid := bytes.Split(parts[0], []byte{'\n'})
	for i, row := range grid {
		for j, val := range row {
			p := Vector{i: i, j: j}
			switch val {
			case '#':
				walls[p] = struct{}{}
			case 'O':
				boxes[p] = struct{}{}
			case '@':
				robot = p
			case '.':
				// ignore
			default:
				lib.Panicf("unexpected char %v", val)
			}
		}
	}
	moves := bytes.Replace(parts[1], []byte{'\n'}, []byte{}, -1)

	return Input{
		walls:  walls,
		boxes:  boxes,
		robot:  robot,
		moves:  moves,
		width:  len(grid),
		height: len(grid[0]),
	}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func disp1(walls, boxes map[Vector]struct{}, robot Vector, width, height int) {
	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte{'.'}, width)
	}

	for wall, _ := range walls {
		grid[wall.i][wall.j] = '#'
	}

	for box, _ := range boxes {
		grid[box.i][box.j] = 'O'
	}

	grid[robot.i][robot.j] = '@'

	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%s", string(val))
		}
		fmt.Println()
	}
}

func disp2(walls, boxes map[Vector]struct{}, robot Vector, width, height int) {
	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte{'.'}, width)
	}

	for wall, _ := range walls {
		grid[wall.i][wall.j] = '#'
	}

	for box, _ := range boxes {
		grid[box.i][box.j] = '['
		grid[box.i][box.j+1] = ']'
	}

	grid[robot.i][robot.j] = '@'

	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%s", string(val))
		}
		fmt.Println()
	}
}

func direction(b byte) Vector {
	switch b {
	case '^':
		return Vec(-1, 0)
	case 'v':
		return Vec(+1, 0)
	case '>':
		return Vec(0, +1)
	case '<':
		return Vec(0, -1)
	default:
		lib.Panicf("unexpected direction %v", string(b))
	}

	// unreachable
	return Vec(0, 0)
}

func tryPush(p, dir Vector, walls, boxes map[Vector]struct{}) bool {
	next := VecAdd(p, dir)
	if _, isWall := walls[next]; isWall {
		return false
	}

	if _, isBox := boxes[next]; isBox {
		if !tryPush(next, dir, walls, boxes) {
			return false
		}
		delete(boxes, next)
		boxes[VecAdd(next, dir)] = struct{}{}
		return true
	}

	return true
}

func part1(input Input) {
	walls := lib.CloneMap(input.walls)
	boxes := lib.CloneMap(input.boxes)
	robot := input.robot

	// fmt.Println("Initial State")
	// disp(walls, boxes, robot, input.width, input.height)
	// lib.MustPressEnter()

	for _, mv := range input.moves {

		dir := direction(mv)
		if tryPush(robot, dir, walls, boxes) {
			robot = VecAdd(robot, dir)
		}

		// fmt.Printf("Move %s:\n", string(mv))
		// disp(walls, boxes, robot, width, height)
		// lib.MustPressEnter()
	}

	sol := 0
	for box := range boxes {
		sol += 100*box.i + box.j
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func checkPush(p, dir Vector, walls, boxes map[Vector]struct{}, moved map[Vector]struct{}) bool {
	next := VecAdd(p, dir)
	if _, isWall := walls[next]; isWall {
		return false
	}

	_, isNextMoved := moved[next]
	_, isLeft := boxes[next]
	if !isNextMoved && isLeft {
		moved[next] = struct{}{}

		right := VecAdd(next, Vec(0, +1))
		if !checkPush(next, dir, walls, boxes, moved) || !checkPush(right, dir, walls, boxes, moved) {
			return false
		}
	}

	left := VecAdd(next, Vec(0, -1))
	_, isLeftMoved := moved[left]
	_, isRight := boxes[left]
	if !isLeftMoved && isRight {
		moved[left] = struct{}{}
		if !checkPush(left, dir, walls, boxes, moved) || !checkPush(next, dir, walls, boxes, moved) {
			return false
		}
	}

	return true
}

func part2(input Input) {

	mod := Vec(1, 2)

	walls := make(map[Vector]struct{}, len(input.walls))
	for w := range input.walls {
		left := VecMultElementwise(mod, w)
		walls[left] = struct{}{}

		right := VecAdd(left, Vec(0, 1))
		walls[right] = struct{}{}
	}

	boxes := make(map[Vector]struct{}, len(input.boxes))
	for b := range input.boxes {
		left := VecMultElementwise(mod, b)
		boxes[left] = struct{}{}
	}

	robot := VecMultElementwise(mod, input.robot)

	// width, height := 2*input.width, input.height

	// fmt.Println("Initial State")
	// disp2(walls, boxes, robot, width, height)
	// lib.MustPressEnter()

	for _, mv := range input.moves {

		dir := direction(mv)
		moved := make(map[Vector]struct{})
		if checkPush(robot, dir, walls, boxes, moved) {
			robot = VecAdd(robot, dir)

			for box := range moved {
				_, found := boxes[box]
				lib.MustBeTrue(found)
				delete(boxes, box)
			}

			for box := range moved {
				boxes[VecAdd(box, dir)] = struct{}{}
			}
		}

		// fmt.Printf("Move %s:\n", string(mv))
		// disp2(walls, boxes, robot, width, height)
		// lib.MustPressEnter()

	}

	sol := 0
	for box := range boxes {
		sol += 100*box.i + box.j
	}
	fmt.Println("SOLUTION TO PART 2:", sol)
}
