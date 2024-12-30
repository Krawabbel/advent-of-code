package main

import (
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

type VM struct {
	prog []int

	pc int

	halt bool
}

func (vm *VM) next() int {
	val := vm.prog[vm.pc]
	vm.pc++
	return val
}

func (vm *VM) add(src1, src2, dst int) {
	vm.prog[dst] = vm.prog[src1] + vm.prog[src2]
}

func (vm *VM) mult(src1, src2, dst int) {
	vm.prog[dst] = vm.prog[src1] * vm.prog[src2]
}

func (vm *VM) step() {
	opcode := vm.next()
	switch opcode {
	case 1:
		vm.add(vm.next(), vm.next(), vm.next())
	case 2:
		vm.mult(vm.next(), vm.next(), vm.next())
	case 99:
		vm.halt = true
	}
}

func (vm *VM) run() {
	for !vm.halt {
		vm.step()
	}
}

func part1(in Input) {
	vm := VM{prog: lib.CloneSlice1D(in.prog)}
	vm.prog[1] = 12
	vm.prog[2] = 02
	vm.run()

	sol := vm.prog[0]
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {

	want := 19690720

	for i := range 100 {
		for j := range 100 {
			vm := VM{prog: lib.CloneSlice1D(in.prog)}
			vm.prog[1] = i
			vm.prog[2] = j
			vm.run()

			have := vm.prog[0]
			if have == want {
				sol := 100*i + j
				fmt.Println("SOLUTION TO PART 2:", sol)
				return
			}
		}
	}

}
