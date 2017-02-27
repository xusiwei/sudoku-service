package sudoku

import "log"

const (
	Boxs      = 9
	BoxSize   = 3
	BoxCells  = BoxSize * BoxSize
	NumValues = 10 // 1~9 filled, 0 means blank
)

type Checker struct {
}

// constructor
func NewChecker() *Checker {
	return &Checker{}
}

// methods
func (checker *Checker) CheckPuzzle(puz string) bool {
	return checker.Check(NewPuzzle(puz))
}

func (checker *Checker) Blanks(p *Puzzle) int {
	if p == nil {
		return -1
	}

	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if p.Get(i, j) == 0 {
				count += 1
			}
		}
	}
	return count
}

func (checker *Checker) Check(p *Puzzle) bool {
	if p == nil {
		return false
	}

	// check cell values
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			v := p.Get(i, j)
			if v < 0 || v > 9 {
				log.Printf("invalid value %d at(%d, %d)\n", v, i, j)
				return false
			}
		}
	}

	// check each row
	for i := 0; i < Size; i++ {
		count := [NumValues]int{}
		for j := 0; j < Size; j++ {
			v := p.Get(i, j)
			count[v] += 1
			if v > 0 && count[v] > 1 {
				return false
			}
		}
	}

	// check each cols
	for j := 0; j < Size; j++ {
		count := [NumValues]int{}
		for i := 0; i < Size; i++ {
			v := p.Get(i, j)
			count[v] += 1
			if v > 0 && count[v] > 1 {
				return false
			}
		}
	}

	// check each box
	for b := 0; b < Boxs; b++ {
		br := b / (Size / BoxSize) * BoxSize
		bc := b % (Size / BoxSize) * BoxSize
		count := [NumValues]int{}
		for i := 0; i < BoxCells; i++ {
			v := p.Get(br+i/BoxSize, bc+i%BoxSize)
			count[v] += 1
			if v > 0 && count[v] > 1 {
				return false
			}
		}
	}
	return true
}
