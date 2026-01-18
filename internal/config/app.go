package config

import "errors"

type App struct {
	Port    string `env:"APP_PORT" envDefault:"8080"`
	Env     string `env:"APP_ENV" envDefault:"development"`
	LogPath string `env:"LOG_PATH" envDefault:"stdout"`
}

func (a *App) Validate() error {
	if a.Port == "" {
		return errors.New("app port is required")
	}
	return nil
}
