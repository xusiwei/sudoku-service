package sudoku

import "fmt"
import "log"
import "strconv"
import "testing"

func TestPuzzle(t *testing.T) {
	puzStr := ""
	for i := 0; i < Cells; i++ {
		puzStr += strconv.Itoa(i % 9)
	}
	log.Printf("puzStr:\t%s\n", puzStr)

	// constructor
	p := NewPuzzle(puzStr)

	// ToString
	log.Printf("ToString:\t%s\n", p.ToString())
	log.Printf("pretty: \n%s\n", p.prettyString())

	// Get/Set
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
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

func TestChecker(t *testing.T) {
	puz1 := ""
	for i := 0; i < Cells; i++ {
		puz1 += strconv.Itoa(i % 9)
	}
	exp1 := false

	puz2 := "268597134"
	puz2 += "439618527"
	puz2 += "571423896"
	puz2 += "715239468"
	puz2 += "942856371"
	puz2 += "386174259"
	puz2 += "894365712"
	puz2 += "623741985"
	puz2 += "157982643"
	exp2 := true

	puz3 := "631780000"
	puz3 += "200500000"
	puz3 += "004060938"
	puz3 += "742800506"
	puz3 += "980000001"
	puz3 += "305040002"
	puz3 += "109070203"
	puz3 += "453628710"
	puz3 += "007391654"
	exp3 := true

	pvs := []struct {
		puzzle string
		expect bool
	}{{puz1, exp1}, {puz2, exp2}, {puz3, exp3}}

	c := NewChecker()
	for _, pv := range pvs {
		printPuzzle(pv.puzzle)
		if c.CheckPuzzle(pv.puzzle) != pv.expect {
			t.Errorf("Assert Failed: c.Check(%s) != %v", pv.puzzle, pv.expect)
		}
	}
}

func solve(puz string, t *testing.T) {
	solver := NewSolver()
	puzzle := NewPuzzle(puz)

	printPuzzle(puz)

	solutions := solver.Solve(puzzle)
	if len(solutions) == 0 {
		t.Errorf("solve puzzle %s failed!", puz)
	}
	log.Println("found solutions:", len(solutions))
	for i, s := range solutions {
		log.Printf("solution %d: %s\n%s", i, s.ToString(), s.prettyString())
	}
}

func TestSolveEasy(t *testing.T) {
	puz := "631780000"
	puz += "200500000"
	puz += "004060938"
	puz += "742800506"
	puz += "980000001"
	puz += "305040002"
	puz += "109070203"
	puz += "453628710"
	puz += "007391654"
	solve(puz, t)
}

func TestSolveNormal(t *testing.T) {
	puz := "070050006"
	puz += "500081900"
	puz += "900200070"
	puz += "800020605"
	puz += "009008700"
	puz += "300600010"
	puz += "700906051"
	puz += "040070000"
	puz += "005000800"
	solve(puz, t)
}

func TestSolveHard(t *testing.T) {
	puz := "003007000"
	puz += "804056000"
	puz += "020400900"
	puz += "080000439"
	puz += "236000500"
	puz += "000000000"
	puz += "010200300"
	puz += "000073056"
	puz += "000040001"
	solve(puz, t)
}

func TestSolveExpert(t *testing.T) {
	puz := "090000007"
	puz += "002780040"
	puz += "010003000"
	puz += "000034000"
	puz += "004020305"
	puz += "720001090"
	puz += "008102060"
	puz += "070000002"
	puz += "000070001"
	solve(puz, t)
}

func TestSolveExtreme(t *testing.T) {
	puz := "306510000"
	puz += "000020000"
	puz += "050003400"
	puz += "025000801"
	puz += "008100570"
	puz += "160007002"
	puz += "000034025"
	puz += "000080000"
	puz += "090700100"
	solve(puz, t)
}

func printPuzzle(puz string) {
	c := NewChecker()
	p := NewPuzzle(puz)
	log.Printf("puzzle: %s\n", puz)
	log.Printf("pretty: \n%s", p.prettyString())
	log.Println("Check:", c.Check(p))
	log.Println("CheckPuzzle:", c.CheckPuzzle(puz))
}

func (p *Puzzle) prettyString() string {
	s := "|"
	for j := 0; j < Size; j++ {
		s += fmt.Sprintf(" %d |", p.Get(0, j))
	}
	s += fmt.Sprintf("\n|")
	for j := 0; j < Size; j++ {
		s += fmt.Sprintf("---|")
	}
	s += fmt.Sprintf("\n")
	for i := 1; i < Size; i++ {
		s += fmt.Sprintf("|")
		for j := 0; j < Size; j++ {
			s += fmt.Sprintf(" %d |", p.Get(i, j))
		}
		s += fmt.Sprintf("\n")
	}
	return s
}
