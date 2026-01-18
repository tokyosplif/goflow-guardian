package config

type Limiter struct {
	Requests int `env:"LIMIT_REQUESTS" envDefault:"10"`
	Window   int `env:"LIMIT_WINDOW_SECONDS" envDefault:"60"`
}
