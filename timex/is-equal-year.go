package timex

import "time"

func IsEqualYear(s, d time.Time) bool {
	return s.Local().Format(YYYY) == d.Local().Format(YYYY)
}
