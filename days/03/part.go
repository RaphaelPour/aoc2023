package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

type CellKind int

const (
	CELL_KIND_UNSPECIFIED CellKind = iota
	CELL_KIND_NUMBER
	CELL_KIND_SYMBOL
	CELL_KIND_EMPTY
)

type Cell struct {
	buffer    string
	num       int
	kind      CellKind
	connected bool
	coord     Point
}

func NewCell(in string, coord Point) Cell {
	k := kind(in)
	var num int
	if k == CELL_KIND_NUMBER {
		num = stellar_strings.ToInt(in)
	}

	return Cell{
		buffer:    in,
		num:       num,
		kind:      k,
		connected: k == CELL_KIND_SYMBOL,
		coord:     coord,
	}
}

func kind(in string) CellKind {
	if in == "." {
		return CELL_KIND_EMPTY
	}

	if rune(in[0]) >= '0' && rune(in[0]) <= '9' {
		return CELL_KIND_NUMBER
	}

	return CELL_KIND_SYMBOL
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

func (p Point) Dist(other Point) int {
	return math.Abs(p.x-other.x) + math.Abs(p.y-other.y)
}

type Grid struct {
	data       map[Point]Cell
	maxX, maxY int
}

func (g Grid) String() string {
	out := ""
	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxX; x++ {
			cell, ok := g.data[Point{x, y}]
			if !ok {
				out += "."
			} else {
				out += cell.buffer
			}
		}
		out += "\n"
	}
	return out
}

func (g Grid) Get(p Point) Cell {
	cell, ok := g.data[p]
	if !ok {
		cell = Cell{}
	}
	return cell
}

func (g *Grid) Set(x, y int, cell Cell) {
	if x > g.maxX {
		g.maxX = x
	}

	if y > g.maxY {
		g.maxY = y
	}
	g.data[Point{x, y}] = cell
}

func (g Grid) HasConnections(p, origin Point, depth int) bool {
	if g.Get(p).kind == CELL_KIND_SYMBOL {
		fmt.Printf("%s %s SYMB\n", strings.Repeat(" ", depth), p)
		return true
	}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			cell := g.Get(p.Add(Point{x, y}))

			if cell.coord == origin {
				// don't go back where we came from
				continue
			}

			if cell.kind == CELL_KIND_SYMBOL {
				fmt.Printf("%s %s Neighbor SYMB\n", strings.Repeat(" ", depth), p)
				return true
			}

			// only allow positive number nieghbor lookup so numbers
			// don't acknoledge each other
			/*if x < 0 || y < 0 {
				continue
			}*/

			if cell.kind == CELL_KIND_NUMBER {
				fmt.Printf("%s %s check number neighbor\n", strings.Repeat(" ", depth), p)
				con := g.HasConnections(cell.coord, p, depth+1)
				if con {
					fmt.Printf("%s %s deep neighbor SYMB\n", strings.Repeat(" ", depth), p)
					return con
				}
			}
		}
	}

	fmt.Printf("%s %s nope\n", strings.Repeat(" ", depth), p)
	return false
}

func (g Grid) FindMissing() int {
	cells := make([]Cell, 0)

	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxX; x++ {
			cell, ok := g.data[Point{x, y}]
			if !ok || cell.kind != CELL_KIND_NUMBER {
				continue
			}
			fmt.Println(cell.buffer)
			if g.HasConnections(Point{x, y}, Point{x, y}, 0) {
				cells = append(cells, cell)
				fmt.Print(">", cell.buffer, "<\n")
			}
		}
	}

	sum := 0
	buffer := cells[0].buffer
	for i, cell := range cells[1:] {
		fmt.Printf("x=%d y=%d\n", cell.coord.x, cell.coord.y)
		if cell.coord.y != cells[i].coord.y ||
			math.Abs(cell.coord.x-cells[i].coord.x) > 1 {
			num := stellar_strings.ToInt(buffer)
			fmt.Println(num)
			sum += num
			buffer = ""
		}

		buffer += cell.buffer
	}
	num := stellar_strings.ToInt(buffer)
	fmt.Println(num)
	sum += num

	return sum
}

func part1(data []string) int {
	grid := Grid{make(map[Point]Cell), -1, -1}

	for y, line := range data {
		for x, r := range line {
			ch := string(r)
			grid.Set(x, y, NewCell(ch, Point{x, y}))
		}
	}

	fmt.Println(grid)

	return grid.FindMissing()
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
