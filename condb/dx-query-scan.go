package condb

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// ดึงข้อมูลตามจำนวนแถว(rowsScan>0) ที่ต้องการ
func dxQueryScan(driver string, rows *sql.Rows) (*Rows, error) {

	// columns
	cols, _ := rows.Columns()

	// scan
	pointers := make([]interface{}, len(cols))
	container := make([]interface{}, len(cols))
	for i := range pointers {
		pointers[i] = &container[i]
	}

	// column types "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", and "BIGINT".

	// rowRunning := int64(0)
	rowm := []Map{}
	for rows.Next() {

		// scan
		if ex := rows.Scan(pointers...); ex != nil {
			return nil, fmt.Errorf(`rows.Scan,%v`, ex.Error())
		}

		// make row
		m := Map{}
		for i := 0; i < len(cols); i++ {
			// m[cols[i]] = container[i]
			cn := cols[i]
			vc := container[i]
			// byte array
			if vs, ok := vc.([]uint8); ok && len(vs) == 36 {
				vx, ex := uuid.ParseBytes(vs)
				if ex == nil {
					if vx == uuid.Nil {
						vc = nil
					} else {
						vc = vx
					}
				}
			}
			m[cn] = vc
		}

		// add row
		rowm = append(rowm, m)

		// // check for rows
		// if rowsScan > 0 {
		// 	rowRunning++
		// 	if rowRunning == rowsScan {
		// 		break
		// 	}
		// }

		// clear data for next row
		container = make([]interface{}, len(cols))
		for i := range pointers {
			pointers[i] = &container[i]
		}
	}

	rowx := NewRows()
	rowx.Columns = cols
	rowx.Rows = rowm
	rowx.DriverName = driver
	return rowx, nil
}
