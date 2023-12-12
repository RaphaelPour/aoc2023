package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Pipe int

const (
	EMPTY_PIPE Pipe = iota // .
	START_PIPE
	NORTH_NORTH_PIPE // |
	NORTH_EAST_PIPE  // L
	EAST_WEST_PIPE   //  -
	NORTH_WEST_PIPE  //   J
	SOUTH_WEST_PIPE  //   7
	SOUTH_EAST_PIPE  //   F
)

func (p Pipe) String() string {
	return map[Pipe]string{
		EMPTY_PIPE:       ".",
		START_PIPE:       "S",
		NORTH_NORTH_PIPE: "|",
		NORTH_EAST_PIPE:  "L",
		EAST_WEST_PIPE:   "-",
		NORTH_WEST_PIPE:  "J",
		SOUTH_WEST_PIPE:  "7",
		SOUTH_EAST_PIPE:  "F",
	}[p]
}

func (p Pipe) Neighbor(x, y int) bool {
	fmt.Println(P{x, y}, p)
	if p == START_PIPE {
		return true
	}
	// |
	// S
	// |
	if x == 0 && p == NORTH_NORTH_PIPE {
		return true
	}

	// -S-
	if y == 0 && p == EAST_WEST_PIPE {
		return true
	}
	// F 7
	// S S
	if x == 0 && y < 0 && (p == SOUTH_EAST_PIPE || p == SOUTH_WEST_PIPE) {
		return true
	}

	// S S
	// L J
	if x == 0 && y > 0 && (p == NORTH_EAST_PIPE || p == NORTH_WEST_PIPE) {
		return true
	}

	// FS
	//    LS
	if y == 0 && x < 0 && (p == SOUTH_EAST_PIPE || p == NORTH_EAST_PIPE) {
		return true
	}

	// SJ
	//    S7
	if y == 0 && x > 0 && (p == NORTH_WEST_PIPE || p == SOUTH_WEST_PIPE) {
		return true
	}

	return false
}

func ParseField(in rune) Pipe {
	return map[rune]Pipe{
		'.': EMPTY_PIPE,
		'S': START_PIPE,
		'|': NORTH_NORTH_PIPE,
		'L': NORTH_EAST_PIPE,
		'-': EAST_WEST_PIPE,
		'J': NORTH_WEST_PIPE,
		'7': SOUTH_WEST_PIPE,
		'F': SOUTH_EAST_PIPE,
	}[in]
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

func (p P) Min(other P) P {
	if p.x > other.x {
		p.x = other.x
	}

	if p.y > other.y {
		p.y = other.y
	}
	return p
}

func (p P) Max(other P) P {
	if p.x < other.x {
		p.x = other.x
	}

	if p.y < other.y {
		p.y = other.y
	}
	return p
}

type Map struct {
	w, h    int
	fields  [][]Pipe
	visited map[P]struct{}
	start   P
}

func (m Map) PrintMap(pos P) {
	for y := range m.fields {
		for x := range m.fields[y] {
			if pos.x == x && pos.y == y {
				fmt.Print("x")
			} else {
				fmt.Print(m.fields[y][x])
			}
		}
		fmt.Println("")
	}
}

func Search(pos, from P, m Map) ([]P, bool) {
	// fmt.Println("pos: ", pos)
	// m.PrintMap(pos)
	m.visited[pos] = struct{}{}

	currentField := m.fields[pos.y][pos.x]
	if currentField == START_PIPE && from.x > -1 {
		return []P{pos}, true
	}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			if x != 0 && y != 0 {
				continue
			}

			next := pos.Add(P{x: x, y: y})
			if next.x < 0 || next.y < 0 || next.x >= m.w || next.y >= m.h {
				continue
			}

			if from.Equal(next) {
				continue
			}

			nextField := m.fields[next.y][next.x]
			if nextField == START_PIPE {
				return []P{pos}, true
			}

			if _, alreadyVisited := m.visited[next]; alreadyVisited {
				continue
			}
			if nextField.Neighbor(x, y) {
				if path, ok := Search(next, pos, m); ok {
					return append(path, pos), true
				}
			}
		}
	}
	return nil, false
}

func NewMap(in []string) Map {
	m := Map{}
	m.fields = make([][]Pipe, len(in))
	m.visited = make(map[P]struct{})
	m.h = len(in)
	m.w = len(in[0])

	for y, line := range in {
		m.fields[y] = make([]Pipe, len(line))
		for x, field := range line {
			pipe := ParseField(field)
			if pipe == START_PIPE {
				m.start = P{x: x, y: y}
			}

			m.fields[y][x] = pipe
		}
	}

	return m
}

func part1(data []string) int {
	m := NewMap(data)
	path, ok := Search(m.start, P{-1, -1}, m)
	if !ok {
		fmt.Println("no path found")
		return -1
	}

	return (len(path) / 2) - 1
}

func part2(data []string) int {
	m := NewMap(data)
	path, ok := Search(m.start, P{-1, -1}, m)
	if !ok {
		fmt.Println("no path found")
		return -1
	}

	min := P{100, 100}
	max := P{}
	for _, p := range path {
		min = min.Min(p)
		max = max.Max(p)
	}

	pipeMap := make(map[P]struct{})
	for _, p := range path {
		pipeMap[p] = struct{}{}
	}

	sum := 0
	for y := min.y; y <= max.y; y++ {
		inside := false
		for x := min.x; x <= max.x; x++ {
			if _, ok := pipeMap[P{x, y}]; ok {
				fmt.Print(".")
				inside = !inside
				continue
			}

			if inside && y != min.y && y != max.y {
				fmt.Print("I")
				sum++
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println("")
	}
	return sum
}

func main() {
	data := input.LoadString("input3")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("too low: 7101")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too high: 1472")
	fmt.Println(part2(data))
}
