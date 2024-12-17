package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	a, b, c int
	prog    []int
}

func parseRegister(line string) int {
	parts := strings.Split(line, ": ")
	return lib.MustToInt(parts[1])
}

func parseProg(line string) []int {
	parts := strings.Split(line, ": ")
	return lib.MustToInts(strings.Split(parts[1], ","))
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))
	a := parseRegister(lines[0])
	b := parseRegister(lines[1])
	c := parseRegister(lines[2])
	prog := parseProg(lines[4])
	return Input{a: a, b: b, c: c, prog: prog}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)
	part2(input)
}

type Device struct {
	a, b, c int
	prog    []int
	ip      int

	history lib.Set[string]

	stdout []int
}

func (d *Device) next() (int, int, bool) {
	if d.ip+1 >= len(d.prog) {
		return 0, 0, true
	}
	d.ip += 2
	return d.prog[d.ip-2], d.prog[d.ip-1], false
}

func (d *Device) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return d.a
	case 5:
		return d.b
	case 6:
		return d.c
	case 7:
		panic("reserved")
	default:
	}
	panic("invalid operand")
}

func (d *Device) div(operand int) int {
	return d.a >> d.combo(operand)
}

func (d *Device) adv(operand int) { d.a = d.div(operand) }       // 0
func (d *Device) bxl(operand int) { d.b = d.b ^ operand }        // 1
func (d *Device) bst(operand int) { d.b = d.combo(operand) % 8 } // 2

func (d *Device) jnz(operand int) { // 3
	if d.a != 0 {
		d.ip = operand
	}
}

func (d *Device) bxc(_ int) { d.b = d.b ^ d.c } // 4

func (d *Device) out(operand int) { // 5

	d.stdout = append(d.stdout, d.combo(operand)%8)
}

func (d *Device) bdv(operand int) { d.b = d.div(operand) } // 6
func (d *Device) cdv(operand int) { d.c = d.div(operand) } // 7

func (d *Device) step() bool {
	opcode, operand, halted := d.next()
	if halted {
		return false
	}

	switch opcode {
	case 0:
		d.adv(operand)
	case 1:
		d.bxl(operand)
	case 2:
		d.bst(operand)
	case 3:
		d.jnz(operand)
	case 4:
		d.bxc(operand)
	case 5:
		d.out(operand)
	case 6:
		d.bdv(operand)
	case 7:
		d.cdv(operand)
	default:
		lib.Panicf("unexpected opcode %4d", opcode)
	}
	return true
}

func (d *Device) debug() string {
	return fmt.Sprintf("a: %4b, b: %4b, c: %4b, ip: %4d, prog: %v", d.a, d.b, d.c, d.ip, d.prog[d.ip:min(d.ip+2, len(d.prog))])
}

func printInts(ints []int) string {
	s := make([]string, len(ints))
	for i, o := range ints {
		s[i] = fmt.Sprint(o)
	}
	return strings.Join(s, ",")
}

func (d *Device) output() string {
	return printInts(d.stdout)
}

func RunDevice(a, b, c int, prog []int) string {

	d := Device{
		a:       a,
		b:       b,
		c:       c,
		prog:    prog,
		ip:      0,
		history: lib.MakeSet[string](),
	}

	// fmt.Println(d.debug())

	for d.step() {
		h := d.debug()

		// fmt.Println(h)

		if d.history.Contains(h) {
			fmt.Println(d.debug())
			fmt.Println("Output: ", d.output())
			panic("CYCLE")
		}
		d.history.Insert(h)

	}

	return d.output()
}

func part1(in Input) {
	sol1 := RunDevice(in.a, in.b, in.c, in.prog)
	fmt.Println("SOLUTION TO PART 1:", sol1)
}

func match(a int, want []int) int {
	i := 0

	for a != 0 {

		c := (a & 0b111) ^ 1

		b := c ^ (a >> c)

		a = a >> 3

		have := (b ^ 6) & 0b111

		if have != want[i] {
			return i
		}

		i++
		if i == len(want) && a != 0 {
			return -1
		}

	}

	return i
}

func updateBest(best, next int) int {
	if best == -1 {
		return next
	} else if next == -1 {
		return best
	}
	return min(best, next)
}

func check(nbits, last int, want []int) int {

	best := -1

	for i := 0; i < 8; i++ {
		next := last | (i << nbits)
		m := match(next, want)

		// strNew := fmt.Sprintf("%03b", i)
		// strLast := fmt.Sprintf("%0"+fmt.Sprint(nbits)+"b", last)
		// strNext := fmt.Sprintf("%0"+fmt.Sprint(nbits+3)+"b", next)
		// tmp := slices.Repeat([]string{"x"}, len(want)-m)

		// for j := 0; j < m; j++ {
		// 	tmp = append(tmp, fmt.Sprint(want[len(tmp)]))
		// }
		// strMatch := "[" + strings.Join(tmp, ",") + "]"
		// fmt.Printf("%s | %s -> %s (%10d) : %s\n", strNew, strLast, strNext, next, strMatch)

		if m == len(want) {
			best = updateBest(best, next)
		} else if m > (nbits-7)/3 {
			best = updateBest(best, check(nbits+3, next, want))
		}

	}

	return best
}

func part2(in Input) {
	want := []int{2, 4, 1, 1, 7, 5, 0, 3, 4, 3, 1, 6, 5, 5, 3, 0}

	best := -1

	for a := 0; a < 1024; a++ {

		m := match(a, want)

		if m > 0 {
			best = updateBest(best, check(10, a, want))
		}

	}

	have := RunDevice(best, in.b, in.c, in.prog)
	lib.MustBeEqual(have, printInts(want))
	fmt.Println("SOLUTION TO PART 2:", best)

}
