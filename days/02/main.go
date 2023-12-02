package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern = regexp.MustCompile(`Game (\d+): ([\w\s\d,]+)(;[\w\s\d,]+)*`)
)

type Result struct {
	red, green, blue int
}

func (r Result) IsImpossible() bool {
	return r.red > 12 || r.green > 13 || r.blue > 14
}

type Game struct {
	id      int
	results []Result
}

func (g Game) IsImpossible() bool {
	for _, result := range g.results {
		if result.IsImpossible() {
			return true
		}
	}
	return false
}

func NewGame(in string) (*Game, error) {
	match := pattern.FindStringSubmatch(in)

	if len(match) < 3 {
		fmt.Printf("suspicious match: %s -> %s\n", in, match)
		return nil, errors.New("fail")
	}

	g := new(Game)
	g.id = stellar_strings.ToInt(match[1])
	g.results = make([]Result, 0)

	for _, result := range match[2:] {
		parts := strings.Split(strings.Trim(result, ";"), ",")
		result := Result{}
		for _, part := range parts {
			components := strings.Split(strings.TrimSpace(part), " ")
			if len(components) != 2 {
				fmt.Printf("suspicios components: %s -> %s\n", part, components)
				return nil, errors.New("fail")
			}
			switch components[1] {
			case "red":
				result.red = stellar_strings.ToInt(components[0])
			case "green":
				result.green = stellar_strings.ToInt(components[0])
			case "blue":
				result.blue = stellar_strings.ToInt(components[0])
			default:
				fmt.Printf("unknown color %s (%s)\n", components[1], components)
				return nil, errors.New("fail")
			}
		}
		g.results = append(g.results, result)
	}

	return g, nil
}

func part1(in []string) int {
	result := 0
	for _, line := range in {
		g, err := NewGame(line)
		if err != nil {
			return 0
		}

		if !g.IsImpossible() {
			result += g.id
		}

	}
	return result
}

func part2(_ []string) int {
	return 0
}

func main() {

	data := input.LoadString("input")

	fmt.Println("bad: 2996")
	fmt.Printf("part 1: %d\n", part1(data))
	// fmt.Printf("part 2: %d\n", part2(data))
}
