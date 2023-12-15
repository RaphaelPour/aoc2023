package main

import (
	"strings"
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 1320, part1(strings.Split(input.LoadString("input1")[0], ",")))
}

func TestExamplePart2(t *testing.T) {
	require.Equal(t, 145, part2(strings.Split(input.LoadString("input1")[0], ",")))
}

func TestInputPart1(t *testing.T) {
	require.Equal(t, 516804, part1(strings.Split(input.LoadString("input")[0], ",")))
}

func TestInputPart2(t *testing.T) {
	require.Equal(t, 231844, part2(strings.Split(input.LoadString("input")[0], ",")))
}
