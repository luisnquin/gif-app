package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/models"
)

func (a *Auth) genSignedJWTToken(user models.User) (string, error) {
	user.Password = ""

	claims := models.Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: a.getTokenTimeout().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *Auth) getTokenTimeout() time.Time {
	return time.Now().Add(time.Hour * a.config.Internal.TokenExpirationTime)
}

// The user is saved in the echo context due to the JWT token.
// So, careful or nothing will be provided.
func GetUserFromContext(c echo.Context) (models.User, bool) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return models.User{}, false
	}

	claim, ok := token.Claims.(*models.Claims)
	if !ok {
		return models.User{}, false
	}

	return claim.User, true
}
