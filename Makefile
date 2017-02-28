all: main test

main: setup
	go build -v main

test: sudoku_test dao_test

dao_test: setup
	go test -v dao

sudoku_test: setup
	go test -v sudoku

setup:
	export GOPATH=`pwd`
