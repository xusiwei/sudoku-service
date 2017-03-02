all: main test

main:
	go build -v main

test: sudoku_test dao_test

dao_test:
	go test -v dao

sudoku_test:
	go test -v sudoku

