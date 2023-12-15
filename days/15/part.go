package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern = regexp.MustCompile(`^(\w+)([=-])(\d+)$`)
)

func part1(data []string) int {
	result := 0
	for _, line := range data {
		currentValue := 0
		for _, r := range line {
			currentValue += int(r)
			currentValue *= 17
			currentValue = currentValue % 256
		}
		result += currentValue
	}
	return result
}

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	list []Lens
	keys map[string]struct{}
}

func part2(data []string) int {
	boxes := make([]Box, 0)

	for _, entry := range data {
		match := pattern.FindStringSubmatch(entry)
		if len(match) != 4 {
			fmt.Printf("error matching %s: %v\n", entry, match)
			return -1
		}

		label := match[1]
		value := sstrings.ToInt(match[3])
		if match[2] == "=" {

		}

	}
}

func main() {
	data := strings.Split(input.LoadString("input")[0], ",")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
