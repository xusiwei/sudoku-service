package sudoku

import "math"

type Solver struct {
	input   *Puzzle
	current *Puzzle
	results []Puzzle
	expects int // max solutions expect to solve
	checker *Checker
}

// constructor
func NewSolver() *Solver {
	return &Solver{}
}

// methods
func (solver *Solver) SolveSome(p *Puzzle, expects int) []Puzzle {
	solver.input = p
	solver.expects = expects
	solver.checker = NewChecker()
	solver.current = CopyPuzzle(p) // new puzzle for solver
	solver.results = make([]Puzzle, 0) // slice with new underlying array
	solver.backtrack(0, 0)
	return solver.results
}

func (solver *Solver) Solve(p *Puzzle) []Puzzle {
	return solver.SolveSome(p, math.MaxInt32)
}

func (solver *Solver) backtrack(r, c int) bool {
	if r*Size + c == Cells { // the last cell.
		solution := CopyPuzzle(solver.current)
		solver.results = append(solver.results, *solution)
		return true
	}
	if c == Size { // current row complete, move to next row.
		return solver.backtrack(r+1, 0)
	}

	if solver.current.Get(r, c) != 0 { // skip the given cell.
		return solver.backtrack(r, c+1)
	}

	for v := 1; v <= 9; v++ { // try each value
		solver.current.Set(r, c, v)
		if solver.checker.Check(solver.current) {
			if solver.backtrack(r, c+1) && len(solver.results) >= solver.expects { // valid, go deep!
				return true // stop search more!
			}
		}
		solver.current.Set(r, c, 0) // backtraking
	}
	return false // no solution
}
