package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gca3020/AdventOfCode2023/internal/solvers"
)

func main() {
	var day, part int
	flag.IntVar(&day, "day", 1, "The day of the puzzle")
	flag.IntVar(&part, "part", 1, "The part of the day")
	flag.Parse()

	fmt.Printf("Running puzzle for Day %v, Part %v\n", day, part)

	solver := solvers.Get(day)
	if solver == nil {
		panic(fmt.Sprintln("No solver registered for day", day))
	}

	input, err := os.ReadFile(fmt.Sprintf("input/day%d", day))
	if err != nil {
		panic(fmt.Sprintf("Could not read input file for day %d: %s", day, err))
	}

	solution := 0
	if part == 1 {
		solution = solver.Part1(input)
	} else if part == 2 {
		solution = solver.Part2(input)
	} else {
		panic("Puzzles only have parts 1 and 2")
	}
	fmt.Printf("Day %d Part %d: %d\n", day, part, solution)
}
