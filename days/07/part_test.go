package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	require.Equal(t, 6440, part1(input.LoadString("input1")))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 249390788, part1(input.LoadString("input")))
}

func TestExample2(t *testing.T) {
	require.Equal(t, 5905, part2(input.LoadString("input1")))
}

func TestHand1(t *testing.T) {
	for _, tCase := range []struct {
		name     string
		in       string
		expected HandType
	}{
		{name: "five hands 1", in: "QQQQQ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "five hands 2", in: "QQQQJ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "five hands 3", in: "QQQJJ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "five hands 4", in: "QQJJJ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "five hands 5", in: "QJJJJ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "five hands 6", in: "JJJJJ", expected: FIVE_OF_A_KIND_TYPE},
		{name: "four hands 1", in: "QQQQ1", expected: FOUR_OF_A_KIND_TYPE},
		{name: "four hands 2", in: "QQQJ1", expected: FOUR_OF_A_KIND_TYPE},
		{name: "four hands 3", in: "QQJJ1", expected: FOUR_OF_A_KIND_TYPE},
		{name: "four hands 4", in: "QJJJ1", expected: FOUR_OF_A_KIND_TYPE},
		{name: "three 1", in: "QQQ12", expected: THREE_OF_A_KIND_TYPE},
		{name: "three 2", in: "QQJ12", expected: THREE_OF_A_KIND_TYPE},
		{name: "three 3", in: "QJJ12", expected: THREE_OF_A_KIND_TYPE},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			h := NewHand([]rune(tCase.in), 0, true)
			require.Equal(t, tCase.expected.String(), h.kind.String())
		})
	}
}
