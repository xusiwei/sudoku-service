package service

import "fmt"
import "log"
import "net/http"
import "encoding/json"

import "sudoku"

type SolverService struct {
	solver  *sudoku.Solver
	checker *sudoku.Checker
}

type SolverResult struct {
	Puzzle    string   `json:"puz"`
	Solutions []string `json:"solutions"`
}

// constructor
func NewSolverService() *SolverService {
	return &SolverService{solver: sudoku.NewSolver()}
}

// methods
/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (ss *SolverService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		fmt.Fprintf(writer, "method not support!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "parse form failed!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	puz := request.FormValue(puzKey)
	if puz == "" {
		fmt.Fprintf(writer, "need puz argument!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	puzzle := sudoku.NewPuzzle(puz)
	result := SolverResult{Puzzle: puz}
	if ss.checker.Check(puzzle) {
		solutions := ss.solver.Solve(puzzle)
		result.Solutions = make([]string, len(solutions))
		for i, p := range solutions {
			result.Solutions[i] = p.ToString()
		}
	} else {
		result.Solutions = nil
	}

	jstr, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(writer, "json marshal failed!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("return json: %s", jstr)
	fmt.Fprintf(writer, "%s\r\n", jstr)
	writer.WriteHeader(http.StatusOK)
}
