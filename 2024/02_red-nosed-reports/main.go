package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func isSafeBase(report []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 1 || diff > 3 {
			isIncreasing = false
		}
		if diff > -1 || diff < -3 {
			isDecreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

func isSafeDampened(report []int) bool {
	if isSafeBase(report) {
		return true
	}

	for i := range report {
		tmp := make([]int, len(report)-1)
		for j := range report {
			if j < i {
				tmp[j] = report[j]
			} else if j == i {
				// skip
			} else {
				tmp[j-1] = report[j]
			}
		}
		if isSafeBase(tmp) {
			return true
		}
	}

	return false
}

func main() {
	path := os.Args[1]
	blob, err := os.ReadFile(path)
	must(err)

	r := bytes.NewReader(blob)
	scan := bufio.NewScanner(r)

	nSafePart1 := 0
	nSafePart2 := 0
	for scan.Scan() {
		line := scan.Text()

		report := strings.Split(line, " ")

		rep := make([]int, len(report))
		for i, r := range report {
			rep[i], err = strconv.Atoi(r)
			must(err)
		}

		if isSafeBase(rep) {
			nSafePart1++
		}

		if isSafeDampened(rep) {
			nSafePart2++
		}

	}
	must(scan.Err())

	fmt.Println("ANSWER TO PART 1:", nSafePart1)

	fmt.Println("ANSWER TO PART 2:", nSafePart2)
}
