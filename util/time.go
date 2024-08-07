package util

import "time"

func SimpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func SimpleTime(hours, minutes int) time.Time {
	return time.Date(0, 1, 1, hours, minutes, 0, 0, time.UTC)
}
