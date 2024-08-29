package tox

import (
	"time"
)

// คืนค่าเฉพาะวันที่เท่านั้น
func DateOnly(s time.Time) *time.Time {
	v := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, s.Location())
	return &v
}

func NowPtr() *time.Time {
	v := time.Now()
	return &v
}
