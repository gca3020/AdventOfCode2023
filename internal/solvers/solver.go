package solvers

var solvers = make(map[int]Solver, 0)

type Solver interface {
	Part1(in []byte) int
	Part2(in []byte) int
}

func Get(day int) Solver {
	if s, ok := solvers[day]; ok {
		return s
	}
	return nil
}
