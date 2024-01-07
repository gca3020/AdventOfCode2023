package solvers

import (
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d *Day2) Part1(in []byte) int {
	sum := 0
	for _, line := range toLines(in) {
		game := newGame(line)
		if game.isPossible(12, 13, 14) {
			sum += game.id
		}
	}
	return sum
}

func (d *Day2) Part2(in []byte) int {
	sum := 0
	for _, line := range toLines(in) {
		game := newGame(line)
		r, g, b := game.fewestPossible()
		power := r * g * b
		sum += power
	}
	return sum
}

type draw struct {
	red, green, blue int
}

type game struct {
	id    int
	draws []draw
}

func newGame(line string) game {
	draws := make([]draw, 0)
	fields := strings.Split(line, ":")

	// Parse the Game ID
	gid, _ := strconv.Atoi(strings.TrimPrefix(fields[0], "Game "))

	games := strings.Split(fields[1], ";")
	for _, game := range games {
		colors := strings.Split(game, ",")
		draw := draw{}
		for _, color := range colors {
			counts := strings.Split(strings.TrimSpace(color), " ")
			count, _ := strconv.Atoi(counts[0])
			switch counts[1] {
			case "red":
				draw.red = count
			case "green":
				draw.green = count
			case "blue":
				draw.blue = count
			}
		}
		draws = append(draws, draw)
	}

	return game{id: gid, draws: draws}
}

func (g *game) isPossible(maxRed, maxGreen, maxBlue int) bool {
	for _, draw := range g.draws {
		if draw.red > maxRed || draw.green > maxGreen || draw.blue > maxBlue {
			return false
		}
	}
	return true
}

func (g *game) fewestPossible() (red, green, blue int) {
	for _, draw := range g.draws {
		red = max(draw.red, red)
		blue = max(draw.blue, blue)
		green = max(draw.green, green)
	}
	return
}

func init() {
	solvers[2] = &Day2{}
}
