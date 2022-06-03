package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/auth"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/repository"
)

type HandlerHead struct {
	config   *config.Configuration
	provider *repository.Provider
	app      *echo.Echo
	auth     *auth.Auth
}

func New(app *echo.Echo, config *config.Configuration, provider *repository.Provider) *HandlerHead {
	auth := auth.New(config, provider)

	return &HandlerHead{
		provider: provider,
		config:   config,
		auth:     auth,
		app:      app,
	}
}

func (h *HandlerHead) Mount() {
	h.registerAuthHandlers()
	h.registerInternalHandlers()
	h.registerHandlers()
}
