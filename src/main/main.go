package main

import "os"
import "fmt"
import "log"
import "flag"
import "errors"
import "context"
import "net/http"
import "syscall"
import "os/signal"
import "service"

var port = flag.Int("port", 9988, "The port you want to bind")
var host = flag.String("host", "localhost", "The host you want to bind")
var logfile = flag.String("log", "", "Log file write back")

var ErrHelp = errors.New("flag: help requested")

func main() {

	flag.Parse()

	address := fmt.Sprintf("%s:%d", *host, *port)
	if fp, err := os.Create(*logfile); err == nil {
		fmt.Println("loging to file", *logfile)
		log.SetOutput(fp)
	}
	fmt.Println("listen on address", address)

	mux := http.NewServeMux()
	mux.Handle("/", service.NewHomeService())
	mux.Handle("/puzzles", service.NewPuzzleService())
	mux.Handle("/checker", service.NewCheckerService())
	mux.Handle("/solver", service.NewSolverService())

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	done := make(chan bool, 1)

	go func() {
		server.ListenAndServe()
		done <- true
	}()
	log.Printf("server goroutine started.")

	go func() {
		ctx := context.TODO()
		sch := make(chan os.Signal, 1)
		signal.Notify(sch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		select {
		case s := <-sch:
			log.Printf("signal raised: %v", s)
			server.Shutdown(ctx) // go1.8
		}
	}()
	log.Printf("signal monitor goroutine started.")

	log.Printf("waiting for server done...")
	<-done

	log.Printf("server done")
}
