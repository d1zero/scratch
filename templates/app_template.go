package templates

const AppTemplate = `package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/d1zero/scratch/pkg/config"
	lgr "gitlab.com/d1zero-diplom/backend/common/pkg/logger"
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

	exit := make(chan os.Signal)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
}
`
