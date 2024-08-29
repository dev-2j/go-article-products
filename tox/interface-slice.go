package tox

import (
	"encoding/json"
	"reflect"

	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
)

func Slice(s interface{}) []interface{} {

	// s is nil
	if s == nil {
		return nil
	}

	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		return Slice(reflect.Indirect(reflect.ValueOf(s)).Interface())
	}

	// s is value
	if vx, ok := s.([]interface{}); ok {
		return vx
	}

	// initial
	var rta interface{} = s

	// s is bytes
	if v, ok := rta.([]byte); ok {
		if err := json.Unmarshal(v, &rta); err != nil {
			logx.Alert(err.Error())
			return nil
		}
	}

	// s is string
	if reflect.TypeOf(rta).Kind() == reflect.String {
		var itemx []interface{}
		if err := json.Unmarshal(([]byte)(rta.(string)), &itemx); err != nil {
			logx.Alert(err.Error())
			return nil
		}
		return itemx
	}

	// s is slice ?
	if reflect.TypeOf(rta).Kind() == reflect.Slice {
		var itemx []interface{}
		v := reflect.ValueOf(rta)
		for i := 0; i < v.Len(); i++ {
			itemx = append(itemx, v.Index(i).Interface())
		}
		return itemx
	}

	return nil

}
