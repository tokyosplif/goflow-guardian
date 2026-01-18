package limiter

import (
	"context"

	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

type Storage interface {
	IsAllowed(ctx context.Context, key string, limit domain.Limit) (bool, error)
	Ping(ctx context.Context) error
}
