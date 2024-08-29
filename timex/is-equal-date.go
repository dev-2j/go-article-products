package timex

import "time"

func IsEqualDate(s, d time.Time) bool {
	return s.Local().Format(YYYYsMMsDD) == d.Local().Format(YYYYsMMsDD)
}
