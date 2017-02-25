package sudoku

import "strconv"

const (
	Rows  = 9
	Cols  = 9
	Cells = Rows * Cols
)

type Puzzle struct {
	data []int
}

// constructor
func NewPuzzle(str string) *Puzzle {
	return &Puzzle{puzString2Array(str)}
}

func MakePuzzle(str string) Puzzle {
	return Puzzle{puzString2Array(str)}
}

// methods
func (p *Puzzle) ToString() string {
	return puzArray2String(p.data)
}

func (p *Puzzle) Get(x, y int) int {
	return p.data[x*Cols+y]
}


// private
func puzString2Array(str string) []int {
	data := []int{}
	for _, r := range str {
		n, _ := strconv.Atoi(string(r))
		data = append(data, n)
	}
	if len(data) != Cells {
		return nil
	}
	return data
}

func puzArray2String(data []int) string {
	if len(data) != Cells {
		return ""
	}
	result := ""
	for _, v := range data {
		result += strconv.Itoa(v)
	}
	return result
}
