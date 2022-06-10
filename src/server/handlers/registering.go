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
	app.GET("/health", store.HealthHandler(h.db, h.cache))
	app.POST("/automock", store.AutoMockHandler(h.db))
}

func (h *HandlerHead) registerAPIHandlers(app *echo.Echo) {
	app.GET("/docs", echoredoc.EchoHandler(h.config.Docs))
	app.File("/docs/openapi.yaml", h.config.Docs.SpecFile)
	app.File("/docs/favicon.png", "./docs/favicon.png")

	app.GET("/hi", BHandler(), middleware.JWTWithConfig(h.auth.JWTConfig))

	// app.Group("/:username", middleware.JWTWithConfig(h.auth.JWTConfig))

	/*
		/upload
		/post/<hash>
	*/
}

/*
	- The profile will be created at the same time the user is created

	GET .../:username/profile
	UPDATE .../:username/profile
	DELETE .../:username/profile (with user, the post's are not deleted, just without owner)

	GET .../:username/gifs
	GET .../:username/gif/:hash
	POST .../:username/gif
	PUT .../:username/gif/:hash
	DELETE .../:username/gif/:hash
	DELETE .../:username/gifs (Menu, request body)

	Additions:
	 - Block

	GET .../gifs (query params, recently<bool>, month<int> and/or year<int>)
	GET .../gifs/stats
	GET .../gif/:hash
	POST .../gif/:hash (like and comments)
	POST .../gif/:hash (report)

	GET .../reports (console)
	POST .../reports/:username

	SPONSOR
*/
