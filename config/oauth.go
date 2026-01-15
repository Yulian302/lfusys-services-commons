package config

import (
	"errors"
)

type OAuthConfig struct {
	*GithubConfig
	*GoogleConfig
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
	if cfg.ClientID == "" {
		return errors.New("OAUTH2_GITHUB_CLIENT_ID is required")
	}

	if cfg.ClientSecret == "" {
		return errors.New("OAUTH2_GITHUB_CLIENT_SECRET is required")
	}

	if cfg.RedirectURI == "" {
		return errors.New("OAUTH2_GITHUB_REDIRECT_URI is required")
	}

	if cfg.ExchangeURL == "" {
		return errors.New("OAUTH2_GITHUB_EXCHANGE_URL is required")
	}

	return nil
}

func (cfg *GoogleConfig) ValidateSecrets() error {
	if cfg.ClientID == "" {
		return errors.New("OAUTH2_GOOGLE_CLIENT_ID is required")
	}

	if cfg.ClientSecret == "" {
		return errors.New("OAUTH2_GOOGLE_CLIENT_SECRET is required")
	}

	if cfg.RedirectURI == "" {
		return errors.New("OAUTH2_GOOGLE_REDIRECT_URI is required")
	}

	if cfg.ExchangeURL == "" {
		return errors.New("OAUTH2_GOOGLE_EXCHANGE_URL is required")
	}

	return nil
}
