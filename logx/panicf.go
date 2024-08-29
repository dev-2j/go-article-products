package logx

import (
	"log"

	"github.com/sirupsen/logrus"
)

func Panicf(format string, v ...any) {

	if isLambdaActiveValue {
		log.Panicf(format, v...)
		return
	}

	LoadInit()
	logrus.Panicf(format, v...)
}
