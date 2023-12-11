package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	s_math "github.com/RaphaelPour/stellar/math"
)

type Map struct {
	w, h     int
	fields   [][]bool
	emptyX   []int
	emptyY   []int
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
	return s_math.Abs(p.x-other.x) + s_math.Abs(p.y-other.y)
}

func (p P) Equal(other P) bool {
	return p.x == other.x && p.y == other.y
}

func part1(data []string) int {
	m := Map{}
	m.fields = make([][]bool, len(data))
	m.emptyX = make([]int, 0)
	m.h = len(data)
	m.w = len(data[0])
	m.galaxies = make([]P, 0)

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
			m.emptyX = append(m.emptyX, y)
		}
	}

	for i := len(m.emptyX) - 1; i >= 0; i-- {
		y := m.emptyX[i]
		m.h++
		m.fields = append(append(m.fields[:y], m.fields[y]), m.fields[y:]...)
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
			m.emptyY = append(m.emptyY, x)
		}
	}

	for i := len(m.emptyY) - 1; i >= 0; i-- {
		x := m.emptyY[i]
		m.w++
		for y := len(m.fields) - 1; y >= 0; y-- {
			m.fields[y] = append(append(m.fields[y][:x], false), m.fields[y][x:]...)
		}
	}

	// find galaxies
	for y := range m.fields {
		for x := range m.fields[y] {
			if m.fields[y][x] {
				m.galaxies = append(m.galaxies, P{x, y})
			}
		}
	}

	sum := 0
	for i := 0; i < len(m.galaxies)-1; i++ {
		for j := i + 1; j < len(m.galaxies); j++ {
			sum += m.galaxies[i].Dist(m.galaxies[j])
		}
	}

	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
