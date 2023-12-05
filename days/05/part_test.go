package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestFindKey(t *testing.T) {
	m, err := NewMap(input.LoadString("input1")[2:])
	require.NoError(t, err)

	require.Equal(t, m.FindKey("seed"), Key{from: "seed", to: "soil"})
}

func TestPart1Example(t *testing.T) {
	t.Skip()
	require.Equal(t, 35, part1(input.LoadString("input1")))
}
