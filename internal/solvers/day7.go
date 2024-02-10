package solvers

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

var jokersWild = false

type Day7 struct {
	hands []hand
}

func (d *Day7) parse(lines []string) {
	for _, line := range lines {
		tokens := strings.Fields(line)
		bid, _ := strconv.Atoi(tokens[1])
		d.hands = append(d.hands, hand{cards: tokens[0], bid: bid})
	}
}

func (d *Day7) Part1(in []byte) int {
	jokersWild = false
	d.parse(toLines(in))
	slices.SortFunc(d.hands, compareHands)
	sum := 0
	for i, hand := range d.hands {
		sum += (hand.bid * (i + 1))
	}

	return sum
}

func (d *Day7) Part2(in []byte) int {
	jokersWild = true
	d.parse(toLines(in))
	slices.SortFunc(d.hands, compareHands)
	sum := 0
	for i, hand := range d.hands {
		sum += (hand.bid * (i + 1))
	}

	return sum
}

type hand struct {
	cards string
	bid   int
}

func (h hand) strength() int {
	sorted := sortCards(h.cards)
	jokers := 0
	if jokersWild {
		sorted = strings.ReplaceAll(sorted, "J", "")
		jokers = 5 - len(sorted)
	}
	cardTypes := 1
	mostCards := 0
	currentCards := 0
	if jokers >= 4 {
		return 7
	}
	old := (rune)(sorted[0])
	for _, c := range sorted {
		if c != old {
			old = c
			cardTypes++
			currentCards = 1
		} else {
			currentCards++
			if currentCards > mostCards {
				mostCards = currentCards
			}
		}
	}
	mostCards += jokers
	if cardTypes == 1 {
		// five of a kind
		return 7
	}
	if cardTypes == 2 {
		if mostCards == 4 {
			// four of a kind
			return 6
		}
		if mostCards == 3 {
			// full house
			return 5
		}
	}
	if cardTypes == 3 {
		if mostCards == 3 {
			// three of a kind
			return 4
		}
		if mostCards == 2 {
			// two pairs
			return 3
		}
	}
	if cardTypes == 4 {
		// One pair
		return 2
	}
	return 1
}

func sortCards(in string) string {
	runes := []rune(in)
	slices.SortFunc(runes, func(a, b rune) int {
		return cmp.Compare(value(a), value(b))
	})
	return string(runes)
}

func compareHands(a, b hand) int {
	if a.strength() > b.strength() {
		return 1
	} else if a.strength() < b.strength() {
		return -1
	} else {
		for i, ca := range a.cards {
			cb := rune(b.cards[i])
			if value(ca) > value(cb) {
				return 1
			} else if value(ca) < value(cb) {
				return -1
			}
		}
		return 0
	}
}

func value(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if jokersWild {
			return 1
		}
		return 11
	case 'T':
		return 10
	default:
		v, _ := strconv.Atoi(string(r))
		return v
	}
}

func init() {
	solvers[7] = &Day7{}
}
