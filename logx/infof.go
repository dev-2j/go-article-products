package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Infof(format string, v ...any) {

	if isLambdaActiveValue {
		fmt.Println(fmt.Sprintf(format, v...))
		return
	}

	LoadInit()
	logrus.Infof(format, v...)
}
