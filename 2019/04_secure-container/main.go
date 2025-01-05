package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	a, b int
}

func preprocess(data []byte) Input {
	str := strings.Trim(string(data), "\n")
	parts := strings.Split(str, "-")
	a := lib.MustToInt(parts[0])
	b := lib.MustToInt(parts[1])
	return Input{a: a, b: b}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	solve(input)
}

func to6Digits(n int) []int {
	out := make([]int, 6)
	for i := range 6 {
		out[i] = n % 10
		n /= 10
	}
	return out
}

func explore(as, bs []int, d, l int, hasDoubles bool, bot, top bool, s string, lastReps, maxReps int, hasPair bool) (int, int) {
	//fmt.Println(as, bs, d, l, double, bot, top, s)
	if d < 0 {
		if hasDoubles {
			//fmt.Println(as, bs, d, l, hasDoubles, bot, top, s, lastReps, maxReps, hasPair)
			if hasPair || lastReps == 2 {
				//fmt.Println(s, "x x")
				//lib.MustPressEnter()
				return 1, 1
			}
			//fmt.Println(s, "x o")
			//lib.MustPressEnter()
			return 1, 0
		} else {
			return 0, 0
		}
	}

	// d > 0

	a := as[d]
	b := bs[d]

	lb := 0
	ub := 9
	if l != -1 {
		lb = max(lb, l)
	}
	if bot {
		lb = max(lb, a)
	}
	if top {
		ub = min(ub, b)
	}
	//fmt.Println(lb, ub)

	sum1, sum2 := 0, 0
	for c := lb; c <= ub; c++ {
		lReps := 1
		hPair := hasPair
		if c == l {
			lReps = lastReps + 1
		} else if lastReps == 2 {
			hPair = true
		}
		mReps := max(maxReps, lReps)
		add1, add2 := explore(as, bs, d-1, c, hasDoubles || c == l, bot && c == a, top && c == b, fmt.Sprintf("%s%d", s, c), lReps, mReps, hPair)
		sum1 += add1
		sum2 += add2
	}

	return sum1, sum2
}

func solve(in Input) {
	sol1, sol2 := explore(to6Digits(in.a), to6Digits(in.b), 5, -1, false, true, true, "", 0, 0, false)
	fmt.Println("SOLUTION TO PART 1:", sol1)
	fmt.Println("SOLUTION TO PART 2:", sol2)
}
