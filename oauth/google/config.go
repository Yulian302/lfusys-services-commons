package google

import "errors"

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	ExchangeURL  string
	FrontendURL  string
}

func (cfg *GoogleConfig) ValidateSecrets() error {
	if cfg.ClientID == "" {
		return errors.New("OAUTH2_GOOGLE_CLIENT_ID was not found")
	}

	if cfg.ClientSecret == "" {
		return errors.New("OAUTH2_GOOGLE_CLIENT_SECRET was not found")
	}

	if cfg.RedirectURI == "" {
		return errors.New("OAUTH2_GOOGLE_REDIRECT_URI was not found")
	}

	if cfg.ExchangeURL == "" {
		return errors.New("OAUTH2_GOOGLE_EXCHANGE_URL was not found")
	}

	return nil
}
