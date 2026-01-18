package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/infrastructure/publisher"
	"github.com/tokyosplif/goflow-guardian/internal/infrastructure/storage"
	"github.com/tokyosplif/goflow-guardian/internal/transport/http"
	"github.com/tokyosplif/goflow-guardian/internal/usecase/limiter"
	"github.com/tokyosplif/goflow-guardian/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Init(cfg.App.Env)

	redisStorage := storage.NewRedisLimiter(cfg.Redis)
	kafkaPublisher := publisher.NewKafka(cfg.Kafka)

	guardUC := limiter.NewGuard(redisStorage, kafkaPublisher, cfg.Limiter)

	router := http.NewRouter(guardUC)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	slog.Info("Guardian starting", "port", cfg.App.Port, "env", cfg.App.Env)

	if err := http.Run(ctx, cfg.App, router); err != nil {
		slog.Error("server stopped with error", "error", err)
		os.Exit(1)
	}

	slog.Info("Guardian stopped gracefully")
}
