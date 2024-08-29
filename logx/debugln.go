package logx

import (
	"github.com/sirupsen/logrus"
)

func Debugln(v ...any) {

	if isLambdaActiveValue {
		return
	}

	LoadInit()
	logrus.Debug(v...)
}
