package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellarStrings "github.com/RaphaelPour/stellar/strings"
)

var (
	digits = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
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

	if number == "" {
		return 0
	}

	return stellarStrings.ToInt(number)
}

func replace(in string) string {
	buffer := ""
	out := ""
	for _, ch := range in {
		buffer += string(ch)

		line := buffer
		for i, digit := range digits {
			line = strings.ReplaceAll(line, digit, strconv.Itoa(i+1))
		}
		if line != buffer {
			out += line
			buffer = ""
		}
	}

	return out + buffer
}

func replaceReverse(in string) string {
	buffer := ""
	out := ""
	for _, ch := range stellarStrings.Reverse(in) {
		buffer = string(ch) + buffer

		line := buffer
		for i, digit := range digits {
			line = strings.ReplaceAll(line, digit, strconv.Itoa(i+1))
		}
		if line != buffer {
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
		result += value(replace(line) + replaceReverse(line))
	}
	return result
}

func main() {

	data := input.LoadString("input")

	fmt.Printf("part 1: %d\n", part1(data))
	fmt.Printf("part 2: %d\n", part2(data))
}
