package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type ClientToken struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

type ClientTokenResetPassword struct {
	jwt.RegisteredClaims
	Id  int64 `json:"id"`
	Exp int64 `json:"exp"`
}
