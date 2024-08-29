package logx

import (
	"github.com/sirupsen/logrus"
)

func Debugf(format string, v ...any) {

	if isLambdaActiveValue {
		return
	}

	LoadInit()
	logrus.Debugf(format, v...)
}
