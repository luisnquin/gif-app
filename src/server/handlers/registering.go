//nolint:typecheck
package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/gif-app/src/server/store"
	echoredoc "github.com/luisnquin/go-redoc/echo"
)

func (h *HandlerHead) registerAuthHandlers(app *echo.Echo) {
	app.POST("/signup", h.auth.SignUpHandler())
	app.POST("/logout", h.auth.LogoutHandler())
	app.POST("/login", h.auth.LoginHandler())
}

func (h *HandlerHead) registerInternalHandlers(app *echo.Echo) {
	app.GET("/health", store.HealthHandler(h.store, h.cache))
	app.POST("/automock", store.AutoMockHandler(h.store))
}

func (h *HandlerHead) registerAPIHandlers(app *echo.Echo) {
	app.GET("/docs", echoredoc.EchoHandler(h.config.Docs))
	app.File("/docs/openapi.yaml", h.config.Docs.SpecFile)
	app.File("/docs/favicon.png", "./docs/favicon.png")

	app.GET("/hi", BHandler(), middleware.JWTWithConfig(h.auth.JWTConfig))

	app.GET("/:username/profile", nil)
	app.PUT("/:username/profile", nil)
	app.DELETE("/:username/profile", nil)

	app.GET("/:username/gifs", nil)
	app.GET("/:username/gifs/:hash", nil)
	app.POST("/:username/gif", nil)
	app.PUT("/:username/gifs/:hash", nil)
	app.DELETE("/:username/gifs/:hash", nil)
	app.DELETE("/:username/gifs", nil) // request body (Menu)

	app.GET("/stats", nil)

	app.GET("/gifs", nil)
	app.GET("/gifs/search", nil)
	app.GET("/gifs/:hash", nil)
	app.POST("/gif", nil)
	app.POST("/gifs/:hash", nil) // comments, etc
	app.POST("/gifs/:hash/report", nil)

	// Console
	app.GET("/reports", nil)
	app.POST("/reports/:hash", nil)
}
