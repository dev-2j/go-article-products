package condb

// https://golangbyexample.com/singleton-design-pattern-go/

import (
	"strings"
	"sync"

	// _ "github.com/3dsinteractive/wrkgo"
	// _ "github.com/marcboeker/go-duckdb"  // Register duckdb
	// _ "github.com/SAP/go-hdb/driver"     // Register hdb driver.
	_ "github.com/denisenkom/go-mssqldb" // Register mssql driver.
	_ "github.com/go-sql-driver/mysql"   // Register mysql.
	_ "github.com/lib/pq"                // Register pg driver.
	_ "github.com/mattn/go-sqlite3"      // Register sqlite
)

var (
	_M_DBS = map[string]*DB{}
	_L_DBS = &sync.RWMutex{}
)

const (
	PARAMS_LIMIT_POSTGRES = 65535 // // PostgreSQL only supports 65535 parameters
	POSTGRES              = `postgres`
	SQLSERVER             = `sqlserver`
	MYSQL                 = `mysql`
	HANA                  = `hana`
	SAPDQ                 = `sapdq` // sap dynamic query
	SQLITE3               = `sqlite3`
	DUCKDB                = `duckdb`
	PARGUET               = `parquet`
)

var _SentryActive bool

func SetSentryActive(v bool) {
	_SentryActive = v
}

// func xConnection(dbName, dbKey, driver, dsn string) (*DB, error) {
// 	db, ex := getInstance(os.Getpid(), dbKey, driver, dsn)
// 	if db != nil {
// 		db.dbName = dbName
// 	}
// 	return db, ex
// }

// func getConnect(pid int, dbKey, driver, dsn string) (*DB, error) {
// 	return getInstance(pid, dbKey, driver, dsn)
// }

// สำหรับ connection instance จาก dbKey เลย
func GetConnected(dbKey string) (*DB, bool) {

	dbKey = strings.ToUpper(dbKey)

	// lock read only
	if db, ok := func() (*DB, bool) {
		_L_DBS.RLock()
		defer _L_DBS.RUnlock()
		v, ok := _M_DBS[dbKey]
		return v, ok
	}(); ok {
		return db, ok
	}

	// lock read/write
	_L_DBS.Lock()
	defer _L_DBS.Unlock()
	v, ok := _M_DBS[dbKey]
	return v, ok
}

// func getInstance(pid int, dbKey, driver, dsn string) (*DB, error) {

// 	dbKey = strings.ToUpper(dbKey)

// 	// lock read only
// 	if db, ok := func() (*DB, bool) {
// 		_L_DBS.RLock()
// 		defer _L_DBS.RUnlock()
// 		v, ok := _M_DBS[dbKey]
// 		return v, ok
// 	}(); ok {
// 		return db, nil
// 	}

// 	// lock read/write
// 	_L_DBS.Lock()
// 	defer _L_DBS.Unlock()
// 	v, ok := _M_DBS[dbKey]
// 	if ok {
// 		return v, nil
// 	}

// 	db, ex := sql.Open(driver, dsn)
// 	if ex != nil {
// 		return nil, ex
// 	}

// 	// sqlDB.SetConnMaxLifetime(0)
// 	// sqlDB.SetMaxIdleConns(5)
// 	// sqlDB.SetMaxOpenConns(100)
// 	// logx.Warnf(`[%v:%v] db instance, created !!!`, dbKey, pid)
// 	_M_DBS[dbKey] = &DB{
// 		dbKey:  dbKey,
// 		driver: driver,
// 		db:     db,
// 		// _M_TableEmpty: map[string]Rows{},
// 		// _L_TableEmpty: &sync.RWMutex{},
// 		_M_PrimaryKey: map[string][]TabPrimaryR{},
// 		_L_PrimaryKey: &sync.RWMutex{},
// 	}

// 	return _M_DBS[dbKey], nil
// }

func Cleanups() {
	//
	_L_DBS.Lock()
	defer _L_DBS.Unlock()
	for _, v := range _M_DBS {
		_ = v.Close()
	}
}
