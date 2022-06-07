//nolint:typecheck
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

type host struct {
	*echo.Echo
}

var (
	hosts = make(map[string]*host)
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

	core.ApplyMiddlewares(api, core.Docs())

	handlers.New(app, config, provider, db, cache).Mount()

	hosts["internal.localhost"+port] = &host{internal}
	hosts["api.localhost"+port] = &host{api}
	hosts["localhost"+port] = &host{app}

	server := echo.New()
	core.ApplyMiddlewares(server)

	server.Any("/*", func(c echo.Context) error {
		host := hosts[c.Request().Host]

		if host == nil {
			return echo.ErrNotFound
		}

		host.Echo.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	startup, wait, shutdown := core.GracefulShutdown(server)
	go startup(port)
	defer shutdown()

	wait()
}
