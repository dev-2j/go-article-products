package logx

import (
	"flag"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.dohome.technology/dohome-2020/go-servicex/config"
)

var (
	_LOCK_INIT          = &sync.Mutex{}
	_INIT_LOADED        = false
	isLambdaActiveValue bool
)

func LoadInit() {
	if config.IsEnvDEV {
		// if flag.Lookup("test.v") == nil {
		// 	return false // normal execution
		// } else {
		// 	return true // test execution
		// }
		if !_INIT_LOADED {
			if flag.Lookup("test.v") != nil {
				_LOCK_INIT.Lock()
				defer _LOCK_INIT.Unlock()
				_INIT_LOADED = true
				Init()
			}
		}
	}
}

func Init() {

	if config.IsEnvDEV {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		ForceColors:               true,
		DisableColors:             false,
		DisableQuote:              true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           " 2006-01-02 15:04:05 ",
		// DisableLevelTruncation: false,
		// PadLevelText: true,
	})
}

// true = running in lambda
func SetLambdaActive() {
	isLambdaActiveValue = true
}

// true = running in lambda
func GetLambdaActive() bool {
	return isLambdaActiveValue
}
