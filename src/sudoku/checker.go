package sudoku

const (
	Boxs      = 9
	BoxRows   = 3
	BoxCols   = 3
	BoxCells  = BoxRows * BoxCols
	MinValue  = 0
	MaxValue  = 9
	NumValues = MaxValue - MinValue + 1
)

/*
type Checker interface {
	Check(p *Puzzle) bool
}
*/

type SudokuChecker struct {
}

// constructor
func NewChecker() *SudokuChecker {
	return &SudokuChecker{}
}

// methods
func (checker *SudokuChecker) CheckPuzzle(puz string) bool {
	return checker.Check(NewPuzzle(puz))
}

func (checker *SudokuChecker) Check(p *Puzzle) bool {
	if p == nil {
		return false
	}

	// check cell values
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			v := p.Get(i, j)
			if v < 0 || v > 9 {
				return false
			}
		}
	}

	// check each row
	for i := 0; i < Rows; i++ {
		count := [NumValues]int{}
		for j := 0; j < Cols; j++ {
			v := p.Get(i, j)
			if count[v] > 1 {
				return false
			}
			count[v] += 1
		}
	}

	// check each cols
	for j := 0; j < Cols; j++ {
		count := [NumValues]int{}
		for i := 0; i < Rows; i++ {
			v := p.Get(i, j)
			if count[v] > 1 {
				return false
			}
			count[v] += 1
		}
	}

	// check each box
	for b := 0; b < Boxs; b++ {
		br := b / (Cols / BoxCols) * BoxCols
		bc := b % (Cols / BoxCols) * BoxCols
		count := [NumValues]int{}
		for i := 0; i < BoxCells; i++ {
			v := p.Get(br+i/BoxCols, bc+i%BoxCols)
			if count[v] > 1 {
				return false
			}
			count[v] += 1

		}
	}
	return true
}
