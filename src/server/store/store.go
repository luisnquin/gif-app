package store

import (
	"github.com/go-redis/redis/v8"
	"github.com/luisnquin/gif-app/src/server/config"
)

func New(config *config.Configuration) (Querier, *redis.Client) {
	redisClient, err := initRedisClient(config)
	if err != nil {
		panic(err)
	}

	postgresClient, err := initPostgresClient(config)
	if err != nil {
		panic(err)
	}

	return postgresClient, redisClient
}
