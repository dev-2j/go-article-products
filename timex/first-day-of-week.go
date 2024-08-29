package timex

import "time"

func FirstDayOfWeek(date time.Time) time.Time {
	isoYear, isoWeek := date.ISOWeek()
	return FirstDayOfISOWeek(isoYear, isoWeek, date.Location())
}
