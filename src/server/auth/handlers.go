package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/log"
	"github.com/luisnquin/meow-app/src/server/models"
	"github.com/luisnquin/meow-app/src/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func (a *Auth) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.User

		err := c.Bind(&req)
		if err != nil {
			return echo.ErrBadRequest
		}

		user, err := a.provider.GetUserByEmailOrUsername(c.Request().Context(), req.Username, req.Email)
		if err != nil {
			log.Error(err)

			if errors.Is(err, sql.ErrNoRows) {
				return echo.ErrUnauthorized
			} else if errors.Is(err, context.Canceled) {
				return echo.ErrBadRequest
			}

			return echo.ErrInternalServerError
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			log.Error(err)

			return echo.ErrUnauthorized
		}

		user.Password = ""

		token, err := a.genSignedJWTToken(user)
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		c.SetCookie(&http.Cookie{
			Expires:  a.getTokenTimeout(),
			HttpOnly: true,
			Name:     "token",
			Value:    token,
		})

		return c.JSON(http.StatusOK, models.TokenResponse{
			Token: token,
		})
	}
}

// return c.JSON(http.StatusOK, models.ShortResponse{
// 	Message: "Token now in cookies",
// })

func (a *Auth) RegisterHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User

		err := c.Bind(&user)
		if err != nil {
			return echo.ErrBadRequest
		}

		if user.Username == "" || user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Password == "" {
			return echo.ErrBadRequest
		}

		exists, err := a.provider.UsernameOrEmailExists(c.Request().Context(), user.Username, user.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		if exists {
			return echo.ErrBadRequest
		}

		password, err := utils.GenHashedPassword(user.Password)
		if err != nil {
			return echo.ErrInternalServerError
		}

		user.Role = UserDefaultRole
		user.Password = password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// Add email verification.

		err = a.provider.SaveUser(c.Request().Context(), user)
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}

func (a *Auth) emailIsValid(email string) bool {
	rexp := regexp.MustCompile(a.config.Internal.EmailRegex)
	if !rexp.MatchString(email) {
		return false
	}

	return true
}

/*
	BasicAuthConfig = middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			user, err := userProxy.GetByEmailOrUsername(username, "")
			if err != nil {
				log.Error(err)

				return false, err
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				log.Error(err)

				return false, err
			}

			signedToken, err := genSignedJWTToken(user)
			if err != nil {
				return false, err
			}

			token, err := jwt.Parse(signedToken, func(t *jwt.Token) (any, error) {
				return publicKey, nil
			})

			if err != nil {
				return false, err
			}

			c.Set("user", token)

			return true, nil
		},
		Skipper: func(c echo.Context) bool {
			_, ok := GetUserFromContext(c)

			return ok
		},
	}
*/
