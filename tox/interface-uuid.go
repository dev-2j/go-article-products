package tox

import (
	"reflect"

	"github.com/google/uuid"
)

func Uuid(s interface{}) *uuid.UUID {

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

	// uuid ?
	if vs, ok := s.(uuid.UUID); ok {
		if vs == uuid.Nil {
			return nil
		}
		return &vs
	}

	// byte array
	if vs, ok := s.([]uint8); ok {
		vx, ex := uuid.ParseBytes(vs)
		if ex != nil {
			return nil
		}
		if vx == uuid.Nil {
			return nil
		}
		return &vx
	}

	// string
	if vs, ok := s.(string); ok {
		vx, ex := uuid.Parse(vs)
		if ex != nil {
			return nil
		}
		if vx == uuid.Nil {
			return nil
		}
		return &vx
	}

	return nil
}
