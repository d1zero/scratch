package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log/slog"
	"os"
	"strings"
)

type HostPort struct {
	Host string `koanf:"host" validate:"required"`
	Port string `koanf:"port" validate:"required"`
}

// LoadConfig loads config on given type, also loads data from environment variables and local config.yml file
// and validates it with govalidator
func LoadConfig[T any](cfg T) (T, error) {
	k := koanf.New(".")

	err := k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "")), "_", ".", -1)
	}), nil)
	if err != nil {
		return cfg, err
	}

	err = k.Load(file.Provider("config.yml"), yaml.Parser())
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(fmt.Errorf("unable to open config.yml: %s", err).Error())
	}

	err = k.Unmarshal("", cfg)
	if err != nil {
		return cfg, err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
