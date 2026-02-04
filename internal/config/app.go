package config

import (
	"errors"
	"strings"
)

type App struct {
	Port         string   `env:"APP_PORT" envDefault:"8080"`
	Env          string   `env:"APP_ENV" envDefault:"development"`
	LogPath      string   `env:"LOG_PATH" envDefault:"stdout"`
	AllowOrigins []string `env:"CORS_ORIGINS" envSeparator:"," envDefault:"http://localhost:3000"`
}

func (a *App) Validate() error {
	if a.Port == "" {
		return errors.New("app port is required")
	}

	for i, origin := range a.AllowOrigins {
		a.AllowOrigins[i] = strings.TrimSpace(origin)
	}

	if len(a.AllowOrigins) == 0 {
		return errors.New("at least one allowed origin is required")
	}

	for _, origin := range a.AllowOrigins {
		if origin == "" {
			return errors.New("empty origin not allowed")
		}
	}

	return nil
}
