package solvers

import (
	"slices"
	"strconv"
	"strings"
)

type card struct {
	num     int
	copies  int
	numbers []int
	winning []int
}

type Day4 struct {
	lines []string
	cards []*card
}

func (d *Day4) Part1(in []byte) int {
	d.lines = toLines(in)
	sum := 0
	for _, l := range d.lines {
		c := newCard(l)
		sum += c.value()
	}
	return sum
}

func (d *Day4) Part2(in []byte) int {
	d.lines = toLines(in)
	d.cards = make([]*card, 0)
	for _, l := range d.lines {
		c := newCard(l)
		d.cards = append(d.cards, c)
	}

	for i, card := range d.cards {
		for j := 0; j < card.winningNums(); j++ {
			d.cards[i+j+1].copies += card.copies
		}
	}

	total := 0
	for _, card := range d.cards {
		total += card.copies
	}
	return total
}

func newCard(line string) *card {
	c := strings.Split(line, ":")
	num, _ := strconv.Atoi(c[0][5:])
	numbers := strings.Split(c[1], "|")
	return &card{num: num, copies: 1, winning: parseNumbers(numbers[0]), numbers: parseNumbers(numbers[1])}
}

func parseNumbers(numStr string) []int {
	fs := strings.Fields(strings.TrimSpace(numStr))
	nums := make([]int, 0, len(fs))
	for _, f := range fs {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}
	return nums
}

func (c *card) value() int {
	val := 0
	for _, n := range c.numbers {
		if slices.Contains(c.winning, n) {
			if val == 0 {
				val = 1
			} else {
				val = val * 2
			}
		}
	}
	return val
}

func (c *card) winningNums() int {
	winning := 0
	for _, n := range c.numbers {
		if slices.Contains(c.winning, n) {
			winning += 1
		}
	}
	return winning
}

func init() {
	solvers[4] = &Day4{}
}
