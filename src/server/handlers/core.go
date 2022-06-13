//nolint:typecheck
package handlers

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/auth"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/provider"
)

type HandlerHead struct {
	config   *config.Configuration
	provider *provider.Queries
	store    *sql.DB
	cache    *redis.Client
	auth     *auth.Auth
}

func New(config *config.Configuration, provider *provider.Queries, store *sql.DB, cache *redis.Client) *HandlerHead {
	auth := auth.New(config, provider)

	return &HandlerHead{
		provider: provider,
		config:   config,
		cache:    cache,
		store:    store,
		auth:     auth,
	}
}

func (h *HandlerHead) APIMount(api *echo.Echo) {
	h.registerAuthHandlers(api)
	h.registerAPIHandlers(api)
}

func (h *HandlerHead) InternalMount(internal *echo.Echo) {
	h.registerInternalHandlers(internal)
	h.registerAuthHandlers(internal)
}
