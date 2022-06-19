//nolint:typecheck
package handlers

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/auth"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/controllers"
	"github.com/luisnquin/gif-app/src/server/provider"
)

type HandlerHead struct {
	controller *controllers.ServiceMan
	config     *config.Configuration
	provider   *provider.Queries
	cache      *redis.Client
	auth       *auth.Auth
	store      *sql.DB
}

func New(config *config.Configuration, provider *provider.Queries, store *sql.DB, cache *redis.Client) *HandlerHead {
	controller := controllers.NewServices(store, provider)
	auth := auth.New(config, provider)

	return &HandlerHead{
		controller: controller,
		provider:   provider,
		config:     config,
		cache:      cache,
		store:      store,
		auth:       auth,
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
