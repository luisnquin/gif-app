package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/repository"
)

func Mount(server *echo.Echo, config *config.Configuration, provider *repository.Provider) {
	auth := auth.New(server, config, provider)

	registerHandlers(server, auth)
}

func registerHandlers(server *echo.Echo, auth *auth.Auth) {
	server.POST("/signup", auth.RegisterHandler())
	server.POST("/logout", auth.LogoutHandler())
	server.POST("/login", auth.LoginHandler())

	// rewards
	server.GET("/rewards", nil)
	// info
	server.GET("/ranges", nil)
	// redoc
	server.GET("/docs", nil)

	// certifications
	server.Group("/leaks")
	// posts, history
	server.Group("/:username", middleware.JWTWithConfig(auth.JWTConfig))
	// new, :id, latest
	server.Group("/reports", middleware.JWTWithConfig(auth.JWTConfig))
	// :id - oficial
	server.Group("/oficial/news", middleware.JWTWithConfig(auth.JWTConfig))
	// :id - secret
	server.Group("/secret/news", middleware.JWTWithConfig(auth.JWTConfig))
	// :id - possibly
	server.Group("/post", middleware.JWTWithConfig(auth.JWTConfig))
}
