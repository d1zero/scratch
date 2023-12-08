package templates

import "github.com/d1zero/scratch/internal/models"

type AppTemplateData struct {
	ProjectName string
}

func BuildAppTemplate(flags models.EnabledIntegrations) string {
	result := `package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"`

	if flags.Postgres {
		result += `
	
	"github.com/d1zero/scratch/pkg/config/postgres"`
	}

	if flags.Grpc {
		result += `
	v1 "{{.ProjectName}}/internal/controller/grpc/v1"
	"github.com/d1zero/scratch/pkg/config/grpc"`
	}

	result += `
	"github.com/d1zero/scratch/pkg/config"
	lgr "github.com/d1zero/scratch/pkg/logger"
)

func Run() {
	// getting default config
	cfg := defaultConfig()

	// loading config
	cfg, err := config.LoadConfig(cfg)
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
		return
	}

	// initializing logger
	defaultLogger, err := lgr.New(cfg.Logger)
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
		return
	}
	defaultLogger.Info("basic logger initialized successfully")

	logger := defaultLogger.Sugar()
	logger.Info("sugared logger initialized successfully")
`

	if flags.Postgres {
		result += `
	db, err := postgres.New(logger, cfg.Postgres)
	if err != nil {
		logger.Error("error while connecting postgres: %s", err)
		return
	}
	defer db.Disconnect()

	logger.Infof("postgres connected successfully")

`
	}

	if flags.Grpc {
		result += `
	grpcServer := grpc.New(logger, v1.ErrCodeMap)

	grpcServer.Start(cfg.GRPC)
	defer grpcServer.Stop()

`
	}

	result += `
	exit := make(chan os.Signal)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
}
`
	return result
}
