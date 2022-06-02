package api

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/config"
	"go.uber.org/fx"
)

func Mount(lc fx.Lifecycle, server *echo.Echo, config *config.Configuration) {
	server.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Logger.Fatal(server.Start(config.Internal.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
