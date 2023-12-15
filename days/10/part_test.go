package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 7102, part1(input.LoadString("input")))
}

func TestExample1Part1(t *testing.T) {
	require.Equal(t, 4, part1(input.LoadString("input1")))
}
func TestExample2Part2(t *testing.T) {
	require.Equal(t, 4, part2(input.LoadString("input2")))
}

func TestExample3Part2(t *testing.T) {
	require.Equal(t, 8, part2(input.LoadString("input3")))
}

func TestExample4Part2(t *testing.T) {
	require.Equal(t, 10, part2(input.LoadString("input4")))
}
