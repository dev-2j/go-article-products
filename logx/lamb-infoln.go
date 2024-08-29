package logx

import (
	"fmt"

	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
)

func LambInfoln(v ...any) {

	if isLambdaActiveValue || config.IsEnvDEV {
		fmt.Println(v...)
	}

}
