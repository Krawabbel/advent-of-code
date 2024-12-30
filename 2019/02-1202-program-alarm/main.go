package main

import (
	"aoc/2019/intcode"
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	prog []int
}

func preprocess(data []byte) Input {
	prog := lib.MustToInts(strings.Split(string(data), ","))
	return Input{prog: prog}
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
	prog := lib.CloneSlice1D(in.prog)
	prog[1] = 12
	prog[2] = 02
	vm := intcode.NewVM(prog)
	vm.Run()

	sol := vm.ReadAddr(0)
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {

	want := 19690720

	for i := range 100 {
		for j := range 100 {

			prog := lib.CloneSlice1D(in.prog)
			prog[1] = i
			prog[2] = j
			vm := intcode.NewVM(prog)
			vm.Run()

			have := vm.ReadAddr(0)
			if have == want {
				sol := 100*i + j
				fmt.Println("SOLUTION TO PART 2:", sol)
				return
			}
		}
	}

}
