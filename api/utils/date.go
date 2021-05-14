package utils

import "time"

func PtrDate() *time.Time {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	return &date
}
