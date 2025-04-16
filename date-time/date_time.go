package datetime

import "time"

func Get_yyyy_mm() string {
	return time.Now().Format("2006_01")
}

func Sleep(pause int, start time.Time) {
	sleep := (time.Duration(pause) * time.Second) - time.Now().Sub(start)
	time.Sleep(sleep)
}
