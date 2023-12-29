package main

import (
	"fmt"
	"time"

	"github.com/RaphaelPour/stellar/input"
)

type Row []bool

type Pattern struct {
	fields []Row
}

func (r Row) String() string {
	out := ""
	for _, cell := range r {
		if cell {
			out += "#"
		} else {
			out += "."
		}
	}
	return out
}

func (p Pattern) Print(axisX, axisY int) {
	fmt.Print(" ")
	for x := 0; x < len(p.fields[0]) && axisX != 0; x++ {
		if x == axisX {
			fmt.Print(">")
		} else if x == axisX+1 {
			fmt.Print("<")
		} else {
			fmt.Print((x + 1) % 10)
		}
	}
	fmt.Println("")

	for y := 0; y < len(p.fields); y++ {
		if y == axisY && axisY > 0 {
			fmt.Print("V")
		} else if y == axisY+1 && axisY > 0 {
			fmt.Print("^")
		} else {
			fmt.Print((y + 1) % 10)
		}

		for x := 0; x < len(p.fields[y]); x++ {
			if p.fields[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		if y == axisY && axisY > 0 {
			fmt.Print("V")
		} else if y == axisY+1 && axisY > 0 {
			fmt.Print("^")
		} else {
			fmt.Print((y + 1) % 10)
		}
		fmt.Println("")
	}

	fmt.Print(" ")
	for x := 0; x < len(p.fields[0]) && axisX != 0; x++ {
		if x == axisX {
			fmt.Print(">")
		} else if x == axisX+1 {
			fmt.Print("<")
		} else {
			fmt.Print((x + 1) % 10)
		}
	}
	fmt.Println("")
}

func (p *Pattern) AddRow(in string) {
	row := make(Row, len(in))
	for i, r := range in {
		row[i] = (r == '#')
	}

	p.fields = append(p.fields, row)
}

func (p Pattern) FindMirrorAxis(smudged bool) (int, int) {
	for y := 0; y < len(p.fields)-1; y++ {
		if p.IsYMirrorAxis(y, smudged) {
			return -1, y
		}
	}
	for x := 0; x < len(p.fields[0])-1; x++ {
		if p.IsXMirrorAxis(x, smudged) {
			return x, -1
		}
	}
	return -1, -1
}

func (p Pattern) IsYMirrorAxis(y int, smudged bool) bool {
	if y < 0 || y >= len(p.fields) {
		return false
	}

	yLow := y
	yHigh := y + 1
	diff := 0
	for yLow >= 0 && yHigh < len(p.fields) {
		for x := 0; x < len(p.fields[yLow]); x++ {
			if p.fields[yLow][x] != p.fields[yHigh][x] {
				diff++
			}
		}

		yLow--
		yHigh++
	}

	if smudged {
		return diff == 1
	}
	return diff == 0
}

func (p Pattern) IsXMirrorAxis(x int, smudged bool) bool {
	if x < 0 || x >= len(p.fields[0]) {
		return false
	}

	xLow := x
	xHigh := x + 1
	diff := 0
	for xLow >= 0 && xHigh < len(p.fields[0]) {
		for y := 0; y < len(p.fields); y++ {
			if p.fields[y][xLow] != p.fields[y][xHigh] {
				diff++
			}
		}

		xLow--
		xHigh++
	}
	if smudged {
		return diff == 1
	}
	return diff == 0
}

func NewPatterns(in []string) []Pattern {
	patterns := make([]Pattern, 0)

	currentPattern := Pattern{}
	currentPattern.fields = make([]Row, 0)
	for _, line := range in {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = Pattern{}
			currentPattern.fields = make([]Row, 0)
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
	sum := 0
	for _, pattern := range patterns {
		x, y := pattern.FindMirrorAxis(false)
		sum += x + 1 + (y+1)*100
	}
	return sum
}

func part2(data []string) int {
	patterns := NewPatterns(data)
	sum := 0
	for _, pattern := range patterns {
		x, y := pattern.FindMirrorAxis(true)
		sum += x + 1 + (y+1)*100
	}
	return sum
}

func main() {
	data := input.LoadString("input")

	start := time.Now()
	fmt.Println("== [ PART 1 ] ==")
	fmt.Printf("%d (%s)\n", part1(data), time.Since(start))

	start = time.Now()
	fmt.Println("== [ PART 2 ] ==")
	fmt.Printf("%d (%s)\n", part2(data), time.Since(start))
}
