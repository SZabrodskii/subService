// @title Subscription Service API
// @version 1.0
// @description REST-сервис для управления пользовательскими подписками.
// @host localhost:8040
// @BasePath /api/v1
package main

import (
	"go.uber.org/fx"
	"subService/config"
	"subService/db"
	"subService/handler"
	"subService/logger"
	"subService/repository"
	"subService/service"
)

func main() {
	app := fx.New(
		config.Provide(),
		logger.Provide(),
		db.Provide(),
		repository.Provide(),
		service.Provide(),
		handler.Provide(),
		handler.ProvideRouter(),
	)

	app.Run()
}
