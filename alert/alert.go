package alert

import (
	"fmt"
	"time"

	"github.com/wasuken/todo-jsay/openjtalk"
	"github.com/google/uuid"
)

type IntervalAlert struct {
	Title            string
	Interval_second int
	Count            int
}

var alertMap map[string]IntervalAlert

// 設定されたアラートリストを返却する
func GetAlertMap() (*map[string]IntervalAlert) {
	// アラートファイルを読み込む
	return &alertMap
}

// 本日中に期限を迎えるアラートのリストを返却する
func ListTodayAlert() {
}

// アラートを実行する
func AddAlert(alt IntervalAlert) {
	if alt.Count <= 0 {
		return
	}
	if alertMap == nil{
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
