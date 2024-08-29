package timex

import "time"

func AddingMonths(dateFr, dateTo time.Time, callback func(time.Time) error) error {

	dateFr = dateFr.Local()
	dateTo = dateTo.Local()

	monthEnd := dateTo.Format(YYYYMM) // YYYYMM

	date := dateFr
	for {
		// validate
		if date.Format(YYYYMM) > monthEnd {
			break
		}
		// process
		if ex := callback(date); ex != nil {
			return ex
		}
		// next month
		date = AddMonths(date, 1)
	}

	return nil
}
