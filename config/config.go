package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Port   int
	DbUrl  string
	AppEnv string
}

func Provide() fx.Option {
	return fx.Provide(NewConfig)
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("No config.yaml loaded, using defaults and ENV:", err)
	}
	v.SetDefault("port", 8040)

	cfg := &Config{
		Port:   v.GetInt("PORT"),
		DbUrl:  v.GetString("DB_URL"),
		AppEnv: v.GetString("APP_ENV"),
	}
	if cfg.DbUrl == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}

	return cfg, nil
}
