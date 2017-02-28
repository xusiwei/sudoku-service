package main

import "net/http"
import "service"

func main() {

	mux := http.NewServeMux()
	mux.Handle("/puzzles", service.NewPuzzleService())
	mux.Handle("/checker", service.NewCheckerService())
	mux.Handle("/solver", service.NewSolverService())

	http.ListenAndServe("127.0.0.1:9988", mux)
}
