package timex

import "time"

//หาเดือนย้อนหลัง 1 เดือน
func GetLastMonth(s time.Time) (int, int) {

	// วันที่ 1 ของเดือน now จะได้ nowFirst [1/4/2020]
	date := time.Date(s.Year(), s.Month(), 1, 0, 0, 0, 0, s.Location())

	// หาวันที่สุดท้ายของเดือนก่อนนหน้า nowFirst - 1 = lastDayOfLastMonth [31/3/2020]
	date = date.AddDate(0, 0, -1)

	// หาปีเดือนของ lastDayOfLastMonth
	return date.Year(), int(date.Month())
}
