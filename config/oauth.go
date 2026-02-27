package config

import (
	"fmt"
	"strings"
)

type OAuthConfig struct {
	Github GithubConfig
	Google GoogleConfig
}

type GithubConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	ExchangeURL  string
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	ExchangeURL  string
}

func (cfg *GithubConfig) ValidateSecrets() error {
	var errs []string

	if cfg.ClientID == "" {
		errs = append(errs, "OAUTH2_GITHUB_CLIENT_ID is required")
	}

	if cfg.ClientSecret == "" {
		errs = append(errs, "OAUTH2_GITHUB_CLIENT_SECRET is required")
	}

	if cfg.RedirectURI == "" {
		errs = append(errs, "OAUTH2_GITHUB_REDIRECT_URI is required")
	}

	if cfg.ExchangeURL == "" {
		errs = append(errs, "OAUTH2_GITHUB_EXCHANGE_URL is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (cfg *GoogleConfig) ValidateSecrets() error {
	var errs []string

	if cfg.ClientID == "" {
		errs = append(errs, "OAUTH2_GOOGLE_CLIENT_ID is required")
	}

	if cfg.ClientSecret == "" {
		errs = append(errs, "OAUTH2_GOOGLE_CLIENT_SECRET is required")
	}

	if cfg.RedirectURI == "" {
		errs = append(errs, "OAUTH2_GOOGLE_REDIRECT_URI is required")
	}

	if cfg.ExchangeURL == "" {
		errs = append(errs, "OAUTH2_GOOGLE_EXCHANGE_URL is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}
