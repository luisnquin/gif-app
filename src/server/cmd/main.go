//nolint:typecheck
package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/core"
	"github.com/luisnquin/gif-app/src/server/handlers"
	"github.com/luisnquin/gif-app/src/server/repository"
	"github.com/luisnquin/gif-app/src/server/store"
)

var (
	hosts = make(map[string]*echo.Echo)
	port  string
)

func main() {
	config := config.Load()

	flag.StringVar(&port, "port", config.Internal.Port, ":XXXX")
	flag.Parse()

	db, cache := store.New(config)
	provider := repository.New(db)

	// Hosts
	api, internal, app := echo.New(), echo.New(), echo.New()

	head := handlers.New(config, provider, db, cache)
	head.InternalMount(internal)
	head.APIMount(api)

	hosts["internal.localhost"+port] = internal
	hosts["api.localhost"+port] = api
	hosts["localhost"+port] = app

	server := echo.New()
	server.Use(middleware.CORS(), middleware.Logger(), middleware.Recover())

	server.Any("/*", func(c echo.Context) error {
		host := hosts[c.Request().Host]
		if host == nil {
			return echo.ErrNotFound
		}

		host.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	startup, wait, shutdown := core.GracefulShutdown(server)
	go startup(port)
	defer shutdown()

	wait()
}
