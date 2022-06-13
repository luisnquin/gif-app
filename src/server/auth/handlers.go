//nolint:typecheck
package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/log"
	"github.com/luisnquin/gif-app/src/server/provider"
	"github.com/luisnquin/gif-app/src/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func (a *Auth) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req provider.User

		err := c.Bind(&req)
		if err != nil {
			return echo.ErrBadRequest
		}

		if req.Password == "" || (req.Username == "" && req.Email == "") {
			return echo.ErrBadRequest
		}

		user, err := a.provider.GetUserByUsernameOrEmail(c.Request().Context(), provider.GetUserByUsernameOrEmailParams{
			Username: req.Username,
			Email:    req.Email,
		})

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
			log.Warn(err)

			return echo.ErrUnauthorized
		}

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

		return c.JSON(http.StatusOK, TokenResponse{
			Token: token,
		})
	}
}

func (a *Auth) SignUpHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request provider.User

		err := c.Bind(&request)
		if err != nil {
			return echo.ErrBadRequest
		}

		if request.Username == "" || request.Firstname == "" || request.Lastname == "" || request.Email == "" || request.Password == "" {
			return echo.ErrBadRequest
		}

		exists, err := a.provider.UserExistsByUsernameOrEmail(c.Request().Context(), provider.UserExistsByUsernameOrEmailParams{
			Username: request.Username,
			Email:    request.Email,
		})

		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		if exists {
			return echo.ErrBadRequest
		}

		password, err := utils.GenHashedPassword(request.Password)
		if err != nil {
			log.Error(err)
			return echo.ErrInternalServerError
		}

		request.Role = UserDefaultRole
		request.Password = password
		request.CreatedAt = time.Now()
		request.UpdatedAt = time.Now()

		// Add email verification.

		_, err = a.provider.CreateUser(c.Request().Context(), provider.CreateUserParams{
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Username:  request.Username,
			Password:  request.Password,
			Email:     request.Email,
			Role:      request.Role,
		})

		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}

func (a *Auth) LogoutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			if errors.Is(err, echo.ErrCookieNotFound) || errors.Is(err, http.ErrNoCookie) {
				return echo.ErrBadRequest
			}

			log.Error(err)

			return echo.ErrInternalServerError
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)

		c.SetCookie(cookie)

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
