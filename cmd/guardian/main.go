package main

import (
	"log/slog"
	"os"

	"github.com/tokyosplif/goflow-guardian/internal/app"
	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Init(cfg.App.Env)

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		slog.Error("Guardian stopped with error", "error", err)
		os.Exit(1)
	}

	slog.Info("Guardian stopped gracefully")
}
