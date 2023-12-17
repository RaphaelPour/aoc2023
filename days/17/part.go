package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/stellar/input"
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

type LavaPool struct {
	fields  [][]int
	visited map[P]bool
	goal    P
}

func NewLavaPool(data [][]int) LavaPool {
	l := LavaPool{}
	l.fields = data
	l.visited = make(map[P]bool)
	l.goal = P{
		x: len(data[0]) - 1,
		y: len(data) - 1,
	}

	v := map[P]struct{}{
		l.goal: struct{}{},
	}
	l.Print(nil)
	for y := len(l.fields) - 1; y >= 0; y-- {
		for x := len(l.fields[0]) - 1; x >= 0; x-- {
			min := -1
			for _, p := range []P{P{-1, 0}, P{1, 0}, P{0, -1}, P{0, 1}} {
				p.x += x
				p.y += y
				if l.OutOfBounds(p) {
					continue
				}

				if _, ok := v[p]; ok {
					continue
				}

				if val := l.fields[p.y][p.x]; val < min || min == -1 {
					min = val
				}
				l.fields[p.y][p.x] += min
				v[p] = struct{}{}
			}
		}
	}
	l.Print(nil)

	return l
}

func (l LavaPool) OutOfBounds(p P) bool {
	return p.x < 0 || p.x >= len(l.fields[0]) || p.y < 0 || p.y >= len(l.fields)
}

func (l LavaPool) Search(loss int, current, previous P) ([]P, int, bool) {
	if current.Equal(l.goal) {
		return []P{current}, loss + l.fields[current.y][current.x], true
	}

	if _, visited := l.visited[current]; visited {
		return nil, -1, false
	}

	min := -1
	var candidate P
	var bestPath []P
	for _, rawNeighbor := range []P{P{-1, 0}, P{1, 0}, P{0, -1}, P{0, 1}} {
		neighbor := current.Add(rawNeighbor)
		if l.OutOfBounds(neighbor) {
			continue
		}
		if neighbor.Equal(current) {
			continue
		}
		l.visited[neighbor] = false

		path, result, found := l.Search(
			loss+l.fields[neighbor.y][neighbor.x],
			neighbor,
			current,
		)

		if found && (result < min || min == -1) {
			min = result
			bestPath = path
			candidate = neighbor
		}
	}

	l.visited[candidate] = true
	return append(bestPath, candidate), min, min != -1
}

func (l LavaPool) Print(path map[P]bool) {
	for y := range l.fields {
		for x := range l.fields[y] {
			if isPath, visited := path[P{x, y}]; visited {
				if isPath {
					fmt.Print("  #  ")
				} else {
					fmt.Print("  .  ")
				}
			} else {
				fmt.Printf("%5d", l.fields[y][x])
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func part1(data [][]int) int {
	_ = NewLavaPool(data)
	/*_, loss, found := l.Search(0, P{0, 0}, P{-1, -1})
	if !found {
		fmt.Println("not found")
		return -1
	}*/

	return -1
}

func part2(data []string) int {
	return 0
}

func LoadIntMap(filename string) [][]int {
	content := input.LoadString(filename)
	result := make([][]int, len(content))
	for y, row := range content {
		result[y] = make([]int, len(row))
		for x, col := range row {
			result[y][x], _ = strconv.Atoi(string(col))
		}
	}

	return result
}

func main() {
	data := LoadIntMap("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
