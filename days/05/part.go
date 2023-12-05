package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellarStrings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern  = regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
	EmptyKey = Key{}
	EmptyMap = M{}
	Cache    = map[CacheKey]int{}
)

type Range struct {
	destinationStart int
	sourceStart      int
	_range           int
}

type CacheKey struct {
	seed int
	key  Key
}

type Key struct {
	from, to string
}

type M struct {
	data map[Key][]Range
}

func (m M) FindKey(from string) Key {
	for k := range m.data {
		if k.from == from {
			return k
		}
	}
	return EmptyKey
}

func (m M) Find(seed int, fromKey string) (int, error) {
	key := m.FindKey(fromKey)
	if key == EmptyKey {
		return -1, fmt.Errorf("error finding key for from key %s", fromKey)
	}

	if value, ok := Cache[CacheKey{key: key, seed: seed}]; ok {
		return value, nil
	}

	min := 10000000
	var err error
	for _, r := range m.data[key] {
		if seed < r.sourceStart || seed > r.sourceStart+r._range {
			if seed < min {
				min = seed
			}
			continue
		}

		for i := 1; i < r._range; i++ {
			value := 0
			source := seed - r.sourceStart
			destination := r.destinationStart + source + i
			// fmt.Println(seed, r, source, destination)

			value = destination
			if key.to != "location" {
				value, err = m.Find(value, key.to)
				if err != nil {
					return -1, err
				}
			}
			if value < min {
				min = value
			}

		}
	}

	Cache[CacheKey{key: key, seed: seed}] = min
	return min, nil
}

func NewMap(data []string) (M, error) {
	maps := M{
		data: make(map[Key][]Range),
	}
	currentKey := EmptyKey
	for _, line := range data {
		if line == "" {
			currentKey = EmptyKey
			continue
		}

		if currentKey == EmptyKey {
			match := pattern.FindStringSubmatch(line)
			if len(match) != 3 {
				return EmptyMap, fmt.Errorf("error matching line %s: %s", line, match)
			}

			currentKey.from = match[1]
			currentKey.to = match[2]
			maps.data[currentKey] = make([]Range, 0)
			continue
		}

		rawNumbers := strings.Split(line, " ")
		if len(rawNumbers) != 3 {
			return EmptyMap, fmt.Errorf(
				"error parsing numbers, expected 3 got %s",
				line,
			)
		}

		r := Range{
			destinationStart: stellarStrings.ToInt(rawNumbers[0]),
			sourceStart:      stellarStrings.ToInt(rawNumbers[1]),
			_range:           stellarStrings.ToInt(rawNumbers[2]),
		}
		maps.data[currentKey] = append(maps.data[currentKey], r)
	}

	return maps, nil
}

func part1(data []string) int {
	seeds := make([]int, 0)
	for _, rawNum := range strings.Split(data[0], " ")[1:] {
		seeds = append(seeds, stellarStrings.ToInt(rawNum))
	}

	maps, err := NewMap(data[2:])
	if err != nil {
		fmt.Println(err)
		return -1
	}

	min := 100
	for _, seed := range seeds {
		val, err := maps.Find(seed, "seed")
		if err != nil {
			fmt.Println(err)
			return -1
		}
		if val < min {
			min = val
		}
	}

	return min
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
