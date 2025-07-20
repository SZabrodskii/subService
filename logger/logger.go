package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
)

func NewZapLogger() *zap.Logger {
	var logger *zap.Logger
	var err error

	if val, ok := os.LookupEnv("APP_ENV"); ok && val == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("error creating zap logger" + err.Error())
	}

	return logger

}

func Provide() fx.Option {
	return fx.Provide(NewZapLogger)
}
