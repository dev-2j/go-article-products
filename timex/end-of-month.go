package timex

import (
	"time"
)

func EndOfMonth(d time.Time) time.Time {
	date := FirstOfMonth(d.Local())
	date = date.Local().AddDate(0, 1, 0)
	date = date.Local().AddDate(0, 0, -1)
	return date
}

func IsEndOfMonth(d time.Time) bool {
	return d.Local().Day() == EndOfMonth(d.Local()).Day()
}
