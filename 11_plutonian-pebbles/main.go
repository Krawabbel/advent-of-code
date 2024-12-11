package main

import (
	"aoc2024/internal/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	stones []int
}

func preprocess(data []byte) Input {
	strs := strings.Split(string(data), " ")
	stones := make([]int, len(strs))
	for i, s := range strs {
		stone, err := strconv.Atoi(s)
		util.Must(err)
		stones[i] = stone
	}
	return Input{stones: stones}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	util.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func blink1(orig []int, n int) int {
	curr := make([]int, len(orig))
	fmt.Println(0, curr)
	copy(curr, orig)
	for i := range n {
		next := make([]int, 0)
		for _, stone := range curr {
			s := fmt.Sprint(stone)
			switch {
			case stone == 0:
				next = append(next, 1)
			case len(s)%2 == 0:
				left, err := strconv.Atoi(s[:len(s)/2])
				util.Must(err)
				right, err := strconv.Atoi(s[len(s)/2:])
				util.Must(err)
				next = append(next, left, right)
			default:
				next = append(next, stone*2024)
			}
		}
		curr = next
		fmt.Println(i + 1)
	}
	return len(curr)
}

func blink2(orig []int, n int) int {
	curr := make(map[int]int, 0)
	for _, val := range orig {
		curr[val]++
	}

	for range n {
		next := make(map[int]int, 0)
		for val, count := range curr {
			s := fmt.Sprint(val)
			switch {
			case val == 0:
				next[1] += count

			case len(s)%2 == 0:
				left, err := strconv.Atoi(s[:len(s)/2])
				util.Must(err)
				right, err := strconv.Atoi(s[len(s)/2:])
				util.Must(err)
				next[left] += count
				next[right] += count
			default:
				next[val*2024] += count
			}
		}
		curr = next
	}

	sol := 0
	for _, count := range curr {
		sol += count
	}
	return sol
}

func part1(input Input) {
	fmt.Println("SOLUTION TO PART 1:", blink2(input.stones, 25))
}

func part2(input Input) {
	fmt.Println("SOLUTION TO PART 2:", blink2(input.stones, 75))
}
