package condb

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
)

// check online
func (s *DB) online(icount int, key6 string) {
	counter := 0
	for {
		counter++
		if ex := s.Ping(1); ex != nil {
			logx.Warnf(`[%s|%s] ... ping(%v|%v): %s`, s.dbKey, key6, icount, counter, ex.Error())
			if _SentryActive {
				go func() {
					_ = sentry.CaptureException(fmt.Errorf(`ping:%s`, ex.Error()))
				}()
			}
			time.Sleep(time.Second)
			continue
		}
		if counter > 1 {
			logx.Infof(`[%s|%s] ... ping(%v|%v): back online.`, s.dbKey, key6, icount, counter)
		}
		break
	}
}
