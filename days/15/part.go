package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern = regexp.MustCompile(`^(\w+)([=-])(\d*)$`)
)

func hash(in string) int {
	v := 0
	for _, r := range in {
		v = ((v + int(r)) * 17) % 256
	}
	return v
}

func part1(data []string) int {
	result := 0
	for _, line := range data {
		result += hash(line)
	}
	return result
}

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	id     int
	lenses []Lens
}

func (b *Box) Add(label string, value int) {
	i := b.Index(label)
	if i == -1 {
		b.lenses = append(b.lenses, Lens{label, value})
		return
	}

	b.lenses[i].focalLength = value
}

func (b Box) Index(label string) int {
	for i, lens := range b.lenses {
		if lens.label == label {
			return i
		}
	}
	return -1
}

func (b *Box) Remove(label string) {
	i := b.Index(label)
	if i == -1 {
		return
	}

	b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
}

func (b Box) FocusPower() int {
	result := 0
	for i, lens := range b.lenses {
		result += (b.id + 1) * (i + 1) * lens.focalLength
	}

	return result
}

type HASHMAP map[int]Box

func (h HASHMAP) Add(label string, value int) {
	key := hash(label)
	box, ok := h[key]
	if !ok {
		box = Box{
			id:     key,
			lenses: make([]Lens, 0),
		}
		h[key] = box
	}

	box.Add(label, value)
	h[key] = box
}

func (h HASHMAP) Remove(label string) {
	key := hash(label)
	if box, ok := h[key]; ok {
		box.Remove(label)
		h[key] = box
	}
}

func part2(data []string) int {
	boxes := make(HASHMAP)

	for _, entry := range data {
		match := pattern.FindStringSubmatch(entry)
		if len(match) != 4 {
			fmt.Printf("error matching %s: %v\n", entry, match)
			return -1
		}

		label := match[1]
		operation := match[2]
		if operation == "=" {
			boxes.Add(label, sstrings.ToInt(match[3]))
		} else if operation == "-" {
			boxes.Remove(label)
		} else {
			fmt.Println("Unknown operation %s\n", operation)
			return -1
		}
	}

	result := 0
	for _, box := range boxes {
		result += box.FocusPower()
	}
	return result
}

func main() {
	data := strings.Split(input.LoadString("input")[0], ",")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
