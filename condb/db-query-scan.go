package condb

import (
	"context"
)

func (db *DB) QueryScan(query string, args ...any) (*Rows, error) {
	return db.QueryScanContext(context.Background(), query, args...)
}

func (db *DB) QueryScanContext(ctx context.Context, query string, args ...any) (*Rows, error) {

	// prepare query
	rows, ex := db.QueryContext(ctx, query, args...)
	if ex != nil {
		return nil, ex
	}
	defer rows.Close()

	// query scan all
	rowx, ex := dxQueryScan(db.GetDriver(), rows)
	if ex != nil {
		return nil, ex
	}

	// set driver-name
	rowx.DriverName = db.driver
	return rowx, nil
}
