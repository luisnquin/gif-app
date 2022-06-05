package store

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/utils"
)

func initRedisClient(config *config.Configuration) (*redis.Client, error) {
	var addr string

	if utils.IsRunningInADockerContainer() {
		addr = config.Cache.ContainerAddr
	} else {
		addr = config.Cache.LocalAddr
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: "",
		Password: "",
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
