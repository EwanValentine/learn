package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v9"
)

func New() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.Join(err, errors.New("error pinging redis"))
	}

	return client, nil
}
