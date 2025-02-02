package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards string
	bid   int
	typ1  []int
	typ2  []int
}

type Input struct {
	hands []Hand
}

func classify1(cards string) []int {
	labels := make(map[rune]int)
	for _, r := range cards {
		labels[r]++
	}
	t := make([]int, 5)
	for _, count := range labels {
		t[count-1]++
	}
	return t
}

func classify2(cards string) []int {
	counts := make(map[rune]int)
	for _, r := range cards {
		counts[r]++
	}

	mlabel, mcnt := '0', -1
	for label, cnt := range counts {
		if label != 'J' && cnt > mcnt {
			mcnt = cnt
			mlabel = label
		}
	}

	counts[mlabel] += counts['J']
	delete(counts, 'J')

	t := make([]int, 5)
	for _, count := range counts {
		t[count-1]++
	}
	return t
}

func preprocess(data []byte) Input {
	lines := lib.SplitLines(string(data))

	hs := make([]Hand, len(lines))
	for i, l := range lines {
		ps := strings.Split(l, " ")
		c := ps[0]
		b := lib.MustToInt(ps[1])
		t1 := classify1(c)
		t2 := classify2(c)
		h := Hand{cards: c, bid: b, typ1: t1, typ2: t2}
		hs[i] = h
	}

	return Input{hands: hs}
}

type ByCards struct {
	hands []Hand
	value map[byte]int
	typ   func(h Hand, i int) int
}

func (a ByCards) Len() int      { return len(a.hands) }
func (a ByCards) Swap(i, j int) { a.hands[i], a.hands[j] = a.hands[j], a.hands[i] }

func (a ByCards) Less(i, j int) bool {

	for k := range 5 {
		ti := a.typ(a.hands[i], 4-k)
		tj := a.typ(a.hands[j], 4-k)
		if ti == tj {
			continue
		}
		return ti < tj
	}

	for l := range 5 {
		if a.hands[i].cards[l] == a.hands[j].cards[l] {
			continue
		}
		return a.value[a.hands[i].cards[l]] < a.value[a.hands[j].cards[l]]
	}

	panic("unexpected")
}

func value(order string) map[byte]int {

	value := make(map[byte]int)
	for i := range len(order) {
		value[order[i]] = i
	}
	return value
}

func solve(hands []Hand, order string, tfun func(h Hand, i int) int) int {

	value := value(order)

	sorted := ByCards{
		hands: make([]Hand, len(hands)),
		value: value,
		typ:   tfun,
	}
	copy(sorted.hands, hands)

	sort.Sort(sorted)

	sol := 0
	for i, h := range sorted.hands {
		sol += (i + 1) * h.bid
	}

	return sol
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	fmt.Println("SOLUTION TO PART 1:", solve(input.hands, "23456789TJQKA", func(h Hand, i int) int { return h.typ1[i] }))
	fmt.Println("SOLUTION TO PART 1:", solve(input.hands, "J23456789TQKA", func(h Hand, i int) int { return h.typ2[i] }))
}
