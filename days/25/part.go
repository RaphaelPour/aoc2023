package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/RaphaelPour/stellar/input"
)

/* Histogram
== [ PART 1 ] ==
299900 99.97
lnr 53463
pgt 53438
tjz 51153
vph 50967
zkt 48092
jhq 47707
pvv 29416
fqm 28537
mqd 28513
rmr 27019
pvz 23430
jzd 22611
gfq 20513
jzp 18219
tph 17941
ggn 17933
ssx 17327
qss 17249
djr 16290
cfc 15929
0
11m37.302033634s
*/

type Pair struct {
	Key   string
	Value int
}

func (p Pair) String() string { return fmt.Sprintf("%3s %4d", p.Key, p.Value) }

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Graph struct {
	edges map[string][]string
	hist  map[string]int
}

func NewGraph(data []string) Graph {
	g := Graph{
		edges: make(map[string][]string),
		hist:  make(map[string]int),
	}

	for _, line := range data {
		parts := strings.Split(line, " ")
		from := parts[0][0:3]
		to := parts[1:]

		// add from + neighbor
		if _, ok := g.edges[from]; !ok {
			g.edges[from] = make([]string, 0)
		}
		for _, neighbor := range to {
			g.edges[from] = append(g.edges[from], neighbor)
		}

		// add from to each neighbor
		for _, neighbor := range to {
			if _, ok := g.edges[neighbor]; !ok {
				g.edges[neighbor] = make([]string, 0)
			}
			g.edges[neighbor] = append(g.edges[neighbor], from)
		}
	}

	return g
}

func (g Graph) Find(node, goal string, visited map[string]struct{}, maxDepth, depth int) bool {
	if depth == maxDepth {
		return false
	}

	if node == goal {
		visited[node] = struct{}{}
		g.hist[node] = g.hist[node] + 1
		return true
	}

	if _, ok := visited[node]; ok {
		return false
	}

	visited[node] = struct{}{}
	for _, n := range g.edges[node] {
		if g.Find(n, goal, visited, maxDepth, depth+1) {
			g.hist[n] = g.hist[n] + 1
			return true
		}
	}
	return false
}

type Edge struct {
	from, to string
}

func (g Graph) FindDividerEdges(candidates []string) []Edge {
	e := make([]Edge, 0)
	for i, c := range candidates {
		e = append(e, Edge{
			from: c,
			to: g.FindNearest(
				c,
				append(append([]string{}, candidates[:i], candidates[i+1:])),
			),
		},
		)
	}
	return e
}

func (g Graph) FindNearest(node, candidates []string) string {
	steps := 1000
	best := ""

	for _, c := range candidates {
		depth := 1
		for {
			if g.Find(node, c, map[string]struct{}{}, depth, 0) {
				break
			}
			depth++
		}
	}
	return best
}

func (g Graph) CreateHist() {
	keys := make([]string, len(g.edges))
	i := 0
	for key := range g.edges {
		keys[i] = key
		i++
	}

	rounds := 50000
	for i := 0; i < rounds; i++ {
		if i%100 == 0 {
			fmt.Printf("\r%d %.2f", i, 100.0/float64(rounds)*float64(i))
		}
		a := keys[rand.Intn(len(keys))]
		b := keys[rand.Intn(len(keys))]
		for depth := 1; ; depth++ {
			if g.Find(a, b, map[string]struct{}{}, depth, 0) {
				break
			}
		}
	}
	fmt.Println("")
}

func (g Graph) PrintEdges() {
	for edge, neighbors := range g.edges {
		fmt.Printf("%s -> %v\n", edge, neighbors)
	}
}

func (g Graph) PrintHist() {
	pairs := make(Pairs, len(g.hist))

	i := 0
	for node, count := range g.hist {
		pairs[i] = Pair{Key: node, Value: count}
		i++
	}

	sort.Sort(sort.Reverse(pairs))

	for i := 0; i < len(pairs) && i < 20; i++ {
		fmt.Printf("%s %d\n", pairs[i].Key, pairs[i].Value)
	}
}

func part1(data []string) int {
	g := NewGraph(data)
	g.CreateHist()
	g.PrintHist()
	fmt.Println(g.FindDividerEdges())
	return 0
}

func main() {
	start := time.Now()
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))
	fmt.Println(time.Since(start))
}
