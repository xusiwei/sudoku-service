# sudoku-service
Sudoku web service

## resources

### sovler

* Path: `/sovler`
* solve a sudoku puzzle
	* Method: `GET`
	* Parameters: `puz=`*serialized_sudoku_puzzle*
	* Returns:`{puz: requested_puzzle, solutions: solution_array}`
		* `solution_array`: an array with all solutions,
		* or `null` when the `requested_puzzle` is invalid.

### checker

* Path: `/checker`
* check a sudoku puzzle is valid or not
	* Method: `GET`
	* Parameters: `puz=`*serialized_sudoku_puzzle*
	* Returns: `{puz: requested_puzzle, valid: [true|false], blanks: number_of_blanks}`

### puzzles

* Path: `/puzzles`
* create a new sudoku puzzle
	* Method: `POST`
	* Parameters:
		* `puz=`*serialized_sudoku_puzzle*
		* `level=`puzzle_level
	* Returns: `{puz: requested_puzzle, date:release_date, level:puzzle_level, status:[created|existed|invalid]}`
* query some sudoku puzzle
	* Method: `GET`
	* Parameters:
		* `date=`YYYYmmDD, request a sudoku puzzle released on a specific `date`
		* `level=`*n* level of a puzzle, default value `0`
		* `rand=`[0|1], request a random sudoku puzzle
			* requested with `rand=1`, severr will ignore `date` parameter
	* Returns:
		* `{puzzles: puzzle_array}`, element of `puzzle_array`:
			* `{puz: serialized_sudoku_puzzle, date: release_date, level: n}`
* update an exist sudoku puzzle
	* Method: `PATCH`
	* Parameters:
		* `old=`*serialized_sudoku_puzzle*,
		* `new=`*serialized_sudoku_puzzle*,
	* Returns:
		* `{puz: serialized_sudoku_puzzle, date: release_date, level: n, status:[updated|existed|invalid]}`
* delete an exist sudoku puzzle
	* Method: `DELETE`
	* Parameters:
		* `puz=`*serialized_sudoku_puzzle*
	* Returns: `{puz: serialized_sudoku_puzzle, date: release_date, level: n, status:[deleted|notfound]}`

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
