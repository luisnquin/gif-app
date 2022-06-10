//nolint:typecheck
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/auth"
)

func AHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"mmm": "patas"})
	}
}

func BHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := auth.GetUserFromContext(c)
		if !ok {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{
			"email":    user.Email,
			"username": user.Username,
		})
	}
}
