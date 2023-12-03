package main

import (
	"fmt"

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
	buffer      string
	num         int
	kind        CellKind
	connected   bool
	coord       Point
	symbolCoord Point
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

func (g Grid) HasConnections(p, origin Point, depth int) (Point, bool) {
	if g.Get(p).kind == CELL_KIND_SYMBOL {
		// fmt.Printf("%s %s SYMB\n", strings.Repeat(" ", depth), p)
		return p, true
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
				// fmt.Printf("%s %s Neighbor SYMB\n", strings.Repeat(" ", depth), p)
				return cell.coord, true
			}

			// only allow positive number nieghbor lookup so numbers
			// don't acknoledge each other
			/*if x < 0 || y < 0 {
				continue
			}*/

			if cell.kind == CELL_KIND_NUMBER {
				// fmt.Printf("%s %s check number neighbor\n", strings.Repeat(" ", depth), p)
				symbolCoord, con := g.HasConnections(cell.coord, p, depth+1)
				if con {
					// fmt.Printf("%s %s deep neighbor SYMB\n", strings.Repeat(" ", depth), p)
					return symbolCoord, con
				}
			}
		}
	}

	//  fmt.Printf("%s %s nope\n", strings.Repeat(" ", depth), p)
	return Point{}, false
}

func (g Grid) FindMissing() (int, int) {
	cells := make([]Cell, 0)

	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxX; x++ {
			cell, ok := g.data[Point{x, y}]
			if !ok || cell.kind != CELL_KIND_NUMBER {
				continue
			}
			// fmt.Println(cell.buffer)
			if symbolCoord, ok := g.HasConnections(Point{x, y}, Point{x, y}, 0); ok {
				cell.symbolCoord = symbolCoord
				cells = append(cells, cell)
				fmt.Print(">", cell.buffer, " ", symbolCoord, "<\n")
			}
		}
	}

	symbolCache := make(map[Point][]int)

	sum := 0
	buffer := cells[0].buffer
	for i, cell := range cells[1:] {
		// fmt.Printf("x=%d y=%d %s\n", cell.coord.x, cell.coord.y, cell.symbolCoord)
		if cell.coord.y != cells[i].coord.y ||
			math.Abs(cell.coord.x-cells[i].coord.x) > 1 {
			num := stellar_strings.ToInt(buffer)
			fmt.Println(num)
			sum += num
			buffer = ""

			if _, ok := symbolCache[cells[i].symbolCoord]; !ok {
				symbolCache[cells[i].symbolCoord] = make([]int, 0)
			}

			symbCell, ok := g.data[cells[i].symbolCoord]
			fmt.Println(num, symbCell.buffer, cells[i].symbolCoord)
			if ok && symbCell.buffer == "*" {
				symbolCache[cells[i].symbolCoord] = append(symbolCache[cells[i].symbolCoord], num)
			}
		}

		buffer += cell.buffer
	}
	num := stellar_strings.ToInt(buffer)
	fmt.Println(num)
	sum += num

	if _, ok := symbolCache[cells[len(cells)-1].symbolCoord]; !ok {
		symbolCache[cells[len(cells)-1].symbolCoord] = make([]int, 0)
	}
	symbolCache[cells[len(cells)-1].symbolCoord] = append(symbolCache[cells[len(cells)-1].symbolCoord], num)

	sum2 := 0
	fmt.Println(symbolCache)
	for _, numbers := range symbolCache {
		if len(numbers) != 2 {
			continue
		}
		sum2 += numbers[0] * numbers[1]
	}

	return sum, sum2
}

func part1(data []string) (int, int) {
	grid := Grid{make(map[Point]Cell), -1, -1}

	for y, line := range data {
		for x, r := range line {
			ch := string(r)
			if ch == "." {
				continue
			}
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
	p1, p2 := part1(data)
	fmt.Println(p1)

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(p2)
}
