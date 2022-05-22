package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/models"
)

var (
	Config middleware.JWTConfig

	privateKey *rsa.PrivateKey
)

func init() {
	privCont, err := ioutil.ReadFile("./private.rsa.key")
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privCont)
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

	Config = middleware.JWTConfig{
		Claims:        &models.Claims{},
		SigningKey:    privateKey,
		SigningMethod: jwt.SigningMethodRS256.Alg(),
		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
	}
}

func LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User

		err := c.Bind(&user)
		if err != nil {
			return echo.ErrUnauthorized
		}

		user.Password = ""

		token, err := genSignedJWTToken(user)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": token,
		})
	}
}
