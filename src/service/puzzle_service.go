package service

import "log"
import "fmt"
import "time"
import "strconv"
import "net/http"
import "encoding/json"

import "dao"

type PuzzleService struct {
	puzDao dao.PuzzleDao
}

type PuzzleResult struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

// constructor
func NewPuzzleService() *PuzzleService {
	return &PuzzleService{puzDao: dao.NewPuzzleMemoryDao()}
}

// methods
/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (ps *PuzzleService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {

	}

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "parse form failed!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pzbean interface{}
	var status dao.PuzzleDaoStatus

	puz := request.FormValue("puz")
	date := request.FormValue("date")
	level := request.FormValue("level")
	rand := request.FormValue("rand")
	newPuz := request.FormValue("new")

	log.Printf("puz: %s, date: %s, level: %s, rand: %s, newPuz: %s",
		puz, date, level, rand, newPuz)

	switch request.Method {
	case http.MethodPost:
		{
			if puz == "" {
				fmt.Fprintf(writer, "need puz argement")
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			if date == "" {
				y, m, d := time.Now().Date()
				date = fmt.Sprintf("%04d%02d%02d", y, m, d)
			}
			if level == "" {
				level = "0"
			}

			leVal, err := strconv.Atoi(level)
			if err != nil {
				fmt.Fprintf(writer, "invalid level argument: %s", err.Error())
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			pzbean, status = ps.puzDao.Create(puz, leVal, date)
		}
	case http.MethodGet:
		{
			if randOpt, _ := strconv.Atoi(rand); randOpt > 0 {
				pzbean, status = ps.puzDao.QueryRandom(randOpt)
			} else {
				if level == "" {
					level = "0"
				}

				leVal, err := strconv.Atoi(level)
				if err != nil {
					fmt.Fprintf(writer, "invalid level argument: %s", err.Error())
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				pzbean, status = ps.puzDao.Query(date, leVal)
			}
		}
	case http.MethodPatch:
		{
			if puz == "" {
				fmt.Fprintf(writer, "need puz argement")
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			pzbean, status = ps.puzDao.Update(puz, newPuz)
		}
	case http.MethodDelete:
		{
			if puz == "" {
				fmt.Fprintf(writer, "need puz argement")
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			pzbean, status = ps.puzDao.Delete(puz)
		}
	default:
		fmt.Fprintf(writer, "method not support!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("bean: %v, status: %s", pzbean, status.Text())

	result := PuzzleResult{}
	result.Code = int(status)
	result.Status = status.Text()
	result.Result = pzbean

	jstr, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(writer, "json marshal failed!\r\n")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("return json: %s", jstr)
	fmt.Fprintf(writer, "%s\r\n", jstr)
	writer.WriteHeader(http.StatusOK)
}
