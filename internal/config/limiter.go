package config

import "errors"

type Limiter struct {
	Requests int `env:"LIMIT_REQUESTS" envDefault:"10"`
	Window   int `env:"LIMIT_WINDOW_SECONDS" envDefault:"60"`
}

func (l *Limiter) Validate() error {
	if l.Requests <= 0 {
		return errors.New("limit requests must be positive")
	}
	if l.Window <= 0 {
		return errors.New("limit window must be positive")
	}
	return nil
}
