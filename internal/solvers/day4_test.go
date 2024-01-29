package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d4sample = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func TestDay4_Part1(t *testing.T) {
	d := Day4{}
	assert.Equal(t, 13, d.Part1([]byte(d4sample)))
}

func TestDay4_Part2(t *testing.T) {
	d := Day4{}
	assert.Equal(t, 30, d.Part2([]byte(d4sample)))
}

func TestNewCard(t *testing.T) {
	c := newCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	assert.Equal(t, []int{41, 48, 83, 86, 17}, c.winning)
	assert.Equal(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, c.numbers)
	assert.Equal(t, 1, c.num)

	assert.Equal(t, 8, c.value())
	assert.Equal(t, 4, c.winningNums())
}
