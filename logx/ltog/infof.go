package ltog

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
)

func Infof(format string, v ...any) {

	if logx.GetLambdaActive() {
		fmt.Println(fmt.Sprintf(format, v...))
		return
	}

	if config.IsShowLogs() {
		logx.LoadInit()
		logrus.Infof(format, v...)
	}
}
