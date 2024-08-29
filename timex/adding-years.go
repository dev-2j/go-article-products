package timex

import "time"

func AddingYears(dateFr, dateTo time.Time, callback func(time.Time) error) error {

	dateFr = dateFr.Local()
	dateTo = dateTo.Local()

	yearEnd := dateTo.Format(YYYY) // YYYY

	date := dateFr
	for {
		// validate
		if date.Format(YYYY) > yearEnd {
			break
		}
		// process
		if ex := callback(date); ex != nil {
			return ex
		}
		// next year
		date = AddMonths(date, 12)
	}

	return nil
}
