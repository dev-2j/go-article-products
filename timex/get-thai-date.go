package timex

import (
	"fmt"
	"strconv"
	"time"
)

func GetThaiDate(s time.Time) (*string, error) {

	yint, err := strconv.Atoi(s.Format(YYYY))
	if err != nil {
		return nil, fmt.Errorf("convert string to int : %v", err.Error())
	}

	v := fmt.Sprintf(`%v%v`, s.Format("02-01-"), yint+543)
	return &v, nil
}
