package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

const (
	STATE_FALSE = iota
	STATE_TRUE
	STATE_UNDEFINED
)

func bool2state(b bool) int {
	if b {
		return STATE_TRUE
	}
	return STATE_FALSE
}

func state2bool(s int) bool {
	switch s {
	case STATE_TRUE:
		return true
	case STATE_FALSE:
		return false
	}
	lib.Panicf("unexpected state %d", s)
	return false
}

type Input struct {
	initial map[string]int
	gates   map[string]*Gate
	zmax    int
}

const (
	GATE_AND = iota
	GATE_OR
	GATE_XOR
)

type Gate struct {
	op          int
	left, right string
}

func preprocess(data []byte) Input {
	sections := strings.Split(string(data), "\n\n")

	initial := make(map[string]int)
	for _, line := range lib.SplitLines(sections[0]) {
		parts := strings.Split(line, ": ")
		initial[parts[0]] = lib.MustToInt(parts[1])
	}

	gates := make(map[string]*Gate)
	zmax := -1
	for _, line := range lib.SplitLines(sections[1]) {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		left := parts[0]
		op := operation(parts[1])
		right := parts[2]
		out := parts[4]
		gates[out] = &Gate{op, left, right}

		if strings.Contains(out, "z") {
			id := lib.MustToInt(strings.ReplaceAll(out, "z", ""))
			zmax = max(zmax, id)
		}
	}

	return Input{initial: initial, gates: gates, zmax: zmax}
}

func operation(op string) int {
	switch op {
	case "AND":
		return GATE_AND
	case "OR":
		return GATE_OR
	case "XOR":
		return GATE_XOR
	}
	lib.Panicf("unexpected operation '%s'", op)
	return -1
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func value(id string, states map[string]int, gates map[string]*Gate) bool {

	// fmt.Println(id, states, gates)
	if _, found := states[id]; !found {
		val := valueImpl(id, states, gates)
		states[id] = bool2state(val)
		// fmt.Println(id, "->", val)
	}

	return state2bool(states[id])
}

func valueImpl(id string, states map[string]int, gates map[string]*Gate) bool {
	// fmt.Println("IMPL", id)
	gate, found := gates[id]
	lib.MustBeTrue(found)
	left := value(gate.left, states, gates)
	right := value(gate.right, states, gates)
	switch gate.op {
	case GATE_AND:
		return left && right
	case GATE_OR:
		return left || right
	case GATE_XOR:
		return left != right
	}
	lib.Panicf("unexpected operation %d", gate.op)
	return false
}

func part1(in Input) {
	// fmt.Printf("%+v\n", in)
	states := make(map[string]int)
	for key, val := range in.initial {
		states[key] = val
	}

	sol := 0
	for z := 0; z <= in.zmax; z++ {
		id := fmt.Sprintf("z%02d", z)
		// fmt.Println("***", id, "***")
		val := value(id, states, in.gates)
		// fmt.Println(id, "=", val)
		if val {
			sol |= (1 << z)
		}
	}

	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(in Input) {
	sol := 0
	fmt.Println("SOLUTION TO PART 2:", sol)
}
