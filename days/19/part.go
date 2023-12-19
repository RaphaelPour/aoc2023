package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
	"github.com/fatih/color"
)

var (
	xmasMap = map[string]int{
		"x": 0,
		"m": 1,
		"a": 2,
		"s": 3,
	}

	goodColor = color.New(color.FgGreen)
	badColor  = color.New(color.FgRed)
	goalColor = color.New(color.FgBlue)
	sep       = "|  "
)

type Rule struct {
	isTerminal   bool
	hasOperation bool
	operand1     string
	larger       bool
	operand2     int
	production   string
}

func (r Rule) String() string {
	if !r.hasOperation {
		return fmt.Sprintf("else -> %s", r.production)
	}
	if r.larger {
		return fmt.Sprintf("%sϵ(%d,4000] -> %s", r.operand1, r.operand2, r.production)
	}
	return fmt.Sprintf("%sϵ[1,%d) -> %s", r.operand1, r.operand2, r.production)
}

func NewRules(in string) []Rule {
	rules := make([]Rule, 0)
	for _, cond := range strings.Split(in, ",") {
		r := Rule{}
		parts := strings.Split(cond, ":")

		if len(parts) == 1 {
			r.production = cond
		} else {
			r.hasOperation = true
			r.operand1 = string(parts[0][0])
			r.larger = (parts[0][1] == '>')
			r.operand2 = sstrings.ToInt(parts[0][2:])
			r.production = parts[1]
		}
		r.isTerminal = (strings.ToUpper(r.production) == r.production)
		rules = append(rules, r)
	}
	return rules
}

func (r Rule) Check(val int) bool {
	if !r.hasOperation {
		return false
	}

	if r.operand2 == val {
		return false
	}

	/*
		op := "<"
		if r.larger {
			op = ">"
		}
	*/

	//fmt.Printf("%d %s %d = %t\n", val, op, r.operand2, (r.larger == (val > r.operand2)))
	return r.larger == (val > r.operand2)
}

type Rating struct {
	data map[string]int
}

func NewRating(in string) Rating {
	r := Rating{}
	r.data = make(map[string]int)

	for _, rating := range strings.Split(in[1:len(in)-1], ",") {
		r.data[string(rating[0])] = sstrings.ToInt(rating[2:])
	}

	return r
}

func (r Rating) Sum() int {
	sum := 0
	for _, val := range r.data {
		sum += val
	}
	return sum
}

type RuleSet struct {
	rules   map[string][]Rule
	ratings []Rating
}

func NewRuleset(in []string) RuleSet {
	ruleset := RuleSet{}
	ruleset.rules = make(map[string][]Rule)
	ruleset.ratings = make([]Rating, 0)

	i := 0
	for _, line := range in {
		i++
		if line == "" {
			break
		}
		curly := strings.Index(line, "{")
		ruleset.rules[line[:curly]] = NewRules(line[curly+1 : len(line)-1])
	}

	for _, line := range in[i:] {
		ruleset.ratings = append(ruleset.ratings, NewRating(line))
	}

	return ruleset
}

func (r *RuleSet) EvalAll() int {
	sum := 0
	for _, rating := range r.ratings {
		if _, ok := r.Eval(rating); ok {
			sum += rating.Sum()
		}
	}
	return sum
}

func (r *RuleSet) Eval(rating Rating) ([]string, bool) {
	rule := "in"
	path := make([]string, 0)
	for {
		path = append(path, rule)
		for _, cond := range r.rules[rule] {
			if cond.hasOperation && !cond.Check(rating.data[cond.operand1]) {
				continue
			}

			if cond.isTerminal {
				return append(path, cond.production), cond.production == "A"
			}

			rule = cond.production
			break
		}
	}
	return path, false
}

type Interval struct {
	min, max int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d,%d]", i.min, i.max)
}

func (i Interval) MinMax(other Interval) Interval {
	if other.min > i.min {
		i.min = other.min
	}

	if other.max < i.max {
		i.max = other.max
	}

	return i
}

func (i Interval) Span() int {
	if i.max < i.min {
		return 0
	}
	return i.max - i.min
}

type Combination [4]Interval

func (c Combination) String() string {
	return fmt.Sprintf(
		"x%s m%s a%s s%s",
		c[0], c[1], c[2], c[3],
	)
}

func (c Combination) MinMax(other Combination) Combination {
	newC := Combination{}
	for i := range c {
		newC[i] = c[i].MinMax(other[i])
	}
	return newC
}

func (c Combination) Count() int {
	count := 1
	for i := range c {
		count *= c[i].Span()
	}
	return count
}

func (c Combination) Apply(rule Rule) (Combination, bool) {
	newComb := FromRule(rule)

	for i := range newComb {
		if newComb[i].min > c[i].max {
			fmt.Printf("NOPE: new %s < original %s\n", newComb[i], c[i])
			return Combination{}, false
		}

		if newComb[i].max < c[i].min {
			fmt.Printf("NOPE: new %s > original %s\n", newComb[i], c[i])
			return Combination{}, false
		}

	}

	return c.MinMax(newComb), true
}

func (c Combination) Merge(other Combination) (Combination, bool) {
	for i := range other {
		if other[i].min > c[i].max {
			fmt.Printf("NOPE: new %s < original %s\n", other[i], c[i])
			return Combination{}, false
		}

		if other[i].max < c[i].min {
			fmt.Printf("NOPE: new %s > original %s\n", other[i], c[i])
			return Combination{}, false
		}

	}

	return c.MinMax(other), true
}

func FromRule(rule Rule) Combination {
	i := xmasMap[rule.operand1]
	comb := NewCombination()
	if rule.larger {
		comb[i].min = rule.operand2 - 1
	} else {
		comb[i].max = rule.operand2 + 1
	}
	return comb
}

func NewCombination() Combination {
	c := Combination{}
	for i := range c {
		c[i] = Interval{1, 4000}
	}
	return c
}

func (r *RuleSet) Resolve(ruleKey string, comb Combination, depth int) []Combination {
	product := make([]Combination, 0)
	fmt.Printf("%s=== %s ===\n", strings.Repeat(sep, depth), ruleKey)
	for _, rule := range r.rules[ruleKey] {
		if rule.hasOperation {
			// has operation, no terminal
			if newComb, ok := comb.Apply(rule); ok {
				goodColor.Println(strings.Repeat(sep, depth), rule)
				if rule.isTerminal {
					if rule.production == "A" {
						product = append(product, newComb)
					}
				} else {
					product = append(product, r.Resolve(rule.production, newComb, depth+1)...)
				}
			} else {
				badColor.Println(strings.Repeat(sep, depth), rule)
			}
		} else {
			goodColor.Println(strings.Repeat(sep, depth), rule)
			if rule.isTerminal {
				if rule.production == "A" {
					product = append(product, comb)
				}
			} else {
				product = append(product, r.Resolve(rule.production, comb, depth+1)...)
			}
		}
	}
	fmt.Printf("%s=== %s ===\n", strings.Repeat(sep, depth), ruleKey)
	return product
}

func part1(data []string) int {
	r := NewRuleset(data)
	return r.EvalAll()
}

func part2(data []string) int {
	r := NewRuleset(data)

	c := r.Resolve("in", NewCombination(), 0)
	fmt.Println(len(c))

	applied := true
	for applied {
		applied = false
		for i := 1; i < len(c)-1; i++ {
			newC, ok := c[i].Merge(c[i+1])
			if ok {
				c[i] = newC
				c = append(c[:i+1], c[i+2:]...)
				applied = true
				break
			}
		}
	}

	sum := 0
	for _, comb := range c {
		sum += comb.Count()
	}

	return sum
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
	fmt.Println("167409079868000 (input1)")

}
