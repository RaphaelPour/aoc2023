package main

import (
	"fmt"
	"time"

	"github.com/RaphaelPour/stellar/input"
)

type FieldType int

const (
	EMPTY FieldType = iota
	ROCK
	OBSTACLE
)

var (
	North = P{0, -1}
	West  = P{-1, 0}
	South = P{0, 1}
	East  = P{1, 0}

	Directions = []P{North, West, South, East}
)

func (f FieldType) String() string {
	return map[FieldType]string{
		EMPTY:    ".",
		ROCK:     "O",
		OBSTACLE: "#",
	}[f]
}

func ParseFieldType(r rune) FieldType {
	return map[rune]FieldType{
		'.': EMPTY,
		'O': ROCK,
		'#': OBSTACLE,
	}[r]
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

func (p P) Multiply(other P) P {
	p.x *= other.x
	p.y *= other.y
	return p
}

type Platform struct {
	fields [][]FieldType
	rocks  int
}

func (p Platform) Print() {
	for y := 0; y < len(p.fields); y++ {
		for x := 0; x < len(p.fields); x++ {
			fmt.Print(p.fields[y][x])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (p Platform) TiltOnce(dir P) int {
	changes := 0
	boundReached := 0
	rocks := 0
	for y := 0; y < len(p.fields); y++ {
		for x := 0; x < len(p.fields); x++ {
			off := dir.Add(P{x, y})
			if off.x < 0 || off.x >= len(p.fields[0]) || off.y < 0 || off.y >= len(p.fields) {
				continue
			}

			if p.fields[y][x] == ROCK && p.fields[off.y][off.x] == EMPTY {
				rocks++
				p.fields[y][x] = EMPTY
				p.fields[off.y][off.x] = ROCK
				changes++

				off = off.Multiply(P{2, 2})
				if off.x < 0 || off.x >= len(p.fields[0]) || off.y < 0 || off.y >= len(p.fields) {
					boundReached++
				}

				if rocks == p.rocks {
					return changes - boundReached
				}
			}
		}
	}
	return changes - boundReached
}

func (p Platform) Tilt(dir P) {
	i := 1
	for p.TiltOnce(dir) != 0 {
		/*fmt.Printf("====[ %d ]====\n", i)
		p.Print()*/
		i++
	}
}

func (p Platform) Round() {
	for _, dir := range Directions {
		p.Tilt(dir)
	}
}

func (p Platform) Load() int {
	load := 0
	for y := 0; y < len(p.fields); y++ {
		for x := 0; x < len(p.fields); x++ {
			if p.fields[y][x] != ROCK {
				continue
			}

			load += len(p.fields) - y
		}
	}
	return load
}

func NewPlatform(in []string) Platform {
	p := Platform{}
	p.fields = make([][]FieldType, len(in))

	for y, line := range in {
		p.fields[y] = make([]FieldType, len(line))
		for x, r := range line {
			p.fields[y][x] = ParseFieldType(r)
			if p.fields[y][x] == ROCK {
				p.rocks++
			}
		}
	}
	return p
}

func part1(data []string) int {
	p := NewPlatform(data)
	p.Tilt(North)
	return p.Load()
}

func part2(data []string) int {
	p := NewPlatform(data)
	start := time.Now()
	rounds := 1000000000
	for i := 1; i <= rounds; i++ {
		if i%10000 == 0 {
			fmt.Printf(
				"%f ETA %s\r",
				float64(i)/float64(rounds),
				time.Now().Add(
					time.Duration(int64(float64(i)/float64(time.Since(start).Seconds())*float64(rounds-i))),
				).Format(time.DateTime),
			)
		}
		p.Round()
	}
	return p.Load()
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
