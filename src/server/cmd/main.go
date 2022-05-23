package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/api"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/handlers"
	"github.com/luisnquin/meow-app/src/server/repository"
	"github.com/luisnquin/meow-app/src/server/store"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.New, store.New, echo.New, repository.New),
		fx.Invoke(handlers.Mount, api.Mount),
	)

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
}
