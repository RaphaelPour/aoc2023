package main

import (
	"fmt"
	"regexp"

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

// https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// https://stackoverflow.com/a/3154503
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func part2(data []string) int {
	dir := Direction{
		directions: data[0],
	}
	nodes := make(Nodes)
	startNodes := make([]Node, 0)

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
		nodes[match[1]] = node

		if node.current[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	result := Search(startNodes[0], nodes, dir)
	for _, node := range startNodes[1:] {
		result = lcm(result, Search(node, nodes, dir))
	}
	return result
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
