package main

import (
	"aoc/internal/lib"
	"fmt"
	"math"
	"os"
	"strings"
)

type Input struct {
	ts, ds []int
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))
	ts := lib.SplitToInts(strings.TrimPrefix(lines[0], "Time:"), " ")
	ds := lib.SplitToInts(strings.TrimPrefix(lines[1], "Distance:"), " ")
	return Input{ts: ts, ds: ds}
}

func main() {
	path := os.Args[1]
	blob, err := os.ReadFile(path)
	lib.Must(err)
	input := preprocess(blob)
	part1(input)
	part2(input)
}

func sim(T, t int) int {
	return (T - t) * t
}

func part1(in Input) {
	sol := 1
	for i := range in.ts {
		n := solve(in.ts[i], in.ds[i])
		sol *= n
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func solve(T int, D int) int {
	t := float64(T)
	d := float64(D)
	l := t/2 - math.Sqrt(math.Pow(t, 2)-4*d)/2
	r := t/2 + math.Sqrt(math.Pow(t, 2)-4*d)/2

	L := int(l)
	for sim(T, L) <= D {
		L++
	}

	R := int(r) + 1
	for sim(T, R) <= D {
		R--
	}

	return R - L + 1
}

func part2(in Input) {
	tstr := ""
	dstr := ""
	for i := range in.ts {
		tstr += fmt.Sprint(in.ts[i])
		dstr += fmt.Sprint(in.ds[i])
	}
	sol := solve(lib.MustToInt(tstr), lib.MustToInt(dstr))
	fmt.Println("SOLUTION TO PART 2:", sol)
}
