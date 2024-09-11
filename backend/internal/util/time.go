package util

import "time"

const timeLayout = "2006-01-02"

func ParseTime(dateStr string) (time.Time, error){
	return time.Parse(timeLayout, dateStr)
}