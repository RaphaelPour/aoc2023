package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	DOWN
	LEFT
	UP
)

var (
	EmptyAction  = Action{}
	EmptyDigPlan = DigPlan{}
)

type Action struct {
	dir    Direction
	length int
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

	// fmt.Printf("%s => %d %d\n", raw, a.dir, a.length)

	return a, nil
}

type DigPlan struct {
	actions  []Action
	visited  map[smath.Point]struct{}
	outer    map[smath.Point]struct{}
	interior int
	min, max smath.Point
}

func (d *DigPlan) Dig() {
	pos := smath.Point{0, 0}
	d.visited[pos] = struct{}{}
	for _, a := range d.actions {
		for i := 1; i <= a.length; i++ {
			pos = a.dir.Move(pos)
			d.visited[pos] = struct{}{}
			d.max = pos.Max(d.max)
			d.min = pos.Min(d.min)
			d.interior++
		}
	}

	d.max = d.max.Add(smath.Point{0, 1})
	//area := d.max.X * d.max.Y

	// fmt.Println(d.visited, d.interior, d.min, d.max, area)
}

func (d DigPlan) Print() {
	for y := d.min.Y; y < d.max.Y; y++ {
		for x := d.min.X; x <= d.max.X; x++ {
			if _, ok := d.visited[smath.Point{x, y}]; ok {
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

					if _, ok := d.visited[smath.Point{offX, offY}]; ok {
						found = true
					}
				}
			}
		}
	}

	return sum
}

func (d *DigPlan) Fill(pos smath.Point) {
	if _, ok := d.visited[pos]; ok {
		return
	}
	if _, ok := d.outer[pos]; ok {
		return
	}

	d.outer[pos] = struct{}{}
	d.Fill(pos.Add(smath.Point{1, 0}))
	d.Fill(pos.Add(smath.Point{-1, 0}))
	d.Fill(pos.Add(smath.Point{0, 1}))
	d.Fill(pos.Add(smath.Point{0, -1}))
}

func (d *DigPlan) Count() int {
	inside := false
	sum := 0
	start := time.Now()
	for y := d.min.Y; y < d.max.Y; y++ {
		if y%10 == 0 {
			fmt.Printf("\r%f %d %s", 100.0/float64(d.max.Y)*float64(y), sum, time.Since(start))
		}
		for x := d.min.X; x <= d.max.X; x++ {
			if _, ok := d.visited[smath.Point{x, y}]; ok {
				inside = !inside
			} else if inside {
				sum++
			}
		}
	}
	return sum + len(d.visited)
}

func NewDigPlan(in []string) (DigPlan, error) {
	p := DigPlan{}
	p.actions = make([]Action, len(in))
	p.visited = make(map[smath.Point]struct{})
	p.outer = make(map[smath.Point]struct{})
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
	fmt.Println("INIT")
	plan, err := NewDigPlan(data)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println("NEW")
	plan.Dig()
	fmt.Println("DIG")

	for y := plan.min.Y; y <= plan.max.Y; y++ {
		plan.Fill(smath.Point{plan.min.X, y})
		plan.Fill(smath.Point{plan.max.X, y})
	}
	fmt.Println("FILL Y")
	for x := plan.min.X; x <= plan.max.X; x++ {
		plan.Fill(smath.Point{x, plan.min.Y})
		plan.Fill(smath.Point{x, plan.max.Y})
	}

	fmt.Println("FILL")
	//plan.Print()
	fmt.Println("PRINT")
	return len(plan.visited)
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
