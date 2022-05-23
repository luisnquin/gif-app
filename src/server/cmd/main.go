package main

import (
	"context"

	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/handlers"
	"github.com/luisnquin/meow-app/src/server/repository"
	"github.com/luisnquin/meow-app/src/server/store"
	"go.uber.org/fx"
)

func main() {
	if err := fx.New(
		fx.Provide(config.New, store.New, repository.New, auth.New),
		fx.Invoke(handlers.New),
	).Start(context.Background()); err != nil {
		panic(err)
	}
}
