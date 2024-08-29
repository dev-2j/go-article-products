package logx

import (
	"fmt"

	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
)

func LambInfof(format string, v ...any) {

	if isLambdaActiveValue || config.IsEnvDEV {
		fmt.Println(fmt.Sprintf(format, v...))
	}

}
