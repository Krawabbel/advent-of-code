package intcode

type VM struct {
	prog []int

	pc int

	halt bool
}

func NewVM(prog []int) *VM {
	return &VM{prog: prog}
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

func (vm *VM) Run() {
	for !vm.halt {
		vm.step()
	}
}

func (vm *VM) ReadAddr(addr int) int {
	return vm.prog[addr]
}
