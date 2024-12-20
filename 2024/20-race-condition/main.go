package main

import (
	"aoc/internal/lib"
	"bytes"
	"fmt"
	"os"
	"slices"
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

	solve(input)
}

func build(grid [][]byte, score [][]int, i, j, s int) {
	if !lib.InMatBounds(grid, i, j) {
		return
	}

	if grid[i][j] == '#' {
		return
	}

	if 0 <= score[i][j] && score[i][j] <= s {
		return
	}
	score[i][j] = s

	if grid[i][j] == 'S' {
		return
	}

	build(grid, score, i+1, j, s+1)
	build(grid, score, i-1, j, s+1)
	build(grid, score, i, j+1, s+1)
	build(grid, score, i, j-1, s+1)
}

func getScore(score [][]int, i, j int) (int, bool) {
	if !lib.InMatBounds(score, i, j) || score[i][j] < 0 {
		return -1, false
	}
	return score[i][j], true
}

func check(score [][]int, freq map[int]int, i0, j0, cheat_rule int) {

	s0, valid := getScore(score, i0, j0)
	if !valid {
		return
	}

	for l := 2; l <= cheat_rule; l++ {

		for di := 0; di <= l; di++ {

			dj := lib.Abs(l) - lib.Abs(di)

			signsi := []int{-1, +1}
			if di == 0 {
				signsi = []int{1}
			}

			for _, sdi := range signsi {

				signsj := []int{-1, +1}
				if dj == 0 {
					signsj = []int{1}
				}

				for _, sdj := range signsj {

					i1 := i0 + sdi*di
					j1 := j0 + sdj*dj
					if s1, valid := getScore(score, i1, j1); valid {
						save := s1 - s0 - l
						if save > 0 {
							// fmt.Println(i0, j0, s0, i1, j1, s1, save)
							freq[save]++
						}
					}
				}
			}
		}
	}
}

func explore(score [][]int, cheat_rule int) int {

	freq := make(map[int]int)
	for i, row := range score {
		for j := range row {
			check(score, freq, i, j, cheat_rule)
		}
	}

	saves := make([]int, 0, len(freq))
	for s := range freq {
		idx, _ := slices.BinarySearch(saves, s)
		saves = slices.Insert(saves, idx, s)
	}

	sol := 0
	for _, s := range saves {
		count := freq[s]
		// fmt.Printf(">> %d cheat(s) that save %d picoseconds.\n", count, s)
		if s >= 100 {
			sol += count
		}
	}

	return sol
}

func solve(in Input) {

	score := make([][]int, len(in.grid))
	iend, jend := -1, -1
	for i, row := range in.grid {
		score[i] = make([]int, len(in.grid[0]))
		for j, cell := range row {
			score[i][j] = -1
			if cell == 'E' {
				iend, jend = i, j
			}
		}
	}

	build(in.grid, score, iend, jend, 0)

	fmt.Println("SOLUTION TO PART 1:", explore(score, 2))

	fmt.Println("SOLUTION TO PART 2:", explore(score, 20))
}
