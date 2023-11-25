package templates

import (
	"github.com/d1zero/scratch/internal/models"
)

func BuildConfigTemplate(flags models.EnabledIntegrations) string {
	result := `package app

import (
	"github.com/d1zero/scratch/pkg/config"`

	if flags.Postgres {
		result += `
	"github.com/d1zero/scratch/pkg/config/postgres"`
	}

	result += `
	"github.com/d1zero/scratch/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Logger      logger.Config   ` + "`" + `koanf:"logger" validate:"required"` + "`"
	if flags.Postgres {
		result += `
	Postgres    postgres.Config ` + "`" + `koanf:"postgres" validate:"required"` + "`"
	}
	result += `
	HealthCheck config.HostPort ` + "`" + `koanf:"healthCheck" validate:"required"` + "`" + `
}

func defaultConfig() *Config {
	return &Config{
		Logger: logger.Config{
			Level: zapcore.DebugLevel,
		},
		HealthCheck: config.HostPort{
			Host: "0.0.0.0",
			Port: "7000",
		},`
	if flags.Postgres {
		result += `
		Postgres: postgres.Config{
			DSN:             "postgresql://user:pass@127.0.0.1:5432/user-service?sslmode=disable",
			MaxOpenConns:    10,
			ConnMaxLifetime: 15,
			MaxIdleConns:    20,
			ConnMaxIdleTime: 30,
		},`
	}
	result += `
	}
}
`
	return result
}
