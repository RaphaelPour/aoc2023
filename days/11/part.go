package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Map struct {
	w, h     int
	fields   [][]bool
	emptyX   map[int]struct{}
	emptyY   map[int]struct{}
	galaxies []P
}

func (m Map) PrintMap() {
	for y := range m.fields {
		for x := range m.fields[y] {
			if m.fields[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

type P struct {
	x, y int
}

func (p P) String() string {
	return fmt.Sprintf("%d/%d", p.x, p.y)
}

func (p P) Add(other P) P {
	p.x += other.x
	p.y += other.y
	return p
}

func (p P) Dist(other P) int {
	return (p.x - other.x) + (p.y - other.y)
}

func (p P) Equal(other P) bool {
	return p.x == other.x && p.y == other.y
}

func part1(data []string) int {
	m := Map{}
	m.fields = make([][]bool, len(data))
	m.emptyX = make(map[int]struct{})
	m.h = len(data)
	m.w = len(data[0])
	m.galaxies = make([]P, 0)

	for y, line := range data {
		m.fields[y] = make([]bool, len(data))
		foundGalaxie := false
		for x, field := range line {
			m.fields[y][x] = (field == '#')
			if field == '#' {
				m.galaxies = append(m.galaxies, P{x, y})
				foundGalaxie = true
			}
		}

		if !foundGalaxie {
			m.emptyX[y] = struct{}{}
		}
	}

	m.PrintMap()

	for y := range m.emptyX {
		m.fields = append(append(m.fields[:y], m.fields[y]), m.fields[y:]...)
	}

	m.PrintMap()

	// check empty y lines
	for x := 0; x < m.w; x++ {
		foundGalaxie := false
		for y := 0; y < m.h; y++ {
			if m.fields[y][x] {
				foundGalaxie = true
				break
			}
		}
		if !foundGalaxie {
			m.emptyY[x] = struct{}{}
		}
	}

	for x := range m.emptyY {
		for y := len(m.fields) - 1; y > 0; y-- {
			m.fields[y] = append(append(m.fields[y][:x], false), m.fields[y][x+1:]...)
		}
	}
	m.PrintMap()

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
