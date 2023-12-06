package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/RaphaelPour/stellar/input"
	stellarStrings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern  = regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
	EmptyKey = Key{}
	EmptyMap = M{}
)

type Range struct {
	destinationStart int
	sourceStart      int
	length           int
}

func (r Range) project(in int) (int, bool) {
	// return input iteself it is out-of-range
	if in < r.sourceStart || in >= r.sourceStart+r.length {
		return -1, false
	}

	// map input onto [0,length]
	offset := in - r.sourceStart

	// apply offset onto destination
	return r.destinationStart + offset, true
}

type RangeRange []Range

func (rr RangeRange) project(in int) int {
	for _, r := range rr {
		if value, ok := r.project(in); ok {
			return value
		}
	}

	return in
}

type CacheKey struct {
	seed int
	key  Key
}

type Key struct {
	from, to string
}

type M struct {
	data  map[Key]RangeRange
	cache map[CacheKey]int
}

func NewMap(data []string) (M, error) {
	maps := M{
		data:  make(map[Key]RangeRange),
		cache: make(map[CacheKey]int),
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
			length:           stellarStrings.ToInt(rawNumbers[2]),
		}
		maps.data[currentKey] = append(maps.data[currentKey], r)
	}

	return maps, nil
}

func (m M) FindKey(from string) Key {
	for k := range m.data {
		if k.from == from {
			return k
		}
	}
	return EmptyKey
}

func (m M) String() string {
	out := ""
	for key, ranges := range m.data {
		out += fmt.Sprintf("%s-to-%s map:\n", key.from, key.to)
		for _, r := range ranges {
			out += fmt.Sprintf("%d %d %d\n", r.destinationStart, r.sourceStart, r.length)
		}
		out += "\n"
	}
	return out
}

func (m M) Find(seed int, fromKey, goalKey string, depth int) (int, error) {
	if fromKey == goalKey {
		return seed, nil
	}

	key := m.FindKey(fromKey)

	if key == EmptyKey {
		return -1, fmt.Errorf("error finding key for from key %s", fromKey)
	}

	/*
		if value, ok := m.cache[CacheKey{key: key, seed: seed}]; ok {
			return value, nil
		}
	*/

	min, err := m.Find(m.data[key].project(seed), key.to, goalKey, depth+1)
	if err != nil {
		return -1, err
	}
	//m.cache[CacheKey{key: key, seed: seed}] = min
	return min, nil
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

	min := -1
	for _, seed := range seeds {
		val, err := maps.Find(seed, "seed", "location", 0)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		if val < min || min == -1 {
			min = val
		}
	}

	return min
}

func work(start, length int, maps M) int {
	min := -1
	for j := start; j < start+length; j += 1 {
		val, err := maps.Find(j, "seed", "location", 0)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		if val < min || min == -1 {
			min = val
		}
	}
	return min
}

func part2(data []string) int {
	seedRanges := make([]int, 0)
	for _, rawNum := range strings.Split(data[0], " ")[1:] {
		seedRanges = append(seedRanges, stellarStrings.ToInt(rawNum))
	}

	maps, err := NewMap(data[2:])
	if err != nil {
		fmt.Println(err)
		return -1
	}

	total := 0
	for i := 0; i < len(seedRanges)/2; i++ {
		total += seedRanges[i*2+1]
	}

	var wg sync.WaitGroup
	var m sync.Mutex
	min := -1
	for i := 0; i < len(seedRanges); i += 2 {
		wg.Add(1)
		go func(start, length int) {
			defer wg.Done()
			fmt.Print(">")
			val := work(start, length, maps)

			m.Lock()
			defer m.Unlock()
			if val < min || min == -1 {
				min = val
			}
			fmt.Print("<")
		}(seedRanges[i], seedRanges[i+1])
	}

	wg.Wait()
	return min
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println("too high: 10834441")
	fmt.Println(part2(data))
}
