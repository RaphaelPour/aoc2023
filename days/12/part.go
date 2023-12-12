package main

import (
	"fmt"
	"strings"

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
	fields  []FieldType
	damaged []int
}

func NewRow(in string) Row {
	r := Row{}

	parts := strings.Split(in, " ")

	r.fields = make([]FieldType, len(parts[0]))
	r.damaged = make([]int, len(parts[1])/2+1)

	for i, field := range parts[0] {
		r.fields[i] = ParseFieldType(field)
	}

	for i, damaged := range strings.Split(parts[1], ",") {
		r.damaged[i] = s_strings.ToInt(damaged)
	}

	return r
}

func GoalReached(in []FieldType, goal []int) bool {
	i := 0
	count := 0
	for _, field := range in {
		if field == OPERATIONAL {
			if count > 0 {
				if i >= len(goal) || goal[i] != count {
					return false
				}
				i++

				count = 0
			}
			continue
		}

		if field == DAMAGED {
			count++
			continue
		}

		// UNKNOWN
		return false
	}

	if count > 0 {
		return i == len(goal)-1 && goal[i] == count
	}

	return i >= len(goal)
}

func Find(in []FieldType, goal []int, visited map[string]struct{}, depth int) int {
	if _, ok := visited[fmt.Sprintf("%s", in)]; ok {
		return 0
	}
	visited[fmt.Sprintf("%s", in)] = struct{}{}
	// fmt.Println("candidate:", in)

	if GoalReached(in, goal) {
		//fmt.Println(in)
		return 1
	}

	sum := 0
	for i := range in {
		if in[i] != UNKNOWN {
			continue
		}

		tmp := append([]FieldType{}, in[:i]...)
		sum += Find(append(append(tmp, DAMAGED), in[i+1:]...), goal, visited, depth+1)
		sum += Find(append(append(tmp, OPERATIONAL), in[i+1:]...), goal, visited, depth+1)
	}
	return sum
}

type Springs struct {
	rows []Row
}

func part1(data []string) int {
	sum := 0
	for i, line := range data {
		fmt.Println("======>", i+1, "/", len(data), "<======")
		row := NewRow(line)
		result := Find(row.fields, row.damaged, map[string]struct{}{}, 0)
		fmt.Println(line, "=>", result)
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
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
