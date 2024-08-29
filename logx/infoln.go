package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Infoln(v ...any) {

	if isLambdaActiveValue {
		fmt.Println(v...)
		return
	}

	LoadInit()
	logrus.Infoln(v...)
}
