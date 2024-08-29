package timex

import (
	"time"
)

func FirstOfMonth(d time.Time) time.Time {
	date := time.Date(d.Local().Year(), d.Local().Month(), 1, 0, 0, 0, 0, time.Local)
	return date
}

func IsFirstOfMonth(d time.Time) bool {
	return d.Local().Day() == 1
}
