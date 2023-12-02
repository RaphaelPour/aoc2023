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
	pattern = regexp.MustCompile(`Game (\d+): (.*)`)

	EmptyResult = Result{}
	EmptyGame   = Game{}
)

type Result struct {
	red, green, blue int
}

func (r Result) Add(other Result) Result {
	r.red += other.red
	r.green += other.green
	r.blue += other.blue
	return r
}

func (r Result) Max(other Result) Result {
	max := Result{}
	if r.red > other.red {
		max.red = r.red
	} else {
		max.red = other.red
	}

	if r.green > other.green {
		max.green = r.green
	} else {
		max.green = other.green
	}

	if r.blue > other.blue {
		max.blue = r.blue
	} else {
		max.blue = other.blue
	}

	return max
}

func (r Result) String() string {
	out := make([]string, 0)
	if r.red > 0 {
		out = append(out, fmt.Sprintf("%d red", r.red))
	}
	if r.blue > 0 {
		out = append(out, fmt.Sprintf("%d blue", r.blue))
	}
	if r.green > 0 {
		out = append(out, fmt.Sprintf("%d green", r.green))
	}
	return strings.Join(out, ",")
}

func NewResult(in string) (Result, error) {
	result := Result{}
	for _, part := range strings.Split(in, ",") {
		components := strings.Split(strings.TrimSpace(part), " ")
		if len(components) != 2 {
			fmt.Printf("suspicios components: %s -> %s\n", part, components)
			return EmptyResult, errors.New("fail")
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
			return EmptyResult, errors.New("fail")
		}
	}

	return result, nil
}

func (r Result) IsPossible() bool {
	return r.red <= 12 && r.green <= 13 && r.blue <= 14
}

type Game struct {
	id      int
	results []Result
}

func (g Game) IsPossible() bool {
	megaResult := Result{}
	for _, result := range g.results {
		if !result.IsPossible() {
			return false
		}
		megaResult = megaResult.Add(result)
	}
	return true
	//return megaResult.IsPossible()
}

func (g Game) Power() int {
	max := Result{}
	for _, result := range g.results {
		max = max.Max(result)
	}
	return max.red * max.blue * max.green
}

func (g Game) String() string {
	out := fmt.Sprintf("Game %d: ", g.id)
	for _, result := range g.results {
		out += fmt.Sprintf("%s; ", result)
	}
	return out
}

func NewGame(in string) (Game, error) {
	match := pattern.FindStringSubmatch(in)

	if len(match) != 3 {
		fmt.Printf("suspicious match: %s -> %s\n", in, match)
		return EmptyGame, errors.New("fail")
	}

	g := Game{}
	g.id = stellar_strings.ToInt(match[1])
	g.results = make([]Result, 0)

	for _, rawResult := range strings.Split(match[2], ";") {
		result, err := NewResult(rawResult)
		if err != nil {
			return EmptyGame, err
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

		if g.IsPossible() {
			result += g.id
		}

	}
	return result
}

func part2(in []string) int {
	result := 0
	for _, line := range in {
		g, err := NewGame(line)
		if err != nil {
			return 0
		}
		result += g.Power()
	}
	return result
}

func main() {

	data := input.LoadString("input")

	fmt.Println("bad: 193,1307, 2996")
	fmt.Printf("part 1: %d\n", part1(data))
	fmt.Printf("part 2: %d\n", part2(data))
}
