package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	almanac map[string]Mapping
	seeds   []int
}

type Mapping struct {
	nextKey string
	entries [][]int
}

func preprocess(data []byte) Input {
	parts := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	seeds := lib.MustSplitToInts(strings.TrimPrefix(parts[0], "seeds: "), " ")

	almanac := make(map[string]Mapping)

	for i := 1; i < len(parts); i++ {
		p := parts[i]
		lines := lib.SplitLines(p)
		keys := strings.Split(strings.TrimSuffix(lines[0], " map:"), "-")
		src := keys[0]
		dst := keys[2]
		entries := make([][]int, 0)
		for j := 1; j < len(lines); j++ {
			l := lines[j]
			e := lib.MustSplitToInts(l, " ")
			entries = append(entries, e)
		}
		almanac[src] = Mapping{nextKey: dst, entries: entries}
	}
	return Input{almanac: almanac, seeds: seeds}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func follow(key string, loc int, in Input) (string, int) {
	mapping, found := in.almanac[key]
	lib.MustBeTrue(found)
	nextLoc := -1
	for _, entry := range mapping.entries {
		if loc >= entry[1] && loc < entry[1]+entry[2] {
			nextLoc = (loc - entry[1]) + entry[0]
			break
		}
	}
	if nextLoc < 0 {
		nextLoc = loc
	}
	return mapping.nextKey, nextLoc
}

func part1(in Input) {
	// fmt.Printf("%+v\n", in)
	sol := -1
	for _, s := range in.seeds {
		key := "seed"
		loc := s
		for key != "location" {
			key, loc = follow(key, loc, in)
		}
		// fmt.Println()
		if sol < 0 || loc < sol {
			sol = loc
		}
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func reverseAlmanac(almanac map[string]Mapping) map[string]Mapping {
	rAlmanac := make(map[string]Mapping)
	for src, mapping := range almanac {
		dst := mapping.nextKey
		rAlmanac[dst] = Mapping{nextKey: src, entries: mapping.entries}
	}
	return rAlmanac
}

func rFollow(key string, loc int, rAlmanac map[string]Mapping) (string, int) {
	mapping, found := rAlmanac[key]
	lib.MustBeTrue(found)
	nextLoc := -1
	for _, entry := range mapping.entries {
		if loc >= entry[0] && loc < entry[0]+entry[2] {
			nextLoc = (loc - entry[0]) + entry[1]
			break
		}
	}
	if nextLoc < 0 {
		nextLoc = loc
	}
	return mapping.nextKey, nextLoc
}

func isSeed(s int, seeds []int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if s >= seeds[i] && s < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}

func part2(in Input) {

	rAlmanac := reverseAlmanac(in.almanac)
	for l := 0; true; l++ {
		key := "location"
		loc := l
		// fmt.Printf("(%s, %d)", key, loc)
		for key != "seed" {
			key, loc = rFollow(key, loc, rAlmanac)
		}
		// fmt.Printf(" -> (%s, %d)\n", key, loc)
		if isSeed(loc, in.seeds) {
			fmt.Println("SOLUTION TO PART 2:", l)
			return
		}
	}
}
