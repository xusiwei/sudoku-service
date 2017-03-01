package dao

import "fmt"
import "math/rand"

import "sudoku"

type PuzzleMemoryDao struct {
	checker   *sudoku.Checker
	records   []PuzzleBean
	puzzIndex map[string]int
	dateIndex map[string]map[int]bool
}

// constructor
func NewPuzzleMemoryDao() *PuzzleMemoryDao {
	return &PuzzleMemoryDao{checker: sudoku.NewChecker(),
		records:   make([]PuzzleBean, 0),
		puzzIndex: map[string]int{},
		dateIndex: map[string]map[int]bool{}}
}

func dateLevelIndex(date string, level int) string {
	return fmt.Sprintf("%s %d", date, level)
}

// methods
// Satisfy PuzzleDao
func (dao *PuzzleMemoryDao) Create(puz string, level int, date string) (*PuzzleBean, PuzzleDaoStatus) {
	// validate puzzle
	if !dao.checker.CheckPuzzle(puz) { // invalid
		return nil, StatusInvliad
	}

	// check puzzle exists or not
	if index, existed := dao.puzzIndex[puz]; existed { // existed
		return CopyPuzzleBean(&dao.records[index]), StatusExisted
	}

	// allocate new puzzle record
	prec := NewPuzzleBean(puz, level, date)

	// add to record set
	dao.records = append(dao.records, *prec)
	index := len(dao.records) - 1

	// add puz index
	dao.puzzIndex[puz] = index

	// add (date, level) index
	dlidx := dateLevelIndex(date, level)
	if _, existed := dao.dateIndex[dlidx]; !existed {
		dao.dateIndex[dlidx] = map[int]bool{}
	}
	dao.dateIndex[dlidx][index] = true

	return prec, StatusCreated
}

func (dao *PuzzleMemoryDao) Query(date string, level int) ([]PuzzleBean, PuzzleDaoStatus) {
	if len(dao.records) == 0 {
		return nil, StatusNotFound
	}

	dlidx := dateLevelIndex(date, level)

	results := make([]PuzzleBean, 0)
	for index, _ := range dao.dateIndex[dlidx] {
		results = append(results, dao.records[index])
	}

	if len(results) == 0 {
		return nil, StatusNotFound
	} else {
		return results, StatusFound
	}
}

func (dao *PuzzleMemoryDao) QueryRandom(n int) ([]PuzzleBean, PuzzleDaoStatus) {
	if len(dao.records) == 0 || n < 1 {
		return nil, StatusNotFound
	}

	if n > len(dao.records) {
		n = len(dao.records)
	}

	start := rand.Intn(len(dao.records) - n + 1)
	return dao.records[start : start+n], StatusFound
}

func (dao *PuzzleMemoryDao) Update(oldPuz, newPuz string) (*PuzzleBean, PuzzleDaoStatus) {
	if len(dao.records) == 0 {
		return nil, StatusNotFound
	}
	if _, existed := dao.puzzIndex[oldPuz]; !existed { // oldPuz not found
		return nil, StatusNotFound
	}
	if !dao.checker.CheckPuzzle(newPuz) { // newPuz invalid
		return nil, StatusInvliad
	}
	if _, existed := dao.puzzIndex[newPuz]; existed { // newPuz existed
		return nil, StatusExisted
	}

	prec := &dao.records[dao.puzzIndex[oldPuz]]
	prec.Puzzle = newPuz
	return CopyPuzzleBean(prec), StatusUpdated
}

func (dao *PuzzleMemoryDao) Delete(puz string) (*PuzzleBean, PuzzleDaoStatus) {
	if len(dao.records) == 0 {
		return nil, StatusNotFound
	}

	index, found := dao.puzzIndex[puz]
	if !found { // not found
		return nil, StatusNotFound
	}

	rec := dao.records[index]
	dlidx := dateLevelIndex(rec.Date, rec.Level)
	lastIdx := len(dao.records) - 1
	lastRec := dao.records[lastIdx]
	lastDlidx := dateLevelIndex(lastRec.Date, lastRec.Level)

	// remove index firstly
	delete(dao.puzzIndex, puz)
	delete(dao.dateIndex[dlidx], index)

	if index != lastIdx {
		// move the last record to this slot
		dao.records[index] = lastRec

		// update the index of moved record
		dao.puzzIndex[lastRec.Puzzle] = index
		dao.dateIndex[lastDlidx][index] = true
	}
	// resize record slice
	dao.records = dao.records[:lastIdx]
	return &rec, StatusDeleted
}
