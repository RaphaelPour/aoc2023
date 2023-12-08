package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dougwatson/Go/v3/math/lcm"

	"github.com/RaphaelPour/stellar/input"
)

var (
	pattern = regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)
)

type Node struct {
	left, right, current string
}

type Nodes map[string]Node

type Direction struct {
	pos        int
	directions string
}

func (d *Direction) Next() rune {
	r := rune(d.directions[d.pos%len(d.directions)])
	d.pos++
	return r
}

func Search(current Node, nodes Nodes, dir Direction) int {
	if current.current[2] == 'Z' {
		return 0
	}

	var next Node
	if dir.Next() == 'L' {
		next = nodes[current.left]
	} else {
		next = nodes[current.right]
	}

	if next == current {
		return 0
	}

	return 1 + Search(next, nodes, dir)
}

var p2CurrentNodes []Node
var p2Nodes Nodes
var p2Dir Direction

func Search2() int {
	for {
		d := p2Dir.Next()
		if p2Dir.pos%1000000 == 0 {
			fmt.Println(p2Dir.pos, p2CurrentNodes)
		}
		goalReached := true
		for i, n := range p2CurrentNodes {
			// skip nodes that already have reached their goal
			if n.current[2] != 'Z' {
				goalReached = false
			}

			var nextNode Node
			if d == 'L' {
				nextNode = p2Nodes[n.left]
			} else {
				nextNode = p2Nodes[n.right]
			}

			// skip loops
			if nextNode == n {
				return 0
			}

			p2CurrentNodes[i] = nextNode

			// nextNodes = append(nextNodes, nextNode)
		}

		if goalReached {
			return p2Dir.pos - 1
		}
	}
}

func part1(data []string) int {
	directions := Direction{
		directions: data[0],
	}
	nodes := make(Nodes)

	for _, line := range data[2:] {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 4 {
			fmt.Printf("error matching %s: %s\n", line, match)
			return -1
		}

		nodes[match[1]] = Node{
			current: match[1],
			left:    match[2],
			right:   match[3],
		}
	}

	return Search(nodes["AAA"], nodes, directions)
}

func part2(data []string) int {
	p2Dir = Direction{
		directions: data[0],
	}
	p2Nodes = make(Nodes)
	p2CurrentNodes = make([]Node, 0)

	for _, line := range data[2:] {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 4 {
			fmt.Printf("error matching %s: %s\n", line, match)
			return -1
		}

		node := Node{
			current: match[1],
			left:    match[2],
			right:   match[3],
		}
		p2Nodes[match[1]] = node

		if strings.HasSuffix(node.current, "A") {
			p2CurrentNodes = append(p2CurrentNodes, node)
		}
	}

	result := 0
	for _, node := range p2CurrentNodes {
		lcm.Lcm(int64(result), int64(Search(node, p2Nodes, p2Dir)))
	}
	return result
}

func main() {
	data := input.LoadString("input3")

	//fmt.Println("== [ PART 1 ] ==")
	//fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too low: 12026000000")
	fmt.Println(part2(data))
}
