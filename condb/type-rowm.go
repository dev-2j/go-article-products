package condb

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/dev-2j/libaryx/stringx"
	"gitlab.dohome.technology/dohome-2020/go-servicex/chanx"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Colm struct {
	Index            int    `json:"idx"`
	DatabaseTypeName string `json:"dtn"`
	PrimaryKeyIndex  *int   `json:"pki,omitempty"`
}

type Rowm struct {
	ColumnTypes map[string]Colm `json:"colt"` // key = Column name
	DriverName  string          `json:"drv"`
	Rowm        map[string]Map  `json:"rowi"`
	_L_ROWM     *sync.Mutex
	_M_ROWM     map[string]bool // key ที่เคยแปลง Map แล้ว
}

func NewRowm() *Rowm {
	return &Rowm{
		ColumnTypes: map[string]Colm{},
		Rowm:        map[string]Map{},
		_L_ROWM:     &sync.Mutex{},
		_M_ROWM:     map[string]bool{},
	}
}

func (s *Rowm) New() *Rowm {
	r := NewRowm()
	r.ColumnTypes = s.ColumnTypes
	r.DriverName = s.DriverName
	return r
}

func (s *Rowm) Length() int {
	if s == nil {
		return 0
	}
	return len(s.Rowm)
}

func (s *Rowm) IsExist(kvs ...string) bool {
	k := strings.Join(kvs, `|`)
	_, ok := s.Rowm[k]
	return ok
}

func (s *Rowm) SetRow(m *Map, kvs ...string) {
	if s == nil {
		return
	}
	// key values
	k := strings.Join(kvs, `|`)
	(*s).Rowm[k] = *m
}

func (s *Rowm) GetRow(kvs ...string) *Map {
	if s == nil {
		return nil
	}

	// key values
	k := strings.Join(kvs, `|`)

	// get value
	vm, ok := s.Rowm[k]
	if !ok {
		return nil
	}

	// already converted
	s._L_ROWM.Lock()
	defer s._L_ROWM.Unlock()
	if _, ok := s._M_ROWM[k]; ok {
		return &vm
	}
	s._M_ROWM[k] = true

	// set key value
	for cn, cv := range s.ColumnTypes {
		ci := fmt.Sprintf(`%v`, cv.Index)
		if _, ok := vm[ci]; !ok {
			continue
		}
		vm[cn] = vm[ci]
		if cv.PrimaryKeyIndex != nil {
			if *cv.PrimaryKeyIndex < len(kvs) {
				vm[cn] = kvs[*cv.PrimaryKeyIndex]
			}
		}
		delete(vm, ci)
	}
	return &vm
}

func (s *Rowm) GetColumns() []string {
	cols := []string{}
	for k := range s.ColumnTypes {
		cols = append(cols, k)
	}
	return cols
}

func (s *Rowm) GetRows() []Map {

	rowm := []Map{}
	lock := sync.Mutex{}

	cx, ex := chanx.Create(10, func(a any) error {
		k := a.(string)
		m := s.GetRow(strings.Split(k, `|`)...)
		lock.Lock()
		defer lock.Unlock()
		rowm = append(rowm, *m)
		return nil
	})
	if ex != nil {
		panic(ex)
	}

	if s != nil {
		for k := range s.Rowm {
			if !cx.Send(k) {
				continue
			}
		}
	}

	if ex := cx.Wait(); ex != nil {
		panic(ex)
	}

	return rowm
}

func (s *Rowm) FindRow(match func(*Map) bool) *Map {
	if s == nil {
		return nil
	}
	for k := range s.Rowm {
		m := s.GetRow(strings.Split(k, `|`)...)
		if match(m) {
			return m
		}
	}
	return nil
}

func (s *Rowm) FilterRow(match func(*Map) bool) *Rowm {
	if s == nil {
		return nil
	}
	r := s.New()
	for k, v := range s.Rowm {
		m := s.GetRow(strings.Split(k, `|`)...)
		if match(m) {
			r.Rowm[k] = v
		}
	}
	return r
}

// fillter by keys
func (s *Rowm) FilterNew(keys ...string) *Rowm {
	if s == nil {
		return nil
	}
	r := s.New()
	for _, k := range keys {
		if m := s.GetRow(strings.Split(k, `|`)...); m != nil {
			r.Rowm[k] = *m
		}
	}
	return r
}

// remove by keys
func (s *Rowm) RemoveKey(keys ...string) {
	if s == nil {
		return
	}
	for _, k := range keys {
		if m := s.GetRow(strings.Split(k, `|`)...); m != nil {
			delete(s.Rowm, k)
		}
	}
}

// remove by keys
func (s *Rowm) RemoveNew(keys ...string) *Rowm {
	if s == nil {
		return nil
	}
	r := s.New()
	r.Rowm = s.Rowm
	for _, k := range keys {
		if m := s.GetRow(strings.Split(k, `|`)...); m != nil {
			delete(r.Rowm, k)
		}
	}
	return r
}

// sort by PrimaryKeyIndex
func (s *Rowm) JsonRows() map[string][]Map {
	type Colp struct {
		ColName string
		PKIndex int
	}
	keyp := []Colp{}
	for k, v := range s.ColumnTypes {
		if v.PrimaryKeyIndex != nil {
			keyp = append(keyp, Colp{
				ColName: k,
				PKIndex: *v.PrimaryKeyIndex,
			})
		}
	}
	sort.SliceStable(keyp, func(i, j int) bool {
		return keyp[i].PKIndex < keyp[j].PKIndex
	})
	keys := []string{}
	for _, v := range keyp {
		keys = append(keys, v.ColName)
	}
	return s.JsonSort(strings.Join(keys, `,`))
}

// sort by colsSort
func (s *Rowm) JsonSort(colsSort string) map[string][]Map {
	return map[string][]Map{
		`rows`: s.Sort(colsSort),
	}
}

// ได้ค่า v(pointer of struct)
func (s *Rowm) ToStruct(v any) error {

	rows := s.GetRows()

	// convert row to string
	bytes, ex := json.Marshal(rows)
	if ex != nil {
		return ex
	}

	if ex := json.Unmarshal(bytes, v); ex != nil {
		return ex
	}

	return nil
}

func (s *Rowm) Sort(colsSort string) []Map {

	// colsSort = col1, col2 desc

	if s == nil {
		return nil
	}

	rows := s.GetRows()

	type SortMode struct {
		ColName string
		DESC    bool // true=จากมากไปน้อย ,false=จากน้อยไปมาก
	}

	cols := []SortMode{}
	for _, v := range strings.Split(colsSort, `,`) {
		colx := stringx.Split(v, ` `)
		if len(colx) == 2 {
			cols = append(cols, SortMode{
				ColName: colx[0],
				DESC:    strings.ToLower(colx[1]) == `desc`,
			})
		} else {
			cols = append(cols, SortMode{
				ColName: v,
				DESC:    false,
			})
		}

	}

	thai := collate.New(language.Thai)
	sort.SliceStable(rows, func(i, j int) bool {
		for _, v := range cols {
			thCompare := thai.CompareString(rows[i].String(v.ColName), rows[j].String(v.ColName))
			if thCompare == -1 && !v.DESC {
				return true
			}
		}
		return false
	})

	return rows

}

// to rows
func (s *Rowm) ToRows() *Rows {

	// create new rows
	r := NewRows()
	r.DriverName = s.DriverName

	// for k, _ := range s.ColumnTypes {

	// 	r.Columns = append(r.Columns, k)
	// }

	r.Rows = append(r.Rows, s.GetRows()...)

	return r
}
