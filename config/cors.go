package config

import "errors"

type CorsConfig struct {
	Origins string
}

func (c *CorsConfig) Validate() error {
	if c.Origins == "" {
		return errors.New("ORIGINS is required")
	}

	return nil
}
