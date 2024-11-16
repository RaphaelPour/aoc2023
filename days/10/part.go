package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

var (
	print = true
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
	FILLED
	INSIDE
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
		FILLED:           "#",
		INSIDE:           "I",
	}[p]
}

func (p Pipe) Neighbor(x, y int) bool {
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

type Map struct {
	w, h    int
	fields  [][]Pipe
	visited map[math.Point]struct{}
	start   math.Point
}

func (m Map) String() string {
	out := ""
	for y := range m.fields {
		for x := range m.fields[y] {
			out += m.fields[y][x].String()
		}
		out += "\n"
	}
	return out
}

func (m Map) PrintMap(pos math.Point) {
	if !print {
		return
	}
	posColor := color.New(color.FgHiRed)
	dotColor := color.New(color.FgBlack, color.Bold)
	pipeColor := color.New(color.FgGreen)
	fillColor := color.New(color.FgBlue, color.Bold)
	for y := range m.fields {
		for x := range m.fields[y] {
			if pos.X == x && pos.Y == y {
				posColor.Print("x")
				/*else if m.start.X == x && m.start.Y == y {
					fmt.Print("S")
				} */
			} else if m.fields[y][x] != EMPTY_PIPE {
				pipeColor.Print(m.fields[y][x])
			} else if m.fields[y][x] == FILLED {
				fillColor.Print("#")
			} else {
				dotColor.Print(m.fields[y][x])
			}
		}
		fmt.Println("")
	}
}

func Search(pos, from math.Point, m Map) ([]math.Point, bool) {
	//fmt.Println("pos: ", pos)
	//m.PrintMap(pos)
	m.visited[pos] = struct{}{}

	currentField := m.fields[pos.Y][pos.X]
	if currentField == START_PIPE && from.X > -1 {
		return []math.Point{pos}, true
	}

	for y := -1; y <= 1; y += 1 {
		for x := -1; x <= 1; x += 1 {
			if x == 0 && y == 0 {
				continue
			}

			if x != 0 && y != 0 {
				continue
			}

			next := pos.Add(math.Point{x, y})
			if next.X < 0 || next.Y < 0 || next.X >= m.w || next.Y >= m.h {
				continue
			}

			if from.Equal(next) {
				continue
			}

			nextField := m.fields[next.Y][next.X]
			if nextField == START_PIPE {
				return []math.Point{pos}, true
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

func (m Map) Fill(start math.Point, loop map[math.Point]struct{}) {
	if start.X < 0 || start.X >= m.w || start.Y < 0 || start.Y >= m.h {
		return
	}

	if m.fields[start.Y][start.X] == FILLED {
		return
	}

	if _, ok := loop[math.Point{start.X, start.Y}]; ok {
		return
	}

	m.fields[start.Y][start.X] = FILLED
	m.Fill(start.Add(math.Point{0, 1}), loop)
	m.Fill(start.Add(math.Point{0, -1}), loop)
	m.Fill(start.Add(math.Point{1, 0}), loop)
	m.Fill(start.Add(math.Point{-1, 0}), loop)
}

func Tile(fields [][]Pipe, x, y int) [][]Pipe {
	p := fields[y][x]
	switch p {
	case NORTH_NORTH_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
		}
	case NORTH_EAST_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, NORTH_EAST_PIPE, EAST_WEST_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
		}
	case EAST_WEST_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EAST_WEST_PIPE, EAST_WEST_PIPE, EAST_WEST_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
		}
	case NORTH_WEST_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
			{EAST_WEST_PIPE, NORTH_WEST_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
		}
	case SOUTH_WEST_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EAST_WEST_PIPE, SOUTH_WEST_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
		}
	case SOUTH_EAST_PIPE:
		return [][]Pipe{
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, SOUTH_EAST_PIPE, EAST_WEST_PIPE},
			{EMPTY_PIPE, NORTH_NORTH_PIPE, EMPTY_PIPE},
		}
	case START_PIPE:
		pipes := [][]Pipe{
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, START_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
		}
		if y-1 >= 0 {
			if p2 := fields[y-1][x]; p2 == NORTH_NORTH_PIPE || p2 == SOUTH_WEST_PIPE || p2 == SOUTH_EAST_PIPE {
				pipes[0][1] = NORTH_NORTH_PIPE
			}
		}

		if y+1 < len(fields) {
			if p2 := fields[y+1][x]; p2 == NORTH_NORTH_PIPE || p2 == NORTH_WEST_PIPE || p2 == NORTH_EAST_PIPE {
				pipes[2][1] = NORTH_NORTH_PIPE
			}
		}

		if x-1 >= 0 {
			if p2 := fields[y][x-1]; p2 == EAST_WEST_PIPE || p2 == NORTH_EAST_PIPE || p2 == SOUTH_EAST_PIPE {
				pipes[1][0] = EAST_WEST_PIPE
			}
		}

		if x+1 < len(fields[0]) {
			if p2 := fields[y][x+1]; p2 == EAST_WEST_PIPE || p2 == NORTH_WEST_PIPE || p2 == SOUTH_WEST_PIPE {
				pipes[1][2] = EAST_WEST_PIPE
			}
		}
		return pipes
	default:
		return [][]Pipe{
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
			{EMPTY_PIPE, EMPTY_PIPE, EMPTY_PIPE},
		}
	}
}

func (m *Map) Expand() {
	fields := make([][]Pipe, len(m.fields)*3)
	for y := 0; y < len(m.fields)*3; y++ {
		fields[y] = make([]Pipe, len(m.fields[0])*3)
	}

	for y := range m.fields {
		for x := range m.fields[y] {
			tile := Tile(m.fields, x, y)
			for y1 := 0; y1 < 3; y1++ {
				for x1 := 0; x1 < 3; x1++ {
					fields[3*y+y1][3*x+x1] = tile[y1][x1]
				}
			}
		}
	}

	visited := make(map[math.Point]struct{})
	for p := range m.visited {
		tile := Tile(m.fields, p.Y, p.X)
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				if tile[y][x] == EMPTY_PIPE {
					continue
				}

				visited[math.Point{p.X + x + 1, p.Y + y + 1}] = struct{}{}
			}
		}
	}

	//visited[m.start] = struct{}{}
	m.visited = visited

	m.fields = fields
	// add one, as the start point is in the middle of the tile
	m.start = math.Point{m.start.X*3 + 1, m.start.Y*3 + 1}
	fmt.Println(m.start)
	m.h = len(m.fields)
	m.w = len(m.fields[0])
}

func (m *Map) Clean(path map[math.Point]struct{}) {
	for y := range m.fields {
		for x := range m.fields[y] {
			if m.fields[y][x] == EMPTY_PIPE {
				continue
			}

			if _, ok := path[math.Point{x, y}]; !ok {
				m.fields[y][x] = EMPTY_PIPE
			}
		}
	}
}

func NewMap(in []string) Map {
	m := Map{}
	m.fields = make([][]Pipe, len(in))
	m.visited = make(map[math.Point]struct{})
	m.h = len(in)
	m.w = len(in[0])

	for y, line := range in {
		m.fields[y] = make([]Pipe, len(line))
		for x, field := range line {
			pipe := ParseField(field)
			if pipe == START_PIPE {
				m.start = math.Point{x, y}
			}

			m.fields[y][x] = pipe
		}
	}

	return m
}

func part1(data []string) int {
	m := NewMap(data)
	path, ok := Search(m.start, math.Point{-1, -1}, m)
	if !ok {
		fmt.Println("no path found")
		return -1
	}

	return int(float64(len(path))/2) + 1
}

func path2Map(path []math.Point) map[math.Point]struct{} {
	m := make(map[math.Point]struct{})
	for _, p := range path {
		m[p] = struct{}{}
	}
	return m
}

func part2(data []string) int {
	m := NewMap(data)

	m.Expand()
	path, ok := Search(m.start, math.Point{-1, -1}, m)
	if !ok {
		fmt.Println("no path found")
		return -1
	}

	pathMap := path2Map(path)
	m.Clean(pathMap)
	for y := 0; y < m.h; y++ {
		m.Fill(math.Point{0, y}, pathMap)
		m.Fill(math.Point{m.w - 1, y}, pathMap)
	}

	for x := 0; x < m.w; x++ {
		m.Fill(math.Point{x, 0}, pathMap)
		m.Fill(math.Point{x, m.h - 1}, pathMap)
	}

	m.PrintMap(math.Point{-1, -1})

	// Count filled
	sum := 0
	for y := 1; y < len(m.fields); y += 3 {
		for x := 1; x < len(m.fields[0]); x += 3 {
			if m.fields[y][x] == EMPTY_PIPE {
				m.fields[y][x] = INSIDE
				sum++
			}
		}
	}
	m.PrintMap(math.Point{-1, -1})
	return sum
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("too low: 7101")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too low: 236")
	fmt.Println("bad: 222, 371, 968, 1068")
	fmt.Println("too high: 1019, 1472")
	fmt.Println(part2(data))
}
