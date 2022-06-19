package controllers

import (
	"github.com/luisnquin/gif-app/src/server/provider"
)

type ServiceMan struct {
	provider *provider.Queries
	store    provider.DBTX
}

func NewServices(store provider.DBTX, provider *provider.Queries) *ServiceMan {
	return &ServiceMan{
		provider: provider,
		store:    store,
	}
}
