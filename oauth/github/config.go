package github

import "errors"

type GithubConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	ExchangeURL  string
	FrontendURL  string
}

func (cfg *GithubConfig) ValidateSecrets() error {
	if cfg.ClientID == "" {
		return errors.New("OAUTH2_GITHUB_CLIENT_ID was not found")
	}

	if cfg.ClientSecret == "" {
		return errors.New("OAUTH2_GITHUB_CLIENT_SECRET was not found")
	}

	if cfg.RedirectURI == "" {
		return errors.New("OAUTH2_GITHUB_REDIRECT_URI was not found")
	}

	if cfg.ExchangeURL == "" {
		return errors.New("OAUTH2_GITHUB_EXCHANGE_URL was not found")
	}

	return nil
}
