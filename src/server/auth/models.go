package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/luisnquin/gif-app/src/server/provider"
)

type Claims struct {
	provider.User `json:"user"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}
