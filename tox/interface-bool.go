package tox

import (
	"reflect"
	"strings"
)

var TRUE bool = true

var FALSE bool = false

func Bool(s interface{}) bool {
	v := BoolPtr(s)
	if v == nil {
		return FALSE
	}
	return *v
}

func BoolPtr(s interface{}) *bool {
	// nil ?
	if s == nil {
		return nil
	}
	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		v := reflect.ValueOf(s)
		if v.IsNil() {
			return nil
		}
		s = reflect.Indirect(v).Interface()
	}
	// bool
	if v, ok := s.(bool); ok {
		return &v
	}
	// int
	if v, ok := s.(int); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// uint
	if v, ok := s.(uint); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// int8
	if v, ok := s.(int8); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// uint8
	if v, ok := s.(uint8); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// int16
	if v, ok := s.(int16); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// uint16
	if v, ok := s.(uint16); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// int32
	if v, ok := s.(int32); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// uint32
	if v, ok := s.(uint32); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// int64
	if v, ok := s.(int64); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// uint64
	if v, ok := s.(uint64); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// float32
	if v, ok := s.(float32); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// float64
	if v, ok := s.(float64); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// string
	if v, ok := s.(string); ok {
		switch strings.TrimSpace(strings.ToUpper(v)) {
		case `1`, `T`, `TRUE`, `Y`, `YES`:
			return &TRUE
		default:
			return &FALSE
		}
	}
	return nil
}
