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

	profile := app.Group("/:username/profile", middleware.JWTWithConfig(h.auth.JWTConfig))
	profile.GET("", h.controller.GetFullProfile())
	profile.PUT("", nil)
	profile.DELETE("", nil)

	myGifs := app.Group("/:username/gifs", middleware.JWTWithConfig(h.auth.JWTConfig))
	myGifs.GET("", nil)
	myGifs.GET("/:hash", nil)
	myGifs.POST("", nil)
	myGifs.PUT("/:hash", nil)
	myGifs.DELETE("/:hash", nil)
	myGifs.DELETE("", nil) // request body (Menu)

	stats := app.Group("/stats")
	stats.GET("/", nil)

	allGifs := app.Group("/gifs")
	allGifs.GET("", nil)
	allGifs.GET("/search", nil)
	allGifs.GET("/:hash", nil)
	allGifs.POST("", nil)
	allGifs.POST("/:hash", nil) // comments, etc
	allGifs.POST("/:hash/report", nil)

	// Console
	console := app.Group("/console", middleware.JWTWithConfig(h.auth.JWTConfig))
	console.GET("/reports", nil)
	console.POST("/reports/:hash", nil)
}
