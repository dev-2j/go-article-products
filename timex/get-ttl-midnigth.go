package timex

import "time"

// เที่ยงคืน
func GetTTLMidnight(ttl time.Duration) time.Duration {

	ttx, _ := GetTTLMidnight2(ttl)
	return ttx
}

func GetTTLMidnight2(ttl time.Duration) (time.Duration, time.Time) {

	nowLocal := time.Now().Local()

	// timeClear = เที่ยงคืน อยากให้เคลียร์ cache เวลานี้ (ทั้งหมด)
	timeMidnight := time.Now().Local().Add(time.Hour * 24)
	timeMidnight = GetDateonly(timeMidnight) // เวลาเที่ยงคืน

	if ttl > 0 {
		tx := nowLocal.Add(ttl)
		if tx.Local().After(timeMidnight.Local()) {
			ttl = timeMidnight.Local().Sub(nowLocal) // ไม่เกินเที่ยงคืน
		}
	} else {
		ttl = timeMidnight.Local().Sub(nowLocal) // ไม่เกินเที่ยงคืน
	}
	return ttl, nowLocal
}
