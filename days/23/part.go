package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

var (
	dirMap = map[byte]P{
		'>': P{1, 0},
		'<': P{-1, 0},
		'^': P{0, -1},
		'v': P{0, 1},
	}
)

type P struct {
	x, y int
}

func (p P) Add(other P) P {
	p.x += other.x
	p.y += other.y
	return p
}

func (p P) Equal(other P) bool {
	return p.x == other.x && p.y == other.y
}

func (p P) String() string {
	return fmt.Sprintf("%d/%d", p.x, p.y)
}

type Visitor []P

func (v Visitor) Add(p P) Visitor {
	v2 := make(Visitor, len(v)+1)
	copy(v2, v)
	v2[len(v2)-1] = p
	return v2
}

func (v Visitor) Find(p P) bool {
	for _, visited := range v {
		if visited == p {
			return true
		}
	}
	return false
}

type Trail struct {
	fields       []string
	start, goal  P
	ignoreSlopes bool
}

func NewTrail(data []string, ignoreSlopes bool) Trail {
	t := Trail{
		fields:       data,
		start:        P{strings.Index(data[0], "."), 0},
		goal:         P{strings.Index(data[len(data)-1], "."), len(data) - 1},
		ignoreSlopes: ignoreSlopes,
	}

	if t.start.x == -1 {
		panic("start not found")
	} else if t.goal.x == -1 {
		panic("goal not found")
	}

	return t
}

func (t Trail) IsOutOfBounds(p P) bool {
	return p.x < 0 || p.x >= len(t.fields[0]) || p.y < 0 || p.y >= len(t.fields)
}

func (t *Trail) Find(current P, visited Visitor, depth int) (int, bool) {
	//fmt.Println(current)
	if current.Equal(t.goal) {
		return len(visited), true
	}

	if visited.Find(current) {
		//fmt.Println("hit", visited)
		return 0, false
	}

	visited = visited.Add(current)

	if !t.ignoreSlopes {
		if p, ok := dirMap[t.fields[current.y][current.x]]; ok {
			//fmt.Println("found slope")
			return t.Find(current.Add(p), visited, depth+1)
		}
	}

	maxLength := 0
	for _, p := range []P{P{-1, 0}, P{1, 0}, P{0, -1}, P{0, 1}} {
		//fmt.Println("before:", p)
		p = p.Add(current)
		//fmt.Println("after:", p)
		if t.IsOutOfBounds(p) {
			//fmt.Printf("neighbor %s out-of-bound\n", p)
			continue
		}

		if t.fields[p.y][p.x] == '#' {
			// fmt.Printf("neighbor %s is wall\n", p)
			continue
		}

		// fmt.Printf("visit neighbor %s\n", p)
		if length, ok := t.Find(p, visited, depth+1); ok && length > maxLength {
			maxLength = length
		}
	}
	return maxLength, maxLength > 0
}

func part1(data []string) int {
	t := NewTrail(data, false)
	length, ok := t.Find(t.start, Visitor{}, 0)
	if !ok {
		fmt.Println("no path found")
	}
	return length
}

func part2(data []string) int {
	t := NewTrail(data, true)
	length, ok := t.Find(t.start, Visitor{}, 0)
	if !ok {
		fmt.Println("no path found")
	}
	return length
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
