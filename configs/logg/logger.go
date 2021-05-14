package logg

import (
	"RestFullAPI-todo/configs"
	"go.uber.org/zap"
	"log"
)

var L *zap.Logger

func Init() {
	var err error
	if configs.C.Env == "production" {
		L, err = zap.NewProduction()

		if err != nil {
			log.Panic(err)
		}
	} else {
		L, err = zap.NewDevelopment()
		if err != nil {
			log.Panic(err)
		}
	}
	L.Info("Log initialize")
}
