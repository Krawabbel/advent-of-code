package main

import (
	"aoc2024/internal/lib"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Robot struct {
	px, py, vx, vy int
}

func (r *Robot) predict(w, h int) {
	r.px = (r.px + r.vx + w) % w
	r.py = (r.py + r.vy + h) % h
}

func disp(w, h int, robots ...Robot) {
	lines := make([][]byte, h)
	for i := range lines {
		lines[i] = []byte(strings.Repeat(" ", w))
	}
	for _, r := range robots {
		// fmt.Printf("%+v\n", r)
		lines[r.py][r.px] = '*'
	}
	for _, l := range lines {
		fmt.Println(string(l))
	}
}

func isSet(grid [][]bool, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j]
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func findClosestRobot(grid [][]bool, i, j int) int {
	for d := 1; ; d++ {
		for ii := -d; ii <= d; ii++ {
			jj := d - abs(ii)
			if isSet(grid, i+ii, j+jj) || isSet(grid, i+ii, j-jj) {
				return d
			}

		}
	}
}

func sumDistances(grid [][]bool) int {
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if isSet(grid, i, j) {
				sum += findClosestRobot(grid, i, j)
			}
		}
	}
	return sum
}

func makeGrid(rs []Robot, w, h int) [][]bool {
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}
	for _, r := range rs {
		grid[r.py][r.px] = true
	}
	return grid
}

type Input struct {
	robots        []Robot
	width, height int
}

var reRobot = regexp.MustCompile(`p=([|-]?[0-9]*),([|-]?[0-9]*) v=([|-]?[0-9]*),([|-]?[0-9]*)`)

func preprocess(data []byte, width, height int) Input {
	lines := strings.Split(string(data), "\n")
	robots := make([]Robot, len(lines))
	for i, l := range lines {
		// fmt.Println(l)
		m := reRobot.FindStringSubmatch(l)
		vs := lib.MustToInts(m[1:])
		r := Robot{px: vs[0], py: vs[1], vx: vs[2], vy: vs[3]}
		l2 := fmt.Sprintf("p=%d,%d v=%d,%d", r.px, r.py, r.vx, r.vy)
		lib.MustBeTrue(l2 == l)
		robots[i] = r
	}

	return Input{
		robots: robots,
		width:  width,
		height: height,
	}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	width := lib.MustToInt(os.Args[2])
	height := lib.MustToInt(os.Args[3])

	input := preprocess(blob, width, height)

	part1(input)

	part2(input)
}

func count(x0, x1, y0, y1 int, rs []Robot) int {
	n := 0
	for _, r := range rs {
		if x0 <= r.px && r.px < x1 && y0 <= r.py && r.py < y1 {
			n++
		}
	}
	return n
}

func part1(input Input) {

	rs := make([]Robot, len(input.robots))
	copy(rs, input.robots)

	for range 100 {
		// fmt.Printf("After %d seconds...\n", i+1)
		for j := range rs {
			rs[j].predict(input.width, input.height)
		}
	}

	// disp(input.width, input.height, rs...)

	sol := count(0, input.width/2, 0, input.height/2, rs) *
		count(0, input.width/2, input.height/2+1, input.height, rs) *
		count(input.width/2+1, input.width, 0, input.height/2, rs) *
		count(input.width/2+1, input.width, input.height/2+1, input.height, rs)

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(input Input) {
	rs := make([]Robot, len(input.robots))
	copy(rs, input.robots)

	w, h := input.width, input.height

	best_score := -1

	for i := 0; ; i++ {
		for j := range rs {
			rs[j].predict(w, h)
		}

		grid := makeGrid(rs, w, h)

		score := sumDistances(grid)

		isCandidate := score < best_score || best_score < 0

		if isCandidate {

			disp(w, h, rs...)
			fmt.Printf("Candidate found after %d seconds...\n", i+1)

			fmt.Printf("Score: %d (prev. best: %d)\n", score, best_score)
			best_score = score

			fmt.Println("Press 'Enter' to continue...")
			_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
			lib.Must(err)
		}
	}

}
