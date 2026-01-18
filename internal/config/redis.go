package config

import "errors"

const (
	errMsgRedisAddrRequired = "redis addr is required"
)

type Redis struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASS"`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

func (c *Redis) Validate() error {
	if c.Addr == "" {
		return errors.New(errMsgRedisAddrRequired)
	}
	return nil
}
