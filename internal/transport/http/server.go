package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/tokyosplif/goflow-guardian/internal/config"
)

const (
	defaultShutdownTimeout = 30 * time.Second
	defaultReadTimeout     = 10 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultIdleTimeout     = 60 * time.Second
)

func Run(ctx context.Context, cfg config.App, handler http.Handler) error {
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}

	errChan := make(chan error, 1)

	go func() {
		slog.Info("HTTP server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		slog.Info("Shutdown signal received, stopping HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("Server forced to shutdown", "error", err)
			return fmt.Errorf("server shutdown failed: %w", err)
		}

		slog.Info("HTTP server stopped gracefully")
		return nil

	case err := <-errChan:
		return err
	}
}
