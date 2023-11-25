package templates

const ConfigTemplate = `package app

import (
	"github.com/d1zero/scratch/pkg/config"
	"github.com/d1zero/scratch/pkg/config/postgres"
	"github.com/d1zero/scratch/pkg/logger"
)

type Config struct {
	Logger      logger.Config   ` + "`" + `koanf:"logger" validate:"required"` + "`" + `
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
		},
	}
}
`
