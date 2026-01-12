package common

import "errors"

type RedisConfig struct {
	HOST string
}

func (cfg *RedisConfig) ValidateSecrets() error {
	if cfg.HOST == "" {
		return errors.New("REDIS_HOST is required")
	}
	return nil
}
