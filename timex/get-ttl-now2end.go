package timex

import (
	"time"
)

// เวลาที่เหลืออยู่ นับจากนี้
func GetTTLNow2End(end time.Time) time.Duration {
	return end.Local().Sub(time.Now().Local())
}
