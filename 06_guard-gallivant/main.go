package main

import (
	"aoc2024/internal/lib"
	"fmt"
	"os"
	"strings"
)

type cell int

const (
	CELL_FREE    = cell(0)
	CELL_BLOCKED = cell(1)
)

func tile(r rune) cell {
	switch r {
	case '#':
		return CELL_BLOCKED
	case '.', '^':
		return CELL_FREE
	}
	panic("unexpected tile")
}

func disp(grid [][]cell, visited [][]bool, iPos, jPos, dir int) {
	for i, row := range grid {
		for j, c := range row {
			symbol := "."
			if c == CELL_BLOCKED {
				symbol = "#"
			}
			if visited != nil && visited[i][j] {
				symbol = "X"
			}
			if i == iPos && j == jPos {
				symbol = fmt.Sprint(dir)
			}
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := string(blob)

	lines := strings.Split(input, "\n")

	grid := [][]cell{}
	start_i := -1
	start_j := -1
	for i, l := range lines {

		row := []cell{}
		for j, r := range l {
			row = append(row, tile(r))
			if r == '^' {
				start_i = i
				start_j = j
			}

		}
		grid = append(grid, row)
	}

	part1(grid, start_i, start_j)

	part2(grid, start_i, start_j)
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func predict(i, j, dir int) (int, int) {
	switch dir {
	case UP:
		return (i - 1), j
	case RIGHT:
		return i, (j + 1)
	case DOWN:
		return (i + 1), j
	case LEFT:
		return i, (j - 1)
	}
	panic("unexpected dir")
}

func step(grid [][]cell, i, j, dir int) (int, int, int) {
	ii, jj := predict(i, j, dir)
	for inBounds(grid, ii, jj) && grid[ii][jj] == CELL_BLOCKED {
		dir = (dir + 1) % 4
		ii, jj = predict(i, j, dir)
	}
	return ii, jj, dir
}

func inBounds(grid [][]cell, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}

func part1(grid [][]cell, i, j int) {

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	dir := UP

	// disp(grid, visited, i, j, dir)

	for inBounds(grid, i, j) {
		visited[i][j] = true

		i, j, dir = step(grid, i, j, dir)

		// disp(grid, visited, i, j, dir)
		// time.Sleep(time.Second / 10)
	}

	sol := 0
	for _, row := range visited {
		for _, v := range row {
			if v {
				sol++
			}
		}
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func hash(i, j, dir int) string {
	return fmt.Sprintf("%d+%d+%d", i, j, dir)
}

func part2(grid [][]cell, iStart, jStart int) {

	sol := 0

	for iObst := range grid {
		for jObst := range grid[iObst] {

			if iObst == iStart && jObst == jStart {
				continue
			}

			// fmt.Println(((iObst*len(grid) + jObst) * 10000) / (len(grid) * len(grid[0])))

			prev := map[string]struct{}{}

			newGrid := make([][]cell, len(grid))
			for i, row := range grid {
				newGrid[i] = make([]cell, len(row))
				copy(newGrid[i], row)
			}
			newGrid[iObst][jObst] = CELL_BLOCKED

			i := iStart
			j := jStart
			dir := UP

			for inBounds(newGrid, i, j) {
				h := hash(i, j, dir)
				_, cycle := prev[h]
				if cycle {
					sol++

					// disp(newGrid, nil, i, j, dir)
					break
				}
				prev[h] = struct{}{}

				i, j, dir = step(newGrid, i, j, dir)

			}
		}
	}
	fmt.Println("SOLUTION TO PART 2:", sol)
}
