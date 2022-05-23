package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/repository"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, config *config.Configuration, provider *repository.Provider) any {
	server := echo.New()

	auth := auth.New(config, provider)

	server.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	registerHandlers(server, auth)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Logger.Fatal(server.Start(config.Internal.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return 0
}
