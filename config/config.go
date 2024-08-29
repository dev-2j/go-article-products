// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

package config

import (
	"os"

	"gitlab.dohome.technology/dohome-2020/go-servicex/constantx"
)

const (
	ENV_NAME_SIT = `-sit`
)

// UAT or PRD
var IsProduction = os.Getenv("GO_ENV") == "Production"

var IsEnvDEV = os.Getenv("GO_ENV") != "Production"
var IsEnvSIT = os.Getenv("GO_ENV") == "Production" && GetEnv("ENV_NAME") == ENV_NAME_SIT

var GetServiceName = func() string {
	return _Config.GetServiceName()
}

func SetServiceName(serviceName string) {
	_Config.SetServiceName(serviceName)
}

var ENV_NAME = func() string {
	if IsEnvDEV {
		return `-dev`
	}
	return GetEnv("ENV_NAME")
}()

func IsServiceModule(moduleName string) bool {
	if IsEnvDEV || GetServiceName() == `go-private` || GetServiceName() == `go-`+moduleName {
		return true
	}
	return false
}

func IsShowLogs() bool {
	return GetEnv(constantx.SHOW_LOGS) == `1`
}
