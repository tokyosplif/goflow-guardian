package config

import "errors"

var (
	ErrAppPortRequired = errors.New("app port is required")
)

type App struct {
	Port    string `env:"APP_PORT" envDefault:"8080"`
	Env     string `env:"APP_ENV" envDefault:"development"`
	LogPath string `env:"LOG_PATH" envDefault:"stdout"`
}

func (a *App) Validate() error {
	if a.Port == "" {
		return ErrAppPortRequired
	}
	return nil
}
