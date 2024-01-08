package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d3sample = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestDay3_Part1(t *testing.T) {
	d := Day3{}
	assert.Equal(t, 4361, d.Part1([]byte(d3sample)))
}

func TestDay3_Part2(t *testing.T) {
	d := Day3{}
	assert.Equal(t, 467835, d.Part2([]byte(d3sample)))
}
