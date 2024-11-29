package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	DB       int
	Username string
	Password string
}

type cli struct {
	client *redis.Client
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
	Get(ctx context.Context, key string, value interface{}) (err error)
	RedisClient() *redis.Client
}

func New(conf Config) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		DB:       conf.DB,
		Username: conf.Username,
		Password: conf.Password,
	})

	return &cli{
		client: client,
	}
}

func (r *cli) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *cli) Get(ctx context.Context, key string, value interface{}) (err error) {
	return r.client.Get(ctx, key).Scan(value)
}

func (r *cli) RedisClient() *redis.Client {
	return r.client
}
