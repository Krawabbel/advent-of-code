package main

import (
	"aoc/internal/lib"
	"testing"
)

func TestFirst(t *testing.T) {
	given := 230498
	have := first(given)
	want := given * 64
	lib.ExpectEqual(t, have, want)
}

func TestSecond(t *testing.T) {
	given := 232938747680498
	have := second(given)
	want := given / 32
	lib.ExpectEqual(t, have, want)
}

func TestThird(t *testing.T) {
	given := 2329387476110
	have := third(given)
	want := given * 2048
	lib.ExpectEqual(t, have, want)
}
