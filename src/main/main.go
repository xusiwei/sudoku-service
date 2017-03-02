package main

import "os"
import "fmt"
import "log"
import "flag"
import "errors"
import "net/http"
import "service"

var port = flag.Int("port", 9988, "The port you want to bind")
var logfile = flag.String("log", "run.log", "Log file write back")

var ErrHelp = errors.New("flag: help requested")

func main() {

	flag.Parse()

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	if fp, err := os.Create(*logfile); err == nil {
		fmt.Println("loging to file", *logfile)
		log.SetOutput(fp)
	}
	fmt.Println("listen on address", address)

	mux := http.NewServeMux()
	mux.Handle("/puzzles", service.NewPuzzleService())
	mux.Handle("/checker", service.NewCheckerService())
	mux.Handle("/solver", service.NewSolverService())

	http.ListenAndServe(address, mux)
}
