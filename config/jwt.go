package config

import "errors"

type JWTConfig struct {
	SecretKey        string
	RefreshSecretKey string
}

func (c JWTConfig) ValidateSecrets() error {
	if c.SecretKey == "" {
		return errors.New("JWT_SECRET_KEY is required")
	}

	if c.RefreshSecretKey == "" {
		return errors.New("JWT_REFRESH_SECRET_KEY key is required")
	}

	return nil
}
