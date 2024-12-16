package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := string(blob)

	lines := strings.Split(input, "\n")

	order := [][]int{}
	pages := [][]int{}
	for _, l := range lines {
		if strings.Contains(l, "|") {

			pair := strings.Split(l, "|")
			left, err := strconv.Atoi(pair[0])
			lib.Must(err)
			right, err := strconv.Atoi(pair[1])
			lib.Must(err)
			order = append(order, []int{left, right})

		} else if len(l) > 0 {

			numbers := strings.Split(l, ",")
			page := make([]int, len(numbers))
			for i, n := range numbers {
				page[i], err = strconv.Atoi(n)
				lib.Must(err)
			}
			pages = append(pages, page)

		}

	}

	part1(order, pages)

	part2(order, pages)
}

func first[T comparable](arr []T, element T) int {
	for i, t := range arr {
		if t == element {
			return i
		}
	}
	return len(arr)
}

func last[T comparable](arr []T, element T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == element {
			return i
		}
	}
	return -1
}

func isValid(pageNumbers []int, order [][]int) (bool, int, int) {
	for _, pair := range order {
		idx1 := last(pageNumbers, pair[0])
		idx2 := first(pageNumbers, pair[1])
		if idx1 > idx2 {
			return false, idx1, idx2
		}
	}
	return true, -1, -1
}

func part1(order, pages [][]int) {
	sol := 0
	for _, p := range pages {
		valid, _, _ := isValid(p, order)
		if valid {
			sol += p[len(p)/2]
		}
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(order, pages [][]int) {

	sol := 0
	for _, p := range pages {
		valid, idx1, idx2 := isValid(p, order)
		if valid {
			continue
		}

		for !valid {
			p[idx1], p[idx2] = p[idx2], p[idx1]
			valid, idx1, idx2 = isValid(p, order)
		}
		sol += p[len(p)/2]

	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}
