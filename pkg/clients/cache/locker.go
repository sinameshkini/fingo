package cache

import (
	"github.com/bsm/redislock"
)

type Locker struct {
	*redislock.Client
}

func NewLocker(c Cache) (*Locker, error) {
	client := redislock.New(c.RedisClient())

	return &Locker{
		client,
	}, nil
}

func (l *Locker) SetKeyPrefix(key string) string {
	return "wallet:lock:" + key
}