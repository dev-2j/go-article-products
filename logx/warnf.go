package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Warnf(format string, v ...any) {

	if isLambdaActiveValue {
		fmt.Println(fmt.Sprintf(format, v...))
		return
	}

	LoadInit()
	logrus.Warningf(format, v...)
}
