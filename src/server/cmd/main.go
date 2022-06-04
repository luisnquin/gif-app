package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/handlers"
	"github.com/luisnquin/meow-app/src/server/repository"
	"github.com/luisnquin/meow-app/src/server/store"
)

func main() {
	config := config.New()

	port := flag.String("port", config.Internal.Port, ":XXXX")

	flag.Parse()

	app := echo.New()

	app.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	db, _ := store.New(config)

	provider := repository.New(db)

	handlers.New(app, config, provider, db).Mount()

	app.Logger.Fatal(app.Start(*port))
}
