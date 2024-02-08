package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d6sample = `
Time:      7  15   30
Distance:  9  40  200
`

func TestDay6_Part1(t *testing.T) {
	d := Day6{}
	assert.Equal(t, 288, d.Part1([]byte(d6sample)))
}

func TestDay6_parse(t *testing.T) {
	d := Day6{}
	d.parse(toLines([]byte(d6sample)))
	assert.Equal(t, []race{{7, 9}, {15, 40}, {30, 200}}, d.races)
}

func TestDay6_parse2(t *testing.T) {
	d := Day6{}
	d.parse2(toLines([]byte(d6sample)))
	assert.Equal(t, []race{{71530, 940200}}, d.races)
}

func TestDay6_Part2(t *testing.T) {
	d := Day6{}
	assert.Equal(t, 71503, d.Part2([]byte(d6sample)))
}

func TestRace_WaysToWin(t *testing.T) {
	r := race{30, 200}
	assert.Equal(t, 9, r.waysToWin())
}
