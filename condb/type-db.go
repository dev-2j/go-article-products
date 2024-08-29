package condb

// https://github.com/golang/go/wiki/SQLDrivers
// https://pkg.go.dev/github.com/lib/pq
// https://github.com/denisenkom/go-mssqldb#parameters

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dev-2j/go-article-products/config"
	"github.com/dev-2j/go-article-products/logx"
	"github.com/dev-2j/libaryx/stringx"
	"github.com/getsentry/sentry-go"
)

type DB struct {
	dbName string
	dbKey  string
	driver string
	db     *sql.DB
	// _M_TableEmpty    map[string]Rows
	// _L_TableEmpty    *sync.RWMutex
	InterfaceExecute bool // get command and arg when success

	_M_PrimaryKey map[string][]TabPrimaryR
	_L_PrimaryKey *sync.RWMutex
}

func getServiceQuery(query string) string {
	prefixService := fmt.Sprintf(`/*%s*/`, config.GetServiceName())
	if !strings.HasPrefix(query, prefixService) {
		return fmt.Sprintf(`%s%s`, prefixService, query)
	}
	return query
}

func (s *DB) GetDriver() string {
	return s.driver
}

func (s *DB) GetDatabaseName() string {
	return s.dbName
}

func (s *DB) Begin() (*Tx, error) {
	sqlTx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx: sqlTx}, nil
}

func (s *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	sqlTx, err := s.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf(`[%s]s.db.BeginTx,%s`, s.dbKey, err.Error())
	}
	return &Tx{
		tx: sqlTx,
		db: s,
	}, nil
}

func (s *DB) Close() error {
	return s.db.Close()
}

func (s *DB) Conn(ctx context.Context) (*sql.Conn, error) {
	return s.db.Conn(ctx)
}

func (s *DB) Driver() driver.Driver {
	return s.db.Driver()
}

func (s *DB) Exec(query string, args ...any) (sql.Result, error) {
	return s.ExecContext(context.Background(), query, args...)
}

func (s *DB) ExecLimit(secLimit uint, query string, args ...any) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(secLimit))
	defer cancel()
	resp, ex := s.ExecContext(ctx, query, args...)
	return resp, ex
}

func (s *DB) Ping(secTimeout int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(secTimeout))
	defer cancel()
	return s.db.PingContext(ctx)
}

func (s *DB) PingContext(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *DB) GetInstant() *sql.DB {
	return s.db
}

func (s *DB) Query(query string, args ...any) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), query, args...)
}

func (s *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	start := time.Now()
	key6 := stringx.Rand(6)
	icount := 0
	for {
		resp, err := func() (sql.Result, error) {
			// if s.driver == SQLITE3 {
			// 	// s._L_SQLITE3.Lock()
			// 	// defer s._L_SQLITE3.Unlock()
			// 	query = fmt.Sprintf(`PRAGMA journal_mode=WAL;%s`, query)
			// }
			return s.db.ExecContext(ctx, getServiceQuery(query), args...)
		}()
		if err != nil {
			if _SentryActive {
				go func() {
					_ = sentry.CaptureException(fmt.Errorf(`[%s|%s]s.db.ExecContext,%s`, s.dbKey, key6, err.Error()))
				}()
			}
			if ts, ok := s.retry(err, icount, key6, func() {
				icount++
				if icount == 1 {
					logx.Warnf("[%s|%s]DB:Exec:%s ... retry(%v)\n    %v", s.dbKey, key6, err.Error(), icount, query)
				} else {
					logx.Warnf("[%s|%s]DB:Exec:%s ... retry(%v)", s.dbKey, key6, err.Error(), icount)
				}
			}); ok {
				time.Sleep(time.Millisecond * time.Duration(ts))
				continue
			}
			logx.Alert("[%s|%s]DB:Exec:%s\n    %v", s.dbKey, key6, err.Error(), query)
			// } else {
			// // success
			// if s.InterfaceExecute {
			// 	vx := Exec{
			// 		Query: query,
			// 		Args:  args,
			// 	}
			// 	go func(key6 string, vx Exec) {
			// 		_ = producer.PushMessage(topics.INTERFACE_EXECUTE, key6, vx)
			// 	}(key6, vx)
			// }
		}
		if icount > 0 {
			logx.Infof("[%s|%s]DB:Exec: back success.(%v)", s.dbKey, key6, time.Since(start))
		}
		return resp, err
	}
}

func (s *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	start := time.Now()
	key6 := stringx.Rand(6)
	icount := 0
	for {
		// resp, err := s.db.QueryContext(ctx, getServiceQuery(query), args...)
		resp, err := func() (*sql.Rows, error) {
			return s.db.QueryContext(ctx, getServiceQuery(query), args...)
		}()
		if err != nil {
			if _SentryActive {
				go func() {
					_ = sentry.CaptureException(fmt.Errorf(`[%s|%s]s.db.QueryContext,%s`, s.dbKey, key6, err.Error()))
				}()
			}
			if ts, ok := s.retry(err, icount, key6, func() {
				icount++
				if icount == 1 {
					logx.Warnf("[%s|%s]DB:Query:%s ... retry(%v)\n    %v", s.dbKey, key6, err.Error(), icount, query)
				} else {
					logx.Warnf("[%s|%s]DB:Query:%s ... retry(%v)", s.dbKey, key6, err.Error(), icount)
				}
			}); ok {
				time.Sleep(time.Millisecond * time.Duration(ts))
				continue
			}
			logx.Alert("[%s|%s]DB:Query:%s\n    %v", s.dbKey, key6, err.Error(), query)
		}
		if icount > 0 {
			logx.Infof("[%s|%s]DB:Query: back success.(%v)", s.dbKey, key6, time.Since(start))
		}
		return resp, err
	}
}

func (s *DB) QueryRow(query string, args ...any) *sql.Row {
	return s.db.QueryRow(getServiceQuery(query), args...)
}

func (s *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return s.db.QueryRowContext(ctx, getServiceQuery(query), args...)
}

func (s *DB) SetConnMaxIdleTime(d time.Duration) {
	s.db.SetConnMaxIdleTime(d)
}

func (s *DB) SetConnMaxLifetime(d time.Duration) {
	s.db.SetConnMaxLifetime(d)
}

func (s *DB) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

func (s *DB) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

func (s *DB) Stats() sql.DBStats {
	return s.db.Stats()
}

func (s *DB) CommandInsert(table string, colsList []string) string {

	colsParm := []string{}
	for i := 0; i < len(colsList); i++ {
		colsParm = append(colsParm, fmt.Sprintf(`$%v`, i+1))
	}

	// insert into xxx(c1,c2,c3)values($1,$2,$3)
	return fmt.Sprintf(`INSERT INTO %s(%s)VALUES(%s)`,
		table,
		strings.Join(colsList, `,`),
		strings.Join(colsParm, `,`),
	)

}
