package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	require.Equal(t, 374, part1(input.LoadString("input1")))
}

func TestInput1(t *testing.T) {
	require.Equal(t, 9556712, part1(input.LoadString("input")))
}
