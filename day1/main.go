package main

import (
	"log/slog"
	"strconv"
	"strings"
	"unicode"

	"github.com/gca3020/AdventOfCode2023/internal/parse"
)

func main() {
	slog.Info("Running day1")
	inputs, err := parse.ReadInputs()
	if err != nil {
		slog.Error("could not read inputs", "err", err)
	}
	for _, input := range inputs {
		part1(input)
		part2(input)
	}
}

func part1(input *parse.Input) {
	sum := 0
	for _, line := range input.Lines() {
		sum += calibrationValue(line, false)
	}
	slog.Info("Total calibration value", "part", 1, "input", input.Name, "value", sum)
}

func part2(input *parse.Input) {
	sum := 0
	for _, line := range input.Lines() {
		sum += calibrationValue(line, true)
	}
	slog.Info("Total calibration value", "part", 2, "input", input.Name, "value", sum)
}

func calibrationValue(line string, includeWords bool) int {
	return 10*firstInt(line, includeWords) + lastInt(line, includeWords)
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

func isNumber(line string, idx int, includeWords bool) (int, bool) {
	substr := line[idx:]
	if unicode.IsDigit(rune(substr[0])) {
		a, _ := strconv.Atoi(string(substr[0]))
		return a, true
	}

	if includeWords {
		for key, val := range replacements {
			if strings.HasPrefix(substr, key) {
				return val, true
			}
		}
	}

	return 0, false
}

func firstInt(line string, includeWords bool) int {
	for i := 0; i < len(line); i++ {
		if v, ok := isNumber(line, i, includeWords); ok {
			return v
		}
	}
	return 0
}

func lastInt(line string, includeWords bool) int {
	for i := len(line) - 1; i >= 0; i-- {
		if v, ok := isNumber(line, i, includeWords); ok {
			return v
		}
	}
	return 0
}
