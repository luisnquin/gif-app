package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/handlers"
	"github.com/luisnquin/meow-app/src/server/log"
	"github.com/luisnquin/meow-app/src/server/store"
	"github.com/luisnquin/meow-app/src/server/utils"
)

func main() {
	preLoad()

	server := echo.New()

	server.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	server.POST("/login", auth.LoginHandler())

	server.GET("/unrestricted", handlers.AHandler())
	server.GET("/restricted", handlers.BHandler(), middleware.JWTWithConfig(auth.Config))

	server.Logger.Fatal(server.Start(config.Server.Internal.Port))
}

func preLoad() {
	queue := utils.HandleExit(1)

	config.Load()

	disconnectDb := store.Connect(store.Postgres)

	queue <- func() { log.FatalWithCheck(disconnectDb()) }

	close(queue)
}
