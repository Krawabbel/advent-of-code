package main

import (
	"aoc/internal/lib"
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

	solve(input)

}

func solve(in Input) {

	idxs := make([][]int, len(in.grid))
	idxLU := []int{0}

	for i, row := range in.grid {
		idxs[i] = make([]int, len(row))
		for j := 0; j < len(row); j++ {
			b := row[j]
			if b >= '0' && b <= '9' {
				start := j
				end := start
				for end < len(row) && row[end] >= '0' && row[end] <= '9' {
					end++
				}
				val := lib.MustToInt(string(row[start:end]))
				idxLU = append(idxLU, val)
				for k := start; k < end; k++ {
					idxs[i][k] = len(idxLU) - 1
				}
				j = end
			}
		}
	}

	// lib.PrintMat(grid)

	sol2 := 0

	idxsToAdd := lib.MakeSet[int]()
	for i, row := range in.grid {
		for j, b := range row {
			if b != '.' && (b < '0' || b > '9') {
				for _, m := range []int{-1, 0, 1} {
					for _, n := range []int{-1, 0, 1} {
						if m == 0 && n == 0 {
							continue
						}
						check(idxs, idxsToAdd, i+m, j+n)
					}
				}

			}
			if b == '*' {
				adjIdxs := lib.MakeSet[int]()
				for _, m := range []int{-1, 0, 1} {
					for _, n := range []int{-1, 0, 1} {
						if m == 0 && n == 0 {
							continue
						}
						check(idxs, adjIdxs, i+m, j+n)
					}
				}

				adjIdxs.Delete(0)

				if adjIdxs.Size() == 2 {
					prod := 1
					for idx := range adjIdxs.Elements {
						prod *= idxLU[idx]
					}
					sol2 += prod
				}
			}
		}
	}

	sol1 := 0
	for idx := range idxsToAdd.Elements {
		sol1 += idxLU[idx]
	}

	fmt.Println("SOLUTION TO PART 1:", sol1)
	fmt.Println("SOLUTION TO PART 2:", sol2)
}

func check(idxs [][]int, idxsToAdd lib.Set[int], i, j int) {
	if lib.InMatBounds(idxs, i, j) {
		idxsToAdd.Insert(idxs[i][j])
	}
}
