package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	patterns, designs []string
}

func preprocess(data []byte) Input {
	parts := strings.Split(string(data), "\n\n")

	patterns := strings.Split(parts[0], ", ")

	designs := lib.SplitLines(parts[1])

	return Input{patterns: patterns, designs: designs}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	solve(input)
}

type Node struct {
	children map[byte]*Node
	isEnd    bool
}

func newNode() *Node {
	return &Node{
		children: make(map[byte]*Node),
		isEnd:    false,
	}
}

func (n *Node) insert(s string) {

	if s == "" {
		n.isEnd = true
		return
	}

	if _, exists := n.children[s[0]]; !exists {
		n.children[s[0]] = newNode()
	}

	n.children[s[0]].insert(s[1:])
}

func solve(in Input) {

	root := newNode()
	for _, pattern := range in.patterns {
		root.insert(pattern)
	}

	sol1 := 0
	sol2 := 0
	for _, design := range in.designs {

		curr := map[*Node]int{root: 1}

		for i := 0; i < len(design); i++ {
			b := design[i]
			next := map[*Node]int{}

			for node, count := range curr {

				if node.isEnd {
					if child, exists := root.children[b]; exists {
						next[child] += count
					}
				}

				if child, exists := node.children[b]; exists {
					next[child] += count
				}

			}
			curr = next
		}

		solved := false
		for node, count := range curr {
			if node.isEnd {
				if !solved {
					sol1++
					solved = true
				}
				sol2 += count
			}
		}

	}

	fmt.Println("SOLUTION TO PART 1:", sol1)

	fmt.Println("SOLUTION TO PART 2:", sol2)
}
