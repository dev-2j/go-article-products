package timex

import (
	"strconv"
	"time"
)

// return timestamp ในหน่วย milisec -> string
func GetTimeStamp() string {
	timex := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(timex, 10) // แปลงจาก int64 เป็น string
}
