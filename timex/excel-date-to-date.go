package timex

import (
	"strconv"
	"time"
)

var ExcelEpoch = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)

func ExcelDateToDate(excelDate string) time.Time {
	var days, _ = strconv.Atoi(excelDate)
	return ExcelEpoch.Add(time.Second * time.Duration(days*86400))
}
