package condb

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dev-2j/go-article-products/timex"
	"github.com/dev-2j/go-article-products/tox"
	"github.com/dev-2j/libaryx/stringx"
	"github.com/google/uuid"
)

type Map map[string]any

// หาค่าคอลัมล์ case-insensitivity , snake-case
func (s *Map) Get(col string) any {
	if s == nil {
		return nil
	}
	val, _ := s.getCase(col, true)
	return val
}

func (s *Map) getCase(col string, snakeCase bool) (any, bool) {
	if s == nil {
		return nil, false
	}
	v, ok := func() (any, bool) {
		// หาตามชื่อคอลัมล์ case-sensitivity.
		if v, ok := (*s)[col]; ok {
			return v, true // found column
		}
		// หาตามชื่อคอลัมล์ case-lower
		if v, ok := (*s)[strings.ToLower(col)]; ok {
			return v, true // found column
		}
		// หาตามชื่อคอลัมล์ case-upper
		if v, ok := (*s)[strings.ToUpper(col)]; ok {
			return v, true // found column
		}
		// หาตามชื่อคอลัมล์ ignore-case
		for k, v := range *s {
			if strings.EqualFold(k, col) {
				return v, true // found column
			}
		}
		return nil, false // not found column
	}()
	if ok {
		return v, true // found column
	}
	// SnakeCase
	if snakeCase {
		colSnake := stringx.SnakeCase(col)
		if colSnake != col && colSnake != strings.ToLower(col) {
			return s.getCase(colSnake, false)
		}
	}
	return nil, false // not found column
}

func (s *Map) Copy() Map {
	if s == nil {
		return nil
	}
	x := Map{}
	for k, v := range *s {
		x[k] = v
	}
	return x
}

func (s *Map) IsExist(col string) bool {
	if s == nil {
		return false
	}
	if _, ok := (*s)[col]; ok {
		return true // หาตามชื่อคอลัมล์ case-sensitivity.
	}
	if _, ok := (*s)[strings.ToLower(col)]; ok {
		return true // หาตามชื่อคอลัมล์ case-lower
	}
	if _, ok := (*s)[strings.ToUpper(col)]; ok {
		return true // หาตามชื่อคอลัมล์ case-upper
	}
	if _, ok := (*s)[stringx.SnakeCase(col)]; ok {
		return true // หาตามชื่อคอลัมล์ SnakeCase
	}
	return false
}

func (s *Map) DeleteKey(key string) {
	if s == nil {
		return
	}
	for k := range *s {
		if strings.EqualFold(k, key) {
			delete(*s, k)
		}
	}
}

func (s *Map) FilterCols(cols ...string) *Map {
	if s == nil {
		return nil
	}
	if len(cols) == 1 && strings.Contains(cols[0], `,`) {
		cols = stringx.Split(cols[0], `,`)
	}

	keyx := []string{}
	for k := range *s {
		for _, c := range cols {
			if stringx.IsEqualFold(k, c) {
				keyx = append(keyx, k)
				break
			}
		}
	}

	colx := Map{}
	for k, v := range *s {
		if stringx.IsContains(keyx, k) {
			colx[k] = v
		}
	}

	return &colx
}

// กำหนดค่าคอลัมล์ case-insensitivity , snake-case โดยถ้าไม่พบคอลัมล์ะจะเพิ่มคอลัมล์ให้
// [true:กำหนดค่าคอลัมล์เดิม]
func (s *Map) Set(col string, val any) {
	if s != nil {
		s.setCase(col, val)
	}
}

func (s *Map) setCase(col string, val any) {
	// หาตามชื่อคอลัมล์ case-sensitivity.
	if _, ok := (*s)[col]; ok {
		(*s)[col] = val
		return
	}
	// หาตามชื่อคอลัมล์ case-lower
	colLower := strings.ToLower(col)
	if colLower != col {
		if _, ok := (*s)[colLower]; ok {
			(*s)[colLower] = val
			return
		}
	}
	// หาตามชื่อคอลัมล์ case-upper
	colUpper := strings.ToUpper(col)
	if colUpper != col {
		if _, ok := (*s)[colUpper]; ok {
			(*s)[colUpper] = val
			return
		}
	}
	// SnakeCase
	colSnake := stringx.SnakeCase(col)
	if colSnake != col && colSnake != colLower {
		if _, ok := (*s)[colSnake]; ok {
			(*s)[colSnake] = val
			return
		}
	}
	// เพิ่มคอลัมล์
	(*s)[col] = val
}

func (s *Map) String(col string) string {
	if s == nil {
		return tox.STRING
	}
	return tox.String(s.Get(col))
}

func (s *Map) StringTrim(col string) string {
	if s == nil {
		return tox.STRING
	}
	return strings.TrimSpace(s.String(col))
}

func (s *Map) Split(col, spliter string) []string {
	return stringx.Split(s.StringTrim(col), spliter)
}

func (s *Map) StringTrimLower(col string) string {
	if s == nil {
		return tox.STRING
	}
	return strings.ToLower(strings.TrimSpace(s.String(col)))
}

func (s *Map) StringPtr(col string) *string {
	if s == nil {
		return &tox.STRING
	}
	return tox.StringPtr(s.Get(col))
}

// read string-tz value to time-zone
func (s *Map) StringTZ(col string) string {
	if s == nil {
		return tox.STRING
	}
	v := s.TimePtr(col)
	if v == nil {
		return tox.STRING
	}
	return v.Format(timex.FULLTIMETZ)
}

func (s *Map) Int(col string) int {
	if s == nil {
		return tox.INT
	}
	return tox.Int(s.Get(col))
}

func (s *Map) IntPtr(col string) *int {
	if s == nil {
		return &tox.INT
	}
	return tox.IntPtr(s.Get(col))
}

func (s *Map) Int64(col string) int64 {
	if s == nil {
		return tox.INT64
	}
	return tox.Int64(s.Get(col))
}

func (s *Map) Int64Ptr(col string) *int64 {
	if s == nil {
		return &tox.INT64
	}
	return tox.Int64Ptr(s.Get(col))
}

func (s *Map) UInt64Ptr(col string) *uint64 {
	if s == nil {
		return &tox.UINT64
	}
	return tox.Uint64Ptr(s.Get(col))
}

func (s *Map) UInt64(col string) uint64 {
	if s == nil {
		return tox.UINT64
	}
	return tox.Uint64(s.Get(col))
}

func (s *Map) Float(col string) float64 {
	if s == nil {
		return tox.FLOAT64
	}
	return tox.Float(s.Get(col))
}

func (s *Map) FloatPtr(col string) *float64 {
	if s == nil {
		return &tox.FLOAT64
	}
	return tox.FloatPtr(s.Get(col))
}

func (s *Map) UUID(col string) *uuid.UUID {
	if s == nil {
		return &uuid.UUID{}
	}
	return tox.Uuid(s.Get(col))
}

func (s *Map) Time(col string) time.Time {
	if s == nil {
		return time.Time{}
	}
	return tox.Time(s.Get(col))
}

func (s *Map) TimePtr(col string) *time.Time {
	if s == nil {
		return &time.Time{}
	}
	v := s.Time(col)
	if v.IsZero() {
		return nil
	}
	return &v
}

// ให้ใช้ (s *Map) Time แทน
func (s *Map) TimeSAP(col string) *time.Time {
	if s == nil {
		return &time.Time{}
	}
	return tox.TimeSAP(s.StringTrim(col))
}

// ให้ใช้ (s *Map) Time แทน
func (s *Map) TimeSAPT(cold, colt string) *time.Time {
	if s == nil {
		return &time.Time{}
	}
	return tox.TimeSAPT(s.StringTrim(cold), s.StringTrim(colt))
}

// ให้ใช้ Time,TimePtr แทน
// func (s *Map) TimePOST(col string) *time.Time {
// 	if s == nil {
// 		return &time.Time{}
// 	}
// 	return tox.TimePOST(s.String(col))
// }

// parquet ใช้เวลาเป็น UTC
func (s *Map) TimeUnixMilli(col string) int64 {
	if s == nil {
		return tox.INT64
	}
	v := s.Time(col)
	if v.IsZero() /*|| v.Year() == 9999*/ {
		return tox.INT64
	}
	// return types.TimeToTIMESTAMP_MILLIS(v, true) // UTC, ปี 9999 จะเป็น 1870 ให้ใช้ UnixMilli แทน
	return v.UnixMilli()
}

func (s *Map) Bool(col string) bool {
	if s == nil {
		return tox.FALSE
	}
	return tox.Bool(s.Get(col))
}

func (s *Map) BoolPtr(col string) *bool {
	if s == nil {
		return &tox.FALSE
	}
	return tox.BoolPtr(s.Get(col))
}

func (s *Map) Join(sep string) string {
	if s == nil {
		return tox.STRING
	}
	items := []string{}
	for k, v := range *s {
		items = append(items, fmt.Sprintf(`%s=%s`, k, tox.String(v)))
	}
	return strings.Join(items, sep)
}

// ได้ค่า v(pointer of struct)
func (s *Map) ToStruct(v any) error {
	if s == nil {
		return fmt.Errorf(`map is nil`)
	}

	// convert row to string
	bytes, ex := json.Marshal(s)
	if ex != nil {
		return ex
	}

	if ex := json.Unmarshal(bytes, v); ex != nil {
		return ex
	}

	return nil
}

func (s *Map) SetStruct(v any) error {
	if s == nil {
		return fmt.Errorf(`map is nil`)
	}
	bytes, ex := json.Marshal(v)
	if ex != nil {
		return ex
	}
	var row Map
	if ex := json.Unmarshal(bytes, &row); ex != nil {
		return ex
	}
	for k, v := range row {
		s.Set(k, v)
	}
	return nil
}

// map ลำดับไม่คงที่ ไม่ควร hash
// func (s *Map) HashSum256() (*string, error) {
// 	if s == nil {
// 		return &tox.STRING, nil
// 	}

func (s *Map) OmitEmpty() Map {
	if s == nil {
		return nil
	}
	z := *s
	for k, v := range z {
		if v == nil {
			delete(z, k)
		} else if x := fmt.Sprintf(`%v`, v); x == `` || x == `<nil>` {
			delete(z, k)
		} else if x, ok := v.(string); ok {
			if x == "" {
				delete(z, k)
			}
		} else if x, ok := v.(*string); ok && x != nil {
			if *x == "" {
				delete(z, k)
			}
		} else if x, ok := v.(int); ok {
			if x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(*int); ok && x != nil {
			if *x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(int64); ok {
			if x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(*int64); ok && x != nil {
			if *x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(float64); ok {
			if x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(*float64); ok && x != nil {
			if *x == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(bool); ok {
			if !x {
				delete(z, k)
			}
		} else if x, ok := v.(*bool); ok && x != nil {
			if !*x {
				delete(z, k)
			}
		} else if x, ok := v.(time.Time); ok {
			if x.IsZero() {
				delete(z, k)
			}
		} else if x, ok := v.(*time.Time); ok {
			if x == nil || x.IsZero() {
				delete(z, k)
			}
		} else if x, ok := v.(uuid.UUID); ok {
			if x == uuid.Nil {
				delete(z, k)
			}
		} else if x, ok := v.(*uuid.UUID); ok {
			if x == nil || *x == uuid.Nil {
				delete(z, k)
			}
		} else if x, ok := v.(map[string]any); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.(Map); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]any); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]Map); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]string); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]int); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]int64); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]float64); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		} else if x, ok := v.([]bool); ok {
			if len(x) == 0 {
				delete(z, k)
			}
		}
	}
	return z
}

func (s Map) GetClausePG() ClauseT {
	return s.GetClause(`$`)
}
func (s Map) GetClauseMS() ClauseT {
	return s.GetClause(`@p`)
}
func (s Map) GetClause(prefix string) ClauseT {

	var count int
	var ands []string
	var args []any

	for k, v := range s {
		count++
		ands = append(ands, fmt.Sprintf(`%v = %v%v`, k, prefix, count))
		args = append(args, v)
	}

	return ClauseT{
		Ands: strings.Join(ands, ` AND `),
		Args: args,
	}
}

// fieldIndex to fieldName
func (s Map) Unhash(cols []string) Map {
	if s == nil {
		return nil
	}
	rx := Map{}
	for k, v := range cols {
		vx, ok := s[fmt.Sprintf(`%v`, k)]
		if ok {
			rx[v] = vx
		}
	}
	return rx
}
