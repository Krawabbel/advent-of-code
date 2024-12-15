package lib_test

import (
	"aoc2024/internal/lib"
	"testing"
)

func TestIteration(t *testing.T) {
	elements := map[int]bool{1: false, 2: false, 3: false, 4: false, 5: false}
	s := lib.MakeSet[int]()
	for key := range elements {
		s.Insert(key)
	}

	for e := range s.Elements {
		elements[e] = true
	}

	for key, val := range elements {
		if !val {
			t.Fatal(key)
		}
	}

}
