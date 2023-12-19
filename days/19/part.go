package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	xmasMap = map[string]int{
		"x": 0,
		"m": 1,
		"a": 2,
		"s": 3,
	}
)

type Rule struct {
	isTerminal   bool
	hasOperation bool
	operand1     string
	larger       bool
	operand2     int
	production   string
}

func NewRules(in string) []Rule {
	rules := make([]Rule, 0)
	for _, cond := range strings.Split(in, ",") {
		r := Rule{}
		parts := strings.Split(cond, ":")

		if len(parts) == 1 {
			r.production = cond
		} else {
			fmt.Print(string(parts[0][1]))
			r.hasOperation = true
			r.operand1 = string(parts[0][0])
			r.larger = (parts[0][1] == '>')
			fmt.Println(r.larger)
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

	op := "<"
	if r.larger {
		op = ">"
	}

	fmt.Printf("%d %s %d = %t\n", val, op, r.operand2, (r.larger == (val > r.operand2)))
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
		path, ok := r.Eval(rating)
		if ok {
			fmt.Println("A ", rating, path)
			sum += rating.Sum()
		} else {
			fmt.Println("R ", rating, path)
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

func (i Interval) MinMax(other Interval) Interval {
	if i.min > other.min {
		i.min = other.min
	}

	if i.max < other.max {
		i.max = other.max
	}

	return i
}

func (i Interval) Span() int {
	return i.max - i.min
}

type Combination [4]Interval

func (c Combination) MinMax(other Combination) Combination {
	for i := range c {
		c[i] = c[i].MinMax(other[i])
	}
	return c
}

func (c Combination) Count() int {
	count := 0
	for i := range c {
		if span := c[i].Span(); span > count {
			count = span
		}
	}
	return count
}

func (c Combination) Apply(rule Rule) (Combination, bool) {
	newComb := FromRule(rule)

	for i := range newComb {
		if newComb[i].min < c.max || newComb[i].max > c.min {
			return Combination{}, false
		}
	}

	return c.MinMax(newComb)
}

func FromRule(r Rule) Combination {
	i := xmasMap[rule.operand1]
	comb := NewCombination()
	if rule.lower {
		comb[i].max = r.operand2 - 1
	} else {
		comb[i].min = r.operand2 + 1
	}
	return comb
}

func NewCombination() Combination {
	c := make(Combination, 4)
	for i := range c {
		c[i] = Interval{1, 4000}
	}
	return c
}

func (r *RuleSet) Resolve(ruleKey string, comb Combination) int {
	product := 0
	for _, rule := range r.rules[ruleKey] {
		if !rule.hasOperation {
			if rule.isTerminal {
				if rule.production == "A" {
					product += comb.Count()
				}
			} else {
				product += append(product, r.Resolve(rule.production, comb))
			}
			continue
		}

		// has operation, no terminal
		newComb, ok := comb.Apply(rule)
		if ok {
			product += append(product, r.Resolve(r.production, newComb)...)
		}
	}
	return product
}

func part1(data []string) int {
	r := NewRuleset(data)
	return r.EvalAll()
}

func part2(data []string) int {
	r := NewRuleset(data)
	r.Resolve("in")
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))

}
