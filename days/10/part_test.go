package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 7102, part1(input.LoadString("input")))
}
