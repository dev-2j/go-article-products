package tox

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"gitlab.dohome.technology/dohome-2020/go-servicex/timex"
)

func Time(s interface{}) time.Time {
	vx := TimePtr(s)
	if vx != nil {
		return *vx
	}
	return time.Time{}
}

// DateDDMMYYYY แปลงวันที่ DD-MM-YYYY (Local)
func TimeDMY(s string) *time.Time {
	s = strings.TrimSpace(strings.ReplaceAll(s, "/", "-"))
	d, err := time.ParseInLocation(timex.DDMMYYYY, s, time.Local)
	if err != nil {
		return nil
	}
	return &d

}

// TimeHHMMSS แปลงเวลา HH:MM:SS
func TimeHMS(s string) *time.Time {
	s = strings.TrimSpace(strings.ReplaceAll(s, "-", ":"))
	t, err := time.ParseInLocation(timex.HHMMSS, s, time.Local)
	if err != nil {
		return nil
	}
	return &t
}

// YYYY-MM-DD HH:MI:SS
func TimePOST(s string) *time.Time {
	if len(s) == 10 {
		// YYYY-MM-DD
		t, err := time.ParseInLocation(timex.YYYYMMDD, s, time.Local)
		if err != nil {
			return nil
		}
		return &t
	}

	// YYYY-MM-DD HH:MI:SS
	if len(s) == len(timex.MSSQL) {
		t, err := time.ParseInLocation(timex.MSSQL, s, time.Local)
		if err != nil {
			return nil
		}
		return &t
	}

	// 2006-01-02 15:04:05.000000-0700
	t, err := time.Parse(timex.FULLTIMETZ, s)
	if err != nil {
		return nil
	}

	// x := t.Local()
	// fmt.Println(x)
	// y := t.UTC()
	// fmt.Println(y)
	return &t
}

// yyyymmdd or yyyy-mm-dd
func TimeSAP(s string) *time.Time {

	var yyyy, mm, dd int

	if len(s) == 8 {
		// yyyymmdd
		runes := []rune(s)
		yyyy, _ = strconv.Atoi(string(runes[0:4]))
		mm, _ = strconv.Atoi(string(runes[4:6]))
		dd, _ = strconv.Atoi(string(runes[6:8]))
	} else if len(s) == 10 {
		// yyyy-mm-dd
		s = strings.TrimSpace(strings.ReplaceAll(s, `/`, `-`))
		s = strings.TrimSpace(strings.ReplaceAll(s, `.`, `-`))
		sx := strings.Split(s, `-`)
		if len(sx) == 3 {
			runeSx2 := []rune(sx[2])
			if len(runeSx2) > 2 {
				// dd-mm-yyyy
				yyyy, _ = strconv.Atoi(sx[2])
				mm, _ = strconv.Atoi(sx[1])
				dd, _ = strconv.Atoi(sx[0])
			} else {
				yyyy, _ = strconv.Atoi(sx[0])
				mm, _ = strconv.Atoi(sx[1])
				dd, _ = strconv.Atoi(sx[2])
			}
		}
	}
	if yyyy > 2000 {
		t := time.Date(yyyy, time.Month(mm), dd, 0, 0, 0, 0, time.Local)
		return &t
	}

	return nil
}

// yyyymmdd or yyyy-mm-dd + time
func TimeSAPT(s, t string) *time.Time {

	var yyyy, mm, dd int

	if len(s) == 8 {
		// yyyymmdd
		runes := []rune(s)
		yyyy, _ = strconv.Atoi(string(runes[0:4]))
		mm, _ = strconv.Atoi(string(runes[4:6]))
		dd, _ = strconv.Atoi(string(runes[6:8]))
	} else if len(s) == 10 {
		// yyyy-mm-dd
		s = strings.TrimSpace(strings.ReplaceAll(s, `/`, `-`))
		sx := strings.Split(s, `-`)
		if len(sx) == 3 {
			yyyy, _ = strconv.Atoi(sx[0])
			mm, _ = strconv.Atoi(sx[1])
			dd, _ = strconv.Atoi(sx[2])
		}
	}

	if yyyy > 2000 {

		var hh, mi, ss int
		if len(t) == 6 {
			// hhmiss
			runes := []rune(s)
			hh, _ = strconv.Atoi(string(runes[0:2]))
			mi, _ = strconv.Atoi(string(runes[2:4]))
			ss, _ = strconv.Atoi(string(runes[4:6]))
		} else if len(t) == 8 {
			// hh:mi:ss
			sx := strings.Split(t, `:`)
			if len(sx) == 3 {
				hh, _ = strconv.Atoi(sx[0])
				mi, _ = strconv.Atoi(sx[1])
				ss, _ = strconv.Atoi(sx[2])
			}
		}

		t := time.Date(yyyy, time.Month(mm), dd, hh, mi, ss, 0, time.Local)
		return &t
	}

	return nil
}

func TimePtr(s interface{}) *time.Time {

	// nil ?
	if s == nil {
		return nil
	}

	// pointer ?
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(s)
		if rv.IsNil() {
			return nil
		}
		s = reflect.Indirect(rv).Interface()
	}

	// time
	if vs, ok := s.(time.Time); ok {
		return &vs
	}

	// string
	if vs, ok := s.(string); ok {
		vx, ex := time.Parse(time.RFC3339, vs) // UTC
		if ex == nil {
			return &vx
		}
		vx, ex = time.Parse(timex.FULLTIME, vs) // UTC
		if ex == nil {
			return &vx
		}
		vx, ex = time.Parse(timex.FULLTIMETZ, vs) // TIMEZONE
		if ex == nil {
			return &vx
		}
		if len(vs) == len(timex.SYMD) {
			vs2 := vs
			vs2 = strings.ReplaceAll(vs2, `/`, `-`)
			vs2 = strings.ReplaceAll(vs2, `.`, `-`)
			vx, ex = time.ParseInLocation(timex.SYMD, vs2, time.Local) // YYYYMMDD
			if ex == nil {
				return &vx
			}
			vx, ex = time.ParseInLocation(timex.SDMY, vs2, time.Local) // DDMMYYYY
			if ex == nil {
				return &vx
			}
		}
		if len(vs) == len(timex.YYYYMMDD) {
			vs2 := vs
			vs2 = strings.ReplaceAll(vs2, `/`, `-`)
			vs2 = strings.ReplaceAll(vs2, `.`, `-`)
			vx, ex = time.ParseInLocation(timex.YYYYMMDD, vs2, time.Local) // YYYY-MM-DD
			if ex == nil {
				return &vx
			}
			vx, ex = time.ParseInLocation(timex.DDMMYYYY, vs2, time.Local) // DD-MM-YYYY
			if ex == nil {
				return &vx
			}
		}
		if len(vs) == len(timex.MSSQL) {
			vx, ex = time.ParseInLocation(timex.MSSQL, vs, time.Local) // YYYY-MM-DD HH:MI:SS
			if ex == nil {
				return &vx
			}
		}
		if len(vs) == len(timex.MSSQLT) {
			vx, ex = time.ParseInLocation(timex.MSSQLT, vs, time.Local) // YYYY-MM-DDTHH:MI:SS
			if ex == nil {
				return &vx
			}
		}
		if strings.Contains(vs, `T`) && strings.Contains(vs, `.`) {
			vx, ex = time.ParseInLocation(timex.FULLTIMETM, vs, time.Local) //  YYYY-MM-DDTHH:MI:SS.FFF
			if ex == nil {
				return &vx
			}
		}
		return nil
	}

	return nil

}
