package limiter

import (
	"context"

	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

type Publisher interface {
	PublishViolation(ctx context.Context, v domain.Violation) error
	Ping(ctx context.Context) error
}
