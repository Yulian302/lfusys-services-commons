package config

import (
	"fmt"
	"strings"
)

type JWTConfig struct {
	SecretKey        string
	RefreshSecretKey string
}

func (c JWTConfig) ValidateSecrets() error {
	var errs []string

	if c.SecretKey == "" {
		errs = append(errs, "JWT_SECRET_KEY is required")
	}

	if c.RefreshSecretKey == "" {
		errs = append(errs, "JWT_REFRESH_SECRET_KEY key is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}
