package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Apply(app *echo.Echo) {
	app.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	app.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"api-key", "Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowOrigins: []string{"*"},
	}))
}
