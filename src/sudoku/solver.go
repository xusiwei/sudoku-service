package sudoku

type Solver struct {
	input   *Puzzle
	current *Puzzle
	results []Puzzle
	checker *Checker
}

// constructor
func NewSolver() *Solver {
	return &Solver{}
}

// methods
func (solver *Solver) Solve(p *Puzzle) []Puzzle {
	solver.input = p
	solver.checker = NewChecker()
	solver.current = NewPuzzle(p.ToString())
	solver.results = make([]Puzzle, 0) // slice with new underlying array
	solver.backtrack(0, 0)
	return solver.results
}

func (solver *Solver) backtrack(r, c int) bool {
	if r*Size+c == Cells {
		solution := CopyPuzzle(solver.current)
		solver.results = append(solver.results, *solution)
		return true
	}
	if c == Size {
		return solver.backtrack(r+1, 0)
	}
	if solver.current.Get(r, c) != 0 {
		return solver.backtrack(r, c+1)
	}

	for v := 1; v <= 9; v++ { // try each value
		solver.current.Set(r, c, v)
		if solver.checker.Check(solver.current) {
			solver.backtrack(r, c+1)
		}
		solver.current.Set(r, c, 0) // backtraking
	}
	return false
}
