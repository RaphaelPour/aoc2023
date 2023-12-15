package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

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

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	init   bool
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

func (b Box) FocusPower(id int) int {
	result := 0
	for i, lens := range b.lenses {
		result += (id + 1) * (i + 1) * lens.focalLength
	}

	return result
}

type HASHMAP []Box

func (h HASHMAP) Add(label string, value int) {
	key := hash(label)
	box := h[key]
	if !box.init {
		box = Box{
			init:   true,
			lenses: make([]Lens, 0),
		}
	}

	box.Add(label, value)
	h[key] = box
}

func (h HASHMAP) Remove(label string) {
	key := hash(label)
	if !h[key].init {
		return
	}
	h[key].Remove(label)
}

func part1(data []string) int {
	result := 0
	for _, line := range data {
		result += hash(line)
	}
	return result
}

func part2(data []string) int {
	boxes := make(HASHMAP, 256)

	for _, entry := range data {
		index := strings.Index(entry, "=")
		if index == -1 {
			boxes.Remove(string(entry[:len(entry)-1]))
		} else {
			boxes.Add(string(entry[:index]), sstrings.ToInt(entry[index+1:]))
		}
	}

	result := 0
	for i, box := range boxes {
		if !box.init {
			continue
		}
		result += box.FocusPower(i)
	}
	return result
}

func main() {
	data := strings.Split(input.LoadString("input")[0], ",")

	fmt.Println("== [ PART 1 ] ==")
	start := time.Now()
	fmt.Println(part1(data), time.Since(start))

	fmt.Println("== [ PART 2 ] ==")
	start = time.Now()
	fmt.Println(part2(data), time.Since(start))
}
