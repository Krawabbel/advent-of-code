package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Edge struct {
	left, right string
}

type Input struct {
	edges []Edge
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))

	edges := make([]Edge, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "-")
		edges[i] = Edge{left: parts[0], right: parts[1]}
	}

	return Input{edges: edges}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

type Node struct {
	id    string
	conns lib.Set[string]
}

func NewNode(id string) *Node {
	return &Node{id, lib.MakeSet[string]()}
}

func add(dict map[string]*Node, id string) {
	if _, found := dict[id]; !found {
		dict[id] = NewNode(id)
	}
}

func connect(dict map[string]*Node, first, next string) {
	add(dict, first)
	root := dict[first]
	add(dict, next)
	root.conns.Insert(next)
}

func part1(in Input) {

	dict := make(map[string]*Node)

	for _, e := range in.edges {
		if e.left < e.right {
			connect(dict, e.left, e.right)
		} else if e.left == e.right {
			panic("unexpected")
		} else {
			connect(dict, e.right, e.left)
		}
	}

	// for _, n := range nodes {
	// 	fmt.Println(n.id, len(n.conns))
	// }

	sets := lib.MakeSet[string]()

	for ida, a := range dict {
		for idb := range a.conns.Elements {
			b := dict[idb]
			for idc := range b.conns.Elements {
				if a.conns.Contains(idc) {
					sets.Insert(ida + idb + idc)
				}
			}
		}
	}

	// fmt.Println(len(sets))
	// lib.PrintSlice(sets)

	sol := 0
	for set := range sets.Elements {

		found := false
		for i := 0; i < len(set); i += 2 {
			if set[i] == 't' {
				found = true
			}
		}

		if found {
			// fmt.Println(set)
			sol++
		}

	}

	// 2377 too high
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func hash(strs []string) string {
	tmp := make([]string, len(strs))
	copy(tmp, strs)
	sort.Strings(tmp)
	return strings.Join(tmp, ",")
}

func bronKerbosch(r, p, x lib.Set[string], dict map[string]*Node, report func([]string)) {
	if p.IsEmpty() && x.IsEmpty() {
		report(r.Slice())
		return
	}

	for !p.IsEmpty() {
		v := p.Peek()

		nextR := lib.SetInsert(r, v)

		n := dict[v].conns
		nextP := lib.SetIntersect(p, n)
		nextX := lib.SetIntersect(x, n)

		bronKerbosch(nextR, nextP, nextX, dict, report)

		p.Delete(v)
		x.Insert(v)
	}

}

func part2(in Input) {

	dict := make(map[string]*Node)

	vertices := lib.MakeSet[string]()

	for _, e := range in.edges {
		connect(dict, e.right, e.left)
		connect(dict, e.left, e.right)
		vertices.Insert(e.left, e.right)
	}

	sol := ""
	report := func(clique []string) {
		h := hash(clique)
		if len(h) > len(sol) {
			sol = h
			// fmt.Printf("new maximal clique: %s\n", h)
		}
	}

	bronKerbosch(lib.MakeSet[string](), vertices, lib.MakeSet[string](), dict, report)

	fmt.Printf("SOLUTION TO PART 2: %s\n", sol)

}
