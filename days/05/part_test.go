package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	require.Equal(t, 35, part1(input.LoadString("input1")))
}
