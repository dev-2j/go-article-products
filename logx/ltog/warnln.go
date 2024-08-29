package ltog

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
	"gitlab.dohome.technology/dohome-2020/go-servicex/logx"
)

func Warnln(v ...any) {

	if config.IsShowLogs() {
		fmt.Println(v...)
		return
	}

	if config.IsShowLogs() {
		logx.LoadInit()
		logrus.Warningln(v...)
	}
}
