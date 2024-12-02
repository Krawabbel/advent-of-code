package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	path := os.Args[1]
	blob, err := os.ReadFile(path)
	must(err)

	r := bytes.NewReader(blob)
	scan := bufio.NewScanner(r)

	list0 := make([]int, 0)
	list1 := make([]int, 0)
	for scan.Scan() {
		line := scan.Text()

		entries := strings.Split(line, "   ")

		entry0, err := strconv.Atoi(entries[0])
		must(err)
		list0 = append(list0, entry0)

		entry1, err := strconv.Atoi(entries[1])
		must(err)
		list1 = append(list1, entry1)

		// small sanity check
		if line != fmt.Sprintf("%d   %d", entry0, entry1) {
			panic(fmt.Errorf("mismatch in line %s", line))
		}
	}
	must(scan.Err())

	sort.Ints(list0)
	sort.Ints(list1)

	distance := 0
	for i := range len(list0) {
		distance += abs(list0[i] - list1[i])
	}

	fmt.Println("ANSWER TO PART 1:", distance)

	count0 := make(map[int]int)
	for _, entry0 := range list0 {
		count0[entry0]++
	}

	count1 := make(map[int]int)
	for _, entry1 := range list1 {
		count1[entry1]++
	}

	similarity_score := 0
	for entry, c0 := range count0 {
		c1, found := count1[entry]
		if found {
			similarity_score += entry * c0 * c1
		}
	}

	fmt.Println("ANSWER TO PART 2:", similarity_score)
}
