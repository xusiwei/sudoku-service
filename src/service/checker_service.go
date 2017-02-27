package service

import "fmt"
import "log"
import "net/http"
import "encoding/json"

import "sudoku"

const puzKey = "puz"

type CheckerService struct {
	checker *sudoku.Checker
}

type CheckerResult struct {
	Puzzle string `json:"puz"`
	Valid  bool   `json:"valid"`
	Blanks int    `json:"blanks"`
}

// constructor
func NewCheckerService() *CheckerService {
	return &CheckerService{sudoku.NewChecker()}
}

// methods
/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (cs *CheckerService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
	result := CheckerResult{
		Puzzle: puz,
		Blanks: cs.checker.Blanks(puzzle),
		Valid:  cs.checker.Check(puzzle)}

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
