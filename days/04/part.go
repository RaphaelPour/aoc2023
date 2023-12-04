package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern = regexp.MustCompile(`^Card\s+(\d+): ([\d\s]+) \| ([\d\s]+)$`)
)

type Card struct {
	id        int
	win, have []int
}

func NewCard(in string) (Card, error) {
	match := pattern.FindStringSubmatch(in)
	if len(match) != 4 {
		return Card{}, fmt.Errorf("line has not enough matches: %s", match)
	}

	c := Card{}
	c.id = stellar_strings.ToInt(match[1])
	for _, num := range strings.Split(match[2], " ") {
		if num == "" {
			continue
		}
		c.win = append(c.win, stellar_strings.ToInt(num))
	}

	for _, num := range strings.Split(match[3], " ") {
		if num == "" {
			continue
		}
		c.have = append(c.have, stellar_strings.ToInt(num))
	}

	return c, nil
}

func (c Card) Points() (int, int) {
	points := 0
	count := 0
	for _, have := range c.have {
		for _, win := range c.win {
			if have != win {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
			count += 1

			break
		}
	}

	return points, count
}

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		card, err := NewCard(line)
		if err != nil {
			fmt.Printf("error parsing card %s: %s\n", line, err)
			return -1
		}
		points, _ := card.Points()
		sum += points
	}
	return sum
}

func part2(data []string) int {
	cache := make(map[int][]int)
	queue := make([]int, 0)

	for _, line := range data {
		card, err := NewCard(line)
		if err != nil {
			fmt.Printf("error parsing card %s: %s\n", line, err)
			return -1
		}

		_, count := card.Points()
		newCards := make([]int, 0)
		for i := 1; i <= count; i++ {
			newCards = append(newCards, card.id+i)
		}

		cache[card.id] = newCards
		queue = append(queue, newCards...)
		// fmt.Printf("%d: %d %v\n", card.id, count, newCards)
	}

	fmt.Println("---")
	for i := 0; i < len(queue); i++ {
		//fmt.Println(queue)
		if i%100000 == 0 {
			fmt.Printf("%d/%d\n", i, len(queue))
		}
		subCards, ok := cache[queue[i]]
		if !ok {
			continue
		}
		// fmt.Printf("%d: %d %v\n", queue[i], len(subCards), subCards)
		queue = append(queue, subCards...)
	}

	return len(queue) + len(data)
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
