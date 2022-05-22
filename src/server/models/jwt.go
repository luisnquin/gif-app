package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	User `json:"user"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}