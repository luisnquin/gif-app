package repository

import (
	"github.com/luisnquin/gif-app/src/server/store"
)

type Provider struct {
	db store.Querier
}

func New(store store.Querier) *Provider {
	return &Provider{
		db: store,
	}
}
