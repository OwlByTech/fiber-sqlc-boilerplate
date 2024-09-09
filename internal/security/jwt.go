package security

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func JWTSign(payload jwt.Claims, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %v", err)
	}

	return t, nil
}

func JWTVerify(token string, secret string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenParsed, nil
}

func JWTGetPayload(token string, secret string, payload jwt.Claims) error {
	_, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	return nil
}
