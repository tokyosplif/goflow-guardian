package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/transport/http/handlers"
	"github.com/tokyosplif/goflow-guardian/internal/transport/http/middleware"
	"github.com/tokyosplif/goflow-guardian/internal/usecase/limiter"
)

func NewRouter(uc limiter.UseCase, cfg config.App) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Security())

	if cfg.Env == "production" {
		r.Use(middleware.CORSProduction(cfg.AllowOrigins))
	} else {
		r.Use(middleware.CORS(cfg.AllowOrigins))
	}

	r.Use(middleware.Logger())

	limiterHandler := handlers.NewLimiter(uc)
	healthHandler := handlers.NewHealth(uc)

	r.GET("/health", healthHandler.Check)
	r.POST("/check", limiterHandler.Handle)

	return r
}
