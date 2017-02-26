# sudoku-service
Sudoku web service

## resources

### sovler

* Path: `/sovler`
* solve a sudoku puzzle
	* Method: `GET`
	* Parameters: `puz=`*serailized_sudoku_puzzle*
	* Returns:`{puz: requested_puzzle, solutions: solution_array}`
		* `solution_array`: an array with all solutions,
		* or `null` when the `requested_puzzle` is invalid.

### checker

* Path: `/checker`
* check a sudoku puzzle is valid or not
	* Method: `GET`
	* Parameters: `puz=`*serailized_sudoku_puzzle*
	* Returns: `{puz: requested_puzzle, valid: [true|false]}`

### puzzles

* Path: `/puzzles`
* create a new sudoku puzzle
	* Method: `POST`
	* Parameters: `puz=`*serailized_sudoku_puzzle*, or `id=`puzid
	* Returns: `{puz: requested_puzzle, id:puzid, status:[created|existed|invalid]}`
* query some sudoku puzzles
	* Method: `GET`
	* Parameters:
		* `start=`*puzid* the minial puzzle id(`puzid`) of result set, default 0
		* `limits=`*n* return result limits, default value `5`
			* `start` and `limits` to support page query
		* `level=`*n* level of a puzzle, default value `0`
	* Returns: at most `n` puzzles which `puzid >= requested_puzid`.
		* format: `{level: n, start:puzid, limits:n, puzzles: puzzle_array}`
		* `puzzle_array`: array of puzzle(`{puz:serailized_sudoku_puzzle, id:puzid}`)
* update an exist sudoku puzzle
	* Method: `PATCH`
	* Parameters: `puz=`*serailized_sudoku_puzzle*, and `id=`*puzid*
	* Returns: `{puz: requested_puzzle, id:puzid, status:[updated|existed|invalid]}`
* delete an exist sudoku puzzle
	* Method: `DELETE`
	* Parameters: `puz=`*serailized_sudoku_puzzle*
	* Returns: `{puz: requested_puzzle, id:puzid, status:[deleted|notfound]}`

## sudoku puzzle serailization

* a row scaned one line string
* use `0` to represent emtpy cell

### example

| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
|---|---|---|---|---|---|---|---|---|
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |

It's an invalid sudoku puzzle, serialed as:

`012345678012345678012345678012345678012345678012345678012345678012345678012345678`


**read more about sudoku**: https://en.wikipedia.org/wiki/Sudoku
