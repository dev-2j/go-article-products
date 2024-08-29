package timex

import "time"

// วันที่เมื่อวาน
func GetYesterday() *time.Time {
	year, month, day := time.Now().Local().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	date = date.AddDate(0, 0, -1) // เอาของเมื่อวาน
	return &date
}
