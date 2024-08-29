package timex

import "time"

// เอาวันที่เท่านั้น โดยตัดชั่วโมงนาทีวินาทีออก
func GetDateonly(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
