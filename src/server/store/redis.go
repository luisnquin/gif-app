package store

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/luisnquin/meow-app/src/server/config"
)

func initRedisClient(*config.Configuration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:5820",
		Username: "",
		Password: "",
	})

	status := client.Ping(context.Background())

	if err := status.Err(); err != nil {
		return nil, err
	}

	return client, nil
}
