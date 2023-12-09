package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

func deviate(in []int) ([]int, bool) {
	zero := true
	derivative := make([]int, len(in)-1)
	for i := 0; i < len(in)-1; i++ {
		d := in[i+1] - in[i]
		if d != 0 {
			zero = false
		}
		derivative[i] = d
	}

	return derivative, zero
}

func part1(data [][]int) int {
	result := 0
	for _, line := range data {
		result += line[len(line)-1]
		dev := line
		var ok bool
		for {
			dev, ok = deviate(dev)
			if ok {
				break
			}
			result += dev[len(dev)-1]
		}
	}

	return result
}

func part2(data [][]int) int {
	result := 0
	for _, line := range data {
		dev := line
		var zero bool
		minus := 1
		result += line[0]
		for {
			dev, zero = deviate(dev)
			if zero {
				break
			}
			result -= dev[0] * minus
			minus *= -1
		}
	}

	return result
}

func main() {
	data := input.LoadIntTable("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too high: 1109690227")
	fmt.Println(part2(data))
}
