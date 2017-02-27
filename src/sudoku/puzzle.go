package sudoku

import "strconv"

const (
	Size  = 9
	Cells = Size * Size
)

type Puzzle struct {
	data []int
}

// constructor
func NewPuzzle(str string) *Puzzle {
	arr := puzString2Array(str)
	if arr == nil {
		return nil
	}
	return &Puzzle{arr}
}

func CopyPuzzle(old *Puzzle) *Puzzle {
	dup := &Puzzle{}
	dup.data = make([]int, len(old.data))
	copy(dup.data, old.data)
	return dup
}

func MakePuzzle(str string) Puzzle {
	return Puzzle{puzString2Array(str)}
}

// methods
func (p *Puzzle) ToString() string {
	return puzArray2String(p.data)
}

func (p *Puzzle) Get(x, y int) int {
	return p.data[x*Size+y]
}

func (p *Puzzle) Set(x, y int, v int) {
	p.data[x*Size+y] = v
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
