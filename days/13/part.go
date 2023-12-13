package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

const (
	ASH  = false
	ROCK = true
)

type Pattern struct {
	fields [][]bool
}

func (p Pattern) Print() {
	for y := 0; y < len(p.fields); y++ {
		for x := 0; x < len(p.fields[y]); x++ {
			if p.fields[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (p *Pattern) AddRow(in string) {
	row := make([]bool, len(in))
	for i, r := range in {
		row[i] = (r == '#')
	}

	p.fields = append(p.fields, row)
}

func (p Pattern) IsYMirrorAxis(y int) bool {
	if y < 1 || y >= len(p.fields) {
		return false
	}
	return false
}

func NewPatterns(in []string) []Pattern {
	patterns := make([]Pattern, 0)

	currentPattern := Pattern{}
	currentPattern.fields = make([][]bool, 0)
	for _, line := range in {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = Pattern{}
			currentPattern.fields = make([][]bool, 0)
			continue
		}

		currentPattern.AddRow(line)
	}

	// add last pattern if input doesn't end with an empty line
	if len(currentPattern.fields) > 0 {
		patterns = append(patterns, currentPattern)
	}

	return patterns
}

func part1(data []string) int {
	patterns := NewPatterns(data)
	for _, pattern := range patterns {
	}
	return 0
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
