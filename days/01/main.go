package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

var (
	nm = []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
)

func value(in string) int {
	number := ""
	for _, ch := range in {
		if _, err := strconv.Atoi(string(ch)); err == nil {
			number += string(ch)
		}
	}

	if len(number) > 2 {
		number = string(number[0]) + string(number[len(number)-1])
	} else if len(number) == 1 {
		number += number
	}
	n, _ := strconv.Atoi(number)
	fmt.Println(n)
	return n
}

func reverse(in string) string {
	result := ""
	for _, ch := range in {
		result = string(ch) + result
	}

	return result
}

func replace(in string) string {
	buffer := ""
	out := ""
	for _, ch := range in {
		buffer += string(ch)
		fmt.Printf("[buffer] %s\n", buffer)

		line := buffer
		for i, n := range nm {
			line = strings.ReplaceAll(line, n, strconv.Itoa(i+1))
		}
		if line != buffer {
			fmt.Printf("[replace] %s -> %s\n", buffer, line)
			out += line
			buffer = ""
		}
	}

	return out + buffer
}

func replaceReverse(in string) string {
	buffer := ""
	out := ""
	for _, ch := range reverse(in) {
		buffer = string(ch) + buffer
		fmt.Printf("[buffer] %s\n", buffer)

		line := buffer
		for i, n := range nm {
			line = strings.ReplaceAll(line, n, strconv.Itoa(i+1))
		}
		if line != buffer {
			fmt.Printf("[replace] %s -> %s\n", buffer, line)
			out = line + out
			buffer = ""
		}
	}

	return buffer + out
}

func part1(in []string) int {
	result := 0
	for _, line := range in {
		result += value(line)
	}
	return result
}

func part2(in []string) int {
	result := 0

	for _, line := range in {
		orig := line
		line = replace(line) + replaceReverse(line)
		fmt.Printf("%s -> %s\n", orig, line)
		result += value(line)
	}
	return result
}

func main() {

	data := input.LoadString("input")

	fmt.Println(part1(data))

	fmt.Println("bad: 54607, 54596")
	fmt.Println(part2(data))
}
