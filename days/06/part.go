package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

type Race struct {
	time, distance int
}

func (r Race) Permutations() int {
	ways := 0
	for holdTime := 1; holdTime < r.time; holdTime++ {
		if holdTime*(r.time-holdTime) > r.distance {
			ways++
		}
	}
	return ways
}

func ParseRaces1(data []string) []Race {
	races := make([]Race, 0)
	for _, rawNum := range strings.Split(strings.Split(data[0], ":")[1], " ") {
		if rawNum == "" {
			continue
		}
		races = append(races, Race{time: stellar_strings.ToInt(rawNum)})
	}

	i := 0
	for _, rawNum := range strings.Split(strings.Split(data[1], ":")[1], " ") {
		if rawNum == "" {
			continue
		}
		races[i].distance = stellar_strings.ToInt(rawNum)
		i++
	}

	return races
}

func ParseRaces2(data []string) Race {
	time := stellar_strings.ToInt(strings.ReplaceAll(strings.Split(data[0], ":")[1], " ", ""))
	distance := stellar_strings.ToInt(strings.ReplaceAll(strings.Split(data[1], ":")[1], " ", ""))
	return Race{
		time:     time,
		distance: distance,
	}
}

func part1(data []string) int {
	result := 1

	for _, race := range ParseRaces1(data) {
		result *= race.Permutations()
	}
	return result
}

func part2(data []string) int {
	return ParseRaces2(data).Permutations()
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
