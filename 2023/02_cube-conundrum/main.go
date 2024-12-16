package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Round struct {
	red, green, blue int
}

type Game struct {
	rounds []Round
}

type Input struct {
	games []Game
}

var (
	reRed   = regexp.MustCompile(`([0-9]*) red`)
	reGreen = regexp.MustCompile(`([0-9]*) green`)
	reBlue  = regexp.MustCompile(`([0-9]*) blue`)
)

func match(re *regexp.Regexp, turn string) int {
	m := re.FindStringSubmatch(turn)
	if len(m) == 0 {
		return 0
	}
	return lib.MustToInt(m[1])
}

func preprocess(data []byte) Input {
	lines := strings.Split(string(data), "\n")
	games := make([]Game, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, ": ")
		turns := strings.Split(parts[1], ";")
		rounds := make([]Round, len(turns))
		// fmt.Printf("Game %d: ", i)
		for j, t := range turns {
			rounds[j].red = match(reRed, t)
			rounds[j].green = match(reGreen, t)
			rounds[j].blue = match(reBlue, t)
			// fmt.Printf("%+v;", rounds[j])
		}
		// fmt.Println()
		games[i] = Game{rounds: rounds}
	}
	return Input{games: games}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	solve(input)

}

func solve(input Input) {
	sol1 := 0
	sol2 := 0
	for i, game := range input.games {

		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range game.rounds {
			counts["red"] = max(counts["red"], round.red)
			counts["green"] = max(counts["green"], round.green)
			counts["blue"] = max(counts["blue"], round.blue)
		}

		isPossible := counts["red"] <= 12 && counts["green"] <= 13 && counts["blue"] <= 14
		if isPossible {
			sol1 += i + 1
		}

		sol2 += counts["red"] * counts["green"] * counts["blue"]

	}

	fmt.Println("SOLUTION TO PART 1:", sol1)
	fmt.Println("SOLUTION TO PART 2:", sol2)
}
