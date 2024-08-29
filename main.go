package main

import (
	"log"
	"os"
	"time"

	"github.com/dev-2j/go-article-products/routex"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	ENV_NAME_SIT = `-sit`
	ENV_NAME_PRD = ``
)

var IsEnvDEV = os.Getenv("ENV_NAME") != "-dev"

func Initx() {

	if IsEnvDEV {
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

type InitType struct {
	Location *time.Location
}

const TEMP_PATH string = `./tmp/`

func CreateTempFolder(subFolder ...string) (*string, error) {

	tempFolder := TEMP_PATH

	for _, sf := range subFolder {
		tempFolder += sf + "/"
	}

	if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
		if err := os.Mkdir(tempFolder, os.ModeDir|0755); err != nil {
			return nil, err
		}
	}
	return &tempFolder, nil
}

func Init() (*InitType, error) {

	// logrus
	Initx()

	if IsEnvDEV {
		// โหลดไฟล์ .env ก่อน แล้วเอาไฟล์ .env.local มา merge ทับอีกที
		_ = godotenv.Load(".env")
	}

	// Time Zone
	loc, ex := time.LoadLocation("Asia/Bangkok")
	if ex != nil {
		log.Fatalf("set time zone, %v", ex.Error())
	}
	time.Local = loc

	// create temp
	_, _ = CreateTempFolder()

	return &InitType{
		Location: loc,
	}, nil

}

func main() {

	rx, ex := Init()
	if ex != nil {
		log.Fatalf(`main.Init: %s`, ex.Error())
	}
	time.Local = rx.Location

	// start crons
	// time.AfterFunc(time.Second*4, func() {
	// 	crons.Starter()
	// })

	// routers
	routex.Routex()

}
