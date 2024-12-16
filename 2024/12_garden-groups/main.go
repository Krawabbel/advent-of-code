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

	part1(input)

	part2(input)
}

func inBounds[T any](i, j int, mat [][]T) bool {
	return i >= 0 && i < len(mat) && j >= 0 && j < len(mat[i])
}

func classify(class [][]int, grid [][]byte, i, j, next_class int) int {

	if !inBounds(i, j, grid) {
		return next_class // out-of-bounds
	}

	if class[i][j] != 0 {
		return next_class // already classified
	}

	classifyRegion(class, grid, i, j, grid[i][j], next_class)

	return next_class + 1
}

func classifyRegion(class [][]int, grid [][]byte, i, j int, match_type byte, match_class int) {

	if !inBounds(i, j, grid) {
		return // out-of-bounds
	}

	if class[i][j] != 0 {
		return // already classified
	}

	if grid[i][j] != match_type {
		return // different region
	}

	class[i][j] = match_class
	classifyRegion(class, grid, i+1, j, match_type, match_class)
	classifyRegion(class, grid, i-1, j, match_type, match_class)
	classifyRegion(class, grid, i, j+1, match_type, match_class)
	classifyRegion(class, grid, i, j-1, match_type, match_class)
}

func disp[T any](format string, class [][]T) {
	for _, row := range class {
		for _, c := range row {
			fmt.Printf(format+" ", c)
		}
		fmt.Println()
	}
}

func evalRegion(class [][]int, i, j, c int) (bool, int, int) {

	if !inBounds(i, j, class) {
		return false, 0, 0 // out-of-bounds
	}

	if class[i][j] == -c {
		return true, 0, 0 // already covered
	} else if class[i][j] != c {
		return false, 0, 0 // not part of region
	}

	class[i][j] = -c

	area := 1
	perimeter := 0

	inRegion1, area1, perimeter1 := evalRegion(class, i+1, j, c)
	if !inRegion1 {
		perimeter++ // we have a border to this cell
	}
	area += area1
	perimeter += perimeter1

	inRegion2, area2, perimeter2 := evalRegion(class, i-1, j, c)
	if !inRegion2 {
		perimeter++ // we have a border to this cell
	}
	area += area2
	perimeter += perimeter2

	inRegion3, area3, perimeter3 := evalRegion(class, i, j+1, c)
	if !inRegion3 {
		perimeter++ // we have a border to this cell
	}
	area += area3
	perimeter += perimeter3

	inRegion4, area4, perimeter4 := evalRegion(class, i, j-1, c)
	if !inRegion4 {
		perimeter++ // we have a border to this cell
	}
	area += area4
	perimeter += perimeter4

	return true, area, perimeter
}

func part1(input Input) {

	class := make([][]int, len(input.grid))
	next_class := 1
	for i, row := range input.grid {
		class[i] = make([]int, len(row))
	}

	for i := range class {
		for j := range class[i] {
			next_class = classify(class, input.grid, i, j, next_class)
		}
	}

	max_class := next_class - 1

	// format := "%" + fmt.Sprint(len(fmt.Sprint(max_class))) + "d"
	// disp(format, class)

	sol := 0

class_loop:
	for k := range max_class {
		c := k + 1
		for i := range class {
			for j := range class[i] {
				if class[i][j] == c {
					_, area, perimeter := evalRegion(class, i, j, c)
					// fmt.Println(c, area, perimeter)
					sol += area * perimeter
					continue class_loop
				}
			}
		}
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func compare[T comparable](mat [][]T, i, j int, t T) bool {
	return inBounds(i, j, mat) && mat[i][j] == t
}

func part2(input Input) {

	class := make([][]int, len(input.grid))
	next_class := 1
	for i, row := range input.grid {
		class[i] = make([]int, len(row))
	}

	for i := range class {
		for j := range class[i] {
			next_class = classify(class, input.grid, i, j, next_class)
		}
	}

	max_class := next_class - 1

	// format := "%" + fmt.Sprint(len(fmt.Sprint(max_class))) + "d"
	// disp(format, class)

	top := make([][]bool, len(class))
	right := make([][]bool, len(class))
	bottom := make([][]bool, len(class))
	left := make([][]bool, len(class))
	for i, row := range class {
		top[i] = make([]bool, len(class[i]))
		right[i] = make([]bool, len(class[i]))
		bottom[i] = make([]bool, len(class[i]))
		left[i] = make([]bool, len(class[i]))
		for j, c := range row {
			top[i][j] = !compare(class, i-1, j, c)
			right[i][j] = !compare(class, i, j+1, c)
			bottom[i][j] = !compare(class, i+1, j, c)
			left[i][j] = !compare(class, i, j-1, c)
		}
	}

	// fmt.Println("TOP")
	// disp("%5v", top)

	// fmt.Println("RIGHT")
	// disp("%5v", right)

	// fmt.Println("BOTTOM")
	// disp("%5v", bottom)

	// fmt.Println("LEFT")
	// disp("%5v", left)

	area := make([]int, max_class)
	for _, row := range class {
		for _, c := range row {
			area[c-1]++
		}
	}

	fences := make([]int, max_class)
	for i, row := range class {
		for j, c := range row {
			fences[c-1] += processEdge(top, class, i, j, c, 0, +1)
			fences[c-1] += processEdge(bottom, class, i, j, c, 0, +1)
			fences[c-1] += processEdge(right, class, i, j, c, +1, 0)
			fences[c-1] += processEdge(left, class, i, j, c, +1, 0)
		}
	}

	sol := 0

	for i := range max_class {
		sol += area[i] * fences[i]
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}

func processEdge(isEdge [][]bool, class [][]int, i, j, c, dirI, dirJ int) int {
	found := false

	for ii, jj := i, j; inBounds(ii, jj, class) && class[ii][jj] == c && isEdge[ii][jj]; ii, jj = ii+dirI, jj+dirJ {
		isEdge[ii][jj] = false
		found = true
	}

	for ii, jj := i-dirI, j-dirJ; inBounds(ii, jj, class) && class[ii][jj] == c && isEdge[ii][jj]; ii, jj = ii-dirI, jj-dirJ {
		isEdge[ii][jj] = false
		found = true
	}

	if found {
		return 1
	}
	return 0
}
