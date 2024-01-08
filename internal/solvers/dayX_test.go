package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dXsample = `
`

func TestDayX_Part1(t *testing.T) {
	d := DayX{}
	assert.Equal(t, 0, d.Part1([]byte(dXsample)))
}

func TestDayX_Part2(t *testing.T) {
	d := DayX{}
	assert.Equal(t, 0, d.Part2([]byte(dXsample)))
}
