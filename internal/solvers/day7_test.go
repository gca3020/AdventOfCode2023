package solvers

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d7sample = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestDay7_Part1(t *testing.T) {
	d := Day7{}
	assert.Equal(t, 6440, d.Part1([]byte(d7sample)))
}

func TestDay7_Part2(t *testing.T) {
	d := Day7{}
	assert.Equal(t, 5905, d.Part2([]byte(d7sample)))
}

func TestDay7_Parse(t *testing.T) {
	d := Day7{}
	d.parse(toLines([]byte(d7sample)))
	assert.Equal(t, []hand{{"32T3K", 765}, {"T55J5", 684}, {"KK677", 28}, {"KTJJT", 220}, {"QQQJA", 483}}, d.hands)
}

func TestDay7_CardSort(t *testing.T) {
	jokersWild = false
	assert.Equal(t, "TQQKA", sortCards("TQKAQ"))
	assert.Equal(t, "279JQ", sortCards("729JQ"))

	jokersWild = true
	assert.Equal(t, "J279Q", sortCards("729JQ"))
	jokersWild = false
}

func Test_hand_strength(t *testing.T) {
	tests := []struct {
		name       string
		cards      string
		jokerswild bool
		strength   int
	}{
		{name: "Five of a kind", cards: "AAAAA", strength: 7},
		{name: "Five Jacks", cards: "JJJJJ", strength: 7},
		{name: "Five Jokers (jw)", cards: "JJJJJ", jokerswild: true, strength: 7},
		{name: "Four of a kind", cards: "AAATA", strength: 6},
		{name: "Four of a kind (jw)", cards: "QJJQ2", jokerswild: true, strength: 6},
		{name: "Full House", cards: "32323", strength: 5},
		{name: "Full House (jw)", cards: "32J23", jokerswild: true, strength: 5},
		{name: "Three of a Kind", cards: "A22J2", strength: 4},
		{name: "Three of a Kind (jw)", cards: "A28J2", jokerswild: true, strength: 4},
		{name: "Two Pairs", cards: "T96T9", strength: 3},
		{name: "One Pair", cards: "53831", strength: 2},
		{name: "One Pair (jw)", cards: "5J831", jokerswild: true, strength: 2},
		{name: "High Card", cards: "A2345", strength: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jokersWild = tt.jokerswild
			assert.Equal(t, tt.strength, hand{cards: tt.cards}.strength())
		})
	}
}

func Test_hand_compare(t *testing.T) {
	// First card is higher
	assert.Positive(t, compareHands(hand{cards: "33332"}, hand{cards: "2AAAA"}))
	assert.Positive(t, compareHands(hand{cards: "77888"}, hand{cards: "77788"}))
}

func Test_hand_sort(t *testing.T) {
	hands := []hand{{"32T3K", 1}, {"T55J5", 4}, {"KK677", 3}, {"KTJJT", 2}, {"QQQJA", 5}}
	slices.SortFunc(hands, compareHands)
	assert.Equal(t, []hand{{"32T3K", 1}, {"KTJJT", 2}, {"KK677", 3}, {"T55J5", 4}, {"QQQJA", 5}}, hands)
}
