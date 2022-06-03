package auth

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/models"
	"github.com/luisnquin/meow-app/src/server/repository"
)

const (
	AdminRole       string = "PRIVILEGED"
	UserDefaultRole string = "USER"
)

type Auth struct {
	config     *config.Configuration
	provider   *repository.Provider
	JWTConfig  middleware.JWTConfig
	privateKey *rsa.PrivateKey
}

func New(config *config.Configuration, provider *repository.Provider) *Auth {
	privateKey, publicKey := loadKeys()

	return &Auth{
		privateKey: privateKey,
		provider:   provider,
		config:     config,
		JWTConfig: middleware.JWTConfig{
			SigningMethod: jwt.SigningMethodRS256.Alg(),
			Claims:        &models.Claims{},
			SigningKey:    privateKey,
			KeyFunc: func(token *jwt.Token) (any, error) {
				return publicKey, nil
			},
		},
	}
}

func loadKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privCont, err := ioutil.ReadFile("./private.rsa.key")
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privCont)
	if err != nil {
		panic(err)
	}

	pubCont, err := ioutil.ReadFile("./public.rsa.key")
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubCont)
	if err != nil {
		panic(err)
	}

	return privateKey, publicKey
}
