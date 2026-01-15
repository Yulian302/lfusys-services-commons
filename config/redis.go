package config

import "errors"

type RedisConfig struct {
	HOST string
}

func (c *RedisConfig) ValidateSecrets() error {
	if c.HOST == "" {
		return errors.New("REDIS_HOST is required")
	}
	return nil
}
