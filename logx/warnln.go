package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Warnln(v ...any) {

	if isLambdaActiveValue {
		fmt.Println(v...)
		return
	}

	LoadInit()
	logrus.Warningln(v...)
}
