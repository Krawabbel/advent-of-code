package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"slices"
	"sort"
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

		op := operation(parts[1])
		left := parts[0]
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

	// fmt.Println(id)

	if _, found := states[id]; !found {
		val := valueImpl(id, states, gates)
		states[id] = bool2state(val)
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

func parse(name string, states map[string]int, gates map[string]*Gate) int {
	v := 0

	for id, val := range states {
		if strings.Contains(id, name) {
			bit := lib.MustToInt(strings.ReplaceAll(id, name, ""))
			if val == 1 {
				v |= (1 << bit)
			}
		}
	}

	for id := range gates {
		if strings.Contains(id, name) {
			bit := lib.MustToInt(strings.ReplaceAll(id, name, ""))
			if value(id, states, gates) {
				v |= (1 << bit)
			}
		}
	}

	return v
}

func part1(in Input) {
	sol := parse("z", lib.CloneMap(in.initial), in.gates)
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func makeId(s string, n int) string {
	return fmt.Sprintf(s+"%02d", n)
}

func symbolic(id string, gates map[string]*Gate) string {
	gate, found := gates[id]
	if !found {
		return id
	}

	op := ""
	switch gate.op {
	case GATE_AND:
		op = "&"
	case GATE_OR:
		op = "|"
	case GATE_XOR:
		op = "^"
	default:
		panic("unexpected")
	}

	s1 := symbolic(gate.left, gates)
	s2 := symbolic(gate.right, gates)

	if s1 < s2 {
		return fmt.Sprintf("(%s %s %s)", s1, op, s2)
	} else {
		return fmt.Sprintf("(%s %s %s)", s2, op, s1)
	}
}

func adder(N int) map[string]*Gate {

	gates := make(map[string]*Gate)

	gates["z00"] = &Gate{op: GATE_XOR, left: "x00", right: "y00"}
	gates["d00"] = &Gate{op: GATE_AND, left: "x00", right: "y00"}
	gates["b00"] = &Gate{op: GATE_XOR, left: "x00", right: "y00"}

	gates["b01"] = &Gate{op: GATE_XOR, left: "x01", right: "y01"}
	gates["d01"] = &Gate{op: GATE_AND, left: "x01", right: "y01"}
	gates["z01"] = &Gate{op: GATE_XOR, left: "a01", right: "b01"}
	gates["a01"] = &Gate{op: GATE_AND, left: "x00", right: "y00"}
	gates["c01"] = &Gate{op: GATE_AND, left: "b01", right: "d00"}

	for n := 2; n <= N; n++ {
		gates[makeId("b", n)] = &Gate{op: GATE_XOR, left: makeId("x", n), right: makeId("y", n)}
		gates[makeId("d", n)] = &Gate{op: GATE_AND, left: makeId("x", n), right: makeId("y", n)}
		gates[makeId("c", n)] = &Gate{op: GATE_AND, left: makeId("a", n), right: makeId("b", n)}
		gates[makeId("a", n)] = &Gate{op: GATE_OR, left: makeId("c", n-1), right: makeId("d", n-1)}
		gates[makeId("z", n)] = &Gate{op: GATE_XOR, left: makeId("a", n), right: makeId("b", n)}
	}

	gates[makeId("z", N)] = &Gate{op: GATE_OR, left: makeId("c", N-1), right: makeId("d", N-1)}

	return gates
}

func compare(have_id, want_id string, have, want map[string]*Gate) (bool, string) {

	want_gate, found := want[want_id]
	lib.MustBeTrue(found)

	have_gate, found := have[have_id]
	lib.MustBeTrue(found)

	if want_gate.op != have_gate.op {
		return false, have_id
	}

	shl := symbolic(have_gate.left, have)
	shr := symbolic(have_gate.right, have)

	swl := symbolic(want_gate.left, want)
	swr := symbolic(want_gate.right, want)

	if shl == swl {
		if shr == swr {
			return true, ""
		}
		return compare(have_gate.right, want_gate.right, have, want)
	}

	if shl == swr {
		if shr == swl {
			return true, ""
		}
		return compare(have_gate.right, want_gate.left, have, want)
	}

	return false, have_id
}

func detect_cycle_implementation(have_id string, prev []string, gates map[string]*Gate) bool {

	if slices.Contains(prev, have_id) {
		return true
	}

	g, found := gates[have_id]
	if !found {
		// fmt.Println(id)
		return false
	}

	hist := append([]string{have_id}, prev...)

	return detect_cycle_implementation(g.left, hist, gates) ||
		detect_cycle_implementation(g.right, hist, gates)
}

func detect_cycle(gates map[string]*Gate) bool {

	for id := range gates {
		if detect_cycle_implementation(id, []string{}, gates) {
			return true
		}
	}

	return false
}

func part2(in Input) {

	add := adder(in.zmax)

	gates := lib.CloneMap(in.gates)

	swaps := make([]string, 0)

	for n := 0; n <= in.zmax; n++ {

		zn := makeId("z", n)
		have := symbolic(zn, gates)
		want := symbolic(zn, add)

		if have != want {
			m, id := compare(zn, zn, gates, add)
			lib.MustBeFalse(m)

			// fmt.Println(zn, have)
			// fmt.Println(zn, want)
			// fmt.Println("***")
			// fmt.Println(m, id)

			swapIds := make([]string, 0)

			for swapId := range in.gates {

				if swapId == id {
					continue
				}

				swapped := lib.CloneMap(gates)
				swapped[id], swapped[swapId] = swapped[swapId], swapped[id]

				if detect_cycle(swapped) {
					continue
				}

				ok := true

				for i := 0; i <= n; i++ {
					cand := symbolic(zn, swapped)
					if cand != want {
						ok = false
						break
					}
				}

				if ok {
					swapIds = append(swapIds, swapId)
					// fmt.Println(n, id, "<->", swapId)
				}
			}

			// fmt.Println("want")
			// fmt.Println(want)
			// fmt.Println("have")
			// fmt.Println(have)

			// fmt.Printf("%s: %s <-> %v\n", zn, id, swapIds)
			// lib.MustPressEnter()

			lib.MustBeEqual(len(swapIds), 1)
			gates[id], gates[swapIds[0]] = gates[swapIds[0]], gates[id]
			swaps = append(swaps, swapIds[0], id)

		}

	}

	sort.Strings(swaps)
	fmt.Println("SOLUTION TO PART 2:", strings.Join(swaps, ","))
}
