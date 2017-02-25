package sudoku

const (
	ROWS = 9
	COLS = 9
	CELLS = ROWS*COLS
)

type Puzzle struct {
	data []int
}

// constructor
func NewPuzzle(str string) *Puzzle {
	p := new(Puzzle)
	p.data := parsePuzzle(str)
	return p
}

func NewPuzzle(data []int) *Puzzle {
	if len(data) == CELLS {
		p := new(Puzzle)
		p.data := data
		return p
	}
	return nil
}

func MakePuzzle(str string) Puzzle {
	return Puzzle{parsePuzzle(str)}
}

// methods
func (p *Puzzle) ToString() string {
	// TODO
}



func parsePuzzle(str string) []int {
	// TODO
}

