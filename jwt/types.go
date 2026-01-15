package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenDuration  = 30 * time.Minute
	RefreshTokenDuration = 30 * 24 * time.Hour
	CookiePath           = "/"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type JWTClaims struct {
	Issuer    string `json:"iss"`
	Subject   string `json:"subject"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
	Type      string `json:"type,omitempty"`
	JTI       string `json:"jti"`
}

func (c JWTClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.ExpiresAt, 0)), nil
}

func (c JWTClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.IssuedAt, 0)), nil
}

func (c JWTClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c JWTClaims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

func (c JWTClaims) GetSubject() (string, error) {
	return c.Subject, nil
}

func (c JWTClaims) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (c JWTClaims) GetJTI() (string, error) {
	return c.JTI, nil
}
