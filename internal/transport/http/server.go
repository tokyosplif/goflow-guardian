package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tokyosplif/goflow-guardian/internal/config"
)

const (
	defaultShutdownTimeout = 5 * time.Second
)

func Run(ctx context.Context, cfg config.App, handler http.Handler) error {
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()

	return srv.Shutdown(shutdownCtx)
}
