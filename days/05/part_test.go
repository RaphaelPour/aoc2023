package main

import (
	"fmt"
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
			actual, err := m.Find(testCase.seed, testCase.from, "location", 0)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestExample1(t *testing.T) {
	m := M{
		data: map[Key]RangeRange{
			Key{from: "seed", to: "soil"}: RangeRange{
				Range{50, 98, 2},
				Range{52, 50, 48},
			},
		},
		cache: map[CacheKey]int{},
	}
	fmt.Println(m)

	result, err := m.Find(79, "seed", "soil", 0)
	require.NoError(t, err)
	require.Equal(t, 81, result)

	result, err = m.Find(14, "seed", "soil", 0)
	require.NoError(t, err)
	require.Equal(t, 14, result)

	result, err = m.Find(55, "seed", "soil", 0)
	require.NoError(t, err)
	require.Equal(t, 57, result)

	result, err = m.Find(13, "seed", "soil", 0)
	require.NoError(t, err)
	require.Equal(t, 13, result)
}

func TestProject(t *testing.T) {
	r := Range{
		destinationStart: 60,
		sourceStart:      56,
		length:           37,
	}
	require.Equal(t, 60, r.project(56))
	require.Equal(t, 0, r.project(0))
	require.Equal(t, 55, r.project(55))
	require.Equal(t, 64, r.project(60))
}

func TestPart1Example(t *testing.T) {
	t.Skip()
	require.Equal(t, 35, part1(input.LoadString("input1")))
}
