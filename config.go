package una

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

var (
	configFilename string
	appConf        interface{}

	ErrInvalidParseConfig = errors.New("conf param must be a pointer type")
)

func SetupConfig(filename string, conf interface{}) {
	rv := reflect.ValueOf(conf)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		panic(ErrInvalidParseConfig)
	}

	configFilename = filename
	appConf = conf
	err := parseConfig()
	if err != nil {
		panic(err)
	}
}

func parseConfig() error {
	err := Parse(configFilename, appConf)
	if err == nil {
		parseTime = time.Now().Unix()
	}
	return err
}

func ReloadConfigHandler() {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGHUP)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				sugarLogger := SugarLogger()
				if sugarLogger != nil {
					sugarLogger.Errorw("[recover] reload config handler panic",
						"err", err,
					)
				} else {
					log.Println(err)
				}
			}
		}()
		for {
			<-signChan
			err := parseConfig()
			if err != nil {
				sugarLogger := SugarLogger()
				if sugarLogger != nil {
					SugarLogger().Errorw("reload and parse config failed",
						"err", err,
					)
				} else {
					log.Println(err)
				}
				log.Println("reload config error")
			} else {
				log.Println("reload config")
			}
		}
	}()
}
