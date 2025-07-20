package db

import (
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"subService/config"
)

func Provide() fx.Option {
	return fx.Provide(NewDB)
}

func NewDB(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := cfg.DbUrl
	logger.Sugar().Infow("Connecting to database", "dsn", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := db.AutoMigrate(&model.Subscription{}); err != nil {
		logger.Sugar().Errorw("Failed to migrate database", "error", err)
		return nil, fmt.Errorf("failed to automigrate database: %w", err)
	}
	logger.Sugar().Infow("Database migration complete")

	return db, nil
}
