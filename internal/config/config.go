package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Service ServiceConfig
	Storage StorageConfig
}

type ServiceConfig struct {
	Port string
}

func LoadConfig(path string, logger *zap.Logger) *Config {
	var cfg = &Config{
		Service: ServiceConfig{
			Port: "0000",
		},
		Storage: StorageConfig{
			Host:     "0.0.0.0",
			Port:     "0000",
			Username: "0000",
			Password: "0000",
			Database: "0000",
		},
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Debug("config file does not exist",
			zap.Any("path", path),
		)
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		logger.Debug("failed to read config", zap.Error(err))
	}

	logger.Info("successfully loaded config",
		zap.Any("config", cfg))
	return cfg
}
