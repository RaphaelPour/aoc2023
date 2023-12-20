package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Pulse int

const (
	HIGH_PULSE Pulse = iota
	LOW_PULSE
	NO_PULSE
)

type Module interface{
	func Get() Pulse
	func AddInput(d Module) {
}

type FlipFlop struct{
	state bool
	inputs []Module
}

func (f *FlipFlop) Get()Pulse {
	in := f.inputs[0].Get()

	if in == HIGH_PULSE {
		return NO_PULSE
	} 

	f.state = !f.state
	
	if f.state{
		 return LOW_PULSE
	}	
	return HIGH_PULSE
}

func (f *FlipFlop) AddInput(d Module) {
	f.inputs = append(f.inputs, d)
}

type Conjunction struct{
	recent []Pulse
	inputs []Module
}

func (c *Conjunction) AddInput(d Module) {
	c.inputs = append(c.inputs, d)
	c.recent = append(c.recent, LOW_PULSE)
}

func (c *Conjunction) Get() Pulse {
	pulse := LOW_PULSE
	for _, input :=  range c.inputs {
		if input.Get() == LOW_PULSE {
			pulse = HIGH_PULSE
			// don't break as all inputs need to be Get()ed in order
			// to update the whole circuit
		}
	}
	return pulse
}

type Inv struct {
	inputs []Module
}

func (i *Inv) AddInput(d Module) {
	i.inputs = append(i.inputs, d)
}

func (i Inv) Get() Pulse {
	if i.inputs[0].Get() == HIGH_PULSE {
		return LOW_PULSE
	}
	return HIGH_PULSE
}

type Broadcaster struct {
	inputs []Module
}

func (b *Broadcaster) AddInput(d Module) {
	b.inputs = append(b.inputs, d)
}

func (b Broadcaster) Get() Pulse {
	return b.inputs[0].Get()
}

type Output struct {
	inputs []Module
}

func (o *Output) AddInput(d Module) {
	o.inputs = append(o.inputs, d)
}

func (o Output) Get() Pulse {
	return o.inputs[0].Get()
}

type Button struct {
	pressed bool
}

func (b Button) AddInput(_ Module){}
func (b Button) Get() Pulse {
	return LOW_PULSE
}

func part1(data []string) int {
	modules := make(map[string]Module)
	for _, line := range data {
		
	}
}

func part2(data []string) int {
	return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadDefaultString()

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
