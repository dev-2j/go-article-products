package tox

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

var INT int
var INT64 int64
var UINT64 uint64

func Int(s interface{}) int {
	v := IntPtr(s)
	if v == nil {
		return INT
	}
	return *v
}

func IntPtr(s interface{}) *int {
	v64 := Int64Ptr(s)
	if v64 == nil {
		return nil
	}
	v := int(*v64)
	return &v
}

func Int64(s interface{}) int64 {
	v := Int64Ptr(s)
	if v == nil {
		return INT64
	}
	return *v
}

func Int64Ptr(s interface{}) *int64 {
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
	// int
	if v, ok := s.(int); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(int16); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(int32); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(int8); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(uint); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(uint16); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(uint32); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(uint64); ok {
		vx := int64(v)
		return &vx
	}
	if v, ok := s.(uint8); ok {
		vx := int64(v)
		return &vx
	}
	// []uint8
	if v, ok := s.([]uint8); ok {
		vx, ex := strconv.ParseFloat(string(v), 64)
		if ex != nil {
			return nil
		}
		vs := int64(math.Round(vx))
		return &vs
	}
	// math/big.Rat
	if vs, ok := s.(big.Rat); ok {
		vx, ex := strconv.ParseFloat(vs.FloatString(18), 64)
		if ex != nil {
			return nil
		}
		vs := int64(math.Round(vx))
		return &vs
	}
	// convert to float64
	vx, ex := strconv.ParseFloat(fmt.Sprintf(`%v`, s), 64)
	if ex != nil {
		return nil
	}
	// convert to int
	vs := int64(math.Round(vx))
	return &vs
}
