package config

import (
	"fmt"
	"strings"
)

type CorsConfig struct {
	Origins     string
	Methods     string
	Headers     string
	Credentials bool
}

func (c *CorsConfig) Validate() error {
	var errs []string

	if c.Origins == "" {
		errs = append(errs, "ALLOW_ORIGINS is required")
	}

	if c.Headers == "" {
		errs = append(errs, "ALLOW_HEADERS is required")
	}

	if c.Methods == "" {
		errs = append(errs, "ALLOW_METHODS is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}
