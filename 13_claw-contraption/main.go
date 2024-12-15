package main

import (
	"aoc2024/internal/lib"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	xA, yA, xB, yB, xP, yP int
}

type Input struct {
	games []Game
}

var reButtonA = regexp.MustCompile(`Button A: X\+([0-9]*), Y\+([0-9]*)`)
var reButtonB = regexp.MustCompile(`Button B: X\+([0-9]*), Y\+([0-9]*)`)
var rePrize = regexp.MustCompile(`Prize: X=([0-9]*), Y=([0-9]*)`)

func processLine(s *bufio.Scanner, re *regexp.Regexp) (int, int) {
	if !s.Scan() {
		panic("could not scan")
	}
	line := s.Text()
	matches := re.FindStringSubmatch(line)

	x, err := strconv.Atoi(matches[1])
	lib.Must(err)

	y, err := strconv.Atoi(matches[2])
	lib.Must(err)

	return x, y
}

func preprocess(data []byte) Input {

	games := make([]Game, 0)

	scan := bufio.NewScanner(bytes.NewReader(data))
	for {
		xA, yA := processLine(scan, reButtonA)
		xB, yB := processLine(scan, reButtonB)
		xP, yP := processLine(scan, rePrize)

		g := Game{
			xA: xA,
			yA: yA,
			xB: xB,
			yB: yB,
			xP: xP,
			yP: yP,
		}

		games = append(games, g)

		if !scan.Scan() {
			break
		}

		if txt := scan.Text(); txt != "" {
			fmt.Println(txt)
			panic("unexpected scan")
		}
	}

	lib.Must(scan.Err())

	return Input{games: games}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

func solve(xA, xB, yA, yB, xP, yP int) int {

	det := xA*yB - yA*xB

	if det == 0 {
		panic("det = 0")
	}

	a := yB*xP - xB*yP
	if (a % det) != 0 {
		return 0
	}

	b := -yA*xP + xA*yP
	if (b % det) != 0 {
		return 0
	}

	return (3*a + b) / det
}

func part1(input Input) {
	sol := 0
	for _, g := range input.games {
		sol += solve(g.xA, g.xB, g.yA, g.yB, g.xP, g.yP)
	}
	fmt.Println("SOLUTION TO PART 1:", sol)
}

func part2(input Input) {
	offset := 10000000000000
	sol := 0
	for _, g := range input.games {
		sol += solve(g.xA, g.xB, g.yA, g.yB, g.xP+offset, g.yP+offset)
	}
	fmt.Println("SOLUTION TO PART 2:", sol)
}
