package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/middleware"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/handlers"
	"github.com/luisnquin/meow-app/src/server/repository"
	"github.com/luisnquin/meow-app/src/server/store"
)

func main() {
	config := config.Load()

	port := flag.String("port", config.Internal.Port, ":XXXX")

	flag.Parse()

	app := echo.New()

	middleware.Apply(app)

	db, _ := store.New(config)

	provider := repository.New(db)

	handlers.New(app, config, provider, db).Mount()

	app.Logger.Fatal(app.Start(*port))
}
