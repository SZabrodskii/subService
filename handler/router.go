package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"subService/config"
)

func ProvideRouter() fx.Option {
	return fx.Options(
		fx.Provide(NewRouter),
		fx.Invoke(RegisterRoutes),
		fx.Invoke(RunHTTP))
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func RegisterRoutes(r *gin.Engine, h *SubscriptionHandler) {
	h.Register(r)
}

func RunHTTP(r *gin.Engine, cfg *config.Config) error {
	addr := fmt.Sprintf(":%d", cfg.Port)
	return r.Run(addr)
}
