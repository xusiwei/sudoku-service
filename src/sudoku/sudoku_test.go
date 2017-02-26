package sudoku

import "fmt"
import "log"
import "strconv"
import "testing"

func TestChecker(t *testing.T) {
	c := NewChecker()

	printPuzzle := func(puz string) {
		p := NewPuzzle(puz)
		log.Printf("puzzle: %s\n", puz)
		log.Printf("pretty: \n%s\n", p.PrettyString())
		log.Println("Check:", c.Check(p))
		log.Println("CheckPuzzle:", c.CheckPuzzle(puz))
	}

	puz1 := ""
	for i := 0; i < Cells; i++ {
		puz1 += strconv.Itoa(i % 9)
	}
	printPuzzle(puz1)
	if c.CheckPuzzle(puz1) != false {
		t.Errorf("Assert FAIL: c.Check(%s) != false", puz1)
	}

	puz2 := "268597134"
	puz2 += "439618527"
	puz2 += "571423896"
	puz2 += "715239468"
	puz2 += "942856371"
	puz2 += "386174259"
	puz2 += "894365712"
	puz2 += "623741985"
	puz2 += "157982643"
	printPuzzle(puz2)
	if c.CheckPuzzle(puz2) != true {
		t.Errorf("Assert FAIL: c.Check(%s) == true", puz2)
	}
}

func TestPuzzle(t *testing.T) {
	puzStr := ""
	for i := 0; i < Cells; i++ {
		puzStr += strconv.Itoa(i % 9)
	}
	log.Printf("puzStr:\t%s\n", puzStr)

	// constructor
	p := NewPuzzle(puzStr)

	// ToString
	log.Printf("ToString:\t:%s\n", p.ToString())
	log.Printf("pretty: \n%s\n", p.PrettyString())

	// Get/Set
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			if p.Get(i, j) != j {
				t.Errorf("Assert failed: p.Get(%d, %d) != %d", i, j, j)
			}

			p.Set(i, j, 9)
			if p.Get(i, j) != 9 {
				t.Errorf("Assert failed: p.Get(%d, %d) != %d", i, j, 9)
			}
		}
	}
}

func (p *Puzzle) PrettyString() string {
	s := "|"
	for j := 0; j < Cols; j++ {
		s += fmt.Sprintf(" %d |", p.Get(0, j))
	}
	s += fmt.Sprintf("\n|")
	for j := 0; j < Cols; j++ {
		s += fmt.Sprintf("---|")
	}
	s += fmt.Sprintf("\n")
	for i := 1; i < Rows; i++ {
		s += fmt.Sprintf("|")
		for j := 0; j < Cols; j++ {
			s += fmt.Sprintf(" %d |", p.Get(i, j))
		}
		s += fmt.Sprintf("\n")
	}
	return s
}
