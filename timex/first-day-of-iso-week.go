package timex

import "time"

// ถ้าส่ง timeZone=time.Local มาจะได้ค่าเป็นวันจันทร์ของสัปดาห์นั้นๆ
func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {

	// https://xferion.com/golang-reverse-isoweek-get-the-date-of-the-first-day-of-iso-week/

	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		_, isoWeek = date.ISOWeek()
	}

	return date
}
