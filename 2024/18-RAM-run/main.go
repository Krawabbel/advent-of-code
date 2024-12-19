package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	i, j int
}

func (p *Point) top() Point    { return Point{i: p.i - 1, j: p.j} }
func (p *Point) bottom() Point { return Point{i: p.i + 1, j: p.j} }
func (p *Point) right() Point  { return Point{i: p.i, j: p.j + 1} }
func (p *Point) left() Point   { return Point{i: p.i, j: p.j - 1} }

type Input struct {
	corrupted []Point
	target    Point
	cycles    int
}

func preprocess(data []byte, iTarget, jTarget, cycles int) Input {
	lines := lib.SplitLines(string(data))
	corrupted := make([]Point, len(lines))
	for i, l := range lines {
		coords := strings.Split(l, ",")
		x := lib.MustToInt(coords[0])
		y := lib.MustToInt(coords[1])
		corrupted[i] = Point{i: x, j: y}
	}

	return Input{corrupted: corrupted, target: Point{i: jTarget, j: iTarget}, cycles: cycles}
}

func main() {
	path := os.Args[1]
	width := lib.MustToInt(os.Args[2])
	height := lib.MustToInt(os.Args[3])
	cycles := lib.MustToInt(os.Args[4])

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob, width, height, cycles)

	part1(input)

	part2(input)
}

func inBounds(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func solve(grid map[Point]int, corrupted lib.Set[Point], pt Point, steps int, target Point) {

	if !inBounds(pt.i, pt.j, target.i+1, target.j+1) || corrupted.Contains(pt) {
		return
	}

	if prev, visited := grid[pt]; !visited || prev > steps {
		grid[pt] = steps
	} else if steps >= prev {
		return
	}

	if pt == target {
		return
	}

	solve(grid, corrupted, pt.top(), steps+1, target)
	solve(grid, corrupted, pt.bottom(), steps+1, target)
	solve(grid, corrupted, pt.left(), steps+1, target)
	solve(grid, corrupted, pt.right(), steps+1, target)
}

func part1(in Input) {
	corrupted := lib.MakeSet(in.corrupted[:in.cycles]...)
	grid := make(map[Point]int)

	solve(grid, corrupted, Point{i: 0, j: 0}, 0, in.target)

	sol, found := grid[in.target]
	lib.MustBeTrue(found)
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func isFeasible(in Input, cycles int) bool {
	corrupted := lib.MakeSet(in.corrupted[:cycles]...)
	grid := make(map[Point]int)

	solve(grid, corrupted, Point{i: 0, j: 0}, 0, in.target)

	_, feasible := grid[in.target]
	return feasible
}

func part2(in Input) {

	left := 0
	lib.MustBeTrue(isFeasible(in, left))

	right := len(in.corrupted) - 1
	lib.MustBeFalse(isFeasible(in, right))

	for left+1 < right {

		next := (left + right) / 2
		if isFeasible(in, next) {
			left = next
		} else {
			right = next
		}

		// fmt.Println(left, right)
	}

	lib.MustBeTrue(isFeasible(in, left))
	lib.MustBeFalse(isFeasible(in, right))

	last := in.corrupted[left]
	sol := fmt.Sprintf("%d,%d", last.i, last.j)

	fmt.Println("SOLUTION TO PART 2:", sol)
}
