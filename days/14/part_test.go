package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 136, part1(input.LoadString("input1")))
}

func TestExamplePart2(t *testing.T) {
	require.Equal(t, 64, part2(input.LoadString("input1")))
}

func TestInputPart1(t *testing.T) {
	require.Equal(t, 109665, part1(input.LoadString("input")))
}
