package main

import (
	"aoc/internal/lib"
	"testing"
)

func TestDevice_step_01(t *testing.T) {
	dev := Device{c: 9, prog: []int{2, 6}}
	dev.step()
	lib.ExpectEqual(t, dev.b, 1)
}

func TestDevice_step_02(t *testing.T) {
	dev := Device{a: 10, prog: []int{5, 0, 5, 1, 5, 4}}
	for dev.step() {
	}
	lib.ExpectEqual(t, dev.output(), "0,1,2")
}

func TestDevice_step_03(t *testing.T) {
	dev := Device{a: 2024, prog: []int{0, 1, 5, 4, 3, 0}}
	for dev.step() {
	}
	lib.ExpectEqual(t, dev.output(), "4,2,5,6,7,7,7,7,3,1,0")
	lib.ExpectEqual(t, dev.a, 0)
}

func TestDevice_step_04(t *testing.T) {
	dev := Device{b: 29, prog: []int{1, 7}}
	for dev.step() {
	}
	lib.ExpectEqual(t, dev.b, 26)
}

func TestDevice_step_05(t *testing.T) {
	dev := Device{a: 2024, c: 43690, prog: []int{4, 0}}
	for dev.step() {
	}
	lib.ExpectEqual(t, dev.b, 44354)
}

func TestBinary(t *testing.T) {
	lib.ExpectEqual(t, 0b111, 7)
	lib.ExpectEqual(t, 0b110, 6)
	lib.ExpectEqual(t, (1 << 3), 8)
	lib.ExpectEqual(t, 7&0b111, 7)
	lib.ExpectEqual(t, 1>>63, 7)
}
