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
	t.Skip()
	require.Equal(t, 4, part1(input.LoadString("input1")))
}

func TestExpand(t *testing.T) {
	input := []string{
		"S-7",
		"|.|",
		"L-J",
	}

	m := NewMap(input)
	require.Equal(t, P{0, 0}, m.start)
	require.Equal(t, 3, m.w)
	require.Equal(t, 3, m.h)

	m.Expand()

	require.Equal(t, P{1, 1}, m.start)
	require.Equal(t, 9, m.w)
	require.Equal(t, 9, m.h)
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
