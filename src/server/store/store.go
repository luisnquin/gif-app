package store

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/luisnquin/gif-app/src/server/config"
)

func New(config *config.Configuration) (*sql.DB, *redis.Client) {
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
