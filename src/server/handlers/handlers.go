package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/models"
)

func AHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"mmm": "patas"})
	}
}

func BHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.ErrInternalServerError
		}

		claim, ok := token.Claims.(*models.Claims)
		if !ok {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{
			"email":    claim.Email,
			"username": claim.Username,
		})
	}
}
