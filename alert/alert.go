package alert

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/wasuken/todo-jsay/openjtalk"
)

type IntervalAlert struct {
	Title           string `json:"title"`
	Interval_second int    `json:"interval_second"`
	Count           int    `json:"count"`
}

var alertMap map[string]IntervalAlert

// 設定されたアラートリストを返却する
func GetAlertMap() *map[string]IntervalAlert {
	// アラートファイルを読み込む
	return &alertMap
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func touchHistFile(path string) {
	if !exists(path) {
		f, er := os.Create(path)
		if er != nil {
			panic(er)
		}
		w := csv.NewWriter(f)
		header := []string{"title", "interval_second", "count", "created_at"}
		if err := w.Write(header); err != nil {
			panic(err)
		}
		w.Flush()
		f.Close()
	}
}
func (ia IntervalAlert) ToHistRow() []string {
	t := time.Now()
	s := ""
	s = t.String()
	return []string{ia.Title, strconv.Itoa(ia.Interval_second), strconv.Itoa(ia.Count), s}
}

func WriteAlertHist(alt IntervalAlert) {
	path := "./hist.csv"
	touchHistFile(path)

	f, er := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if er != nil {
		panic(er)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write(alt.ToHistRow()); err != nil {
		panic(err)
	}
	w.Flush()
}

// アラートを実行する
func AddAlert(alt IntervalAlert) {
	if alt.Count <= 0 {
		return
	}
	if alertMap == nil {
		alertMap = map[string]IntervalAlert{}
	}
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	done := make(chan bool)
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	uid := uuidObj.String()
	alertMap[uid] = alt

	go func() {
		for {
			select {
			case <-done:
				delete(alertMap, uid)
				openjtalk.Jsay(alt.Title)
				return
			}
		}
	}()

	time.Sleep(time.Duration(alt.Interval_second) * time.Second)
	ticker.Stop()
	done <- true
	fmt.Printf("Alert! %s, cnt: %d \n", alt.Title, alt.Count)
	alt.Count -= 1
	AddAlert(alt)
}
