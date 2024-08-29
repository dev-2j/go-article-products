package timex

import (
	"time"
)

func AddingDays(dateFr, dateTo time.Time, callback func(time.Time) error) error {

	dateFr = GetDateonly(dateFr.Local())
	dateTo = GetDateonly(dateTo.Local())

	date := dateFr
	for {
		// validate
		if date.Local().After(dateTo.Local()) {
			break
		}
		// process
		if ex := callback(date); ex != nil {
			return ex
		}
		// next day
		date = date.AddDate(0, 0, 1).Local()
	}

	return nil

}
