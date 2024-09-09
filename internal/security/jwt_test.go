package security

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const Secret = "secret"

// You must provide the token type struct
type UserToken struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func TestSignToken(t *testing.T) {
	user := &UserToken{
		UserId: 1,
	}

	token, err := JWTSign(user, Secret)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	var userVerified UserToken
	err = JWTGetPayload(token, Secret, &userVerified)
	assert.Nil(t, err)

	assert.Equal(t, user.UserId, userVerified.UserId)
}
