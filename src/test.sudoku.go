package main

import "fmt"
import "strconv"
import "sudoku"

func assert(expr bool, text string) {
	if !expr {
		panic(text)
	}
}

func main() {
	defer fmt.Println("done")

	puzzleTest()
	checkerTest()
}

func checkerTest() {
	c := sudoku.NewChecker()

	checkPuzzle := func(puz string, expected bool) {
		p := sudoku.NewPuzzle(puz)

		printPuzzle(p)
		fmt.Printf("puzzle: %s\n", puz)
		fmt.Println("valid:", c.Check(p))

		assert(c.Check(p) == expected, "Checker.Check() assertion failed")
	}

	checkPuzzle(func() string {
		puz := ""
		for i := 0; i < sudoku.Cells; i++ {
			puz += strconv.Itoa(i % 9)
		}
		return puz
	}(), false)

	checkPuzzle(func() string {
		puz := ""
		puz += "268597134"
		puz += "439618527"
		puz += "571423896"
		puz += "715239468"
		puz += "942856371"
		puz += "386174259"
		puz += "894365712"
		puz += "623741985"
		puz += "157982643"
		return puz
	}(), true)
}

func puzzleTest() {
	puzStr := ""
	for i := 0; i < sudoku.Cells; i++ {
		puzStr += strconv.Itoa(i % 9)
	}
	fmt.Printf("puzStr:\t%s\n", puzStr)

	// constructor
	p := sudoku.NewPuzzle(puzStr)

	// ToString
	fmt.Printf("ToString:\t:%s\n", p.ToString())

	printPuzzle(p)

	// Get/Set
	for i := 0; i < sudoku.Rows; i++ {
		for j := 0; j < sudoku.Cols; j++ {
			assert(p.Get(i, j) == j,
				fmt.Sprintf("assertion failed on (%d, %d)", i, j))

			p.Set(i, j, 9)
			assert(p.Get(i, j) == 9,
				fmt.Sprintf("assertion failed on (%d, %d)", i, j))
		}
	}
}

func printPuzzle(p *sudoku.Puzzle) {
	fmt.Printf("\n|")
	for j := 0; j < sudoku.Cols; j++ {
		fmt.Printf(" %d |", p.Get(0, j))
	}
	fmt.Printf("\n|")
	for j := 0; j < sudoku.Cols; j++ {
		fmt.Printf("---|")
	}
	fmt.Printf("\n")
	for i := 1; i < sudoku.Rows; i++ {
		fmt.Printf("|")
		for j := 0; j < sudoku.Cols; j++ {
			fmt.Printf(" %d |", p.Get(i, j))
		}
		fmt.Printf("\n")
	}
}
