package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/repository"
)

func Mount(server *echo.Echo, config *config.Configuration, provider *repository.Provider) any {
	auth := auth.New(server, config, provider)

	registerHandlers(server, auth)

	return 0
}

func registerHandlers(server *echo.Echo, auth *auth.Auth) {
	server.POST("/login", auth.LoginHandler())
	server.POST("/register", auth.RegisterHandler())

	server.GET("/unrestricted", AHandler())
	server.GET("/restricted", BHandler(), middleware.JWTWithConfig(auth.JWTConfig))
}
