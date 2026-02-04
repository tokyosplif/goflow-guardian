package handlers

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	key := c.ClientIP()
	if req.Data != "" {
		key = req.Data
	}

	allowed, err := h.uc.Handle(c.Request.Context(), key)
	if err != nil {
		if errors.Is(err, domain.ErrLimitExceeded) {
			c.JSON(http.StatusTooManyRequests, dto.Response{Message: domain.StatusRejected})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	if !allowed {
		c.JSON(http.StatusTooManyRequests, dto.Response{Message: domain.StatusRejected})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: domain.StatusAllowed})
}
