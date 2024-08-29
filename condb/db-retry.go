package condb

import (
	"strings"

	"github.com/getsentry/sentry-go"
)

// error will retry
func (s *DB) retry(err error, icount int, key6 string, onAlert func()) (int, bool) {
	ok := func() bool {
		if err != nil {
			et := err.Error()
			// Transaction (Process ID 422) was deadlocked on lock resources with another process and has been chosen as the deadlock victim. Rerun the transaction.
			// Transaction (Process ID 2807) was deadlocked on lock | communication buffer resources with another process and has been chosen as the deadlock victim. Rerun the transaction.
			if strings.Contains(et, `was deadlocked on lock`) && strings.Contains(et, `the deadlock victim`) && strings.Contains(et, `Rerun the transaction`) {
				return true // Rerun the transaction.
			}
			// EOF
			if strings.TrimSpace(et) == `EOF` {
				return true
			}
			//  the database system is in recovery mode
			if strings.Contains(et, `the database system is in recovery mode`) {
				return true
			}
			// the database system is starting up
			if strings.Contains(et, `the database system is starting up`) {
				return true
			}
			// cannot execute INSERT in a read-only transaction
			if strings.Contains(et, `cannot execute`) && strings.Contains(et, `in a read-only transaction`) {
				return true // เกิดขึ้นกรณีที่มีการ set failover rds เปลี่ยน instance primary
			}
			// read tcp 10.200.32.62:37946->10.200.7.123:5432: read: connection reset by peer
			if strings.Contains(et, `read tcp`) && strings.Contains(et, `connection reset by peer`) {
				return true
			}
			// dial tcp 10.200.7.123:5432: connect: connection refused
			if strings.Contains(et, `dial tcp`) && strings.Contains(et, `connection refused`) {
				return true
			}
			// pq: terminating connection because backend initialization completed past seamless quiet point
			if strings.Contains(et, `terminating connection`) && strings.Contains(et, `past seamless quiet point`) {
				return true
			}
			// sqlite: database is locked
			if s.driver == SQLITE3 && strings.Contains(et, `database is locked`) {
				return true
			}
			// server closed idle connection
			if strings.Contains(et, `server closed idle connection`) {
				return true
			}
			// unable to open tcp connection with host 'database.dohome.co.th:1433': dial tcp 192.168.8.88:1433: i/o timeout
			if strings.Contains(et, `unable to open tcp connection with host`) && strings.Contains(et, `i/o timeout`) {
				return true
			}
			// connection timed out
			if strings.Contains(et, `connection timed out`) {
				return true
			}
		}
		return false
	}()
	if ok {
		// alert
		onAlert()
		// sentry
		if _SentryActive {
			go func() {
				_ = sentry.CaptureException(err)
			}()
		}
		// check online
		s.online(icount, key6)
	}
	// retry : sleep for database
	/*
		x	= icount
		a	= 30
		b	= -30
		c	= 3000
		----
		y 	= a*x*x + b*x + c
	*/
	// y := (30 * icount * icount) + (-30 * icount) + 3000
	y := 3000 // 3วิ
	return y, ok
}
