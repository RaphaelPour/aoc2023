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

func TestFind(t *testing.T) {
	m, err := NewMap(input.LoadString("input1")[2:])
	require.NoError(t, err)

	for _, testCase := range []struct {
		seed     int
		expected int
		name     string
		from     string
	}{
		{seed: 79, expected: 82, from: "seed", name: "1st seed"},
		{seed: 14, expected: 43, from: "seed", name: "2nd seed"},
		{seed: 55, expected: 86, from: "seed", name: "3rd seed"},
		{seed: 13, expected: 35, from: "seed", name: "4th seed"},
		{seed: 78, expected: 82, from: "humidity", name: "humidity 1"},
		{seed: 43, expected: 43, from: "humidity", name: "humidity 2"},
		{seed: 82, expected: 86, from: "humidity", name: "humidity 3"},
		{seed: 35, expected: 35, from: "humidity", name: "humidity 4"},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := m.Find(testCase.seed, testCase.from)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestPart1Example(t *testing.T) {
	t.Skip()
	require.Equal(t, 35, part1(input.LoadString("input1")))
}
