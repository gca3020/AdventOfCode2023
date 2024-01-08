package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d1sample1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestDay1_Part1(t *testing.T) {
	d := Day1{}
	assert.Equal(t, 142, d.Part1([]byte(d1sample1)))
}

var d1sample2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestDay1_Part2(t *testing.T) {
	d := Day1{}
	assert.Equal(t, 281, d.Part2([]byte(d1sample2)))
}
