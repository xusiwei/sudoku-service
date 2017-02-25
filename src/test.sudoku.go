package main

import "fmt"
import "strconv"
import "sudoku"

func assert(expr bool, text string) {
	if ! expr {
		panic(text)
	}
}

func main() {
	defer fmt.Println("done")

	puzStr := ""
	for i := 0; i < sudoku.Cells; i++ {
		puzStr += strconv.Itoa(i%9)
	}
	fmt.Printf("puzStr:\t%s\n", puzStr)

	// constructor
	p := sudoku.NewPuzzle(puzStr)

	// ToString
	fmt.Printf("ToString:\t:%s\n", p.ToString())

	// Get
	for i := 0; i < sudoku.Rows; i++ {
		for j := 0; j < sudoku.Cols; j++ {
			assert(p.Get(i, j) == j,
				fmt.Sprintf("assertion failed on (%d, %d)", i, j))
		}
	}
}
