package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinations(t *testing.T) {
	c1 := NewCombination()
	c1[0].max = 100
	c2 := NewCombination()
	c2[1].min = 2000
	c3 := c1.MinMax(c2)

	require.Equal(t, 1, c3[0].min)
	require.Equal(t, 100, c3[0].max)

	require.Equal(t, 2000, c3[1].min)
	require.Equal(t, 4000, c3[1].max)
}

func TestInterval(t *testing.T) {
	i1 := Interval{1, 100}
	i2 := Interval{1, 4000}
	i3 := i1.MinMax(i2)

	require.Equal(t, 1, i3.min)
	require.Equal(t, 100, i3.max)
}
