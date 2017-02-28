package dao

type PuzzleBean struct {
	Puzzle string `json:"puz"` // primary key
	Level  int    `json:"level"`
	Date   string `json:"date"`
}

func NewPuzzleBean(puz string, level int, date string) *PuzzleBean {
	return &PuzzleBean{Puzzle: puz, Level: level, Date: date}
}

func CopyPuzzleBean(p *PuzzleBean) *PuzzleBean {
	return &PuzzleBean{Puzzle: p.Puzzle, Level: p.Level, Date: p.Date}
}

type PuzzleDao interface {
	Create(puz string, level int, date string) (*PuzzleBean, PuzzleDaoStatus)
	Query(date string, level int) ([]PuzzleBean, PuzzleDaoStatus)
	QueryRandom(n int) ([]PuzzleBean, PuzzleDaoStatus)
	Update(oldPuz, newPuz string) (*PuzzleBean, PuzzleDaoStatus)
	Delete(puz string) (*PuzzleBean, PuzzleDaoStatus)
}

type PuzzleDaoStatus int

const (
	_              = iota
	StatusFound    = iota
	StatusCreated  = iota
	StatusUpdated  = iota
	StatusDeleted  = iota
	errorBase      = 400
	StatusInvliad  = iota + errorBase
	StatusExisted  = iota + errorBase
	StatusNotFound = iota + errorBase
)

// satisfy error interface
func (ec PuzzleDaoStatus) Text() string {
	statusText := map[int]string{
		StatusInvliad:  "invalid",
		StatusExisted:  "existed",
		StatusCreated:  "created",
		StatusUpdated:  "updated",
		StatusDeleted:  "deleted",
		StatusFound:    "found",
		StatusNotFound: "notfound"}
	return statusText[int(ec)]
}
