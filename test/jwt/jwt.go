package jwt

import (
	"testing"
	"time"

	appjwt "github.com/Yulian302/lfusys-services-commons/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func MakeTestJWT(t *testing.T, secretKey string, email string) string {
	t.Helper()

	claims := appjwt.JWTClaims{
		Issuer:    "lfusys",
		Subject:   email,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secretKey))
	require.NoError(t, err)

	return signed
}
