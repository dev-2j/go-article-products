package condb

import (
	"context"
	"database/sql"
	"time"

	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
	"gitlab.dohome.technology/dohome-2020/go-servicex/stringx"
)

type Tx struct {
	tx *sql.Tx
	db *DB
}

func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.ExecContext(context.Background(), query, args...)
}

func (tx *Tx) ExecLimit(secLimit uint, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(secLimit))
	defer cancel()
	resp, ex := tx.ExecContext(ctx, query, args...)
	return resp, ex
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	key6 := stringx.Rand(6)
	icount := 0
	for {
		resp, err := tx.tx.ExecContext(ctx, query, args...)
		if err != nil {
			if ts, ok := tx.db.retry(err, icount, key6, func() {
				icount++
				if icount == 1 {
					logx.Warnf("[%s|%s]TX:Exec:%s ... retry(%v)\n    %v", tx.db.dbKey, key6, err.Error(), icount, query)
				} else {
					logx.Warnf("[%s|%s]TX:Exec:%s ... retry(%v)", tx.db.dbKey, key6, err.Error(), icount)
				}
			}); ok {
				time.Sleep(time.Millisecond * time.Duration(ts))
				continue
			}
			logx.Alert("[%s|%s]TX:Exec:%s\n    %v", tx.db.dbKey, key6, err.Error(), query)
		}
		if icount > 0 {
			logx.Infof("[%s|%s]TX:Exec: back success.(%v)", tx.db.dbKey, key6, time.Since(start))
		}
		return resp, err
	}
}

func (tx *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tx.QueryContext(context.Background(), query, args...)
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	key6 := stringx.Rand(6)
	icount := 0
	for {
		resp, err := tx.tx.QueryContext(ctx, query, args...)
		if err != nil {
			if ts, ok := tx.db.retry(err, icount, key6, func() {
				icount++
				if icount == 1 {
					logx.Warnf("[%s|%s]TX:Query:%s ... retry(%v)\n    %v", tx.db.dbKey, key6, err.Error(), icount, query)
				} else {
					logx.Warnf("[%s|%s]TX:Query:%s ... retry(%v)", tx.db.dbKey, key6, err.Error(), icount)
				}
			}); ok {
				time.Sleep(time.Millisecond * time.Duration(ts))
				continue
			}
			logx.Alert("[%s|%s]TX:Query:%s\n    %v", tx.db.dbKey, key6, err.Error(), query)
		}
		if icount > 0 {
			logx.Infof("[%s|%s]TX:Query: back success.(%v)", tx.db.dbKey, key6, time.Since(start))
		}
		return resp, err
	}
}

func (tx *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return tx.tx.QueryRow(query, args...)
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return tx.tx.QueryRowContext(ctx, query)
}

func (tx *Tx) Rollback() error {
	if tx == nil || tx.tx == nil {
		return nil
	}
	return tx.tx.Rollback()
}

func (tx *Tx) Stmt(stmt *sql.Stmt) *sql.Stmt {
	return tx.tx.Stmt(stmt)
}

func (tx *Tx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	return tx.tx.StmtContext(ctx, stmt)
}

func (tx *Tx) CommandInsert(table string, colsList []string) string {
	return tx.db.CommandInsert(table, colsList)
}
