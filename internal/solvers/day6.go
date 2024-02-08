package solvers

import (
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

type Day6 struct {
	races []race
}

func (d *Day6) parse(lines []string) {
	ts := vals(lines[0])
	ds := vals(lines[1])
	for i, t := range ts {
		d.races = append(d.races, race{time: t, distance: ds[i]})
	}
}

func (d *Day6) parse2(lines []string) {
	ts := vals(strings.ReplaceAll(lines[0], " ", ""))
	ds := vals(strings.ReplaceAll(lines[1], " ", ""))
	for i, t := range ts {
		d.races = append(d.races, race{time: t, distance: ds[i]})
	}
}

func vals(line string) []int {
	ret := make([]int, 0)
	tokens := strings.Fields(strings.Split(line, ":")[1])
	for _, token := range tokens {
		i, _ := strconv.Atoi(token)
		ret = append(ret, i)
	}
	return ret
}

func (d *Day6) Part1(in []byte) int {
	d.parse(toLines(in))
	ways := 1
	for _, race := range d.races {
		ways *= race.waysToWin()
	}
	return ways
}

func (d *Day6) Part2(in []byte) int {
	d.parse2(toLines(in))
	return d.races[0].waysToWin()
}

func (r race) waysToWin() int {
	startWay, stopWay := 0, 0
	for i := 0; i < r.time; i++ {
		if r.isWinning(i) {
			startWay = i
			break
		}
	}

	for i := r.time; i >= 0; i-- {
		if r.isWinning(i) {
			stopWay = i
			break
		}
	}
	return stopWay - startWay + 1
}

func (r race) isWinning(holdTime int) bool {
	speed := holdTime
	timeLeft := r.time - holdTime
	return timeLeft*speed > r.distance
}

func init() {
	solvers[6] = &Day6{}
}
