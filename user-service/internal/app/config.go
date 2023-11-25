package app

import (
	"github.com/d1zero/scratch/pkg/config"
	"github.com/d1zero/scratch/pkg/config/postgres"
	"github.com/d1zero/scratch/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Logger      logger.Config   `koanf:"logger" validate:"required"`
	Postgres    postgres.Config `koanf:"postgres" validate:"required"`
	HealthCheck config.HostPort `koanf:"healthCheck" validate:"required"`
}

func defaultConfig() *Config {
	return &Config{
		Logger: logger.Config{
			Level: zapcore.DebugLevel,
		},
		HealthCheck: config.HostPort{
			Host: "0.0.0.0",
			Port: "7000",
		},
		Postgres: postgres.Config{
			DSN:             "postgresql://user:pass@127.0.0.1:5432/user-service?sslmode=disable",
			MaxOpenConns:    10,
			ConnMaxLifetime: 15,
			MaxIdleConns:    20,
			ConnMaxIdleTime: 30,
		},
	}
}
