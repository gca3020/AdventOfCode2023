package solvers

import (
	"strconv"
	"strings"
	"unicode"
)

type Day1 struct {
	includeWords bool
}

func (d *Day1) Part1(in []byte) int {
	d.includeWords = false

	sum := 0
	for _, line := range toLines(in) {
		sum += d.calibrationValue(line)
	}
	return sum
}

func (d *Day1) Part2(in []byte) int {
	d.includeWords = true

	sum := 0
	for _, line := range toLines(in) {
		sum += d.calibrationValue(line)
	}
	return sum
}

func (d *Day1) calibrationValue(line string) int {
	return 10*d.firstInt(line) + d.lastInt(line)
}

var replacements = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func (d *Day1) isNumber(line string, idx int) (int, bool) {
	substr := line[idx:]
	if unicode.IsDigit(rune(substr[0])) {
		a, _ := strconv.Atoi(string(substr[0]))
		return a, true
	}

	if d.includeWords {
		for key, val := range replacements {
			if strings.HasPrefix(substr, key) {
				return val, true
			}
		}
	}

	return 0, false
}

func (d *Day1) firstInt(line string) int {
	for i := 0; i < len(line); i++ {
		if v, ok := d.isNumber(line, i); ok {
			return v
		}
	}
	return 0
}

func (d *Day1) lastInt(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if v, ok := d.isNumber(line, i); ok {
			return v
		}
	}
	return 0
}

func init() {
	solvers[1] = &Day1{}
}
