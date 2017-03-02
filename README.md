# sudoku-service
Sudoku web service

## quick start

1. download source code here
2. build and run test

```sh
GOPATH=`pwd` make
```

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
		* `date=`YYYYmmDD, Optional
	* Returns: `{result: puzzle, status:[created|existed|invalid]}`, puzzle is:
		* `{puz: serialized_sudoku_puzzle, date: release_date, level: n}`
		* or `null`, if `status` is `invalid`
* query some sudoku puzzle
	* Method: `GET`
	* Parameters:
		* `date=`YYYYmmDD, request a sudoku puzzle released on a specific `date`
		* `level=`*n*, level of a puzzle, default value `0`
		* `rand=`[0|1], Optional, request a random sudoku puzzle
			* requested with `rand=1`, severr will ignore `date` parameter
	* Returns:
		* `{result: puzzle_array, status:[found|notfound]}`, element of `puzzle_array`:
		* `{puz: serialized_sudoku_puzzle, date: release_date, level: n}`
		* or `null`, if `status` is `notfound`
* update an exist sudoku puzzle
	* Method: `PATCH`
	* Parameters:
		* `puz=`*serialized_sudoku_puzzle*,
		* `new=`*serialized_sudoku_puzzle*,
	* Returns:
		* `{result: puzzle, status:[updated|notfound|existed|invalid]}`, puzzle is:
		* `{puz: serialized_sudoku_puzzle, date: release_date, level: n}`
		* or `null`, if `status` is `invalid`
* delete an exist sudoku puzzle
	* Method: `DELETE`
	* Parameters:
		* `puz=`*serialized_sudoku_puzzle*
	* Returns: `{result: puzzle, status:[deleted|notfound]}`, puzzle is:
		* `{puz: serialized_sudoku_puzzle, date: release_date, level: n}`
		* or `null`, if `status` is `notfound`

## sudoku puzzle serailization

* a row scaned one line string
* use `0` to represent emtpy cell

### example

| 1 | 6 | 3 | 9 | 2 | 7 | 8 | 4 | 5 |
|---|---|---|---|---|---|---|---|---|
| 8 | 9 | 4 | 3 | 5 | 6 | 1 | 7 | 2 |
| 7 | 2 | 5 | 4 | 1 | 8 | 9 | 6 | 3 |
| 5 | 8 | 1 | 7 | 6 | 2 | 4 | 3 | 9 |
| 2 | 3 | 6 | 8 | 9 | 4 | 5 | 1 | 7 |
| 4 | 7 | 9 | 5 | 3 | 1 | 6 | 2 | 8 |
| 6 | 1 | 7 | 2 | 8 | 5 | 3 | 9 | 4 |
| 9 | 4 | 8 | 1 | 7 | 3 | 2 | 5 | 6 |
| 3 | 5 | 2 | 6 | 4 | 9 | 7 | 8 | 1 |

It's a full filled sudoku, will serialized as:

`163927845894356172725418963581762439236894517479531628617285394948173256352649781`


**read more about sudoku**: https://en.wikipedia.org/wiki/Sudoku
