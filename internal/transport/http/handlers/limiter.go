package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tokyosplif/goflow-guardian/internal/domain"
	"github.com/tokyosplif/goflow-guardian/internal/transport/http/dto"
	"github.com/tokyosplif/goflow-guardian/internal/usecase/limiter"
)

type Limiter struct {
	uc limiter.UseCase
}

func NewLimiter(uc limiter.UseCase) *Limiter {
	return &Limiter{uc: uc}
}

func (h *Limiter) Handle(c *gin.Context) {
	var req dto.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowed, err := h.uc.Handle(c.Request.Context(), c.ClientIP())
	if err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		return
	}

	if !allowed {
		c.JSON(http.StatusTooManyRequests, gin.H{"status": domain.StatusRejected})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": domain.StatusAllowed})
}
