package main

import (
	"aoc/internal/lib"
	"bytes"
	"fmt"
	"os"
)

type Input struct {
	keys, locks []Schematic
}

type Schematic struct {
	heights []int
}

func preprocess(data []byte) Input {

	schematics := bytes.Split(data, []byte("\n\n"))

	keys := make([]Schematic, 0)
	locks := make([]Schematic, 0)

	for _, s := range schematics {
		schematic := bytes.Split(s, []byte("\n"))

		isLock := !bytes.ContainsFunc(schematic[0], func(r rune) bool { return r != '#' })
		isKey := !bytes.ContainsFunc(schematic[0], func(r rune) bool { return r != '.' })

		if isLock { // LOCK
			h := make([]int, len(schematic[0]))
			for c := range schematic[0] {
				for r := 1; r < len(schematic); r++ {
					v := schematic[r][c]
					if v == '#' {
						h[c]++
					}
				}
			}
			locks = append(locks, Schematic{h})

		} else if isKey { // KEY
			h := make([]int, len(schematic[0]))
			for c := range schematic[0] {
				h[c] = -1
				for r := 1; r < len(schematic); r++ {
					v := schematic[r][c]
					if v == '#' {
						h[c]++
					}
				}
			}
			keys = append(keys, Schematic{h})

		} else {
			panic("unexpected schematic")
		}

		fmt.Println("***")
	}

	return Input{keys: keys, locks: locks}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func part1(in Input) {
	fmt.Printf("%+v\n", in)
	sol := 0
	for _, key := range in.keys {

		for _, lock := range in.locks {
			fit := true
			for c := range key.heights {
				if key.heights[c]+lock.heights[c] >= 6 {
					fit = false
				}
			}
			if fit {
				sol++
			}
		}
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	fmt.Println("SOLUTION TO PART 2:", "nothing to do :)")
}
