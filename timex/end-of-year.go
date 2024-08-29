package timex

import "time"

func EndOfYear(d time.Time) time.Time {
	date := time.Date(d.Year(), time.Month(12), 31, 0, 0, 0, 0, d.Location())
	return date
}

func IsEndOfYear(d time.Time) bool {
	return d.YearDay() == EndOfYear(d).YearDay()
}
