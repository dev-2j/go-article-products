package tox

// https://stackoverflow.com/questions/38185916/how-to-convert-interface-to-map

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func _ScanToM(s interface{}) map[string]interface{} {

	// nil ?
	if s == nil {
		return nil
	}

	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		return _ScanToM(reflect.Indirect(reflect.ValueOf(s)).Interface())
	}

	// initial
	var rta interface{} = s

	// s is bytes
	if v, ok := rta.([]byte); ok {
		if err := json.Unmarshal(v, &rta); err != nil {
			// logx.Alert(`interface-map#30, %v`, err.Error())
			return nil
		}
	}

	// s is Map
	if reflect.TypeOf(rta).Kind() == reflect.Map {
		itemx := map[string]interface{}{}
		v := reflect.ValueOf(rta)
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			itemx[fmt.Sprintf(`%v`, key.Interface())] = strct.Interface()
		}
		return itemx
	}

	// s is string
	if reflect.TypeOf(rta).Kind() == reflect.String {
		itemx := map[string]interface{}{}
		if err := json.Unmarshal(([]byte)(rta.(string)), &itemx); err != nil {
			// logx.Alert(`interface-map#50, %v`, err.Error())
			return nil
		}
		return itemx
	}

	// s is struct
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		bytes, ex := json.Marshal(s)
		if ex != nil {
			return nil
		}
		var m map[string]interface{}
		if ex := json.Unmarshal(bytes, &m); ex != nil {
			return nil
		}
		return m
	}

	return nil
}
