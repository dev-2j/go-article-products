package timex

import "time"

func FirstOfYear(d time.Time) time.Time {
	date := time.Date(d.Year(), time.Month(1), 1, 0, 0, 0, 0, d.Location())
	return date
}

func IsFirstOfYear(d time.Time) bool {
	return d.YearDay() == 1
}
