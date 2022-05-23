package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/auth"
)

func registerHandlers(server *echo.Echo, auth *auth.Auth) {
	server.POST("/login", auth.LoginHandler())
	server.POST("/register", auth.RegisterHandler())

	server.GET("/unrestricted", AHandler())
	server.GET("/restricted", BHandler(), middleware.JWTWithConfig(auth.JWTConfig))
}

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
