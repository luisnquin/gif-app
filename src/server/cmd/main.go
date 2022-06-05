package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/handlers"
	"github.com/luisnquin/gif-app/src/server/middleware"
	"github.com/luisnquin/gif-app/src/server/repository"
	"github.com/luisnquin/gif-app/src/server/store"
)

func main() {
	config := config.Load()

	port := flag.String("port", config.Internal.Port, ":XXXX")

	flag.Parse()

	app := echo.New()
	db, _ := store.New(config)
	provider := repository.New(db)

	middleware.Apply(app)
	handlers.New(app, config, provider, db).Mount()

	app.Logger.Fatal(app.Start(*port))
}
