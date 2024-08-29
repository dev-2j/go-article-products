package timex

import (
	"crypto/rand"
	"math/big"
	"time"
)

func Sleeper(minSec, maxSec int) {

	/*
		minSec = 2, maxSec = 5  : ให้สุ่มถ่วงเวลาตั้งแต่ 2วิ ถึง 5วิ ในระดับ milliseconds
	*/

	// rand maxSec-minSec
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(maxSec-minSec)*1000))

	//time.Sleep(time.Second)  [3s++]
	time.Sleep(time.Millisecond * time.Duration((int64(minSec)*1000)+(n.Int64()+1)))

}
