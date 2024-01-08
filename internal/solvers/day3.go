package solvers

import (
	"strconv"
	"unicode"
)

type gearPos struct {
	row, col int
}

type Day3 struct {
	lines []string
	gears map[gearPos][]int
}

func (d *Day3) Part1(in []byte) int {
	d.lines = toLines(in)
	d.gears = make(map[gearPos][]int)
	sum := 0
	for _, pn := range d.findPartNumbers() {
		sum += pn
	}

	return sum
}

func (d *Day3) Part2(in []byte) int {
	d.lines = toLines(in)
	d.gears = make(map[gearPos][]int)
	d.findPartNumbers()
	sum := 0
	for _, gear := range d.gears {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}
	return sum
}

func (d *Day3) findPartNumbers() []int {
	numbers := make([]int, 0)

	for row := 0; row < len(d.lines); row++ {
		rowStr := d.lines[row]
		startCol := -1
		for col := 0; col < len(rowStr); col++ {
			if startCol < 0 && unicode.IsDigit(rune(rowStr[col])) {
				startCol = col
			}
			if startCol >= 0 {
				if col == len(rowStr)-1 || !unicode.IsDigit(rune(rowStr[col+1])) {
					num, _ := strconv.Atoi(rowStr[startCol : col+1])
					if d.isPartNumber(row, startCol, col, num) {
						numbers = append(numbers, num)
					}
					startCol = -1
				}
			}
		}
	}
	return numbers
}

func (d *Day3) isPartNumber(rowStart int, colStart, colEnd int, num int) bool {
	pn := false
	for row := max(rowStart-1, 0); row <= min(rowStart+1, len(d.lines)-1); row++ {
		rowStr := d.lines[row]
		for col := max(colStart-1, 0); col <= min(colEnd+1, len(d.lines[row])-1); col++ {
			if isSymbol(rowStr[col]) {
				if rowStr[col] == '*' {
					gp := gearPos{row: row, col: col}
					gearNums := d.gears[gp]
					if gearNums == nil {
						d.gears[gp] = []int{num}
					} else {
						d.gears[gp] = append(d.gears[gp], num)
					}
				}
				pn = true
			}
		}
	}
	return pn
}

func isSymbol(char byte) bool {
	if char == '.' || unicode.IsDigit(rune(char)) {
		return false
	}
	return true
}

func init() {
	solvers[3] = &Day3{}
}
