package app

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/infrastructure/publisher"
	"github.com/tokyosplif/goflow-guardian/internal/infrastructure/storage"
	"github.com/tokyosplif/goflow-guardian/internal/transport/http"
	"github.com/tokyosplif/goflow-guardian/internal/usecase/limiter"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Run() error {
	redisStorage := storage.NewRedisLimiter(a.cfg.Redis)
	kafkaPublisher := publisher.NewKafka(a.cfg.Kafka)

	guardUC := limiter.NewGuard(redisStorage, kafkaPublisher, a.cfg.Limiter)

	router := http.NewRouter(guardUC)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	slog.Info("Guardian starting",
		slog.String("port", a.cfg.App.Port),
		slog.String("env", a.cfg.App.Env),
	)

	if err := http.Run(ctx, a.cfg.App, router); err != nil {
		return err
	}

	return nil
}
