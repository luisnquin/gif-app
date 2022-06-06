package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/core"
	"github.com/luisnquin/gif-app/src/server/handlers"
	"github.com/luisnquin/gif-app/src/server/repository"
	"github.com/luisnquin/gif-app/src/server/store"
)

func main() {
	config := config.Load()

	port := flag.String("port", config.Internal.Port, ":XXXX")

	flag.Parse()

	app := echo.New()
	db, cache := store.New(config)
	provider := repository.New(db)

	core.ApplyMainMiddlewares(app)
	handlers.New(app, config, provider, db, cache).Mount()

	startup, wait, shutdown := core.GracefulShutdown(app)
	go startup(*port)
	defer shutdown()

	wait()
}
