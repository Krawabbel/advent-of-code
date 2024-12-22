package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
)

type Input struct {
	start []int
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))
	return Input{start: lib.MustToInts(lines)}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func mix(u, v int) int {
	return u ^ v
}

func prune(u int) int {
	return u & 0xFFFFFF
}

func first(secret int) int {
	return secret << 6
}

func second(secret int) int {
	return secret >> 5
}

func third(secret int) int {
	return secret << 11
}

func next(last int) int {
	a := prune(mix(last, first(last)))
	b := prune(mix(a, second(a)))
	c := prune(mix(b, third(b)))
	return c
}

func price(secret int) int {
	return secret % 10
}

func part1(in Input) {
	// fmt.Printf("%+v\n", in)
	sol := 0
	for _, start := range in.start {
		secret := start
		for range 2000 {
			secret = next(secret)
		}
		// fmt.Printf("%d: %d\n", start, secret)
		sol += secret
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

type Sequence struct {
	changes [4]int
}

func NewSequence(buf []int) Sequence {
	lib.MustBeEqual(len(buf), 4)
	seq := Sequence{}
	copy(seq.changes[:], buf)
	return seq
}

func part2(in Input) {

	// long-term memory
	ltm := make(map[Sequence]int)

	for _, start := range in.start {
		secret := start
		buf := make([]int, 4)

		last := price(secret)
		// fmt.Printf("%d: %d\n", secret, last)
		for i := range 4 {
			secret = next(secret)
			next := price(secret)
			diff := next - last
			buf[i] = diff
			last = next
			// fmt.Printf("%d: %d (%d)\n", secret, last, diff)
		}

		// short-term memory
		stm := make(map[Sequence]int)

		stm[NewSequence(buf)] = last

		// fmt.Println(memo)

		for range 2000 - 4 {
			secret = next(secret)
			next := price(secret)
			diff := next - last
			buf = append(buf[1:], diff)
			seq := NewSequence(buf)
			if _, exists := stm[seq]; !exists {
				stm[seq] = next
			}
			last = next
		}

		for key, val := range stm {
			ltm[key] += val
		}
	}

	sol := -1
	// seq := Sequence{}
	for _, val := range ltm {
		if val > sol {
			sol = val
			// seq = key
		}
	}

	// fmt.Println(seq, sol)

	fmt.Println("SOLUTION TO PART 2:", sol)
}
