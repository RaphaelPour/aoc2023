package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	smath "github.com/RaphaelPour/stellar/math"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

type Direction int

func ParseDirection(in string) Direction {
	return map[string]Direction{
		"R": RIGHT,
		"L": LEFT,
		"U": UP,
		"D": DOWN,
	}[in]
}

func (d Direction) Move(p smath.Point) smath.Point {
	return map[Direction]smath.Point{
		RIGHT: smath.Point{p.X + 1, p.Y},
		LEFT:  smath.Point{p.X - 1, p.Y},
		UP:    smath.Point{p.X, p.Y - 1},
		DOWN:  smath.Point{p.X, p.Y + 1},
	}[d]
}

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

var (
	EmptyAction  = Action{}
	EmptyDigPlan = DigPlan{}
)

type Action struct {
	dir    Direction
	length int
	color  string
}

func NewAction(in string) (Action, error) {
	a := Action{}
	parts := strings.Split(in, " ")
	if len(parts) != 3 {
		return EmptyAction, fmt.Errorf("error parsing input '%s': %s", in, parts)
	}
	raw := strings.Trim(parts[2], "()")[1:]
	val, err := strconv.ParseInt(raw[:5], 16, 64)
	if err != nil {
		return EmptyAction, err
	}
	a.length = int(val)
	a.dir = Direction(sstrings.ToInt(string(raw[5])))
	return a, nil
}

type DigPlan struct {
	actions  []Action
	visited  Cache
	interior int
	min, max smath.Point
}

func (d *DigPlan) Dig() {
	pos := smath.Point{0, 0}
	d.visited.Add(pos)
	for _, a := range d.actions {
		for i := 1; i <= a.length; i++ {
			pos = a.dir.Move(pos)
			d.visited.Add(pos)
			d.max = pos.Max(d.max)
			d.min = pos.Min(d.min)
			d.interior++
		}
	}

	d.max = d.max.Add(smath.Point{0, 1})
	area := d.max.X * d.max.Y

	fmt.Println(d.visited, d.interior, d.min, d.max, area)
}

func (d DigPlan) Print() {
	for y := d.min.Y; y < d.max.Y; y++ {
		for x := d.min.X; x <= d.max.X; x++ {
			if d.visited.Contains(smath.Point{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (d DigPlan) Area() int {
	sum := 0
	for y := d.min.Y; y < d.max.Y; y++ {
		for x := d.min.X; x <= d.max.X; x++ {
			found := false
			for offY := y - 1; !found && offY < y+1; offY++ {
				for offX := x - 1; !found && offX < x+1; offX++ {

					if d.visited.Contains(smath.Point{offX, offY}) {
						found = true
					}
				}
			}
		}
	}

	return sum
}

type Range struct {
	from, to smath.Point
}

func (r Range) IsSameXAxis(p smath.Point) bool {
	return (r.from.X == r.to.X && r.to.X == p.X)
}

func (r Range) IsSameYAxis(p smath.Point) bool {
	return (r.from.Y == r.to.Y && r.to.Y == p.Y)
}

func (r Range) IsYAxisNeighbor(p smath.Point) bool {
	return r.from.Y-1 == p.Y || r.to.Y+1 == p.Y
}

func (r Range) IsXAxisNeighbor(p smath.Point) bool {
	return r.from.X-1 == p.X || r.to.X+1 == p.X
}

func (r Range) AddXNeighbor(p smath.Point) Range {
	if r.from.X-1 == p.X {
		r.from.X--
	} else {
		r.to.X++
	}

	return r
}

func (r Range) AddYNeighbor(p smath.Point) Range {
	if r.from.Y-1 == p.Y {
		r.from.Y--
	} else {
		r.to.Y++
	}

	return r
}

func (r Range) Contains(p smath.Point) bool {
	return (r.IsSameXAxis(p) && p.Y >= r.from.Y && p.Y <= r.to.Y) ||
		(r.IsSameYAxis(p) && p.X >= r.from.X && p.X <= r.to.X)
}

func (r Range) Area() int {
	return (r.to.X + 1 - r.from.X) * (r.to.Y + 1 - r.from.Y)
}

type Cache struct {
	visited map[Range]struct{}
}

func (c Cache) Add(p smath.Point) {
	for r := range c.visited {
		if r.IsSameXAxis(p) && r.IsYAxisNeighbor(p) {
			delete(c.visited, r)
			c.visited[r.AddYNeighbor(p)] = struct{}{}
			return
		} else if r.IsSameYAxis(p) && r.IsXAxisNeighbor(p) {
			delete(c.visited, r)
			c.visited[r.AddXNeighbor(p)] = struct{}{}
			return
		}
	}
	c.visited[Range{p, p}] = struct{}{}
}

func (c Cache) Contains(p smath.Point) bool {
	for r := range c.visited {
		if r.Contains(p) {
			return true
		}
	}
	return false
}

func (c Cache) Area() int {
	sum := 0
	for r := range c.visited {
		sum += r.Area()
	}
	return sum
}

func (d *DigPlan) Fill(pos smath.Point) {
	if d.visited.Contains(pos) {
		return
	}

	d.visited.Add(pos)

	d.Fill(pos.Add(smath.Point{1, 0}))
	d.Fill(pos.Add(smath.Point{-1, 0}))
	d.Fill(pos.Add(smath.Point{0, 1}))
	d.Fill(pos.Add(smath.Point{0, -1}))
}

func NewDigPlan(in []string) (DigPlan, error) {
	p := DigPlan{}
	p.actions = make([]Action, len(in))
	p.visited = Cache{
		visited: make(map[Range]struct{}),
	}
	p.min = smath.Point{100, 100}

	for i := range in {
		action, err := NewAction(in[i])
		if err != nil {
			return EmptyDigPlan, err
		}
		p.actions[i] = action
	}
	return p, nil
}

func part1(data []string) int {
	plan, err := NewDigPlan(data)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	plan.Dig()
	plan.Fill(smath.Point{1, 1})
	plan.Print()
	return plan.visited.Area()
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
