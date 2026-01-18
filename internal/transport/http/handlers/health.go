package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tokyosplif/goflow-guardian/internal/domain"
	"github.com/tokyosplif/goflow-guardian/internal/usecase/limiter"
)

type Health struct {
	uc limiter.UseCase
}

func NewHealth(uc limiter.UseCase) *Health {
	return &Health{uc: uc}
}

func (h *Health) Check(c *gin.Context) {
	status := h.uc.CheckHealth(c.Request.Context())

	httpStatus := http.StatusOK
	if status.Status == domain.StatusDown {
		httpStatus = http.StatusServiceUnavailable
	}

	c.JSON(httpStatus, status)
}
