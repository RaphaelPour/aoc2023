package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

type HandType int

const (
	HIGH_CARD_TYPE HandType = iota
	ONE_PAIR_TYPE
	TWO_PAIR_TYPE
	THREE_OF_A_KIND_TYPE
	FULL_HOUSE_TYPE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var (
	kinds = []string{
		"2", "3", "4", "5", "6", "7", "8", "9",
		"T", "J", "Q", "K", "A",
	}
)

type Hand struct {
	cards string
	bid   int
}

func (h Hand) Type() HandType {
	t := HIGH_CARD_TYPE

}

func part1(data []string) int {
	hands := make([]Hand, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		hands[i] = Hand{
			cards: parts[0],
			bid:   stellar_strings.ToInt(parts[1]),
		}
	}

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
