//nolint:typecheck
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/auth"
	"github.com/luisnquin/gif-app/src/server/core"
)

func (h *HandlerHead) AHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, core.StandardResponse{
			APIVersion: core.APIVersion,
			Context:    "test",
			Method:     c.Request().Method,
			Data: echo.Map{
				"ip": c.RealIP(),
			},
		})
	}
}

func (h *HandlerHead) BHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := auth.GetUserFromContext(c)
		if !ok {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, core.StandardResponse{
			APIVersion: core.APIVersion,
			Context:    "test",
			Method:     c.Request().Method,
			Data: echo.Map{
				"username": user.Username,
				"email":    user.Email,
			},
		})
	}
}
