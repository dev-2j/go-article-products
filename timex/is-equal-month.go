package timex

import "time"

// เดือนและปี ตรงกัน โดยไม่สนใจวัน
func IsEqualMonthYear(s, d time.Time) bool {
	return s.Local().Format(YYYYsMM) == d.Local().Format(YYYYsMM)
}

// เดือนและวัน ตรงกัน โดยไม่สนใจปี
func IsEqualMonthDay(s, d time.Time) bool {
	return s.Local().Format(MMsDD) == d.Local().Format(MMsDD)
}

// เดือนเท่านั้น ตรงกัน โดยไม่สนใจปีและวัน
func IsEqualMonthOnly(s, d time.Time) bool {
	return s.Local().Format(MM) == d.Local().Format(MM)
}
