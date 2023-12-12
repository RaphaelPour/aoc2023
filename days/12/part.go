package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type FieldType int

func (f FieldType) String() string {
	return map[FieldType]string{
		UNKNOWN:     "?",
		OPERATIONAL: ".",
		DAMAGED:     "#",
	}[f]
}

func ParseFieldType(in rune) FieldType {
	return map[rune]FieldType{
		'?': UNKNOWN,
		'.': OPERATIONAL,
		'#': DAMAGED,
	}[in]
}

const (
	UNKNOWN FieldType = iota
	OPERATIONAL
	DAMAGED
)

type Row struct {
	fields  string
	damaged []int
}

func NewRow(in string) Row {
	r := Row{}

	parts := strings.Split(in, " ")

	r.fields = parts[0]
	r.damaged = make([]int, 0)

	for _, damaged := range strings.Split(parts[1], ",") {
		r.damaged = append(r.damaged, s_strings.ToInt(damaged))
	}

	return r
}

func GoalReached(in string, goal []int) bool {
	if strings.Index(in, "#") >= 0 && len(goal) == 0 {
		return false
	}

	i := 0
	count := 0
	for _, field := range in {
		if field == '.' {
			if count > 0 {
				if i >= len(goal) || goal[i] != count {
					return false
				}
				i++

				count = 0
			}
			continue
		}

		if field == '#' {
			count++
			continue
		}

		// UNKNOWN
		return false
	}

	if count > 0 {
		return i == len(goal)-1 && goal[i] == count
	}

	return i == len(goal)
}

func Find(in string, goal []int, depth int) int {
	// fmt.Println("candidate:", in)

	if GoalReached(in, goal) {
		//fmt.Println(in)
		return 1
	}

	sum := 0
	idx := strings.Index(in, "?")
	if idx >= 0 {
		sum += Find(fmt.Sprintf("%s#%s", in[:idx], in[idx+1:]), goal, depth+1)
		sum += Find(fmt.Sprintf("%s.%s", in[:idx], in[idx+1:]), goal, depth+1)
	}

	return sum
}

type Springs struct {
	rows []Row
}

func part1(data []string) int {
	sum := 0
	start := time.Now()
	lap := time.Now()
	for i, line := range data {
		fmt.Println("======>", i+1, "/", len(data), "<======")
		row := NewRow(line)
		result := Find(row.fields, row.damaged, 0)
		newLap := time.Now()
		fmt.Printf("%s => %d in %.2fs (%.2f total)\n", line, result, newLap.Sub(lap).Seconds(), newLap.Sub(start).Seconds())
		lap = newLap

		sum += result
	}
	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("too low: 7395")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
