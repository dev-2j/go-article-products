package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Alert(format string, v ...any) {

	if isLambdaActiveValue {
		fmt.Println(fmt.Sprintf(format, v...))
		return
	}

	LoadInit()
	logrus.Errorf(format, v...)

}
