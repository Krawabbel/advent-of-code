package main

import (
	"aoc2024/internal/lib"
	"bytes"
	"fmt"
	"os"
)

type Input struct {
	grid [][]byte
}

func preprocess(data []byte) Input {
	grid := bytes.Split(data, []byte{'\n'})

	return Input{grid: grid}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	run(input)
}

type int2D = lib.Vector[int]

var (
	DIR_NORTH = lib.Vec(-1, 0)
	DIR_SOUTH = lib.Vec(+1, 0)
	DIR_WEST  = lib.Vec(0, -1)
	DIR_EAST  = lib.Vec(0, +1)

	DIRS = []int2D{DIR_NORTH, DIR_EAST, DIR_SOUTH, DIR_WEST}
)

func turnLeft(idxDir int) int {
	return (idxDir - 1 + 4) % 4
}

func turnRight(idxDir int) int {
	return (idxDir + 1) % 4
}

func explore(grid [][]byte, scores [][][]int, score int, last int2D, idxDir int) {

	dir := DIRS[idxDir]
	next := lib.VecAdd(last, dir)
	i, j := next.Get(0), next.Get(1)

	// fmt.Println("score", score, "last", last, "dir", idxDir, "next", next)

	if !lib.InMatBounds(grid, i, j) {
		// fmt.Println("out-of-bounds")
		return
	}

	if grid[i][j] == '#' {
		// fmt.Println("blocked")
		return
	}

	prevScore := scores[i][j][idxDir]

	if -1 < prevScore && prevScore <= score {
		// fmt.Println(prevScore, "<", score, "worse")
		return
	}

	scores[i][j][idxDir] = score

	// visuScore(scores)
	// lib.MustPressEnter()

	explore(grid, scores, score+1, next, idxDir)
	explore(grid, scores, score+1001, next, turnLeft(idxDir))
	explore(grid, scores, score+1001, next, turnRight(idxDir))
}

func bestScore(scores []int) int {

	ret := -1
	for _, s := range scores {
		if ret < 0 || (-1 < s && s < ret) {
			ret = s
		}
	}
	return ret
}

func visuScore(scores [][][]int) {
	for _, row := range scores {
		for _, cell := range row {
			fmt.Print("[")
			for _, s := range cell {
				if s == -1 {
					fmt.Print("*")
				} else {
					fmt.Print(s)
				}
			}
			fmt.Print("] ")
		}
		fmt.Println()
	}
}

func run(in Input) {
	i, j := len(in.grid)-2, 1
	lib.MustBeEqual(in.grid[i][j], 'S')

	scores := make([][][]int, len(in.grid))
	for i := range scores {
		scores[i] = make([][]int, len(in.grid[i]))
		for j := range scores[i] {
			scores[i][j] = make([]int, 4)
			for k := range scores[i][j] {
				scores[i][j][k] = -1
			}
		}
	}

	explore(in.grid, scores, 0, lib.Vec(i, j-1), 1)

	ii, jj := 1, len(in.grid[2])-2
	lib.MustBeEqual(in.grid[ii][jj], 'E')

	sol1 := bestScore(scores[ii][jj])

	// visuScore(scores)

	fmt.Println("SOLUTION TO PART 1:", sol1)

	path := lib.Mat(len(in.grid), len(in.grid[0]), false)
	construct(in.grid, scores, 0, lib.Vec(i, j-1), 1, path)

	// visuPath(in.grid, path)

	sol2 := 0
	for _, row := range path {
		for _, cell := range row {
			if cell {
				sol2++
			}
		}
	}
	fmt.Println("SOLUTION TO PART 2:", sol2)
}

func construct(grid [][]byte, scores [][][]int, score int, last int2D, idxDir int, path [][]bool) bool {

	dir := DIRS[idxDir]
	next := lib.VecAdd(last, dir)
	i, j := next.Get(0), next.Get(1)

	if !lib.InMatBounds(grid, i, j) {
		return false
	}

	if grid[i][j] == '#' {
		return false
	}

	if score != scores[i][j][idxDir] {
		return false
	}

	if i == 1 && j == len(grid[2])-2 && score == bestScore(scores[i][j]) {
		path[i][j] = true
		return true
	}

	optimalStraight := construct(grid, scores, score+1, next, idxDir, path)

	optimalLeft := construct(grid, scores, score+1001, next, turnLeft(idxDir), path)

	optimalRight := construct(grid, scores, score+1001, next, turnRight(idxDir), path)

	optimal := optimalStraight || optimalLeft || optimalRight

	if optimal {
		path[i][j] = true
	}

	return optimal
}

func visuPath(grid [][]byte, path [][]bool) {
	for i, row := range grid {
		for j, cell := range row {
			switch cell {
			case '.', 'S', 'E':
				if path[i][j] {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			case '#':
				fmt.Print(string(cell))
			default:
				lib.Panicf("unexpected cell %s", string(cell))
			}
		}
		fmt.Println()
	}
}
