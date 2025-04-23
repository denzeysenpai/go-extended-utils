package datetime

import (
	"fmt"
	"time"
)

func Get_yyyy_mm() string {
	return time.Now().Format("2006_01")
}

func SleepAt(pause int, start time.Time) {
	sleep := (time.Duration(pause) * time.Second) - time.Now().Sub(start)
	time.Sleep(sleep)
}

func SleepNow(pause int) {
	sleep := (time.Duration(pause) * time.Second) - time.Now().Sub(time.Now())
	time.Sleep(sleep)
}

func DateCode() string {
	var strTemp string
	defer func() {
		if r := recover(); r != nil {
			strTemp = ""
		}
	}()
	now := time.Now()
	calcDateTime := now.Year() - 2016
	calcDateTime = (calcDateTime * 12) + int(now.Month())
	calcDateTime = (calcDateTime * 32) + now.Day()
	strTemp = fmt.Sprintf("%04d", calcDateTime)

	return strTemp
}
