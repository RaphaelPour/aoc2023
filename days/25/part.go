package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

type Component struct {
	name string
	adj  map[string]struct{}
}

func (c Component) String() string {
	return fmt.Sprintf("%s -> %v", c.name, c.adj)
}

type Diagram struct {
	components map[string]Component
	candidates []string
}

func NewDiagram(data []string) Diagram {
	d := Diagram{}
	d.components = make(map[string]Component)
	for _, line := range data {
		parts := strings.Split(line, " ")
		name := parts[0][:3]
		c := Component{
			name: name,
			adj:  map[string]struct{}{},
		}

		for _, n := range parts[1:] {
			c.adj[n] = struct{}{}
		}
		d.components[name] = c
	}

	for name, c := range d.components {
		if name == "" {
			fmt.Println("init err")
		}
		for neighbor := range c.adj {
			if neighbor == "" {
				fmt.Println("init err2")
			}
			c2 := d.components[neighbor]
			if len(c2.adj) == 0 {
				c2.adj = make(map[string]struct{})
			}
			c2.adj[name] = struct{}{}
			d.components[neighbor] = c2
		}
	}

	d.candidates = make([]string, 0)
	for _, c := range d.components {
		d.candidates = append(d.candidates, c.name)
	}

	return d
}

func (d Diagram) Find(current, goal string, visited, hist map[string]int) bool {
	if current == goal {
		// fmt.Printf("%q", current)
		hist[current] = hist[current] + 1
		return true
	}

	if current == "" {
		fmt.Println("DONG DONG DONG")
	}
	visited[current] = 0

	for n := range d.components[current].adj {
		if _, ok := visited[n]; ok {
			continue
		}
		if d.Find(n, goal, visited, hist) {
			return true
		}
	}
	return false
}

type Element struct {
	name   string
	visits int
}

func (e Element) String() string {
	return fmt.Sprintf("%d %s", e.visits, e.name)
}

type Elements struct {
	list []Element
}

func (e Elements) Len() int {
	return len(e.list)
}

func (e Elements) Less(i, j int) bool {
	return e.list[i].visits < e.list[j].visits
}

func (e Elements) Swap(i, j int) {
	e.list[i], e.list[j] = e.list[j], e.list[i]
}

func ElementsFromHist(hist map[string]int) Elements {
	e := Elements{
		list: make([]Element, len(hist)),
	}
	i := 0
	for name, visits := range hist {
		if name == "" {
			fmt.Println("DING DING DING")
		}
		e.list[i] = Element{name: name, visits: visits}
		i++
	}

	sort.Sort(e)

	return e
}

func (d Diagram) FindAll() map[string]int {
	hist := make(map[string]int)

	// start := time.Now()

	for i := 0; i < len(d.candidates)-1; i++ {
		/*
			fmt.Printf(
				"\r%d/%d (%f) ETA %.2fs",
				i,
				len(d.candidates),
				100.0/float64(len(d.candidates))*float64(i),
				time.Since(start).Seconds()/float64(i)*float64((len(d.candidates)-i)),
			)*/
		for j := i + 1; j < len(d.candidates); j++ {
			d.Find(
				d.candidates[i],
				d.candidates[j],
				map[string]int{},
				hist,
			)
		}
	}
	fmt.Println("")
	return hist
}

func part1(data []string) int {
	d := NewDiagram(data)
	hist := d.FindAll()

	e := ElementsFromHist(hist)
	for _, e := range e.list[0:7] {
		fmt.Println(e)
	}

	return len(hist)
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
