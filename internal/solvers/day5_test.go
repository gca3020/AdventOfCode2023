package solvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d5sample = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestDay5_Part1(t *testing.T) {
	d := Day5{}
	assert.Equal(t, 35, d.Part1([]byte(d5sample)))
}

func TestDay5_Part2(t *testing.T) {
	d := Day5{}
	assert.Equal(t, 46, d.Part2([]byte(d5sample)))
}

func TestParse(t *testing.T) {
	d := Day5{}
	d.parse(toLines([]byte(d5sample)))
	assert.Equal(t, []int{79, 14, 55, 13}, d.seeds)
	assert.Equal(t, []int{
		79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92,
		55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	}, d.getSeeds())
	assert.Equal(t, &seedMap{src: "temperature", dest: "humidity", ranges: []seedRange{{0, 69, 1}, {1, 0, 69}}}, d.maps["temperature"])
	assert.Equal(t, &seedMap{src: "humidity", dest: "location", ranges: []seedRange{{60, 56, 37}, {56, 93, 4}}}, d.maps["humidity"])
}
