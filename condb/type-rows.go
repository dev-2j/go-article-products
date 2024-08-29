package condb

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/dev-2j/go-article-products/tox"
	"github.com/dev-2j/libaryx/stringx"
	"github.com/dev-2j/libaryx/validx"
)

type Rows struct {
	TableName  string            `json:"tbl,omitempty"` // TableEmpty
	Columns    []string          `json:"cols,omitempty"`
	DriverName string            `json:"drv,omitempty"`
	Rows       []Map             `json:"rows,omitempty"`
	RMap       map[string][]int  `json:"rmap,omitempty"`
	LockRows   *sync.Mutex       `json:"-"`
	LockRMap   *sync.Mutex       `json:"-"`
	_FMap      func(*Map) string `json:"-"`
}

func NewRows() *Rows {
	return (&Rows{}).New()
}

func (s *Rows) New() *Rows {
	return &Rows{
		TableName: s.TableName,
		// ColumnTypes: s.ColumnTypes,
		Columns:    s.Columns,
		DriverName: s.DriverName,
		Rows:       []Map{},
		RMap:       map[string][]int{},
		LockRows:   &sync.Mutex{},
		LockRMap:   &sync.Mutex{},
	}
}

func (s *Rows) Length() int {
	if s == nil {
		return 0
	}
	return len(s.Rows)
}

func (s *Rows) GetRows(idx int) *Rows {
	if idx >= len(s.Rows) || idx < 0 {
		return nil
	}
	return &Rows{
		// ColumnTypes: s.ColumnTypes,
		Columns:    s.Columns,
		Rows:       []Map{s.Rows[idx]},
		DriverName: s.DriverName,
	}
}

func (s *Rows) GetRow(idx int) *Map {
	if idx >= len(s.Rows) || idx < 0 {
		return nil
	}
	return &s.Rows[idx]
}

func (s *Rows) RemoveIndex(indexes ...int) int {

	if len(indexes) == 0 {
		return 0
	}

	s.LockRows.Lock()
	defer s.LockRows.Unlock()

	rows := []Map{}
	for k, v := range s.Rows {
		found := false
		for _, idx := range indexes {
			if k == idx {
				found = true
				break
			}
		}
		if found {
			continue // remove index
		}
		rows = append(rows, v)
	}
	s.Rows = rows

	return len(indexes)
}

// อ่านค่าคอลัมล์ ตาม colType
func (s *Rows) GetValue(m Map, col string) any {

	var val any

	if v, ok := m[col]; ok {
		val = v // หาตามชื่อคอลัมล์ case-sensitivity.
	} else if v, ok := m[strings.ToLower(col)]; ok {
		val = v // หาตามชื่อคอลัมล์ case-lower
	} else if v, ok := m[strings.ToUpper(col)]; ok {
		val = v // หาตามชื่อคอลัมล์ case-upper
	} else {
		colSnake := stringx.SnakeCase(col)
		if col == colSnake {
			val = v // SnakeCase
		}
	}

	// if val != nil {
	// for k, v := range s.Columns {
	// 	if strings.EqualFold(v, col) {
	// 		return tox.BoolPtr(val)

	// 		// if len(s.ColumnTypes) > k {
	// 		// 	colType := s.ColumnTypes[k]

	// 		// }
	// 	}
	// }
	// }

	return val
}

func (s *Rows) RemoveFind(match func(*Map) bool) int {
	// ฟังก์ชั่นที่จะบอกว่ารายการนี้ลบ
	idxs := []int{}
	for i := range s.Rows {
		if match(&s.Rows[i]) {
			idxs = append(idxs, i)
		}
	}
	return s.RemoveIndex(idxs...)
}

// หารายการแรกที่เจอ
func (s *Rows) FindRow(match func(*Map) bool) *Map {
	if s == nil {
		return nil
	}
	for i := range s.Rows {
		if match(&s.Rows[i]) {
			return &s.Rows[i]
		}
	}
	return nil
}

// กรองเอารายการที่ตรงตาม match
func (s *Rows) Filter(match func(*Map) bool) *Rows {
	rows := s.New()
	for i := range s.Rows {
		if match(&s.Rows[i]) {
			rows.Rows = append(rows.Rows, s.Rows[i])
		}
	}
	return rows
}

// สร้างรายการค้นหา
func (s *Rows) BuildMap(genkey func(*Map) string) *Rows {

	s.LockRMap.Lock()
	defer s.LockRMap.Unlock()

	s.RMap = map[string][]int{}
	s._FMap = genkey

	for i := range s.Rows {
		key := s._FMap(&s.Rows[i])
		if _, ok := s.RMap[key]; !ok {
			s.RMap[key] = []int{}
		}
		s.RMap[key] = append(s.RMap[key], i)
	}

	return s
}

// กรองเอารายการที่ตรงตาม match
func (s *Rows) FilterMap(key string) *Rows {

	rows := s.New()

	if ids, ok := s.RMap[key]; ok {
		for _, v := range ids {
			if v < len(s.Rows) {
				rows.Rows = append(rows.Rows, s.Rows[v])
			}
		}
	}

	return rows
}

// ค้นหาจากรายการที่สร้างไว้
func (s *Rows) FindMap(key string) *Map {
	if s.RMap == nil {
		return nil
	}
	if idx, ok := s.RMap[key]; ok {
		if len(idx) > 0 {
			if idx[0] < len(s.Rows) {
				return &s.Rows[idx[0]] // เอารายการแรกที่เจอ
			}
		}
	}
	return nil
}

// ค้นหาจากรายการที่สร้างไว้ case-insensitivity.
func (s *Rows) FindMapEqualFold(key string) *Map {
	if s.RMap == nil {
		return nil
	}
	// match case
	if idx, ok := s.RMap[key]; ok {
		if len(idx) > 0 {
			if idx[0] < len(s.Rows) {
				return &s.Rows[idx[0]]
			}
		}
	}
	// small case
	if idx, ok := s.RMap[strings.ToLower(key)]; ok {
		if len(idx) > 0 {
			if idx[0] < len(s.Rows) {
				return &s.Rows[idx[0]]
			}
		}
	}
	// upper case
	if idx, ok := s.RMap[strings.ToUpper(key)]; ok {
		if len(idx) > 0 {
			if idx[0] < len(s.Rows) {
				return &s.Rows[idx[0]]
			}
		}
	}
	return nil
}

// append with build map
func (s *Rows) Append(m Map) {

	s.Rows = append(s.Rows, m)

	if s.RMap != nil && s._FMap != nil {

		s.LockRMap.Lock()
		defer s.LockRMap.Unlock()

		key := s._FMap(&m)
		if _, ok := s.RMap[key]; !ok {
			s.RMap[key] = []int{}
		}
		s.RMap[key] = append(s.RMap[key], len(s.Rows)-1)

	}

}

// กรองเอาเฉพาะคอลัมล์ที่ต้องการ
func (s *Rows) FilterCols(cols ...string) *Rows {

	if len(cols) == 1 && strings.Contains(cols[0], `,`) {
		cols = stringx.Split(cols[0], `,`)
	}

	r := s.New()

	// column in rows
	for _, row := range s.Rows {
		m := Map{}
		for _, colx := range cols {

			var colFr, colTo string
			if strings.Contains(colx, `:`) {
				coli := strings.Split(colx, `:`)
				colFr = coli[0]
				colTo = coli[1]
			} else {
				colFr = colx
				colTo = colx
			}

			m.Set(colTo, row.Get(colFr))
		}
		r.Rows = append(r.Rows, m)
	}

	// column in columns
	// r.ColumnTypes = []ColumnType{}
	r.Columns = []string{}
	for _, col := range cols {
		for i := 0; i < len(s.Columns); i++ {
			if strings.EqualFold(col, s.Columns[i]) {
				// r.ColumnTypes = append(r.ColumnTypes, s.ColumnTypes[i])
				r.Columns = append(r.Columns, s.Columns[i])
			}
		}
	}

	return r
}

func (s *Rows) ToStructs(v any) error {
	// return tox.ScanTo(s.Rows, v)
	if validx.IsNil(v) {
		return fmt.Errorf(`ToStructs: v is nil`)
	}
	bytes, ex := json.Marshal(s.Rows)
	if ex != nil {
		return ex
	}
	if ex := json.Unmarshal(bytes, v); ex != nil {
		return ex
	}
	return nil
}

func (s *Rows) AddStructs(v any) error {

	// pointer ?
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return fmt.Errorf(`v: value is nil`)
		}
		v = reflect.Indirect(rv).Interface()
	}

	setValue := func(vr any) error {
		row := Map{}
		if ex := row.SetStruct(vr); ex != nil {
			return ex
		}
		s.Rows = append(s.Rows, row)
		return nil
	}

	rk := reflect.TypeOf(v).Kind()
	if rk == reflect.Slice || rk == reflect.Array {
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.Len(); i++ {
			if ex := setValue(rv.Index(i).Interface()); ex != nil {
				return ex
			}
		}
	} else if rk == reflect.Struct {
		if ex := setValue(v); ex != nil {
			return ex
		}
	} else {
		return fmt.Errorf(`v, not support type: %v`, rk)
	}
	return nil
}

func (s *Rows) JsonRows() map[string]any {
	return map[string]any{
		`rows`: s.Rows,
	}
}

func (s *Rows) Sum(col string) float64 {
	var amt float64
	for i := range s.Rows {
		amt += s.Rows[i].Float(col)
	}
	return amt
}

func (s *Rows) ToRowm(keys ...string) *Rowm {

	r := NewRowm()
	if s == nil {
		return r
	}
	r.DriverName = s.DriverName

	for _, v := range s.Rows {
		rm := Map{}
		kv := make([]string, len(keys))
		for cn, cv := range r.ColumnTypes {
			ci := fmt.Sprintf(`%v`, cv.Index)
			rm[ci] = v.Get(cn)
			if cv.PrimaryKeyIndex != nil {
				kv[*cv.PrimaryKeyIndex] = tox.String(rm[ci])
				rm[ci] = *cv.PrimaryKeyIndex
			}
		}
		r.Rowm[strings.Join(kv, `|`)] = rm
	}

	return r
}

// func isColumnTime(colType string) bool {
// 	return strings.Contains(strings.ToLower(colType), `date`) || strings.Contains(strings.ToLower(colType), `time`)
// }

func (s *Rows) TimeLocal2UTC() {

	// convert from ms to pg (time on sqlserver is +07)
	colsTime := []string{}
	// for _, v := range s.ColumnTypes {
	// 	if isColumnTime(strings.ToLower(v.DatabaseTypeName)) {
	// 		colsTime = append(colsTime, v.Name)
	// 	}
	// }
	// if len(colsTime) == 0 {
	// 	return
	// }

	for i := 0; i < len(s.Rows); i++ {
		for _, colName := range colsTime {
			v1 := s.Rows[i].Time(colName)
			if v1.IsZero() {
				s.Rows[i].Set(colName, nil)
			} else {
				s.Rows[i].Set(colName, v1.UTC())
			}
		}
	}

}

func (s *Rows) TimeUTC2Local() {

	// convert from pg to ms (time on sqlserver is +07)
	colsTime := []string{}
	// for _, v := range s.ColumnTypes {
	// 	if isColumnTime(strings.ToLower(v.DatabaseTypeName)) {
	// 		colsTime = append(colsTime, v.Name)
	// 	}
	// }
	if len(colsTime) == 0 {
		return
	}

	for i := 0; i < len(s.Rows); i++ {
		for _, colName := range colsTime {
			v1 := s.Rows[i].Time(colName)
			if v1.IsZero() {
				s.Rows[i].Set(colName, nil)
			} else {
				s.Rows[i].Set(colName, v1.Local())
			}
		}
	}

}

// calculate hash sum of rows
func (s *Rows) HashSum256(columns ...string) (*string, error) {

	// row not found
	if s == nil {
		return &tox.STRING, nil
	}
	if s.Rows == nil {
		return &tox.STRING, nil
	}

	// columns target
	colsTarget := s.Columns
	if len(columns) > 0 {
		colsTarget = columns
	}

	// create object to hash
	rowx := [][]string{}
	for _, row := range s.Rows {
		rowi := []string{}
		for _, col := range colsTarget {

			val := row.Get(col)

			if val != nil {
				// pointer ?
				if reflect.TypeOf(val).Kind() == reflect.Ptr {
					rv := reflect.ValueOf(val)
					val = reflect.Indirect(rv).Interface()
				}
			}

			rowi = append(rowi, tox.String(val))
		}
		rowx = append(rowx, rowi)
	}

	// marshal to bytes
	bytes, ex := json.Marshal(rowx)
	if ex != nil {
		return nil, ex
	}

	// get check sum from value
	h := fmt.Sprintf(`%x`, sha256.Sum256(bytes))
	return &h, nil

}

func (s *Rows) ToMapAny() []map[string]any {
	rows := []map[string]any{}
	for _, row := range s.Rows {
		rows = append(rows, row)
	}
	return rows
}

func (s *Rows) GetValues(idx int) Map {
	if idx >= len(s.Rows) || idx < 0 {
		return Map{} // ไม่พบรายการ
	}
	mx := Map{}
	for _, col := range s.Columns {
		mx[strings.ToLower(col)] = s.GetValue(s.Rows[idx], col)
	}
	return mx
}

// fieldName to fieldIndex
func (s *Rows) Hash(idx int, cols []string) Map {
	if s == nil {
		return nil
	}
	rx := Map{}
	mx := s.GetValues(idx)
	for k, v := range cols {
		rx[fmt.Sprintf(`%v`, k)] = mx[strings.ToLower(v)]
	}
	return rx
}
