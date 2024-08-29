package tox

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

var FLOAT64 float64

func Float(s interface{}) float64 {
	v := FloatPtr(s)
	if v == nil {
		return FLOAT64
	}
	return *v
}

func FloatPtr(s interface{}) *float64 {
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
	// float64
	if v, ok := s.(float64); ok {
		return &v
	}
	// float32
	if v, ok := s.(float32); ok {
		vx := float64(v)
		return &vx
	}
	// int
	if v, ok := s.(int); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(int16); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(int32); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(int64); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(int8); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(uint); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(uint16); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(uint32); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(uint64); ok {
		vx := float64(v)
		return &vx
	}
	if v, ok := s.(uint8); ok {
		vx := float64(v)
		return &vx
	}
	// []uint8
	if v, ok := s.([]uint8); ok {
		vx, ex := strconv.ParseFloat(string(v), 64)
		if ex != nil {
			return nil
		}
		return &vx
	}
	// math/big.Rat
	if vs, ok := s.(big.Rat); ok {
		vx, ex := strconv.ParseFloat(vs.FloatString(18), 64)
		if ex != nil {
			return nil
		}
		return &vx
	}
	// convert
	vx, ex := strconv.ParseFloat(fmt.Sprintf(`%v`, s), 64)
	if ex != nil {
		return nil
	}
	return &vx
}
