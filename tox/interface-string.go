package tox

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.dohome.technology/dohome-2020/go-servicex/timex"
)

var STRING string

func StringDefault(s interface{}, d string) string {
	v := StringTrim(s)
	if v == STRING {
		return d
	}
	return v
}

func StringTrim(s interface{}) string {
	return strings.TrimSpace(String(s))
}

func StringTrimPtr(s interface{}) *string {
	p := StringPtr(s)
	if p == nil {
		return nil
	}
	v := strings.TrimSpace(*p)
	return &v
}

func StringEmptNil(s interface{}) *string {
	v := StringTrim(s)
	if v == `` {
		return nil
	}
	return &v
}

func String(s interface{}) string {
	v := StringPtr(s)
	if v == nil {
		return STRING
	}
	return *v
}

func StringPtr(s interface{}) *string {

	// nil ?
	if s == nil {
		return nil
	}

	// string
	if v, ok := s.(string); ok {
		return &v
	}

	// pointer ?
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(s)
		if rv.IsNil() {
			return nil
		}
		s = reflect.Indirect(rv).Interface()
	}

	// time.time
	if vt, ok := s.(time.Time); ok {
		// https://docs.aws.amazon.com/redshift/latest/dg/automatic-recognition.html
		// https://stackoverflow.com/questions/33119748/convert-time-time-to-string
		// StampMicro
		// https://docs.aws.amazon.com/redshift/latest/dg/r_DATEFORMAT_and_TIMEFORMAT_strings.html
		vs := vt.Format(timex.FULLTIMETZ) // ให้ได้ค่าตาม timezone ของ value จริงๆ
		return &vs
	}

	// math/big.Rat
	if vs, ok := s.(big.Rat); ok {
		vx, _ := strconv.ParseFloat(vs.FloatString(18), 64)
		sv := fmt.Sprintf(`%v`, vx)
		return &sv
	}

	// uuid ?
	if vs, ok := s.(uuid.UUID); ok {
		if vs == uuid.Nil {
			return nil
		}
		vx := vs.String()
		return &vx
	}

	// byte array
	if vs, ok := s.([]uint8); ok {
		vx := string(vs)
		return &vx
	}

	// kind
	kind := reflect.TypeOf(s).Kind()

	// pointer ?
	if kind == reflect.Ptr {
		rv := reflect.ValueOf(s)
		if rv.IsNil() {
			return nil
		}
		s = reflect.Indirect(rv).Interface()
	}

	// float32, float64
	if kind == reflect.Float32 || kind == reflect.Float64 {
		// https://stackoverflow.com/questions/31289409/format-a-float-to-n-decimal-places-and-no-trailing-zeros
		// 10370090.000000  -> 10370090
		var v string
		if vs, ok := s.(float32); ok {
			v = strconv.FormatFloat(float64(vs), 'f', -1, 64)
		} else if vs, ok := s.(float64); ok {
			v = strconv.FormatFloat(vs, 'f', -1, 64)
		} else {
			v = fmt.Sprintf(`%f`, s)
		}
		return &v
	}

	// default
	v := fmt.Sprintf(`%v`, s)
	return &v
}
