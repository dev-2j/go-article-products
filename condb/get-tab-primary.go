package condb

import (
	"fmt"
	"strings"
)

type TabPrimaryR struct {
	ConstraintName string `json:"constraint_name"`
	ColumnName     string `json:"column_name"`
	ColumnType     string `json:"column_type"`
}

func (db *DB) GetTabPrimarys(tabname string) ([]string, error) {

	rx, ex := db.GetTabPrimary(tabname)
	if ex != nil {
		return nil, ex
	}
	items := []string{}
	for _, v := range rx {
		items = append(items, v.ColumnName)
	}
	return items, nil
}

func (db *DB) GetTabPrimary(tabname string) ([]TabPrimaryR, error) {

	tabname = strings.ToLower(tabname)

	// lock read only
	if v, ok := func() ([]TabPrimaryR, bool) {
		db._L_PrimaryKey.RLock()
		defer db._L_PrimaryKey.RUnlock()
		v, ok := db._M_PrimaryKey[tabname]
		return v, ok
	}(); ok {
		return v, nil
	}

	// lock read/write
	db._L_PrimaryKey.Lock()
	defer db._L_PrimaryKey.Unlock()
	if v, ok := db._M_PrimaryKey[tabname]; ok {
		return v, nil
	}

	// get table primary from database
	var items []TabPrimaryR
	var err error
	if db.driver == POSTGRES {
		items, err = db.getTabPrimaryPostgres(tabname)
	} else if db.driver == SQLSERVER {
		items, err = db.getTabPrimaryMssql(tabname)
	} else if db.driver == SQLITE3 {
		items, err = db.getTabPrimarySQLite(tabname)
	} else {
		return nil, fmt.Errorf(`not support driver:%s`, db.driver)
	}
	if err != nil {
		return nil, err
	}
	db._M_PrimaryKey[tabname] = items
	return items, nil
}

func (db *DB) getTabPrimaryMssql(tabname string) ([]TabPrimaryR, error) {

	rows, ex := db.QueryScan(`SELECT T.CONSTRAINT_NAME as constraint_name, 
		C.COLUMN_NAME as column_name
	FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS T  WITH (NOLOCK)
	JOIN INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE C  WITH (NOLOCK)
		ON C.CONSTRAINT_NAME = T.CONSTRAINT_NAME
	WHERE C.TABLE_NAME = @p1
	AND T.CONSTRAINT_TYPE = 'PRIMARY KEY'`, tabname)
	if ex != nil {
		return nil, ex
	}

	items := []TabPrimaryR{}
	for _, v := range rows.Rows {
		items = append(items, TabPrimaryR{
			ConstraintName: v.String(`constraint_name`),
			ColumnName:     v.String(`column_name`),
		})
	}
	return items, nil
}

func (db *DB) getTabPrimaryPostgres(tabname string) ([]TabPrimaryR, error) {

	rows, ex := db.QueryScan(`select tco.constraint_name,
	kcu.column_name
from information_schema.table_constraints tco
join information_schema.key_column_usage kcu 
  on kcu.constraint_name = tco.constraint_name
  and kcu.constraint_schema = tco.constraint_schema
  and kcu.constraint_name = tco.constraint_name
where 1=1
and tco.constraint_type = 'PRIMARY KEY'
and kcu.table_schema  = current_schema()
and kcu.table_name = $1
order by kcu.table_name,
	  kcu.ordinal_position`, tabname)
	if ex != nil {
		return nil, ex
	}

	items := []TabPrimaryR{}
	for _, v := range rows.Rows {
		items = append(items, TabPrimaryR{
			ConstraintName: v.String(`constraint_name`),
			ColumnName:     v.String(`column_name`),
		})
	}

	return items, nil
}

func (db *DB) getTabPrimarySQLite(tabname string) ([]TabPrimaryR, error) {

	rows, ex := db.QueryScan(fmt.Sprintf(`PRAGMA table_info('%s')`, tabname))
	if ex != nil {
		return nil, ex
	}

	items := []TabPrimaryR{}
	for _, v := range rows.Rows {
		if ok := v.Int(`pk`) > 0; ok {
			items = append(items, TabPrimaryR{
				ConstraintName: `PrimaryKey`,
				ColumnName:     v.String(`name`),
				ColumnType:     v.String(`type`),
			})
		}
	}

	return items, nil
}
