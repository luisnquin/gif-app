package handlers

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/gif-app/src/server/store"
)

func (h *HandlerHead) registerAuthHandlers() {
	h.app.POST("/signup", h.auth.SigInHandler())
	h.app.POST("/logout", h.auth.LogoutHandler())
	h.app.POST("/login", h.auth.LoginHandler())
}

func (h *HandlerHead) registerInternalHandlers() {
	h.app.GET("/health", store.HealthHandler(h.db))
	h.app.POST("/automock", store.AutoMockHandler(h.db))
}

func (h *HandlerHead) registerHandlers() {
	h.app.GET("/hi", BHandler(), middleware.JWTWithConfig(h.auth.JWTConfig))

	// rewards
	h.app.GET("/rewards", nil)
	// info
	h.app.GET("/ranges", nil)
	// redoc
	h.app.GET("/docs", nil)

	// certifications
	h.app.Group("/leaks")
	// posts, history
	h.app.Group("/:username", middleware.JWTWithConfig(h.auth.JWTConfig))
	// new, :id, latest
	h.app.Group("/reports", middleware.JWTWithConfig(h.auth.JWTConfig))
	// :id - oficial
	h.app.Group("/oficial/news", middleware.JWTWithConfig(h.auth.JWTConfig))
	// :id - secret
	h.app.Group("/secret/news", middleware.JWTWithConfig(h.auth.JWTConfig))
	// :id - possibly
	h.app.Group("/post", middleware.JWTWithConfig(h.auth.JWTConfig))
}
