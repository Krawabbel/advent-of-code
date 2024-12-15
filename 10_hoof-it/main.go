package main

import (
	"aoc2024/internal/lib"
	"bytes"
	"fmt"
	"os"
)

type Point struct {
	i, j int
}

type Input struct {
	grid  [][]int
	heads []Point
	peaks []Point
}

func preprocess(data []byte) Input {
	rows := bytes.Split(data, []byte{'\n'})

	grid := make([][]int, len(rows))
	heads := make([]Point, 0)
	peaks := make([]Point, 0)
	for i, row := range rows {
		grid[i] = make([]int, len(row))
		for j, b := range row {
			if b == '.' {
				grid[i][j] = -100
			} else {
				grid[i][j] = int(b - '0')
				if b == '0' {
					heads = append(heads, Point{i: i, j: j})
				}
				if b == '9' {
					peaks = append(peaks, Point{i: i, j: j})
				}
			}
		}
	}

	return Input{grid: grid, heads: heads, peaks: peaks}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func inBounds(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}

func explore1(grid [][]int, visited [][]bool, i, j, ii, jj int) {

	if !inBounds(grid, ii, jj) {
		return
	}

	if visited[ii][jj] {
		return
	}

	if grid[i][j]+1 != grid[ii][jj] {
		return
	}

	visit1(grid, visited, ii, jj)
}

func visit1(grid [][]int, visited [][]bool, i, j int) {

	visited[i][j] = true
	explore1(grid, visited, i, j, i+1, j)
	explore1(grid, visited, i, j, i-1, j)
	explore1(grid, visited, i, j, i, j+1)
	explore1(grid, visited, i, j, i, j-1)

}

func part1(input Input) {

	sol := 0

	for _, head := range input.heads {
		visited := make([][]bool, len(input.grid))
		for i, row := range input.grid {
			visited[i] = make([]bool, len(row))
		}

		visit1(input.grid, visited, head.i, head.j)

		for _, peak := range input.peaks {
			if visited[peak.i][peak.j] {
				sol++
			}
		}

	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func explore2(grid [][]int, i, j, ii, jj int) int {

	if !inBounds(grid, ii, jj) {
		return 0
	}

	if grid[i][j]+1 != grid[ii][jj] {
		return 0
	}

	if grid[ii][jj] == 9 {
		return 1
	}

	return visit2(grid, ii, jj)

}

func visit2(grid [][]int, i, j int) int {
	return explore2(grid, i, j, i+1, j) +
		explore2(grid, i, j, i-1, j) +
		explore2(grid, i, j, i, j+1) +
		explore2(grid, i, j, i, j-1)
}

func part2(input Input) {
	sol := 0

	for _, head := range input.heads {
		sol += visit2(input.grid, head.i, head.j)
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}
