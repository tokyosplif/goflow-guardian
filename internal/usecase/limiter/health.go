package limiter

import (
	"context"

	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

func (g *Guard) CheckHealth(ctx context.Context) domain.HealthStatus {
	components := make(map[string]string)
	status := domain.StatusOK

	if err := g.storage.Ping(ctx); err != nil {
		components["storage"] = domain.StatusDown
		status = domain.StatusDown
	} else {
		components["storage"] = domain.StatusOK
	}

	if err := g.publisher.Ping(ctx); err != nil {
		components["publisher"] = domain.StatusDown
		status = domain.StatusDown
	} else {
		components["publisher"] = domain.StatusOK
	}

	return domain.HealthStatus{
		Status:     status,
		Components: components,
	}
}
