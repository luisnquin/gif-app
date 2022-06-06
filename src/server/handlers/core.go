package handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/auth"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/repository"
	"github.com/luisnquin/gif-app/src/server/store"
)

type HandlerHead struct {
	config   *config.Configuration
	provider *repository.Provider
	db       store.Querier
	cache    *redis.Client
	app      *echo.Echo
	auth     *auth.Auth
}

func New(app *echo.Echo, config *config.Configuration, provider *repository.Provider, db store.Querier, cache *redis.Client) *HandlerHead {
	auth := auth.New(config, provider)

	return &HandlerHead{
		provider: provider,
		config:   config,
		cache:    cache,
		auth:     auth,
		app:      app,
		db:       db,
	}
}

func (h *HandlerHead) Mount() {
	h.registerAuthHandlers()
	h.registerInternalHandlers()
	h.registerHandlers()
}
