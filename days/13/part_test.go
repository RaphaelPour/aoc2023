package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	require.Equal(t, 405, part1(input.LoadString("input1")))
}

func TestPart1Input(t *testing.T) {
	require.Equal(t, 34202, part1(input.LoadString("input")))
}

func TestPart2Example(t *testing.T) {
	require.Equal(t, 400, part2(input.LoadString("input1")))
}

func TestPart1Mirror1(t *testing.T) {
	data := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	require.Equal(t, 400, part1(data))
}

func TestPart2Mirror3(t *testing.T) {
	data := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}

	require.Equal(t, 300, part2(data))
}

func TestPart2Mirror1(t *testing.T) {
	data := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	require.Equal(t, 100, part2(data))
}

func TestMirror1(t *testing.T) {
	data := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	p := NewPatterns(data)[0]
	x, y := p.FindMirrorAxis()
	require.Equal(t, -1, x)
	require.Equal(t, 3, y)

}

func TestMirror2(t *testing.T) {
	data := []string{
		"#...##..#",
		"#...##..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	p := NewPatterns(data)[0]
	x, y := p.FindMirrorAxis()
	require.Equal(t, -1, x)
	require.Equal(t, 0, y)

}
