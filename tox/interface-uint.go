package tox

import (
	"fmt"
	"reflect"
	"strconv"
)

var DEFAULT_UINT uint
var DEFAULT_UINT64 uint64

func Uint(s interface{}) uint {
	v := UintPtr(s)
	if v == nil {
		return DEFAULT_UINT
	}
	return *v
}

func UintPtr(s interface{}) *uint {
	v64 := Uint64Ptr(s)
	if v64 == nil {
		return nil
	}
	v := uint(*v64)
	// if v == DEFAULT_UINT {
	// 	return nil // ไม่ใช้เพราะจะทำให้ค่าเป็น null
	// }
	return &v
}

func Uint64(s interface{}) uint64 {
	v := Uint64Ptr(s)
	if v == nil {
		return DEFAULT_UINT64
	}
	return *v
}

func Uint64Ptr(s interface{}) *uint64 {
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
	// convert
	vx, ex := strconv.ParseUint(fmt.Sprintf(`%v`, s), 10, 64)
	if ex != nil {
		return nil
	}
	// if vx == DEFAULT_UINT64 {
	// 	return nil // ไม่ใช้เพราะจะทำให้ค่าเป็น null
	// }
	return &vx
}

func Uint8Ptr(s interface{}) *uint8 {
	v := uint8(Uint(s))
	return &v
}
