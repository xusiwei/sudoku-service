package dao

import "log"
import "fmt"
import "time"
import "testing"

var pDao PuzzleDao

func TestInit(t *testing.T) {
	pDao = NewPuzzleMemoryDao()
	if pDao == nil {
		t.Error("create PuzzleMemoryDao failed!")
	}
}

func currentDateStr() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%04%02d%02d", y, m, d)
}

func TestCreate(t *testing.T) {
	createAssert := func(puz string, level int, expect PuzzleDaoStatus) {
		p, s := pDao.Create(puz, level, currentDateStr())
		log.Printf("p: %v, s: %s", p, s.Text())
		if s != expect {
			t.Errorf("create assert failed, expect: %v, puz: %s", expect.Text(), puz)
		}
	}
	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391654", 3, StatusCreated)
	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391654", 3, StatusExisted)

	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391650", 3, StatusCreated)
	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391600", 3, StatusCreated)
	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391000", 3, StatusCreated)
	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007390000", 3, StatusCreated)

	createAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391655", 3, StatusInvliad)
}

func TestQuery(t *testing.T) {
	queryAssert := func(date string, level int, expect PuzzleDaoStatus, size int) {
		ps, s := pDao.Query(date, level)
		log.Printf("ps: %v, s: %s", ps, s.Text())
		if s != expect {
			t.Errorf("query status assert failed, date: %s, level: %d, expect: %s, size: %d", date, level, expect.Text(), size)
		}
		if len(ps) != size {
			t.Errorf("query results assert failed, date: %s, level: %d, expect: %s, size: %d", date, level, expect.Text(), size)
		}
	}
	queryAssert(currentDateStr(), 3, StatusFound, 5)
	queryAssert(currentDateStr(), 999, StatusNotFound, 0)
}

func queryAssert(t *testing.T, n int, expect PuzzleDaoStatus, size int) {
	ps, s := pDao.QueryRandom(n)
	log.Printf("ps: %v, s: %s", ps, s.Text())
	if s != expect {
		t.Errorf("query status assert failed, n: %d, expect: %s, size: %d", n, expect.Text(), size)
	}
	if len(ps) != size {
		t.Errorf("query status assert failed, n: %d, expect: %s, size: %d", n, expect.Text(), size)
	}
}

func TestQueryRandom(t *testing.T) {
	queryAssert(t, 1, StatusFound, 1)
	queryAssert(t, 0, StatusNotFound, 0)
	queryAssert(t, 10, StatusFound, 5)
}

func TestUpdate(t *testing.T) {
	updateAssert := func(oldPuz, newPuz string, expect PuzzleDaoStatus) {
		p, s := pDao.Update(oldPuz, newPuz)
		log.Printf("p: %v, s: %s", p, s.Text())
		if s != expect {
			t.Errorf("update assert failed, oldPuz: %s, newPuz: %s", oldPuz, newPuz)
		}
	}
	updateAssert("631780000200500000004060938742800506980000001305040002109070203453628710007390000",
		"631780000200500000004060938742800506980000001305040002109070203453628710007390050", StatusUpdated)
	updateAssert("631780000200500000004060938742800506980000001305040002109070203453628710007390004",
		"631780000200500000004060938742800506980000001305040002109070203453628710007390600", StatusNotFound)
	updateAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391654",
		"631780000200500000004060938742800506980000001305040002109070203453628710007391000", StatusExisted)
	updateAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391654",
		"631780000200500000004060938742800506980000001305040002109070203453628710007391666", StatusInvliad)
}

func TestDelete(t *testing.T) {
	deleteAssert := func(puz string, expect PuzzleDaoStatus) {
		p, s := pDao.Delete(puz)
		log.Printf("p: %v, s: %s", p, s.Text())
		if s != expect {
			t.Errorf("delete assert failed, puz: %s", puz)
		}
	}
	deleteAssert("631780000200500000004060938742800506980000001305040002109070203453628710007390004", StatusNotFound)
	deleteAssert("631780000200500000004060938742800506980000001305040002109070203453628710007391654", StatusDeleted)
	queryAssert(t, 10, StatusFound, 4)
}
