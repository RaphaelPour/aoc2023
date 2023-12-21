package main

import (
	"fmt"
	"time"

	"github.com/RaphaelPour/stellar/input"
)

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

func (p P) Modulo(mod P) P {
	p.x = (p.x%mod.x + mod.x) % mod.x
	p.y = (p.y%mod.y + mod.y) % mod.y
	return p
}

type Garden struct {
	rocks      map[P]struct{}
	start      P
	dimensions P
}

func NewGarden(data []string) Garden {
	rocks := make(map[P]struct{})
	start := P{}
	for y, line := range data {
		for x, field := range line {
			if field == '#' {
				rocks[P{x, y}] = struct{}{}
			} else if field == 'S' {
				start = P{x, y}
			}
		}
	}

	return Garden{
		rocks:      rocks,
		start:      start,
		dimensions: P{len(data[0]), len(data)},
	}
}

func (g Garden) AllSteps() int {
	positions := map[P]struct{}{
		g.start: struct{}{},
	}

	start := time.Now()
	for i := 1; i <= 5000; i++ {
		newPos := make(map[P]struct{})

		if i%100 == 0 {
			fmt.Printf("\r%s", time.Since(start))
		}

		for field := range positions {
			for _, neighbor := range []P{
				P{-1, 0}, P{1, 0}, P{0, -1}, P{0, 1},
			} {
				p := field.Add(neighbor)
				if _, isRock := g.rocks[p.Modulo(g.dimensions)]; isRock {
					continue
				}
				newPos[p] = struct{}{}
			}
		}
		positions = newPos
	}
	fmt.Println("")
	return len(positions)
}

func part1(data []string) int {
	g := NewGarden(data)
	fmt.Println(g.dimensions)
	return g.AllSteps()
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
