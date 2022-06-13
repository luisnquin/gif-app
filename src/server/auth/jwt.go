//nolint:typecheck
package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/provider"
)

func (a *Auth) genSignedJWTToken(user provider.User) (string, error) {
	user.Password = ""

	claims := Claims{
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
func GetUserFromContext(c echo.Context) (provider.User, bool) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return provider.User{}, false
	}

	claim, ok := token.Claims.(*Claims)
	if !ok {
		return provider.User{}, false
	}

	return claim.User, true
}
