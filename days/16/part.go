package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Direction int

func (d Direction) TurnClockwise() Direction {
	return (d + 1) % 4
}

func (d Direction) TurnAntiClockwise() Direction {
	return (d + 3) % 4
}

/*    ^      v
 *    ^      v
 * >>>/   <<</
 */
func (d Direction) Rotate(mirror string) Direction {
	if mirror == "/" {
		switch d {

		case LEFT:
			return DOWN
		case DOWN:
			return LEFT
		case RIGHT:
			return UP
		case UP:
			return RIGHT
		}
	} else if mirror == `\` {
		switch d {
		case LEFT:
			return UP
		case DOWN:
			return RIGHT
		case RIGHT:
			return DOWN
		case UP:
			return LEFT
		}
	}
	return -1
}

const (
	LEFT Direction = iota
	DOWN
	RIGHT
	UP
)

var (
	dirMap = map[Direction]P{
		LEFT:  P{-1, 0, 0},
		DOWN:  P{0, 1, 0},
		RIGHT: P{1, 0, 0},
		UP:    P{0, -1, 0},
	}
)

type P struct {
	x, y int
	dir  Direction
}

func (p P) Mask() P {
	p.dir = -1
	return p
}

func (p P) String() string {
	return fmt.Sprintf("%d,%d %d", p.x, p.y, p.dir)
}

func (p P) Add(other P) P {
	p.x += other.x
	p.y += other.y
	return p
}

func (p P) Next() P {
	return p.Add(dirMap[p.dir])
}

func PrintMap(data []string, visited map[P]struct{}) {
	for y := range data {
		for x := range data[y] {
			if _, ok := visited[P{x, y, -1}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

type Beams struct {
	visited       map[P]struct{}
	points        []P
	buffer        []P
	width, height int
}

func NewBeams(width, height int) Beams {
	start := P{x: 0, y: 0, dir: RIGHT}

	visited := Beams{}
	visited.visited = map[P]struct{}{start.Mask(): struct{}{}}
	visited.points = []P{start}
	visited.buffer = make([]P, 0)
	visited.width = width
	visited.height = height
	return visited
}

func (b Beams) Empty() bool {
	return len(b.points) == 0
}

func (b *Beams) Add(p P) {
	if p.x < 0 || p.x > b.width || p.y < 0 || p.y > b.height {
		fmt.Println(p, "OoB")
		return
	}

	if _, ok := b.visited[p.Mask()]; ok {
		fmt.Println(p, "visited")
		return
	}

	fmt.Println("add", p)

	b.buffer = append(b.buffer, p)
	b.Visit(p)
}

func (b Beams) Visit(p P) {
	if p.x < 0 || p.x > b.width || p.y < 0 || p.y > b.height {
		fmt.Println(p, "map OoB")
		return
	}
	b.visited[p.Mask()] = struct{}{}
}

func (b *Beams) Commit() {
	b.points = b.buffer
	b.buffer = make([]P, 0)
}

func part1(data []string) int {
	beams := NewBeams(len(data[0]), len(data))
	for !beams.Empty() {
		for _, beam := range beams.points {
			next := beam.Next()
			if next.x < 0 || next.x >= len(data[0]) || next.y < 0 || next.y >= len(data) {
				continue
			}

			field := string(data[next.y][next.x])
			fmt.Println("next field:", field)
			switch field {
			case "|": // split or pass-through
				if beam.dir == UP || beam.dir == DOWN {
					// just follow the direction
					beams.Add(next)
				} else {
					beams.Visit(next)

					up := next.Add(dirMap[UP])
					up.dir = UP
					beams.Add(up)

					down := next.Add(dirMap[DOWN])
					down.dir = DOWN
					beams.Add(down)
				}
			case "-":
				if beam.dir == LEFT || beam.dir == RIGHT {
					// just follow the direction
					beams.Add(next)
				} else {
					beams.Visit(next)

					left := next.Add(dirMap[LEFT])
					left.dir = LEFT
					beams.Add(left)

					right := next.Add(dirMap[RIGHT])
					right.dir = RIGHT
					beams.Add(right)
				}
			case `\`, "/": // rotate based on mirror
				beam.dir = beam.dir.Rotate(field)
				beams.buffer = append(beams.buffer, beam)
			case ".": // just follow the direction
				beams.Add(next)
			default:
				fmt.Printf("Unkown field %s\n", field)
				return -1
			}
		}

		beams.Commit()
	}

	PrintMap(data, beams.visited)
	fmt.Println(beams.visited)

	return len(beams.visited)
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
