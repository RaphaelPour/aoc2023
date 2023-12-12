package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	s_math "github.com/RaphaelPour/stellar/math"
)

type Map struct {
	w, h   int
	fields [][]bool
	emptyX map[int]struct{}
	emptyY map[int]struct{}
}

func NewMap(data []string) Map {
	m := Map{}
	m.fields = make([][]bool, len(data))
	m.emptyX = make(map[int]struct{})
	m.emptyY = make(map[int]struct{})
	m.h = len(data)
	m.w = len(data[0])

	for y, line := range data {
		m.fields[y] = make([]bool, len(data))
		foundGalaxie := false
		for x, field := range line {
			m.fields[y][x] = (field == '#')
			if field == '#' {
				foundGalaxie = true
			}
		}

		if !foundGalaxie {
			m.emptyX[y] = struct{}{}
		}
	}

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

	return m
}

func (m Map) FindGalaxies() []P {
	galaxies := make([]P, 0)
	// find galaxies
	for y := range m.fields {
		for x := range m.fields[y] {
			if m.fields[y][x] {
				galaxies = append(galaxies, P{x, y})
			}
		}
	}

	return galaxies
}

func (m Map) Sum(expansion int) int {
	galaxies := m.FindGalaxies()

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].Dist(galaxies[j])

			for y := range m.emptyX {
				if within(y, galaxies[i].y, galaxies[j].y) {
					sum += expansion
				}
			}

			for x := range m.emptyY {
				if within(x, galaxies[i].x, galaxies[j].x) {
					sum += expansion
				}
			}
		}
	}
	return sum
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
	return s_math.Abs(p.x-other.x) + s_math.Abs(p.y-other.y)
}

func (p P) Equal(other P) bool {
	return p.x == other.x && p.y == other.y
}

func part1(data []string) int {
	return NewMap(data).Sum(1)
}

func within(n, b1, b2 int) bool {
	return (n >= b1 && n <= b2) || (n >= b2 && n <= b1)
}

func part2(data []string) int {
	return NewMap(data).Sum(1000000 - 1)
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
