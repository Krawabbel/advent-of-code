package main

import (
	"aoc2024/internal/util"
	"fmt"
	"os"
	"slices"
	"strings"
)

func solve(input string) {
	fs := make([]int, 0)

	isFile := true
	id := 0
	for _, r := range input {
		n := int(r - '0')
		if isFile {
			fs = append(fs, slices.Repeat([]int{id}, n)...)
			id++
		} else {
			fs = append(fs, slices.Repeat([]int{-1}, n)...)
		}
		isFile = !isFile
	}
	maxId := id - 1

	part1(cloneFS(fs))

	part2(cloneFS(fs), maxId)
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	util.Must(err)

	input := string(blob)

	for _, line := range strings.Split(input, "\n") {
		solve(line)
	}

}

func cloneFS(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func printFS(fs []int) {
	for _, f := range fs {
		if f == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(f)
		}
	}
	fmt.Println()
}

func checksum(fs []int) uint64 {
	sol := uint64(0)
	for k, f := range fs {
		if f != -1 {
			sol += uint64(k) * uint64(f)
		}
	}
	return sol
}

func part1(fs []int) {

	// printFS(fs)

	i := 0
	j := len(fs) - 1
	for {
		for i < j && fs[i] != -1 {
			i++
		}
		for fs[j] == -1 {
			j--
		}
		if i < j {
			fs[i], fs[j] = fs[j], fs[i]
		} else {
			break
		}
	}

	// printFS(fs)

	fmt.Println("SOLUTION TO PART 1:", checksum(fs))
}

func part2(fs []int, maxId int) {
	// printFS(fs)

	j := len(fs) - 1
	for id := maxId; id >= 0; id-- {

		for j >= 0 && fs[j] != id {
			j--
		}
		fEnd := j + 1

		for j >= 0 && fs[j] == id {
			j--
		}
		fStart := j + 1

		fSize := fEnd - fStart

		i := 0
		for i <= j {

			for i <= j && fs[i] != -1 {
				i++
			}
			gStart := i

			for i <= j && fs[i] == -1 {
				i++
			}
			gEnd := i

			gSize := gEnd - gStart

			if gSize < fSize {
				continue
			}

			for k := 0; k < fSize; k++ {
				if fs[gStart+k] != -1 {
					panic("this should not happen a")
				}

				if fs[fStart+k] == -1 {
					panic("this should not happen b")
				}
				fs[gStart+k] = id
				fs[fStart+k] = -1
			}

			break
		}

	}

	// printFS(fs)

	fmt.Println("SOLUTION TO PART 2:", checksum(fs))
}
