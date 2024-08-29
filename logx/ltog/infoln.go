package ltog

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
)

func Infoln(v ...any) {

	if logx.GetLambdaActive() {
		fmt.Println(v...)
		return
	}

	if config.IsShowLogs() {
		logx.LoadInit()
		logrus.Infoln(v...)
	}
}
