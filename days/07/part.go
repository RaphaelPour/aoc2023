package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
)

type HandType int

func max(a, b HandType) HandType {
	if a > b {
		return a
	}
	return b
}

const (
	HIGH_CARD_TYPE HandType = iota
	ONE_PAIR_TYPE
	TWO_PAIR_TYPE
	THREE_OF_A_KIND_TYPE
	FULL_HOUSE_TYPE
	FOUR_OF_A_KIND_TYPE
	FIVE_OF_A_KIND_TYPE
)

func (h HandType) String() string {
	return map[HandType]string{
		HIGH_CARD_TYPE:       "high card",
		ONE_PAIR_TYPE:        "one pair",
		TWO_PAIR_TYPE:        "two pair",
		THREE_OF_A_KIND_TYPE: "three of a kind",
		FULL_HOUSE_TYPE:      "full house",
		FOUR_OF_A_KIND_TYPE:  "four of a kind",
		FIVE_OF_A_KIND_TYPE:  "five of a kind",
	}[h]
}

var (
	cardOrderP1 = []rune("AKQJT98765432")
	cardOrderP2 = []rune("AKQT98765432J")
	//                    3210987654321
	//                    111
	joker = len(cardOrderP2) - 1
)

type Cards []int

func (c Cards) String() string {
	out := ""
	for _, card := range c {
		out += string(cardOrderP2[card])
	}
	return out
}

type Hand struct {
	id    string
	cards Cards
	bid   int
	kind  HandType
	joker bool
}

func (h Hand) String() string {
	return fmt.Sprintf("%s (%s) %d", h.id, h.kind, h.bid)
}

func NewHand(cards []rune, bid int, p2 bool) Hand {
	h := Hand{id: string(cards), bid: bid, joker: p2}
	var order []rune
	if p2 {
		order = cardOrderP2
	} else {
		order = cardOrderP1
	}

	crds := make(Cards, len(cards))
	for i, card1 := range cards {
		for j, card2 := range order {
			if card1 == card2 {
				crds[i] = j
			}
		}
	}

	h.cards = crds

	if p2 {
		h.kind = Perm(crds)
	} else {
		h.kind = Type(crds)
	}
	return h
}

func Perm(cards Cards) HandType {
	hist := make(map[int]int)
	for _, card := range cards {
		hist[card] = hist[card] + 1
	}

	if hist[joker] == 0 {
		return Type(cards)
	}

	kind := Type(cards)
	for i, card := range cards {
		if card != joker {
			continue
		}

		for card2 := range cardOrderP2 {
			if card2 == joker {
				continue
			}
			kind = max(kind, Perm(append(append(append(Cards{}, cards[:i]...), card2), cards[i+1:]...)))
		}

		if kind == FIVE_OF_A_KIND_TYPE {
			return kind
		}
	}

	return kind
}

func Type(cards Cards) HandType {
	hist := make(map[int]int)
	for _, card := range cards {
		hist[card] = hist[card] + 1
	}

	// check Hive
	for _, count := range hist {
		if count == 5 {
			return FIVE_OF_A_KIND_TYPE
		}
	}

	// check Four
	for _, count := range hist {
		if count == 4 {
			return FOUR_OF_A_KIND_TYPE
		}
	}

	// check full house
	pair := -1
	trio := -1
	for card, count := range hist {
		if count == 2 {
			pair = card
		} else if count == 3 {
			trio = card
		}
	}

	if pair >= 0 && trio >= 0 {
		return FULL_HOUSE_TYPE
	}

	// check three
	for _, count := range hist {
		if count == 3 {
			return THREE_OF_A_KIND_TYPE
		}
	}

	// check Two pairs
	pair1 := -1
	pair2 := -1
	for card, count := range hist {
		if count == 2 {
			if pair1 == -1 {
				pair1 = card
			} else {
				pair2 = card
				break
			}
		}
	}

	if pair1 >= 0 && pair2 >= 0 {
		return TWO_PAIR_TYPE
	}

	// check One Pair from previous check
	if pair1 >= 0 {
		return ONE_PAIR_TYPE
	}

	return HIGH_CARD_TYPE
}

type Hands []Hand

func (h Hands) Less(i, j int) bool {
	if h[i].kind == h[j].kind {
		for k := range h[i].cards {
			if h[i].cards[k] != h[j].cards[k] {
				return h[i].cards[k] > h[j].cards[k]
			}
		}
	}

	return h[i].kind < h[j].kind
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Len() int {
	return len(h)
}

func part1(data []string) int {
	hands := make(Hands, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		hands[i] = NewHand(
			[]rune(parts[0]),
			stellar_strings.ToInt(parts[1]),
			false,
		)
	}

	result := 0
	sort.Sort(hands)
	for i, hand := range hands {
		// fmt.Printf("%s * %d\n", hand, i+1)
		result += hand.bid * (i + 1)
	}

	return result
}

func part2(data []string) int {
	hands := make(Hands, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		hands[i] = NewHand(
			[]rune(parts[0]),
			stellar_strings.ToInt(parts[1]),
			true,
		)
	}

	result := 0
	sort.Sort(hands)
	for i, hand := range hands {
		// fmt.Printf("%s * %d\n", hand, i+1)
		result += hand.bid * (i + 1)
	}

	return result
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too low: 248662219")
	fmt.Println("    bad: ")
	fmt.Println("         247986162")
	fmt.Println("         248064906")
	fmt.Println("         248252141")
	fmt.Println("         248652582")
	fmt.Println("         248801590")
	fmt.Println("         248844779")
	fmt.Println("         249075763")
	fmt.Println("         249614258")
	fmt.Printf("   goal: %d\n", part2(data))
}
