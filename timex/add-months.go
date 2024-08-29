package timex

import "time"

func AddMonths(d time.Time, months int) time.Time {

	// วันแรกของเดือน
	date := FirstOfMonth(d.Local())

	// เพิ่มลดเดือน
	date = date.Local().AddDate(0, months, 0)

	// สิ้นเดือน ของเดือนที่ได้
	edate := EndOfMonth(date).Local()

	// ถ้าวันของที่ส่งมา(d) เป็นวันสิ้นเดือน จะคืนค่าสิ้นเดือนด้วย
	if d.Local().Day() == EndOfMonth(d).Local().Day() {
		return edate
	}

	// ถ้าวันของที่ส่งมา(d) เกินกว่าวันสิ้นเดือนของเดือนที่ได้(date) จะคืนว่าไม่เกินสิ้นเดือน
	if d.Local().Day() > edate.Local().Day() {
		return edate
	}

	// วันเดียวกันของเดือนที่ได้
	return time.Date(date.Year(), date.Month(), d.Day(), 0, 0, 0, 0, d.Location())

}
