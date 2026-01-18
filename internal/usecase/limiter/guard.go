package limiter

import (
	"context"
	"log/slog"
	"time"

	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

type UseCase interface {
	Handle(ctx context.Context, key string) (bool, error)
	CheckHealth(ctx context.Context) domain.HealthStatus
}

type Guard struct {
	storage   Storage
	publisher Publisher
	cfg       config.Limiter
}

func NewGuard(s Storage, p Publisher, cfg config.Limiter) *Guard {
	return &Guard{storage: s, publisher: p, cfg: cfg}
}

func (g *Guard) Handle(ctx context.Context, key string) (bool, error) {
	limit := domain.Limit{
		Requests: g.cfg.Requests,
		Window:   time.Duration(g.cfg.Window) * time.Second,
	}

	allowed, err := g.storage.IsAllowed(ctx, key, limit)
	if err != nil {
		slog.ErrorContext(ctx, "storage error", "err", err)
		return true, nil
	}

	if !allowed {
		go g.notify(key)
		return false, domain.ErrLimitExceeded
	}

	return true, nil
}

func (g *Guard) notify(ip string) {
	go func() {
		err := g.publisher.PublishViolation(context.Background(), domain.Violation{
			Key:       ip,
			Reason:    domain.ReasonRateLimitExceeded,
			Timestamp: time.Now(),
		})
		if err != nil {
			slog.Error("failed to publish violation", "err", err, "ip", ip)
		}
	}()
}
